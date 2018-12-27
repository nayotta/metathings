package metathings_deviced_connection

import (
	"fmt"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type saramaBridgeOption struct {
	Brokers []string
}

type saramaBridge struct {
	opt    *saramaBridgeOption
	logger log.FieldLogger

	id     string
	symbol string

	producer sarama.SyncProducer
	consumer *cluster.Consumer

	recv_err error
	send_err error
}

func (self *saramaBridge) init_producer() error {
	var err error

	if self.producer != nil {
		return nil
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	if self.producer, err = sarama.NewSyncProducer(self.opt.Brokers, config); err != nil {
		return err
	}
	self.logger.WithField("topic", self.symbol).Debugf("init producer")

	return nil
}

func (self *saramaBridge) init_consumer() error {
	var err error

	if self.consumer != nil {
		return nil
	}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false

	if self.consumer, err = cluster.NewConsumer(self.opt.Brokers, self.Id(), []string{self.symbol}, config); err != nil {
		return err
	}
	self.logger.WithFields(log.Fields{
		"group": self.Id(),
		"topic": self.symbol,
	}).Debugf("init consumer")

	return nil
}

func (self *saramaBridge) Id() string { return self.id }

func (self *saramaBridge) Send(buf []byte) error {
	if self.send_err != nil {
		return self.send_err
	}

	logger := self.logger.WithField("#event", "send")

	err := self.init_producer()
	if err != nil {
		logger.WithError(err).Debugf("failed to init producer")
		self.send_err = err
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     self.symbol,
		Value:     sarama.ByteEncoder(buf),
		Partition: -1,
	}

	partition, offset, err := self.producer.SendMessage(msg)
	if err != nil {
		self.send_err = err
		logger.WithError(err).Debugf("failed to send message")
		return err
	}
	logger.WithFields(log.Fields{
		"topic":     self.symbol,
		"partition": partition,
		"offset":    offset,
	}).Debugf("send msg")

	return nil
}

func (self *saramaBridge) Recv() ([]byte, error) {
	var msg *sarama.ConsumerMessage
	var err error
	var ok bool

	if self.recv_err != nil {
		return nil, err
	}

	logger := self.logger.WithField("#event", "recv")

	if err = self.init_consumer(); err != nil {
		logger.WithError(err).Debugf("failed to init consumer")
		self.recv_err = err
		return nil, err
	}

	select {
	case err = <-self.consumer.Errors():
		self.recv_err = err
		logger.WithError(err).Debugf("failed to recv msg")
		return nil, err
	case msg, ok = <-self.consumer.Messages():
		if ok {
			self.consumer.MarkOffset(msg, "")
			logger.WithFields(log.Fields{
				"#event": "recv",
				"topic":  self.symbol,
			}).Debugf("recv msg")
			return msg.Value, nil
		}
	}

	return nil, ErrUnexpectedResponse
}

func (self *saramaBridge) Close() error {
	if self.producer != nil {
		self.producer.Close()
		self.producer = nil
		self.send_err = ErrBridgeClosed
	}

	if self.consumer != nil {
		self.consumer.Close()
		self.consumer = nil
		self.recv_err = ErrBridgeClosed
	}

	return nil
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
		Brokers: self.opt.Brokers,
	}

	br := &saramaBridge{
		opt:    opt,
		logger: self.logger.WithField("bridge", id),

		id:     id,
		symbol: bridge_id_to_symbol(id),
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
