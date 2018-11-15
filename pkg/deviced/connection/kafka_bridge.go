package metathings_deviced_connection

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
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

	producer, err := kafka.NewProducer(&cfg)
	if err != nil {
		return err
	}

	self.producer = producer

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

	consumer, err := kafka.NewConsumer(&cfg)
	if err != nil {
		return err
	}

	self.consumer = consumer

	return nil
}

func (self *kafkaBridge) Id() string {
	return self.id
}

func (self *kafkaBridge) Send(buf []byte) error {
	self.init_producer()

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &self.symbol,
			Partition: kafka.PartitionAny,
		},
		Value: buf,
	}

	self.producer.ProduceChannel() <- msg

	return nil
}

func (self *kafkaBridge) Recv() ([]byte, error) {
	self.init_consumer()

	for {
		ev := <-self.consumer.Events()
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			self.consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			self.consumer.Unassign()
		case kafka.Error:
			return nil, ErrUnexpectedResponse
		case kafka.PartitionEOF:
		case kafka.OffsetsCommitted:
		case *kafka.Message:
			return e.Value, nil
		default:
			return nil, ErrUnexpectedResponse
		}
	}
}

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
	buf := fmt.Sprintf("device.%v.session.%v", device_id, session)
	hash := md5.New()
	id := hex.EncodeToString(hash.Sum([]byte(buf)))

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
	var key string
	var ok bool
	var logger log.FieldLogger

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	opt := &kafkaBridgeFactoryOption{
		ProducerConfig: map[string]string{
			"queue.buffering.max.ms": "100",
		},
		ConsumerConfig: map[string]string{
			"topic.metadata.refresh.interval.ms": "3000",
		},
	}

	for i := 0; i < len(args); i += 2 {
		if key, ok = args[i].(string); !ok {
			return nil, ErrInvalidArgument
		}

		switch key {
		case "logger":
			if logger, ok = args[i+1].(log.FieldLogger); !ok {
				return nil, ErrInvalidArgument
			}
		}
	}

	return &kafkaBridgeFactory{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_bridge_factory_factory("kafka", new_kafka_bridge_factory)
}
