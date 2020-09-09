package metathings_deviced_connection

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	nonce_helper "github.com/nayotta/metathings/pkg/common/nonce"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
)

type redisStreamChannelOption struct {
	Id   string
	Side Side

	ReadStreamGroupBlockTime time.Duration
	StreamExpireTime         time.Duration
	ReceiveTimeout           time.Duration
}

func newRedisStreamChannelOption(id string, side Side) *redisStreamChannelOption {
	return &redisStreamChannelOption{
		Id:                       id,
		Side:                     side,
		ReadStreamGroupBlockTime: 3 * time.Second,
		StreamExpireTime:         3 * time.Minute,
		ReceiveTimeout:           31 * time.Second,
	}
}

func newRedisStreamChannel(opt *redisStreamChannelOption, pool pool_helper.Pool, logger log.FieldLogger) *redisStreamChannel {
	return &redisStreamChannel{
		opt:    opt,
		logger: logger,
		closed: make(chan struct{}),
		pool:   pool,
	}
}

type redisStreamChannel struct {
	opt    *redisStreamChannelOption
	logger log.FieldLogger

	async_recv_once sync.Once
	receiving       chan []byte
	client_mutex    sync.Mutex

	closed chan struct{}

	pool   pool_helper.Pool
	client client_helper.RedisClient
}

func (c *redisStreamChannel) get_logger() log.FieldLogger {
	return c.logger
}

func (c *redisStreamChannel) get_client() (client_helper.RedisClient, error) {
	c.client_mutex.Lock()
	defer c.client_mutex.Unlock()

	if c.client == nil {
		cli, err := c.pool.Get()
		if err != nil {
			c.get_logger().WithError(err).Debugf("failed to get redis client")
			return nil, err
		}
		c.get_logger().WithField("pool_size", c.pool.Size()).Debugf("get client from pool")

		// TODO(Peer): check client alive or not.
		c.client = cli.(client_helper.RedisClient)
	}

	return c.client, nil
}

func (c *redisStreamChannel) get_context() context.Context {
	return context.TODO()
}

func (c *redisStreamChannel) another_side() Side {
	if c.opt.Side == NORTH_SIDE {
		return SOUTH_SIDE
	} else {
		return NORTH_SIDE
	}
}

func (c *redisStreamChannel) op_key(id string, from, to Side) string {
	return fmt.Sprintf("mt.devd.ch.%v.%v.%v", id, from, to)
}

func (c *redisStreamChannel) sending_key() string {
	return c.op_key(c.opt.Id, c.opt.Side, c.another_side())
}

func (c *redisStreamChannel) receiving_key() string {
	return c.op_key(c.opt.Id, c.another_side(), c.opt.Side)
}

func (c *redisStreamChannel) Send(buf []byte) error {
	ctx := c.get_context()
	key := c.sending_key()
	logger := c.get_logger().WithFields(log.Fields{
		"key":   key,
		"side":  c.opt.Side,
		"event": "send",
	})

	client, err := c.get_client()
	if err != nil {
		return err
	}

	err = client.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		Values: map[string]interface{}{
			"value": buf,
		},
	}).Err()
	if err != nil {
		logger.WithError(err).Debugf("failed to send msg")
		return err
	}
	err = client.Expire(ctx, key, c.opt.StreamExpireTime).Err()
	if err != nil {
		logger.WithError(err).Warningf("failed to set expire for stream")
	}

	return nil
}

func (c *redisStreamChannel) Recv() ([]byte, error) {
	select {
	case buf := <-c.AsyncRecv():
		return buf, nil
	case <-time.After(c.opt.ReceiveTimeout):
		return nil, ErrReceiveTimeout
	case <-c.closed:
		return nil, ErrChannelClosed
	}
}

func (c *redisStreamChannel) AsyncSend() chan<- []byte {
	panic("unimplemented")
}

func (c *redisStreamChannel) init_async_recv() {
	key := c.receiving_key()

	logger := c.get_logger().WithFields(log.Fields{
		"key":   key,
		"group": c.opt.Id,
		"side":  c.opt.Side,
		"event": "recv",
	})

	c.receiving = make(chan []byte)
	go func() {
		defer func() {
			close(c.receiving)
			logger.Debugf("loop closed")
		}()

		client, err := c.get_client()
		if err != nil {
			return
		}

		ctx := c.get_context()

		err = client.XGroupCreateMkStream(ctx, key, c.opt.Id, "$").Err()
		if err != nil {
			logger.WithError(err).Debugf("failed to create redis stream group")
			return
		}
		err = client.Expire(ctx, key, c.opt.StreamExpireTime).Err()
		if err != nil {
			logger.WithError(err).Warningf("failed to set expire for stream")
		}

		for {
			select {
			case <-c.closed:
				return
			default:
			}

			// refresh stream expire time
			err = client.Expire(ctx, key, c.opt.StreamExpireTime).Err()
			if err != nil {
				logger.WithError(err).Warningf("failed to set expire for stream")
			}

			vals, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    c.opt.Id,
				Consumer: c.opt.Id,
				Streams:  []string{key, ">"},
				Count:    1,
				Block:    c.opt.ReadStreamGroupBlockTime,
				NoAck:    true,
			}).Result()

			switch err {
			case redis.Nil:
				continue
			case nil:
			default:
				// TODO(Peer): handle error
				logger.WithError(err).Debugf("recv error")
				return

			}

			for _, val := range vals {
				for _, msg := range val.Messages {
					if buf, ok := msg.Values["value"]; ok {
						c.receiving <- []byte(buf.(string))
					}
				}
			}
		}
	}()

	logger.Debugf("init async recv")
}

