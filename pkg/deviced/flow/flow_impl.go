package metathings_deviced_flow

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	stpb "github.com/golang/protobuf/ptypes/struct"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	mongo_helper "github.com/nayotta/metathings/pkg/common/mongo"
	nonce_helper "github.com/nayotta/metathings/pkg/common/nonce"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type flowOption struct {
	*FlowOption

	ReadStreamGroupBlockTime time.Duration
	StreamExpireTime         time.Duration
	StreamTrimLimit          int64
	StreamTrimProb           float32
}

func newFlowOption(o *FlowOption) *flowOption {
	return &flowOption{
		FlowOption:               o,
		ReadStreamGroupBlockTime: 3 * time.Second,
		StreamExpireTime:         30 * time.Minute,
		StreamTrimLimit:          15,
		StreamTrimProb:           0.001,
	}
}

type flow struct {
	opt    *flowOption
	logger log.FieldLogger

	mgo_db *mongo.Database
	rs_cli client_helper.RedisClient

	close_cbs  []func() error
	close_once sync.Once
	closed     bool

	err error
}

func (f *flow) mongo_collection() *mongo.Collection {
	return f.mgo_db.Collection("mtflw." + f.Id())
}

func (f *flow) redis_stream_key() string {
	return "mt.flw." + f.Id()
}

func (f *flow) context() context.Context {
	return context.TODO()
}

func (f *flow) Id() string {
	return f.opt.FlowId
}

func (f *flow) Device() string {
	return f.opt.DeviceId
}

func (f *flow) Close() error {
	var err error

	f.close_once.Do(func() {
		for _, cb := range f.close_cbs {
			if cerr := cb(); err != nil {
				f.logger.WithError(err).Debugf("failed to close callback")
				if err == nil {
					err = cerr
				}
			}
		}
		f.closed = true
	})

	return err
}

func (f *flow) Err() error {
	return f.err
}

func (f *flow) PushFrame(frm *pb.Frame) error {
	ts := pb_helper.Now()
	frm.Ts = &ts

	err := f.push_frame_to_mgo(frm)
	if err != nil {
		f.logger.WithError(err).Errorf("failed to push frame to mgo")
		return err
	}

	// TODO(Peer): dont push frame to redis stream when noone pull frame.
	err = f.push_frame_to_redis_stream(frm)
	if err != nil {
		f.logger.WithError(err).Errorf("failed to push frame to redis stream")
		return err
	}

	f.logger.Debugf("push frame")

	return nil
}

func (f *flow) push_frame_to_redis_stream(frm *pb.Frame) error {
	ctx := f.context()

	frm_txt, err := grpc_helper.JSONPBMarshaler.MarshalToString(frm)
	if err != nil {
		return err
	}

	if err := f.rs_cli.XAdd(ctx, &redis.XAddArgs{
		Stream: f.redis_stream_key(),
		Values: map[string]interface{}{
			"frame": frm_txt,
		},
	}).Err(); err != nil {
		return err
	}

	if err = f.rs_cli.Expire(ctx, f.redis_stream_key(), f.opt.StreamExpireTime).Err(); err != nil {
		f.logger.WithError(err).Debugf("failed to expire stream")
	}

	if rand_helper.Float32() < f.opt.StreamTrimProb {
		if err = f.rs_cli.XTrimApprox(ctx, f.redis_stream_key(), f.opt.StreamTrimLimit).Err(); err != nil {
			f.logger.WithError(err).Debugf("failed to trim stream")
		}
	}

	return nil
}

func (f *flow) push_frame_to_mgo(frm *pb.Frame) error {
	frm_dat := frm.GetData()
	frm_dat_txt, err := grpc_helper.JSONPBMarshaler.MarshalToString(frm_dat)
	if err != nil {
		return err
	}

	frm_dat_buf := bson.M{}
	err = bson.UnmarshalExtJSON([]byte(frm_dat_txt), true, &frm_dat_buf)
	if err != nil {
		return err
	}
	frm_dat_buf["#ts"] = pb_helper.ToTime(frm.GetTs()).UnixNano()

	coll := f.mongo_collection()
	_, err = coll.InsertOne(f.context(), frm_dat_buf)
	if err != nil {
		return err
	}

	return nil
}

