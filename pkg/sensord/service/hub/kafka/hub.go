package kafka_hub

import (
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
)

type kafkaHub struct {
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
	panic("unimplemented")
}

func (self *kafkaHub) Publisher(opt opt_helper.Option) (hub.Publisher, error) {
	panic("unimplemented")
}

func (self *kafkaHub) Close(sp hub.SubPub) error {
	panic("unimplemented")
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

func NewHub(opt opt_helper.Option) (hub.Hub, error) {
	return &kafkaHub{
		glock:  new(sync.Mutex),
		logger: opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{"#module": "hub", "#driver": "kafka"}),
	}, nil
}

func init() {
	hub.XXX_RegisterHub("kafka", NewHub)
}
