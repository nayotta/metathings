package metathings_deviced_flow

import (
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	nonce_helper "github.com/nayotta/metathings/pkg/common/nonce"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type RedisStreamFlowSetOption struct {
	Id string

	ReadStreamGroupBlockTime time.Duration
	StreamExpireTime         time.Duration
	StreamTrimLimit          int64
	StreamTrimProb           float32
}

func NewRedisStreamFlowSetOption(id string) *RedisStreamFlowSetOption {
	return &RedisStreamFlowSetOption{
		Id:                       id,
		ReadStreamGroupBlockTime: 3 * time.Second,
		StreamExpireTime:         30 * time.Minute,
		StreamTrimLimit:          15,
		StreamTrimProb:           0.001,
	}
}

type RedisStreamFlowSet struct {
	opt    *RedisStreamFlowSetOption
	logger log.FieldLogger

	rs_cli *redis.Client

	close_cbs  []func() error
	close_once sync.Once
	closed     bool

	err error
}

func (rsfs *RedisStreamFlowSet) get_logger() log.FieldLogger {
	return rsfs.logger
}

func (rsfs *RedisStreamFlowSet) redis_stream_key() string {
	return "mt.flwst." + rsfs.Id()
}

func (rsfs *RedisStreamFlowSet) Id() string {
	return rsfs.opt.Id
}

func (rsfs *RedisStreamFlowSet) Err() error {
	return rsfs.err
}

func (rsfs *RedisStreamFlowSet) Close() (err error) {
	rsfs.close_once.Do(func() {
		for _, cb := range rsfs.close_cbs {
			if cerr := cb(); err != nil {
				rsfs.get_logger().WithError(err).Debugf("failed to call close callback")
				if err == nil {
					err = cerr
				}
			}
			rsfs.closed = true
		}
	})

	return err
}

func (rsfs *RedisStreamFlowSet) PushFrame(flwst_frm *FlowSetFrame) error {
	ts := pb_helper.Now()
	flwst_frm.Frame.Ts = &ts

	dev_txt, err := json_encoder.MarshalToString(flwst_frm.Device)
	if err != nil {
		return err
	}

	frm_txt, err := json_encoder.MarshalToString(flwst_frm.Frame)
	if err != nil {
		return err
	}

	if err = rsfs.rs_cli.XAdd(&redis.XAddArgs{
		Stream: rsfs.redis_stream_key(),
		Values: map[string]interface{}{
			"device": dev_txt,
			"frame":  frm_txt,
		},
	}).Err(); err != nil {
		return err
	}

	if err = rsfs.rs_cli.Expire(rsfs.redis_stream_key(), rsfs.opt.StreamExpireTime).Err(); err != nil {
		rsfs.logger.WithError(err).Debugf("failed to expire stream")
	}

	if rand_helper.Float32() < rsfs.opt.StreamTrimProb {
		if err = rsfs.rs_cli.XTrimApprox(rsfs.redis_stream_key(), rsfs.opt.StreamTrimLimit).Err(); err != nil {
			rsfs.logger.WithError(err).Debugf("failed to trim stream")
		}
	}

	return nil
}

func (rsfs *RedisStreamFlowSet) PullFrame() <-chan *FlowSetFrame {
	frm_ch := make(chan *FlowSetFrame)

	go rsfs.pull_frame_from_redis_stream_loop(frm_ch)

	return frm_ch
}

func (rsfs *RedisStreamFlowSet) pull_frame_from_redis_stream_loop(frm_ch chan<- *FlowSetFrame) {
	defer close(frm_ch)

	var err error
	cli := rsfs.rs_cli
	key := rsfs.redis_stream_key()
	logger := rsfs.get_logger()

	nonce := nonce_helper.GenerateNonce()
	if err = cli.XGroupCreateMkStream(key, nonce, "$").Err(); err != nil {
		logger.WithError(err).Debugf("failed to create redis stream group")
		return
	}
	defer func() {
		if err = cli.XGroupDestroy(key, nonce).Err(); err != nil {
			logger.WithError(err).Debugf("failed to destory redis stream group")
		}
	}()

	for !rsfs.closed {
		if err = cli.Expire(key, rsfs.opt.StreamExpireTime).Err(); err != nil {
			logger.WithError(err).Debugf("failed to set expire stream")
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
			logger.WithError(err).Debugf("failed to read redis stream")
			rsfs.err = err
			return
		}

		for _, val := range vals {
			for _, msg := range val.Messages {
				var frm pb.Frame
				var dev pb.Device

				if buf, ok := msg.Values["frame"]; ok {
					if err = json_decoder.Unmarshal(strings.NewReader(buf.(string)), &frm); err != nil {
						logger.WithError(err).Warningf("failed to unmarshal message to frame")
						return
					}
				}

				if buf, ok := msg.Values["device"]; ok {
					if err = json_decoder.Unmarshal(strings.NewReader(buf.(string)), &dev); err != nil {
						logger.WithError(err).Warningf("failed to unmarshal message to device")
					}
				}

				frm_ch <- &FlowSetFrame{
					Frame:  &frm,
					Device: &dev,
				}

				logger.Debugf("pull frame")
			}
		}
	}
}

type RedisStreamFlowSetFactoryOption struct {
	RedisStreamAddr        string
	RedisStreamDB          int
	RedisStreamPassword    string
	RedisStreamPoolInitial int
	RedisStreamPoolMax     int
}

type RedisStreamFlowSetFactory struct {
	opt    *RedisStreamFlowSetFactoryOption
	logger log.FieldLogger

	redis_stream_pool pool_helper.Pool
}

func (fty *RedisStreamFlowSetFactory) New(opt *FlowSetOption) (FlowSet, error) {
	cli, err := fty.redis_stream_pool.Get()
	if err != nil {
		fty.logger.WithError(err).Debugf("failed to get redis stream client in pool")
		return nil, err
	}

	return &RedisStreamFlowSet{
		opt:    NewRedisStreamFlowSetOption(opt.FlowSetId),
		rs_cli: cli.(*redis.Client),
		logger: fty.logger,
		close_cbs: []func() error{
			func() error { return fty.redis_stream_pool.Put(cli) },
		},
		closed: false,
	}, nil
}

func new_redis_stream_flow_set_factory(args ...interface{}) (FlowSetFactory, error) {
	var logger log.FieldLogger
	opt := new(RedisStreamFlowSetFactoryOption)

	opt.RedisStreamPoolInitial = 5
	opt.RedisStreamPoolMax = 23

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"redis_stream_addr":         opt_helper.ToString(&opt.RedisStreamAddr),
		"redis_stream_db":           opt_helper.ToInt(&opt.RedisStreamDB),
		"redis_stream_password":     opt_helper.ToString(&opt.RedisStreamPassword),
		"redis_stream_pool_initial": opt_helper.ToInt(&opt.RedisStreamPoolInitial),
		"redis_stream_pool_max":     opt_helper.ToInt(&opt.RedisStreamPoolMax),
		"logger":                    opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	rspool, err := pool_helper.NewPool(opt.RedisStreamPoolInitial, opt.RedisStreamPoolMax, func() (pool_helper.Client, error) {
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

	fty := &RedisStreamFlowSetFactory{
		opt:               opt,
		logger:            logger,
		redis_stream_pool: rspool,
	}

	return fty, nil
}

func init() {
	registry_flow_set_factory("redis", new_redis_stream_flow_set_factory)
}
