package metathings_deviced_connection

import (
	"fmt"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type saramaChannelOption struct {
	Id      string
	Side    Side
	Brokers []string
}

func newSaramaChannel(opt *saramaChannelOption, logger log.FieldLogger) Channel {
	return &saramaChannel{
		opt:    opt,
		logger: logger,
	}
}

type saramaChannel struct {
	opt    *saramaChannelOption
	logger log.FieldLogger

	producer sarama.AsyncProducer
	consumer *cluster.Consumer
}

func (self *saramaChannel) producer_topic() string {
	return fmt.Sprintf("metathings.deviced.channel.%v.side.%v", self.opt.Id, self.opt.Side)
}

func (self *saramaChannel) consumer_topic() string {
	var s Side
	if self.opt.Side == NORTH_SIDE {
		s = SOUTH_SIDE
	} else {
		s = NORTH_SIDE
	}

	return fmt.Sprintf("metathings.deviced.channel.%v.side.%v", self.opt.Id, s)
}

func (self *saramaChannel) init_producer() error {
	var err error

	if self.producer != nil {
		return nil
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	if self.producer, err = sarama.NewAsyncProducer(self.opt.Brokers, config); err != nil {
		return err
	}
	self.logger.WithField("topic", self.producer_topic).Debugf("init producer")

	return nil
}

func (self *saramaChannel) init_consumer() error {
	var err error

	if self.consumer != nil {
		return nil
	}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false

	if self.consumer, err = cluster.NewConsumer(self.opt.Brokers, self.opt.Id, []string{self.consumer_topic()}, config); err != nil {
		return err
	}
	self.logger.WithFields(log.Fields{
		"group": self.opt.Id,
		"topic": self.consumer_topic,
	}).Debugf("init consumer")

	return nil
}

func (self *saramaChannel) AsyncSend() chan<- []byte {
	self.init_producer()

	ch := make(chan []byte)
	go func() {
		for {
			select {
			case buf := <-ch:
				msg := &sarama.ProducerMessage{
					Topic:     self.producer_topic(),
					Value:     sarama.ByteEncoder(buf),
					Partition: -1,
				}
				self.producer.Input() <- msg
				self.logger.Debugf("send msg")
			}
		}
	}()

	return ch
}

func (self *saramaChannel) AsyncRecv() <-chan []byte {
	self.init_consumer()

	ch := make(chan []byte)
	go func() {
		for {
			select {
			case msg, ok := <-self.consumer.Messages():
				if ok {
					self.consumer.MarkOffset(msg, "")
					self.logger.Debugf("recv msg")
					ch <- msg.Value
				}
			}
		}

	}()

	return ch
}

func (self *saramaChannel) Send(buf []byte) error {
	select {
	case self.AsyncSend() <- buf:
		return nil
	case err := <-self.producer.Errors():
		return err
	}
}

func (self *saramaChannel) Recv() ([]byte, error) {
	select {
	case buf := <-self.AsyncRecv():
		return buf, nil
	case err := <-self.consumer.Errors():
		return nil, err
	}
}

func (self *saramaChannel) Close() error {
	var err error

	if err = self.producer.Close(); err != nil {
		self.logger.WithError(err).Warningf("failed to close producer")
		return err
	}

	if err = self.consumer.Close(); err != nil {
		self.logger.WithError(err).Warningf("failed to close consumer")
		return err
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

	north Channel
	south Channel
}

func (self *saramaBridge) Id() string {
	return self.opt.Id
}

func (self *saramaBridge) North() Channel {
	self.init_north()
	return self.north
}

func (self *saramaBridge) South() Channel {
	self.init_south()
	return self.south
}

func (self *saramaBridge) Close() error {
	if err := self.North().Close(); err != nil {
		return err
	}

	if err := self.South().Close(); err != nil {
		return err
	}

	return nil
}

func (self *saramaBridge) init_north() {
	if self.north != nil {
		return
	}

	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    NORTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.north = newSaramaChannel(opt, self.logger)
}

func (self *saramaBridge) init_south() {
	if self.south != nil {
		return
	}

	opt := &saramaChannelOption{
		Id:      self.Id(),
		Side:    SOUTH_SIDE,
		Brokers: self.opt.Brokers,
	}

	self.south = newSaramaChannel(opt, self.logger)
}

type saramaBridgeFactoryOption struct {
	Brokers []string
}

type saramaBridgeFactory struct {
	opt    *saramaBridgeFactoryOption
	logger log.FieldLogger
}

func (self *saramaBridgeFactory) BuildBridge(device_id string, session int32) (Bridge, error) {
	id := id_helper.NewNamedId(fmt.Sprintf("device.%v.session.%v", device_id, session))
	return self.GetBridge(id)
}

func (self *saramaBridgeFactory) GetBridge(id string) (Bridge, error) {
	opt := &saramaBridgeOption{
		Id:      id,
		Brokers: self.opt.Brokers,
	}

	br := &saramaBridge{
		opt:    opt,
		logger: self.logger.WithField("bridge", id),
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