func (f *flow) pull_frame_from_redis_stream() <-chan *pb.Frame {
	frm_ch := make(chan *pb.Frame)

	go f.pull_frame_from_redis_stream_loop(frm_ch)

	return frm_ch
}

func (f *flow) pull_frame_from_redis_stream_loop(frm_ch chan<- *pb.Frame) {
	defer close(frm_ch)

	var err error
	cli := f.rs_cli
	key := f.redis_stream_key()
	ctx := f.context()

	nonce := nonce_helper.GenerateNonce()
	if err = cli.XGroupCreateMkStream(ctx, key, nonce, "$").Err(); err != nil {
		f.logger.WithError(err).Debugf("failed to create redis stream group")
		return
	}
	defer func() {
		if err = cli.XGroupDestroy(ctx, key, nonce).Err(); err != nil {
			f.logger.WithError(err).Debugf("failed to destory redis stream group")
		}

	}()

	for !f.closed {
		if err = cli.Expire(ctx, key, f.opt.StreamExpireTime).Err(); err != nil {
			f.logger.WithError(err).Debugf("failed to expire stream")
		}

		vals, err := cli.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    nonce,
			Consumer: nonce,
			Streams:  []string{key, ">"},
			Count:    1,
			Block:    f.opt.ReadStreamGroupBlockTime,
			NoAck:    true,
		}).Result()
		switch err {
		case redis.Nil:
			continue
		case nil:
		default:
			f.logger.WithError(err).Debugf("failed to read redis stream")
			f.err = err
			return
		}

		for _, val := range vals {
			for _, msg := range val.Messages {
				if buf, ok := msg.Values["frame"]; ok {
					var frm pb.Frame
					if err = grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(buf.(string)), &frm); err != nil {
						f.logger.WithError(err).Warningf("failed to unmarshal message to frame")
						return
					}

					frm_ch <- &frm
					f.logger.Debugf("pull frame")
				}
			}
		}
	}
}

func (f *flow) PullFrame() <-chan *pb.Frame {
	return f.pull_frame_from_redis_stream()
}

func (f *flow) QueryFrame(flrs ...*FlowFilter) ([]*pb.Frame, error) {
	var ret []*pb.Frame

	coll := f.mongo_collection()

	for _, flr := range flrs {
		frms, err := f.query_frame(coll, flr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, frms...)
	}

	return ret, nil
}

func (f *flow) query_frame(coll *mongo.Collection, flr *FlowFilter) ([]*pb.Frame, error) {
	flr_buf := bson.M{}

	if !flr.BeginAt.Equal(time.Time{}) {
		flr_buf["$gte"] = flr.BeginAt.UnixNano()
	}

	if !flr.EndAt.Equal(time.Time{}) {
		flr_buf["$lte"] = flr.EndAt.UnixNano()
	}

	cur, err := coll.Find(f.context(), bson.M{"#ts": flr_buf})
	if err != nil {
		return nil, err
	}

	var frms []*pb.Frame
	for cur.Next(f.context()) {
		var res_buf bson.M
		if err = cur.Decode(&res_buf); err != nil {
			return nil, err
		}

		ts, ok := res_buf["#ts"]
		if !ok {
			ts = nil
		}
		delete(res_buf, "#ts")
		delete(res_buf, "_id")

		ts_int64, ok := ts.(int64)
		if !ok {
			ts_int64 = 0
		}

		frm_dat_txt, err := json.Marshal(res_buf)
		if err != nil {
			return nil, err
		}

		var frm_dat stpb.Struct
		err = grpc_helper.JSONPBUnmarshaler.Unmarshal(bytes.NewReader(frm_dat_txt), &frm_dat)
		if err != nil {
			return nil, err
		}

		frm_ts := pb_helper.FromTimestamp(ts_int64)
		frm := &pb.Frame{
			Ts:   &frm_ts,
			Data: &frm_dat,
		}

		frms = append(frms, frm)
	}

	return frms, nil
}

type flowFactoryOption struct {
	MongoUri         string
	MongoDatabase    string
	MongoPoolInitial int
	MongoPoolMax     int

	RedisStreamAddr        string
	RedisStreamDB          int
	RedisStreamPassword    string
	RedisStreamPoolInitial int
	RedisStreamPoolMax     int
}

