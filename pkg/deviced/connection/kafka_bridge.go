package metathings_deviced_connection

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type kafkaBridgeOption struct {
	ProducerConfig map[string]string
	ConsumerConfig map[string]string
}

type kafkaBridge struct {
	opt    *kafkaBridgeOption
	logger log.FieldLogger

	consumer *kafka.Consumer
	producer *kafka.Producer

	id     string
	symbol string
}

func (self *kafkaBridge) init_producer() error {
	if self.producer != nil {
		return nil
	}

	cfg := kafka.ConfigMap{}
	for key, val := range self.opt.ProducerConfig {
		cfg.SetKey(key, val)
	}

	cfg["queue.buffering.max.ms"] = 30

	producer, err := kafka.NewProducer(&cfg)
	if err != nil {
		return err
	}

	self.producer = producer
	self.logger.WithField("config", cfg).Debugf("init producer")

	return nil
}

func (self *kafkaBridge) init_consumer() error {
	if self.consumer != nil {
		return nil
	}

	cfg := kafka.ConfigMap{}
	for key, val := range self.opt.ConsumerConfig {
		cfg.SetKey(key, val)
	}
	cfg["group.id"] = self.Id()
	// cfg["topic.metadata.refresh.interval.ms"] = 100
	cfg["session.timeout.ms"] = 6000
	cfg["socket.blocking.max.ms"] = 300
	cfg["go.events.channel.enable"] = true
	cfg["go.application.rebalance.enable"] = true
	cfg["default.topic.config"] = kafka.ConfigMap{"auto.offset.reset": "latest"}

	consumer, err := kafka.NewConsumer(&cfg)
	if err != nil {
		return err
	}

	err = consumer.SubscribeTopics([]string{self.symbol}, nil)
	if err != nil {
		return err
	}

	self.consumer = consumer
	self.logger.WithFields(log.Fields{
		"topic":  self.symbol,
		"group":  self.Id(),
		"config": cfg,
	}).Debugf("init consumer")

	return nil
}

func (self *kafkaBridge) Id() string {
	return self.id
}

func (self *kafkaBridge) Send(buf []byte) error {
	var err error

	if err = self.init_producer(); err != nil {
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &self.symbol,
			Partition: kafka.PartitionAny,
		},
		Value: buf,
	}

	logger := self.logger.WithField("#event", "send")

	self.producer.ProduceChannel() <- msg
	logger.WithField("topic", self.symbol).Debugf("send msg")

	return nil
}

func (self *kafkaBridge) Recv() ([]byte, error) {
	var err error

	if err = self.init_consumer(); err != nil {
		return nil, err
	}

	logger := self.logger.WithField("#event", "recv")

	for {
		ev := <-self.consumer.Events()
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			logger.Debugf("assigned partitions")
			self.consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			logger.Debugf("revoked partitions")
			self.consumer.Unassign()
		case kafka.Error:
			logger.WithError(e).Debugf("kafka error")
			return nil, ErrUnexpectedResponse
		case kafka.PartitionEOF:
			logger.Debugf("partition eof")
		case kafka.OffsetsCommitted:
			logger.Debugf("offsets committed")
		case *kafka.Message:
			logger.WithField("topic", self.symbol).Debugf("recv msg")
			return e.Value, nil
		default:
			logger.Debugf("unexpected kafka event")
			return nil, ErrUnexpectedResponse
		}
	}
}

func (self *kafkaBridge) Close() error { return nil }

type kafkaBridgeFactoryOption struct {
	ProducerConfig map[string]string
	ConsumerConfig map[string]string
}

type kafkaBridgeFactory struct {
	opt    *kafkaBridgeFactoryOption
	logger log.FieldLogger
}

func bridge_id_to_symbol(id string) string {
	return fmt.Sprintf("metathings.deviced.connection.bridge.%v", id)
}

func (self *kafkaBridgeFactory) BuildBridge(device_id string, session int32) (Bridge, error) {
	id := id_helper.NewNamedId(fmt.Sprintf("device.%v.session.%v", device_id, session))
	return self.GetBridge(id)
}

func (self *kafkaBridgeFactory) GetBridge(id string) (Bridge, error) {
	opt := &kafkaBridgeOption{
		ProducerConfig: self.opt.ProducerConfig,
		ConsumerConfig: self.opt.ConsumerConfig,
	}

	br := &kafkaBridge{
		opt:    opt,
		logger: self.logger.WithField("bridge", id),

		id:     id,
		symbol: bridge_id_to_symbol(id),
	}

	return br, nil
}

func new_kafka_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	var ok bool
	var logger log.FieldLogger
	var err error

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	opt := &kafkaBridgeFactoryOption{
		ProducerConfig: map[string]string{},
		ConsumerConfig: map[string]string{},
	}

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": func(key string, val interface{}) error {
			logger, ok = val.(log.FieldLogger)
			if !ok {
				return ErrInvalidArgument
			}
			return nil
		},
		"brokers": func(key string, val interface{}) error {
			var vals []interface{}
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

			servers := strings.Join(brokers, ",")
			opt.ProducerConfig["bootstrap.servers"] = servers
			opt.ConsumerConfig["bootstrap.servers"] = servers

			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return &kafkaBridgeFactory{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_bridge_factory_factory("kafka", new_kafka_bridge_factory)
}
