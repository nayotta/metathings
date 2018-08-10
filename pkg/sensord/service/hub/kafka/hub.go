package kafka_hub

import (
	"fmt"
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
)

type kafkaHub struct {
	opt    opt_helper.Option
	logger log.FieldLogger
	glock  *sync.Mutex
}

func symbol(opt opt_helper.Option) string {
	sensor_id := opt.GetString("sensor_id")
	if sensor_id == "" {
		sensor_id = "*"
	}

	core_id := opt.GetString("core_id")
	if core_id == "" {
		core_id = "*"
	}

	entity_name := opt.GetString("entity_name")
	if entity_name == "" {
		entity_name = "*"
	}

	owner_id := opt.GetString("owner_id")
	if owner_id == "" {
		owner_id = "*"
	}

	return fmt.Sprintf("sensor.%v.core.%v.entity.%v.user.%v", sensor_id, core_id, entity_name, owner_id)
}

func (self *kafkaHub) Subscriber(opt opt_helper.Option) (hub.Subscriber, error) {
	sub_id := id_helper.NewUint64Id()
	brokers := self.opt.GetStrings("brokers")
	group_id := self.opt.GetString("group_id")
	if group_id == "" {
		group_id = fmt.Sprintf("group.sensord.%v", sub_id)
	}
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":               strings.Join(brokers, ","),
		"group.id":                        group_id,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "latest"},
	}
	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	sub := &kafkaSubscriber{
		logger:   self.logger,
		id:       sub_id,
		opt:      opt,
		consumer: consumer,
	}

	return sub, nil
}

func (self *kafkaHub) Publisher(opt opt_helper.Option) (hub.Publisher, error) {
	pub_id := id_helper.NewUint64Id()
	brokers := self.opt.GetStrings("brokers")
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(brokers, ","),
	}
	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	pub := &kafkaPublisher{
		logger:   self.logger,
		id:       pub_id,
		opt:      opt,
		producer: producer,
	}

	return pub, nil
}

type kafkaSubscriber struct {
	logger   log.FieldLogger
	id       uint64
	opt      opt_helper.Option
	consumer *kafka.Consumer
}

func (self *kafkaSubscriber) Subscribe() (*sensord_pb.SensorData, error) {
	for {
		ev := <-self.consumer.Events()
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			self.consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			self.consumer.Unassign()
		case kafka.Error:
			self.logger.WithError(e).Errorf("failed to subscribe from kafka")
			return nil, hub.ErrUnsubscribable
		case *kafka.Message:
			var data sensord_pb.SensorData
			err := proto.Unmarshal(e.Value, &data)
			if err != nil {
				return nil, err
			}
			return &data, nil
		}
	}
}

func (self *kafkaSubscriber) Id() uint64 {
	return self.id
}

func (self *kafkaSubscriber) Symbol() string {
	return symbol(self.opt)
}

func (self *kafkaSubscriber) Close() error {
	err := self.consumer.Close()
	if err != nil {
		return err
	}

	return nil
}

type kafkaPublisher struct {
	logger   log.FieldLogger
	id       uint64
	opt      opt_helper.Option
	producer *kafka.Producer
}

func (self *kafkaPublisher) Publish(dat *sensord_pb.SensorData) error {
	topic := self.Symbol()
	val, err := proto.Marshal(dat)
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: val,
	}
	self.producer.ProduceChannel() <- msg

	return nil
}

func (self *kafkaPublisher) Id() uint64 {
	return self.id
}

func (self *kafkaPublisher) Symbol() string {
	return symbol(self.opt)
}

func (self *kafkaPublisher) Close() error {
	self.producer.Close()

	return nil
}

func NewHub(opt opt_helper.Option) (hub.Hub, error) {
	return &kafkaHub{
		opt:    opt,
		glock:  new(sync.Mutex),
		logger: opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{"#module": "hub", "#driver": "kafka"}),
	}, nil
}

func init() {
	hub.XXX_RegisterHub("kafka", NewHub)
}
