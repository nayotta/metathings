package hub

import (
	"errors"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

var (
	ErrBadHubName     = errors.New("bad hub name")
	ErrUnsubscribable = errors.New("unsubscribable")
	ErrSubPubNotFound = errors.New("not found")
	ErrUnexpected     = errors.New("unexpected")
)

type SubPub interface {
	Id() uint64
	Path() string
}

type Hub interface {
	Subscriber(string) (Subscriber, error)
	Publisher(string) (Publisher, error)
	Close(SubPub) error
}

type Subscriber interface {
	SubPub
	Subscribe() (*sensord_pb.SensorData, error)
}

type Publisher interface {
	SubPub
	Publish(*sensord_pb.SensorData) error
}

var (
	hubs = make(map[string]func(opt_helper.Option) (Hub, error))
)

func XXX_RegisterHub(name string, fn func(opt_helper.Option) (Hub, error)) {
	hubs[name] = fn
}

func NewHub(opt opt_helper.Option) (Hub, error) {
	name := opt.GetString("name")
	if name == "" {
		return nil, ErrBadHubName
	}

	hub_maker, ok := hubs[name]
	if !ok {
		return nil, ErrBadHubName
	}

	return hub_maker(opt)
}
