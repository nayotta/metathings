package hub

import (
	"errors"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
)

var (
	ErrBadHubName     = errors.New("bad hub name")
	ErrUnsubscribable = errors.New("unsubscribable")
)

type Hub interface {
	Subscriber(string) Subscriber
	Publisher(string) Publisher
}

type Subscriber interface {
	Id() uint64
	Subscribe() (*sensor_pb.SensorData, error)
}

type Publisher interface {
	Id() uint64
	Publish(*sensor_pb.SensorData) error
}

var (
	hubs map[string]func(opt_helper.Option) (Hub, error)
)

func XXX_RegisterHub(name string, fn func(opt_helper.Option) (Hub, error)) {
	hubs[name] = fn
}

func NewHub(opt opt_helper.Option) (Hub, error) {
	name := opt.GetString("name")
	if name == "" {
		return nil, ErrBadHubName
	}
	return nil, errors.New("unimplemented")
}
