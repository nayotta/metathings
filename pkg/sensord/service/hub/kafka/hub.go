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
	return fmt.Sprintf("sensor.%v.core.%v.entity.%v.user.%v", opt.GetString("sensor_id"), opt.GetString("core_id"), opt.GetString("entity_name"), opt.GetString("owner_id"))
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

type kafkaSubscriber struct{}

func (self *kafkaSubscriber) Subscribe() (*sensord_pb.SensorData, error) {
	panic("unimplemented")
}

func (self *kafkaSubscriber) Id() uint64 {
	panic("unimplemented")
}

func (self *kafkaSubscriber) Symbol() string {
	panic("unimplemented")
}

type kafkaPublisher struct {
	id       uint64
	sym      string
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
	panic("unimplemented")
}

func (self *kafkaPublisher) Symbol() string {
	panic("unimplemented")
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
