package metathings_deviced_connection

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	log "github.com/sirupsen/logrus"
)

type redisChannelOption struct {
	Id       string
	Side     Side
	Addr     string
	Password string
	Db       int
}

type redisChannel struct {
	opt    *redisChannelOption
	logger log.FieldLogger

	client *redis.Client

	sending   chan []byte
	receiving chan []byte

	client_once    *sync.Once
	sending_once   *sync.Once
	receiving_once *sync.Once
}

func newRedisChannel(opt *redisChannelOption, logger log.FieldLogger) Channel {
	return &redisChannel{
		opt:            opt,
		logger:         logger,
		client_once:    new(sync.Once),
		sending_once:   new(sync.Once),
		receiving_once: new(sync.Once),
	}
}

func (self *redisChannel) sender_key() string {
	return fmt.Sprintf("metathings.deviced.channel.%v.side.%v", self.opt.Id, self.opt.Side)
}

func (self *redisChannel) receiver_key() string {
	var s Side
	if self.opt.Side == NORTH_SIDE {
		s = SOUTH_SIDE
	} else {
		s = NORTH_SIDE
	}

	return fmt.Sprintf("metathings.deviced.channel.%v.side.%v", self.opt.Id, s)
}

func (self *redisChannel) init_client() {
	opts := &redis.Options{
		Addr: self.opt.Addr,
	}

	if self.opt.Db != 0 {
		opts.DB = self.opt.Db
	}

	if self.opt.Password != "" {
		opts.Password = self.opt.Password
	}

	self.client = redis.NewClient(opts)

	return
}

func (self *redisChannel) init_sending() {
	self.client_once.Do(self.init_client)

	logger := self.logger.WithFields(log.Fields{
		"key":   self.sender_key(),
		"side":  self.opt.Side,
		"event": "send",
	})

	self.sending = make(chan []byte)
	go func() {
		defer func() {
			close(self.sending)
			logger.Debugf("loop closed")
		}()

		for {
			select {
			case buf := <-self.sending:
				self.client.Publish(self.sender_key(), buf)
				logger.Debugf("send msg")
			}
		}
	}()
}

func (self *redisChannel) init_receiving() {
	self.client_once.Do(self.init_client)

	logger := self.logger.WithFields(log.Fields{
		"key":   self.receiver_key(),
		"side":  self.opt.Side,
		"event": "recv",
	})

	self.receiving = make(chan []byte)
	go func() {
		defer func() {
			close(self.receiving)
			logger.Debugf("loop closed")
		}()

		subpub := self.client.Subscribe(self.receiver_key())

		for {
			select {
			case msg, ok := <-subpub.Channel():
				if !ok {
					return
				}
				self.receiving <- []byte(msg.Payload)
				self.logger.Debugf("recv msg")
			}
		}
	}()

}

func (self *redisChannel) AsyncSend() chan<- []byte {
	self.sending_once.Do(self.init_sending)
	return self.sending
}

func (self *redisChannel) AsyncRecv() <-chan []byte {
	self.receiving_once.Do(self.init_receiving)
	return self.receiving
}

func (self *redisChannel) Send(buf []byte) error {
	self.AsyncSend() <- buf
	return nil
}

func (self *redisChannel) Recv() ([]byte, error) {
	buf := <-self.AsyncRecv()
	return buf, nil
}

func (self *redisChannel) Close() error {
	err := self.client.Close()
	if err != nil {
		return err
	}

	return nil
}

type redisBridgeOption struct {
	Id       string
	Addr     string
	Password string
	Db       int
}

type redisBridge struct {
	opt    *redisBridgeOption
	logger log.FieldLogger

	north Channel
	south Channel

	north_once *sync.Once
	south_once *sync.Once
}

func (self *redisBridge) init_north() {
	opt := &redisChannelOption{
		Id:       self.opt.Id,
		Addr:     self.opt.Addr,
		Password: self.opt.Password,
		Db:       self.opt.Db,
	}

	self.north = newRedisChannel(opt, self.logger)
}

func (self *redisBridge) init_south() {
}

func (self *redisBridge) Id() string {
	return self.opt.Id
}

func (self *redisBridge) North() Channel {
	self.north_once.Do(self.init_north)
	return self.north
}

func (self *redisBridge) South() Channel {
	self.south_once.Do(self.init_south)
	return self.south
}

func (self *redisBridge) Close() error {
	var err error

	if self.north != nil {
		if err = self.north.Close(); err != nil {
			return err
		}

	}

	if self.south != nil {
		if err = self.south.Close(); err != nil {
			return err
		}
	}

	return nil
}

type redisBridgeFactoryOption struct {
	Id       string
	Addr     string
	Password string
	Db       int
}

type redisBridgeFactory struct {
	opt    *redisBridgeFactoryOption
	logger log.FieldLogger
}

func (self *redisBridgeFactory) BuildBridge(device string, session int32) (Bridge, error) {
	return self.GetBridge(parse_bridge_id(device, session))
}

func (self *redisBridgeFactory) GetBridge(id string) (Bridge, error) {
	opt := &redisBridgeOption{
		Id:       id,
		Addr:     self.opt.Addr,
		Password: self.opt.Password,
		Db:       self.opt.Db,
	}

	return &redisBridge{
		opt:    opt,
		logger: self.logger.WithField("bridge", id),
	}, nil
}

func new_redis_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	var ok bool
	var logger log.FieldLogger
	var err error

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	opt := &redisBridgeFactoryOption{}

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": func(k string, v interface{}) error {
			logger, ok = v.(log.FieldLogger)
			if !ok {
				return ErrInvalidArgument
			}
			logger = logger.WithField("#bridge_driver", "redis")
			return nil
		},
		"addr": func(k string, v interface{}) error {
			opt.Addr, ok = v.(string)
			if !ok {
				return ErrInvalidArgument
			}
			return nil
		},
		"password": func(k string, v interface{}) error {
			opt.Password, ok = v.(string)
			if !ok {
				return ErrInvalidArgument
			}
			return nil
		},
		"db": func(k string, v interface{}) error {
			opt.Db, ok = v.(int)
			if !ok {
				return ErrInvalidArgument
			}
			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return &redisBridgeFactory{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_bridge_factory_factory("redis", new_redis_bridge_factory)
}
