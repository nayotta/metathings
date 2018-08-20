package pubsub

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

type PubSubManager interface {
	GetPublisherManager(id uint64) (PublisherManager, error)
	ListPublisherManagers() (map[uint64]PublisherManager, error)
	GetSubscriberManager(id uint64) (SubscriberManager, error)
	ListSubscriberManagers() (map[uint64]SubscriberManager, error)
}

type PublisherManager interface {
	Id() uint64
	NewPublisher(opt_helper.Option) (Publisher, error)
	GetPublisher(opt_helper.Option) (Publisher, error)
	Close() error
	Closed() bool
}

type SubscriberManager interface {
	Id() uint64
	NewSubscriber(opt_helper.Option) (Subscriber, error)
	GetSubscriber(opt_helper.Option) (Subscriber, error)
	Close() error
	Closed() bool
}

type PubSub interface {
	Symbol() string
	Close() error
}

type Subscriber interface {
	PubSub
	Subscribe() (*sensord_pb.SensorData, error)
}

type Publisher interface {
	PubSub
	Publish(*sensord_pb.SensorData) error
}

var (
	mgrs = make(map[string]func(opt_helper.Option) (PubSubManager, error))
)

func XXX_RegisterManager(name string, fn func(opt_helper.Option) (PubSubManager, error)) {
	mgrs[name] = fn
}

func NewManager(opt opt_helper.Option) (PubSubManager, error) {
	name := opt.GetString("name")
	if mgr_fty, ok := mgrs[name]; !ok {
		return nil, ErrUnregisterManagerName
	} else {
		return mgr_fty(opt)
	}
}
