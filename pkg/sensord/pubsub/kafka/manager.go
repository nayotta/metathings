package kafka_manager

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
	"github.com/nayotta/metathings/pkg/sensord/pubsub"
)

const (
	UUID_REGEX = "[0-9a-z]{32}"
	NAME_REGEX = "[a-zA-Z0-9]+"
)

type option struct {
	Brokers        []string
	ProducerConfig map[string]string `json:"producer_config"`
	ConsumerConfig map[string]string `json:"consumer_config"`
}

func default_option() *option {
	opt := &option{
		ProducerConfig: map[string]string{
			"queue.buffering.max.ms": "100",
		},
		ConsumerConfig: map[string]string{
			// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
			"topic.metadata.refresh.interval.ms": "3000",
		},
	}
	return opt
}

type kafkaPubSubManager struct {
	opt    *option
	logger log.FieldLogger
	glock  *sync.Mutex
	wg     sync.WaitGroup

	pub_mgrs map[uint64]*kafkaPublisherManager
	sub_mgrs map[uint64]*kafkaSubscriberManager
}

func (self *kafkaPubSubManager) newPublisherManager(id uint64) (*kafkaPublisherManager, error) {
	pub := &kafkaPublisherManager{
		id:     id,
		opt:    self.opt,
		logger: self.logger.WithField("pub_mgr_id", id),
		mlock:  &sync.Mutex{},
		closed: false,
		pubs:   make(map[string]*kafkaPublisher),
	}
	pub.close_callback = func() error {
		self.glock.Lock()
		defer self.glock.Unlock()

		id := pub.Id()
		if _, ok := self.pub_mgrs[id]; ok {
			delete(self.pub_mgrs, id)
		}

		return nil
	}
	self.pub_mgrs[id] = pub

	self.logger.WithField("id", id).Debugf("new publisher manager")
	return pub, nil
}

func (self *kafkaPubSubManager) newSubscriberManager(id uint64) (*kafkaSubscriberManager, error) {
	sub := &kafkaSubscriberManager{
		id:     id,
		opt:    self.opt,
		logger: self.logger.WithField("sub_mgr_id", id),
		mlock:  &sync.Mutex{},
		closed: false,
		subs:   make(map[string]*kafkaSubscriber),
	}
	sub.close_callback = func() error {
		self.glock.Lock()
		defer self.glock.Unlock()

		id := sub.Id()
		if _, ok := self.sub_mgrs[id]; ok {
			delete(self.sub_mgrs, id)
		}

		return nil
	}
	self.sub_mgrs[id] = sub

	self.logger.WithField("id", id).Debugf("new subscriber manager")
	return sub, nil
}

func (self *kafkaPubSubManager) GetPublisherManager(id uint64) (pubsub.PublisherManager, error) {
	var mgr *kafkaPublisherManager
	var ok bool
	var err error
	if mgr, ok = self.pub_mgrs[id]; !ok {
		mgr, err = self.newPublisherManager(id)
		if err != nil {
			return nil, err
		}
	} else {
		if mgr.Closed() {
			mgr, err = self.newPublisherManager(id)
			if err != nil {
				return nil, err
			}
		}
	}
	return mgr, nil
}

func (self *kafkaPubSubManager) ListPublisherManagers() (map[uint64]pubsub.PublisherManager, error) {
	mgrs := make(map[uint64]pubsub.PublisherManager)

	for id, mgr := range self.pub_mgrs {
		if mgr.Closed() {
			delete(self.pub_mgrs, id)
		} else {
			mgrs[id] = mgr
		}
	}

	return mgrs, nil
}

func (self *kafkaPubSubManager) GetSubscriberManager(id uint64) (pubsub.SubscriberManager, error) {
	var mgr *kafkaSubscriberManager
	var ok bool
	var err error
	if mgr, ok = self.sub_mgrs[id]; !ok {
		mgr, err = self.newSubscriberManager(id)
		if err != nil {
			return nil, err
		}
	} else {
		if mgr.Closed() {
			mgr, err = self.newSubscriberManager(id)
			if err != nil {
				return nil, err
			}
		}
	}
	return mgr, nil
}

func (self *kafkaPubSubManager) ListSubscriberManagers() (map[uint64]pubsub.SubscriberManager, error) {
	mgrs := make(map[uint64]pubsub.SubscriberManager)

	for id, mgr := range self.sub_mgrs {
		if mgr.Closed() {
			delete(self.sub_mgrs, id)
		} else {
			mgrs[id] = mgr
		}
	}

	return mgrs, nil
}

