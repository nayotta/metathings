package metathings_deviced_connection

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type saramaChannelOption struct {
	Id      string
	Side    Side
	Brokers []string
}

func newSaramaChannel(opt *saramaChannelOption, client sarama.Client, logger log.FieldLogger) Channel {
	return &saramaChannel{
		opt:             opt,
		logger:          logger,
		client:          client,
		async_recv_once: new(sync.Once),
		sync_send_once:  new(sync.Once),
	}
}

type saramaChannel struct {
	opt    *saramaChannelOption
	logger log.FieldLogger

	client   sarama.Client
	producer sarama.SyncProducer
	consumer *cluster.Consumer

	sending   chan []byte
	receiving chan []byte

	recv_err error

	async_recv_once *sync.Once
	sync_send_once  *sync.Once
}

func (self *saramaChannel) producer_topic() string {
	var another Side
	if self.opt.Side == NORTH_SIDE {
		another = SOUTH_SIDE
	} else {
		another = NORTH_SIDE
	}
	return fmt.Sprintf("metathings.deviced.channel.%v.from.%v.to.%v", self.opt.Id, self.opt.Side, another)
}

func (self *saramaChannel) consumer_topic() string {
	var another Side
	if self.opt.Side == NORTH_SIDE {
		another = SOUTH_SIDE
	} else {
		another = NORTH_SIDE
	}
	return fmt.Sprintf("metathings.deviced.channel.%v.from.%v.to.%v", self.opt.Id, another, self.opt.Side)
}

func (self *saramaChannel) init_producer() {
	var err error

	if self.producer, err = sarama.NewSyncProducerFromClient(self.client); err != nil {
		panic(err)
	}
	self.logger.WithField("topic", self.producer_topic()).Debugf("init producer")
}

func (self *saramaChannel) init_consumer() {
	var err error

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Group.Heartbeat.Interval = 30 * time.Millisecond

	if self.consumer, err = cluster.NewConsumer(self.opt.Brokers, self.opt.Id, []string{self.consumer_topic()}, config); err != nil {
		self.logger.WithError(err).Warningf("failed to init consumer")
		return
	}
	self.logger.WithFields(log.Fields{
		"group": self.opt.Id,
		"topic": self.consumer_topic(),
	}).Debugf("init consumer")

	return
}

func (self *saramaChannel) init_async_recv() {
	self.init_consumer()

	logger := self.logger.WithFields(log.Fields{
		"topic": self.consumer_topic(),
		"side":  self.opt.Side,
		"event": "recv",
	})

	self.receiving = make(chan []byte)
	go func() {
		defer func() {
			close(self.receiving)
			logger.Debugf("loop closed")
		}()

		for {
			logger.Debugf("waiting")
			select {
			case msg, ok := <-self.consumer.Messages():
				if ok {
					self.consumer.MarkOffset(msg, "")
					self.receiving <- msg.Value
					logger.Debugf("recv msg")
				}
			case err := <-self.consumer.Errors():
				self.recv_err = err
				logger.WithError(err).Debugf("recv error")
				return
			case ntf := <-self.consumer.Notifications():
				logger.WithField("notification", ntf).Debugf("rebalanced")
			case prt := <-self.consumer.Partitions():
				logger.WithField("partitions", prt).Debugf("partitions")
			}
		}
	}()
	logger.Debugf("init async recv")
}

func (self *saramaChannel) AsyncSend() chan<- []byte {
	panic("unimplemented")
}

func (self *saramaChannel) AsyncRecv() <-chan []byte {
	self.async_recv_once.Do(self.init_async_recv)
	return self.receiving
}

func (self *saramaChannel) Send(buf []byte) error {
	self.sync_send_once.Do(self.init_producer)

	topic := self.producer_topic()
	logger := self.logger.WithFields(log.Fields{
		"topic": topic,
		"side":  self.opt.Side,
		"event": "send",
	})

	msg := &sarama.ProducerMessage{
		Topic:     self.producer_topic(),
		Value:     sarama.ByteEncoder(buf),
		Partition: -1,
	}

	_, _, err := self.producer.SendMessage(msg)
	if err != nil {
		logger.WithError(err).Debugf("failed to send msg")
		return err
	}
	logger.Debugf("send msg")

	return nil
}

