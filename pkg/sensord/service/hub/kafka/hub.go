package kafka_hub

import (
	"fmt"
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
)

const (
	UUID_REGEX = "[0-9a-z]{32}"
	NAME_REGEX = "[a-zA-Z0-9]+"
)

type option struct {
	Brokers []string
}

type kafkaHub struct {
	opt    option
	logger log.FieldLogger
	glock  *sync.Mutex
}

func (self *kafkaHub) Subscriber(opt opt_helper.Option) (hub.Subscriber, error) {
	sub_id := id_helper.NewUint64Id()
	brokers := self.opt.Brokers
	group_id := fmt.Sprintf("group.sensord.%v", sub_id)
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

	err = consumer.SubscribeTopics([]string{sub.Symbol()}, nil)
	if err != nil {
		return nil, err
	}

	self.logger.WithFields(log.Fields{
		"symbol": sub.Symbol(),
	}).Debugf("create subscriber")

	return sub, nil
}

func (self *kafkaHub) Publisher(opt opt_helper.Option) (hub.Publisher, error) {
	pub_id := id_helper.NewUint64Id()
	brokers := strings.Join(self.opt.Brokers, ",")
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
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
		quit:     make(chan interface{}),
	}

	go pub.loop()

	self.logger.WithFields(log.Fields{
		"symbol": pub.Symbol(),
	}).Debugf("create publisher")

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
			self.logger.WithField("symbol", self.Symbol()).Debugf("subscribe data")
			return &data, nil
		}
	}
}

func (self *kafkaSubscriber) Id() uint64 {
	return self.id
}

func (self *kafkaSubscriber) Symbol() string {
	sensor_id := self.opt.GetString("sensor_id")
	if sensor_id == "" {
		sensor_id = UUID_REGEX
	}

	core_id := self.opt.GetString("core_id")
	if core_id == "" {
		core_id = UUID_REGEX
	}

	entity_name := self.opt.GetString("entity_name")
	if entity_name == "" {
		entity_name = NAME_REGEX
	}

	owner_id := self.opt.GetString("owner_id")
	if owner_id == "" {
		owner_id = UUID_REGEX
	}

	return fmt.Sprintf("^sensor.%v.core.%v.entity.%v.user.%v$", sensor_id, core_id, entity_name, owner_id)
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
	quit     chan interface{}
}

func (self *kafkaPublisher) loop() {
	defer close(self.quit)
	for {
		select {
		case <-self.quit:
			return
		case e := <-self.producer.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					self.logger.WithError(ev.TopicPartition.Error).Debugf("failed to delivery message")
				}
			}

		}
	}
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
	self.logger.WithField("symbol", topic).Debugf("publish data")

	return nil
}

func (self *kafkaPublisher) Id() uint64 {
	return self.id
}

func (self *kafkaPublisher) Symbol() string {
	sensor_id := self.opt.GetString("sensor_id")
	core_id := self.opt.GetString("core_id")
	entity_name := self.opt.GetString("entity_name")
	owner_id := self.opt.GetString("owner_id")
	return fmt.Sprintf("sensor.%v.core.%v.entity.%v.user.%v", sensor_id, core_id, entity_name, owner_id)
}

func (self *kafkaPublisher) Close() error {
	self.quit <- nil
	self.producer.Close()

	return nil
}

func NewHub(opt opt_helper.Option) (hub.Hub, error) {
	var opts option
	err := opt.Get("options").(*viper.Viper).Unmarshal(&opts)
	if err != nil {
		return nil, err
	}

	hub := &kafkaHub{
		opt:    opts,
		glock:  new(sync.Mutex),
		logger: opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{"#module": "hub", "#driver": "kafka"}),
	}
	return hub, nil
}

func init() {
	hub.XXX_RegisterHub("kafka", NewHub)
}