type kafkaPublisherManager struct {
	id             uint64
	opt            *option
	logger         log.FieldLogger
	mlock          *sync.Mutex
	close_callback func() error
	closed         bool
	wg             sync.WaitGroup

	pubs map[string]*kafkaPublisher
}

func (self *kafkaPublisherManager) Id() uint64 {
	return self.id
}

func (self *kafkaPublisherManager) NewPublisher(opt opt_helper.Option) (pubsub.Publisher, error) {
	self.mlock.Lock()
	defer self.mlock.Unlock()

	if _, ok := self.pubs[(&kafkaPublisher{opt: opt}).Symbol()]; ok {
		return nil, pubsub.ErrExistedPublisher
	}

	brokers := strings.Join(self.opt.Brokers, ",")
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
	}
	for key, val := range self.opt.ProducerConfig {
		cfg.SetKey(key, val)
	}
	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	pub := &kafkaPublisher{
		logger:   self.logger,
		opt:      opt,
		producer: producer,
		quit:     make(chan interface{}),
	}
	go pub.loop()
	pub.close_callback = func() error {
		self.mlock.Lock()
		defer self.mlock.Unlock()

		sym := pub.Symbol()
		if _, ok := self.pubs[sym]; ok {
			delete(self.pubs, sym)
		}
		self.wg.Done()

		return nil
	}
	self.wg.Add(1)

	self.pubs[pub.Symbol()] = pub

	self.logger.WithFields(log.Fields{
		"config": cfg,
		"symbol": pub.Symbol(),
	}).Debugf("create publisher")

	return pub, nil
}

func (self *kafkaPublisherManager) GetPublisher(opt opt_helper.Option) (pubsub.Publisher, error) {
	var pub *kafkaPublisher
	var ok bool
	sym := (&kafkaSubscriber{opt: opt}).Symbol()
	if pub, ok = self.pubs[sym]; !ok {
		return nil, pubsub.ErrNotFoundPublisher
	}
	return pub, nil
}

func (self *kafkaPublisherManager) Close() error {
	var err error

	self.wg.Wait()

	self.mlock.Lock()
	defer self.mlock.Unlock()

	if err = self.close_callback(); err != nil {
		return err
	}

	self.closed = true

	self.logger.WithField("id", self.id).Debugf("publisher manager closed")
	return nil
}

func (self *kafkaPublisherManager) Closed() bool {
	return self.closed
}

type kafkaSubscriberManager struct {
	id             uint64
	opt            *option
	logger         log.FieldLogger
	mlock          *sync.Mutex
	close_callback func() error
	closed         bool
	wg             sync.WaitGroup

	subs map[string]*kafkaSubscriber
}

func (self *kafkaSubscriberManager) Id() uint64 {
	return self.id
}

func (self *kafkaSubscriberManager) NewSubscriber(opt opt_helper.Option) (pubsub.Subscriber, error) {
	self.mlock.Lock()
	defer self.mlock.Unlock()

	if _, ok := self.subs[(&kafkaSubscriber{opt: opt}).Symbol()]; ok {
		return nil, pubsub.ErrExistedSubscriber
	}

	sub_id := id_helper.NewUint64Id()
	brokers := self.opt.Brokers
	group_id := fmt.Sprintf("group.sensord.%v", sub_id)
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":               strings.Join(brokers, ","),
		"group.id":                        group_id,
		"session.timeout.ms":              6000,
		"socket.blocking.max.ms":          100,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "latest"},
	}
	for key, val := range self.opt.ConsumerConfig {
		cfg.SetKey(key, val)
	}
	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	sub := &kafkaSubscriber{
		logger:   self.logger,
		opt:      opt,
		consumer: consumer,
	}
	sub.close_callback = func() error {
		self.mlock.Lock()
		defer self.mlock.Unlock()

		sym := sub.Symbol()
		if _, ok := self.subs[sym]; ok {
			delete(self.subs, sym)
		}
		self.wg.Done()

		return nil
	}
	self.wg.Add(1)

	err = consumer.SubscribeTopics([]string{sub.Symbol()}, nil)
	if err != nil {
		sub.Close()
		return nil, err
	}

	self.subs[sub.Symbol()] = sub

	self.logger.WithFields(log.Fields{
		"config": cfg,
		"symbol": sub.Symbol(),
	}).Debugf("create subscriber")

	return sub, nil
}

