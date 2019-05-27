package metathings_deviced_connection

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
)

type saramaChannelOption struct {
	Id      string
	Side    Side
	Brokers []string
}

func newSaramaChannel(
	opt *saramaChannelOption,
	produce_client sarama.Client,
	consumer_client_pool pool_helper.Pool,
	logger log.FieldLogger,
) Channel {
	return &saramaChannel{
		opt:                  opt,
		logger:               logger,
		produce_client:       produce_client,
		consumer_client_pool: consumer_client_pool,
		async_recv_once:      new(sync.Once),
		sync_send_once:       new(sync.Once),
	}
}

type saramaChannel struct {
	opt    *saramaChannelOption
	logger log.FieldLogger

	produce_client       sarama.Client
	consumer_client      sarama.Client
	consumer_client_pool pool_helper.Pool
	producer             sarama.SyncProducer
	consumer             sarama.ConsumerGroup

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

	if self.producer, err = sarama.NewSyncProducerFromClient(self.produce_client); err != nil {
		panic(err)
	}
	self.logger.WithField("topic", self.producer_topic()).Debugf("init producer")
}

func (self *saramaChannel) init_consumer() {
	cli, err := self.consumer_client_pool.Get()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to init sarama consumer client")
		return
	}
	self.consumer_client = cli.(sarama.Client)

	self.consumer, err = sarama.NewConsumerGroupFromClient(self.opt.Id, self.consumer_client)
	if err != nil {
		self.logger.WithError(err).Warningf("failed to init sarama consumer group")
		return
	}

	self.logger.WithFields(log.Fields{
		"group": self.opt.Id,
		"topic": self.consumer_topic(),
	}).Debugf("init consumer")
}

type saramaChannelConsumerGroupHandler struct {
	receiving chan []byte
	logger    log.FieldLogger
}

func (h *saramaChannelConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	h.logger.Debugf("consumer setup")

	return nil
}

func (h *saramaChannelConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	h.logger.Debugf("consumer cleanup")

	return nil
}

func (h *saramaChannelConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		sess.MarkMessage(msg, "")
		h.receiving <- msg.Value
		h.logger.Debugf("recv msg")
	}

	return nil
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
			err := self.consumer.Consume(context.TODO(), []string{self.consumer_topic()}, &saramaChannelConsumerGroupHandler{
				receiving: self.receiving,
				logger:    logger,
			})
			if err != nil {
				logger.WithError(err).Debugf("recv error")
				return
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

		if err = self.consumer_client_pool.Put(self.consumer_client); err != nil {
			self.logger.WithError(err).Warningf("failed to put client back to pool")
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

	producer_client      sarama.Client
	consumer_client_pool pool_helper.Pool

	north Channel
	south Channel

	north_once *sync.Once
	south_once *sync.Once
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

func (self *saramaBridge) init_producer_client() {
	var err error

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	if self.producer_client, err = sarama.NewClient(self.opt.Brokers, cfg); err != nil {
		self.logger.WithError(err).Fatalf("failed to create sarama client")
	}
}

func (self *saramaBridge) init_client() {
	self.init_producer_client()
}

func (self *saramaBridge) init_north() {
	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    NORTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.north = newSaramaChannel(opt, self.producer_client, self.consumer_client_pool, self.logger)
}

func (self *saramaBridge) init_south() {
	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    SOUTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.south = newSaramaChannel(opt, self.producer_client, self.consumer_client_pool, self.logger)
}

type saramaBridgeFactoryOption struct {
	Brokers  []string
	Consumer struct {
		Initial int
		Max     int
	}
}

type saramaBridgeFactory struct {
	opt           *saramaBridgeFactoryOption
	prod_cli      sarama.Client
	cons_cli_pool pool_helper.Pool
	init_once     sync.Once

	logger log.FieldLogger
}

func (self *saramaBridgeFactory) init_producer_client() {
	var err error

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner

	if self.prod_cli, err = sarama.NewClient(self.opt.Brokers, cfg); err != nil {
		self.logger.WithError(err).Fatalf("failed to init sarama producer client")
	}
}

func (self *saramaBridgeFactory) init_consumer_client_pool() {
	var err error

	if self.cons_cli_pool, err = pool_helper.NewPool(self.opt.Consumer.Initial, self.opt.Consumer.Max, func() (pool_helper.Client, error) {
		cfg := sarama.NewConfig()
		cfg.Consumer.Group.Heartbeat.Interval = 100 * time.Millisecond
		cfg.Version = sarama.V2_0_0_0
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

		return sarama.NewClient(self.opt.Brokers, cfg)
	}); err != nil {
		self.logger.WithError(err).Fatalf("failed to init sarama consumer client pool")
	}
}

func (self *saramaBridgeFactory) init_client() {
	self.init_producer_client()
	self.init_consumer_client_pool()
}

func (self *saramaBridgeFactory) BuildBridge(device_id string, session int64) (Bridge, error) {
	self.init_once.Do(self.init_client)

	return self.GetBridge(parse_bridge_id(device_id, session))
}

func (self *saramaBridgeFactory) GetBridge(id string) (Bridge, error) {
	self.init_once.Do(self.init_client)

	opt := &saramaBridgeOption{
		Id:      id,
		Brokers: self.opt.Brokers,
	}

	br := &saramaBridge{
		opt:                  opt,
		logger:               self.logger.WithField("bridge", id),
		producer_client:      self.prod_cli,
		consumer_client_pool: self.cons_cli_pool,
		north_once:           new(sync.Once),
		south_once:           new(sync.Once),
	}

	return br, nil
}

func new_sarama_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	var ok bool
	var logger log.FieldLogger
	var err error

	opt := &saramaBridgeFactoryOption{}
	opt.Consumer.Initial = 5
	opt.Consumer.Max = 23

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

var register_sarama_bridge_factory_factory_once sync.Once

func init() {
	register_sarama_bridge_factory_factory_once.Do(func() {
		register_bridge_factory_factory("sarama", new_sarama_bridge_factory)
	})
}