type flowFactory struct {
	opt    *flowFactoryOption
	logger log.FieldLogger

	redis_stream_pool pool_helper.Pool
	mongo_pool        pool_helper.Pool
}

func (ff *flowFactory) context() context.Context {
	return context.TODO()
}

func (ff *flowFactory) get_alive_redis_stream_client() (client_helper.RedisClient, error) {
	// TODO(Peer): max retry should greater than redis stream pool max size, magic number here.
	for i := 0; i < 6; i++ {
		ctx := ff.context()

		cli, err := ff.redis_stream_pool.Get()
		if err != nil {
			ff.logger.WithError(err).Debugf("failed to get redis stream client in pool")
			return nil, err
		}

		rs_cli := cli.(client_helper.RedisClient)
		if err = rs_cli.Ping(ctx).Err(); err != nil {
			defer rs_cli.Close()
			continue
		}

		return rs_cli, nil
	}

	return nil, ErrGetAliveRedisClientMaxRetry
}

func (ff *flowFactory) New(opt *FlowOption) (Flow, error) {
	cli, err := ff.mongo_pool.Get()
	if err != nil {
		ff.logger.WithError(err).Debugf("failed to get mongo client in pool")
		return nil, err
	}

	mgo_cli := cli.(*mongo_helper.MongoClientWrapper)
	mgo_db := mgo_cli.Database(ff.opt.MongoDatabase)

	rs_cli, err := ff.get_alive_redis_stream_client()
	if err != nil {
		ff.logger.WithError(err).Debugf("failed to get alive redis client")
		return nil, err
	}

	return &flow{
		opt:    newFlowOption(opt),
		mgo_db: mgo_db,
		rs_cli: rs_cli,
		logger: ff.logger,
		close_cbs: []func() error{
			func() error {
				return ff.mongo_pool.Put(mgo_cli)
			},
			func() error {
				return ff.redis_stream_pool.Put(rs_cli)
			},
		},
		closed: false,
	}, nil
}

func new_default_flow_factory(args ...interface{}) (FlowFactory, error) {
	var logger log.FieldLogger
	opt := new(flowFactoryOption)

	opt.MongoPoolInitial = 2
	opt.MongoPoolMax = 5
	opt.RedisStreamPoolInitial = 5
	opt.RedisStreamPoolMax = 23

	redis_config, err := cfg_helper.FoldConfigOption(args, "redis")
	if err != nil {
		return nil, err
	}

	mongo_config, err := cfg_helper.FoldConfigOption(args, "mongo")
	if err != nil {
		return nil, err
	}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	redis_args := cfg_helper.FlattenConfigOption(redis_config)
	mongo_args := cfg_helper.FlattenConfigOption(mongo_config)

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"pool_initial": opt_helper.ToInt(&opt.RedisStreamPoolInitial),
		"pool_max":     opt_helper.ToInt(&opt.RedisStreamPoolMax),
	}, opt_helper.SetSkip(true))(redis_args...); err != nil {
		return nil, err
	}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"pool_initial": opt_helper.ToInt(&opt.MongoPoolInitial),
		"pool_max":     opt_helper.ToInt(&opt.MongoPoolMax),
		"uri":          opt_helper.ToString(&opt.MongoUri),
		"database":     opt_helper.ToString(&opt.MongoDatabase),
	}, opt_helper.SetSkip(true))(mongo_args...); err != nil {
		return nil, err
	}

	rspool, err := pool_helper.NewPool(opt.RedisStreamPoolInitial, opt.RedisStreamPoolMax, func() (pool_helper.Client, error) {
		return client_helper.NewRedisClient(redis_args...)
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to new redis stream pool")
		return nil, err
	}

	mgopool, err := pool_helper.NewPool(opt.MongoPoolInitial, opt.MongoPoolMax, func() (pool_helper.Client, error) {
		return mongo_helper.NewMongoClient(opt.MongoUri)
	})
	if err != nil {
		logger.WithError(err).Debugf("failed to new mongo pool")
		return nil, err
	}

	fty := &flowFactory{
		opt:               opt,
		logger:            logger,
		redis_stream_pool: rspool,
		mongo_pool:        mgopool,
	}

	return fty, nil
}

func init() {
	register_flow_factory("default", new_default_flow_factory)
}