func (self *kafkaSubscriberManager) GetSubscriber(opt opt_helper.Option) (pubsub.Subscriber, error) {
	var sub *kafkaSubscriber
	var ok bool
	sym := (&kafkaSubscriber{opt: opt}).Symbol()
	if sub, ok = self.subs[sym]; !ok {
		return nil, pubsub.ErrNotFoundSubscriber
	}
	return sub, nil
}

func (self *kafkaSubscriberManager) Close() error {
	var err error

	self.wg.Wait()

	self.mlock.Lock()
	defer self.mlock.Unlock()

	if err = self.close_callback(); err != nil {
		return err
	}

	self.closed = true

	self.logger.WithField("id", self.id).Debugf("subscriber manager closed")
	return nil
}

func (self *kafkaSubscriberManager) Closed() bool {
	return self.closed
}

type kafkaSubscriber struct {
	logger         log.FieldLogger
	opt            opt_helper.Option
	consumer       *kafka.Consumer
	close_callback func() error

	_symbol string
}

func (self *kafkaSubscriber) Subscribe() (*sensord_pb.SensorData, error) {
	for {
		ev := <-self.consumer.Events()
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			self.consumer.Assign(e.Partitions)
			self.logger.WithField("symbol", self.Symbol()).Debugf("assign partitions to consumer")
		case kafka.RevokedPartitions:
			self.consumer.Unassign()
			self.logger.WithField("symbol", self.Symbol()).Debugf("unassign partitions form consumer")
		case kafka.Error:
			self.logger.WithField("symbol", self.Symbol()).WithError(e).Errorf("failed to subscribe from kafka")
			return nil, pubsub.ErrUnsubscribable
		case *kafka.Message:
			var data sensord_pb.SensorData
			err := proto.Unmarshal(e.Value, &data)
			if err != nil {
				return nil, err
			}
			self.logger.WithField("symbol", self.Symbol()).Debugf("subscribe data")
			return &data, nil
		case kafka.PartitionEOF:
			self.logger.WithField("symbol", self.Symbol()).Debugf("partition eof")
		case kafka.OffsetsCommitted:
			self.logger.WithField("symbol", self.Symbol()).Debugf("offsets committed")
		default:
			self.logger.Debugln(ev)
			return nil, pubsub.ErrUnsubscribable
		}
	}
}

func (self *kafkaSubscriber) Symbol() string {
	if self._symbol == "" {
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

		self._symbol = fmt.Sprintf("^sensor.%v.core.%v.entity.%v.user.%v$", sensor_id, core_id, entity_name, owner_id)
	}

	return self._symbol
}

func (self *kafkaSubscriber) Close() error {
	err := self.consumer.Close()
	if err != nil {
		return err
	}
	self.close_callback()

	self.logger.WithField("symbol", self.Symbol()).Debugf("subscriber closed")
	return nil
}

type kafkaPublisher struct {
	logger         log.FieldLogger
	opt            opt_helper.Option
	producer       *kafka.Producer
	quit           chan interface{}
	close_callback func() error
}

func (self *kafkaPublisher) loop() {
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

func (self *kafkaPublisher) Symbol() string {
	sensor_id := self.opt.GetString("sensor_id")
	core_id := self.opt.GetString("core_id")
	entity_name := self.opt.GetString("entity_name")
	owner_id := self.opt.GetString("owner_id")
	return fmt.Sprintf("sensor.%v.core.%v.entity.%v.user.%v", sensor_id, core_id, entity_name, owner_id)
}

func (self *kafkaPublisher) Close() error {
	self.quit <- nil
	defer close(self.quit)
	self.producer.Close()
	self.close_callback()

	self.logger.WithField("symbol", self.Symbol()).Debugf("publisher closed")
	return nil
}

func NewManager(opt opt_helper.Option) (pubsub.PubSubManager, error) {
	o := default_option()
	err := opt.Get("options").(*viper.Viper).Unmarshal(o)
	if err != nil {
		return nil, err
	}

	mgr := &kafkaPubSubManager{
		opt:      o,
		glock:    &sync.Mutex{},
		logger:   opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{"#module": "pubsubmanager", "#driver": "kafka"}),
		pub_mgrs: make(map[uint64]*kafkaPublisherManager),
		sub_mgrs: make(map[uint64]*kafkaSubscriberManager),
	}

	return mgr, nil
}

func init() {
	pubsub.XXX_RegisterManager("kafka", NewManager)
}