func (self *saramaChannel) Recv() ([]byte, error) {
	if self.recv_err != nil {
		return nil, self.recv_err
	}

	select {
	case buf := <-self.AsyncRecv():
		if self.recv_err != nil {
			return nil, self.recv_err
		}
		return buf, nil
	}
}

func (self *saramaChannel) Close() error {
	var err error

	if self.producer != nil {
		if err = self.producer.Close(); err != nil {
			self.logger.WithError(err).Warningf("failed to close producer")
			return err
		}
	}

	if self.consumer != nil {
		if err = self.consumer.Close(); err != nil {
			self.logger.WithError(err).Warningf("failed to close consumer")
			return err
		}
	}

	self.logger.Debugf("channel closed")
	return nil
}

type saramaBridgeOption struct {
	Id      string
	Brokers []string
}

type saramaBridge struct {
	opt    *saramaBridgeOption
	logger log.FieldLogger

	client sarama.Client

	north Channel
	south Channel

	client_once *sync.Once
	north_once  *sync.Once
	south_once  *sync.Once
}

func (self *saramaBridge) Id() string {
	return self.opt.Id
}

func (self *saramaBridge) North() Channel {
	self.north_once.Do(self.init_north)
	return self.north
}

func (self *saramaBridge) South() Channel {
	self.south_once.Do(self.init_south)
	return self.south
}

func (self *saramaBridge) Close() error {
	var err error

	if err = self.North().Close(); err != nil {
		return err
	}

	if err = self.South().Close(); err != nil {
		return err
	}

	return nil
}

func (self *saramaBridge) init_client() {
	var err error

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	if self.client, err = sarama.NewClient(self.opt.Brokers, cfg); err != nil {
		self.logger.WithError(err).Fatalf("failed to create sarama client")
	}
}

func (self *saramaBridge) init_north() {
	self.client_once.Do(self.init_client)

	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    NORTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.north = newSaramaChannel(opt, self.client, self.logger)
}

func (self *saramaBridge) init_south() {
	self.client_once.Do(self.init_client)

	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    SOUTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.south = newSaramaChannel(opt, self.client, self.logger)
}

type saramaBridgeFactoryOption struct {
	Brokers []string
}

type saramaBridgeFactory struct {
	opt    *saramaBridgeFactoryOption
	logger log.FieldLogger
}

func (self *saramaBridgeFactory) BuildBridge(device_id string, session int32) (Bridge, error) {
	return self.GetBridge(parse_bridge_id(device_id, session))
}

func (self *saramaBridgeFactory) GetBridge(id string) (Bridge, error) {
	opt := &saramaBridgeOption{
		Id:      id,
		Brokers: self.opt.Brokers,
	}

	br := &saramaBridge{
		opt:         opt,
		logger:      self.logger.WithField("bridge", id),
		client_once: new(sync.Once),
		north_once:  new(sync.Once),
		south_once:  new(sync.Once),
	}

	return br, nil
}

func new_sarama_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	var ok bool
	var logger log.FieldLogger
	var err error

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	opt := &saramaBridgeFactoryOption{}

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": func(key string, val interface{}) error {
			logger, ok = val.(log.FieldLogger)
			if !ok {
				return ErrInvalidArgument
			}
			logger = logger.WithField("#bridge_driver", "sarama")
			return nil
		},
		"brokers": func(key string, val interface{}) error {
			var vals []interface{}
			var ok bool

			vals, ok = val.([]interface{})
			if !ok {
				return ErrInvalidArgument
			}

			var broker string
			var brokers []string
			for _, v := range vals {
				broker, ok = v.(string)
				if !ok {
					return ErrInvalidArgument
				}
				brokers = append(brokers, broker)
			}
			opt.Brokers = brokers

			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return &saramaBridgeFactory{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_bridge_factory_factory("sarama", new_sarama_bridge_factory)
}