func (c *redisStreamChannel) AsyncRecv() <-chan []byte {
	c.async_recv_once.Do(c.init_async_recv)
	return c.receiving
}

func (c *redisStreamChannel) Close() error {
	c.client_mutex.Lock()
	defer c.client_mutex.Unlock()

	close(c.closed)
	if err := c.pool.Put(c.client); err != nil {
		c.get_logger().WithError(err).Debugf("failed to put client back to pool")
		return err
	}
	c.get_logger().WithField("pool_size", c.pool.Size()).Debugf("put client to pool")

	c.get_logger().Debugf("channel closed")

	return nil
}

type redisStreamBridgeOption struct {
	Id string
}

type redisStreamBridge struct {
	opt    *redisStreamBridgeOption
	logger log.FieldLogger

	pool pool_helper.Pool

	north Channel
	south Channel

	north_once sync.Once
	south_once sync.Once
	mtx        sync.Mutex
}

func (f *redisStreamBridge) Id() string {
	return f.opt.Id
}

func (f *redisStreamBridge) init_north() {
	opt := newRedisStreamChannelOption(f.Id(), NORTH_SIDE)
	f.north = newRedisStreamChannel(opt, f.pool, f.logger.WithFields(log.Fields{
		"side": "north",
	}))
}

func (f *redisStreamBridge) North() Channel {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	f.north_once.Do(f.init_north)
	return f.north
}

func (f *redisStreamBridge) init_south() {
	opt := newRedisStreamChannelOption(f.Id(), SOUTH_SIDE)
	f.south = newRedisStreamChannel(opt, f.pool, f.logger.WithFields(log.Fields{
		"side": "south",
	}))
}

func (f *redisStreamBridge) South() Channel {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	f.south_once.Do(f.init_south)
	return f.south
}

func (f *redisStreamBridge) Close() error {
	var err error

	f.mtx.Lock()
	defer f.mtx.Unlock()

	if f.north != nil {
		if err = f.north.Close(); err != nil {
			return err
		}
	} else {
		f.logger.Debugf("north channel is empty")
	}

	if f.south != nil {
		if err = f.south.Close(); err != nil {
			return err
		}
	} else {
		f.logger.Debugf("south channel is empty")
	}

	f.logger.Debugf("bridge closed")
	return nil
}

type redisStreamFactoryOption struct {
	Redis struct {
		Args []interface{}
	}
	Pool struct {
		Initial int
		Max     int
	}
}

type redisStreamBridgeFactory struct {
	opt    *redisStreamFactoryOption
	logger log.FieldLogger

	pool pool_helper.Pool

	init_once sync.Once
}

func (f *redisStreamBridgeFactory) init_pool() {
	var err error

	f.pool, err = pool_helper.NewPool(f.opt.Pool.Initial, f.opt.Pool.Max, func() (pool_helper.Client, error) {
		return client_helper.NewRedisClient(f.opt.Redis.Args...)
	})

	if err != nil {
		f.logger.WithError(err).Fatalf("failed to init pool")
	}
}

func (f *redisStreamBridgeFactory) BuildBridge(device_id string, sess int64) (Bridge, error) {
	id := parse_bridge_id(device_id, sess)
	defer f.logger.WithFields(log.Fields{
		"bridge":  id,
		"session": sess,
	}).Debugf("build bridge")
	return f.get_bridge(id)
}

func (f *redisStreamBridgeFactory) get_bridge(id string) (Bridge, error) {
	f.init_once.Do(f.init_pool)

	opt := &redisStreamBridgeOption{
		Id: id,
	}

	br := &redisStreamBridge{
		opt: opt,
		logger: f.logger.WithFields(log.Fields{
			"bridge": id,
			"nonce":  nonce_helper.GenerateNonce(),
		}),
		pool: f.pool,
	}

	return br, nil
}

func (f *redisStreamBridgeFactory) GetBridge(id string) (Bridge, error) {
	f.logger.WithField("bridge", id).Debugf("get bridge")
	return f.get_bridge(id)
}

func new_redis_stream_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	var logger log.FieldLogger
	var err error

	opt := &redisStreamFactoryOption{}
	opt.Pool.Initial = 5
	opt.Pool.Max = 23
	opt.Redis.Args = args

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":       opt_helper.ToLogger(&logger),
		"pool_initial": opt_helper.ToInt(&opt.Pool.Initial),
		"pool_max":     opt_helper.ToInt(&opt.Pool.Max),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	logger = logger.WithField("#bridge_driver", "redis-stream")

	return &redisStreamBridgeFactory{
		opt:    opt,
		logger: logger,
	}, nil
}

var register_redis_stream_bridge_factory_factory_once sync.Once

func init() {
	register_redis_stream_bridge_factory_factory_once.Do(func() {
		register_bridge_factory_factory("redis-stream", new_redis_stream_bridge_factory)
	})
}
