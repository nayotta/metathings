package metathings_deviced_flow

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/jsonpb"
	struct_ "github.com/golang/protobuf/ptypes/struct"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	mongo_helper "github.com/nayotta/metathings/pkg/common/mongo"
	nonce_helper "github.com/nayotta/metathings/pkg/common/nonce"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	json_encoder = jsonpb.Marshaler{}
	json_decoder = jsonpb.Unmarshaler{}
)

type flow struct {
	opt       *FlowOption
	close_cbs []func() error
	logger    log.FieldLogger

	mgo_db *mongo.Database
	rs_cli *redis.Client
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

	for _, cb := range f.close_cbs {
		if err = cb(); err != nil {
			f.logger.WithError(err).Debugf("failed to close callback")
		}
	}

	return err
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
	frm_txt, err := json_encoder.MarshalToString(frm)
	if err != nil {
		return err
	}

	if err := f.rs_cli.XAdd(&redis.XAddArgs{
		Stream: f.redis_stream_key(),
		Values: map[string]interface{}{
			"value": frm_txt,
		},
	}).Err(); err != nil {
		return err
	}

	return nil
}

func (f *flow) push_frame_to_mgo(frm *pb.Frame) error {
	frm_dat := frm.GetData()
	frm_dat_txt, err := json_encoder.MarshalToString(frm_dat)
	if err != nil {
		return err
	}

	frm_dat_buf := bson.M{}
	err = bson.UnmarshalExtJSON([]byte(frm_dat_txt), true, &frm_dat_buf)
	if err != nil {
		return err
	}
	frm_dat_buf["#ts"] = pb_helper.ToTime(*frm.GetTs()).UnixNano()

	coll := f.mongo_collection()
	_, err = coll.InsertOne(f.context(), frm_dat_buf)
	if err != nil {
		return err
	}

	return nil
}

func (f *flow) pull_frame_from_redis_stream() (<-chan *pb.Frame, <-chan struct{}) {
	frm_ch := make(chan *pb.Frame)
	quit_ch := make(chan struct{})

	go f.pull_frame_from_redis_stream_loop(frm_ch, quit_ch)

	return frm_ch, quit_ch
}

func (f *flow) pull_frame_from_redis_stream_loop(frm_ch chan<- *pb.Frame, quit_ch chan struct{}) {
	defer close(frm_ch)

	var err error
	cli := f.rs_cli
	key := f.redis_stream_key()

	nonce := nonce_helper.GenerateNonce()

	if err = cli.XGroupCreateMkStream(key, nonce, "$").Err(); err != nil {
		f.logger.WithError(err).Debugf("failed to create redis stream group")
		return
	}
	defer func() {
		if err = cli.XGroupDestroy(key, nonce).Err(); err != nil {
			f.logger.WithError(err).Debugf("failed to destory redis stream group")
		}

	}()

	for {
		select {
		case <-quit_ch:
			f.logger.Debugf("catch quit signal from outside")
			return
		default:
		}

		vals, err := cli.XReadGroup(&redis.XReadGroupArgs{
			Group:    nonce,
			Consumer: nonce,
			Streams:  []string{key, ">"},
			Count:    1,
			Block:    3 * time.Second,
			NoAck:    true,
		}).Result()

		switch err {
		case redis.Nil:
			continue
		case nil:
		default:
			f.logger.WithError(err).Debugf("failed to read redis stream")
			return
		}

		for _, val := range vals {
			for _, msg := range val.Messages {
				if buf, ok := msg.Values["value"]; ok {
					var frm pb.Frame
					if err = json_decoder.Unmarshal(strings.NewReader(buf.(string)), &frm); err != nil {
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

func (f *flow) PullFrame() (<-chan *pb.Frame, <-chan struct{}) {
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

	dec := jsonpb.Unmarshaler{}
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

		var frm_dat struct_.Struct
		err = dec.Unmarshal(bytes.NewReader(frm_dat_txt), &frm_dat)
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

func (ff *flowFactory) New(opt *FlowOption) (Flow, error) {
	cli, err := ff.mongo_pool.Get()
	if err != nil {
		ff.logger.WithError(err).Debugf("failed to get mongo client in pool")
		return nil, err
	}

	mgo_cli := cli.(*mongo_helper.MongoClientWrapper)
	mgo_db := mgo_cli.Database(ff.opt.MongoDatabase)

	cli, err = ff.redis_stream_pool.Get()
	if err != nil {
		ff.logger.WithError(err).Debugf("failed to get redis stream client in pool")
		return nil, err
	}

	rs_cli := cli.(*redis.Client)

	return &flow{
		opt:    opt,
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
	}, nil
}

func new_default_flow_factory(args ...interface{}) (FlowFactory, error) {
	var logger log.FieldLogger
	opt := new(flowFactoryOption)

	opt.MongoPoolInitial = 2
	opt.MongoPoolMax = 5
	opt.RedisStreamPoolInitial = 5
	opt.RedisStreamPoolMax = 23

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"redis_stream_addr":     opt_helper.ToString(&opt.RedisStreamAddr),
		"redis_stream_db":       opt_helper.ToInt(&opt.RedisStreamDB),
		"redis_stream_password": opt_helper.ToString(&opt.RedisStreamPassword),
		"mongo_uri":             opt_helper.ToString(&opt.MongoUri),
		"mongo_database":        opt_helper.ToString(&opt.MongoDatabase),
		"logger":                opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	rspool, err := pool_helper.NewPool(opt.RedisStreamPoolInitial, opt.MongoPoolMax, func() (pool_helper.Client, error) {
		rdopt := &redis.Options{
			Addr:     opt.RedisStreamAddr,
			DB:       opt.RedisStreamDB,
			Password: opt.RedisStreamPassword,
		}

		return redis.NewClient(rdopt), nil
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
