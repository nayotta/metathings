package metathings_deviced_connection

import (
	"errors"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

var (
	ErrUnknownBridgeDriver = errors.New("unknown bridge driver")
)

type Side string

const (
	NORTH_SIDE Side = "north"
	SOUTH_SIDE Side = "south"
)

type Channel interface {
	Send([]byte) error
	Recv() ([]byte, error)
	AsyncSend() chan<- []byte
	AsyncRecv() <-chan []byte
	Close() error
}

type Bridge interface {
	Id() string
	North() Channel
	South() Channel
	Close() error
}

type BridgeFactory interface {
	BuildBridge(device_id string, sess int64) (Bridge, error)
	GetBridge(br_id string) (Bridge, error)
}

var bridge_factory_factries map[string]func(...interface{}) (BridgeFactory, error)
var bridge_factory_factries_once sync.Once

func register_bridge_factory_factory(name string, fn func(args ...interface{}) (BridgeFactory, error)) {
	bridge_factory_factries_once.Do(func() {
		bridge_factory_factries = make(map[string]func(...interface{}) (BridgeFactory, error))
	})
	bridge_factory_factries[name] = fn
}

func NewBridgeFactory(name string, args ...interface{}) (BridgeFactory, error) {
	bri_fty_fty, ok := bridge_factory_factries[name]
	if !ok {
		return nil, ErrUnknownBridgeDriver
	}

	fty, err := bri_fty_fty(args...)
	if err != nil {
		return nil, err
	}

	return fty, nil
}

func ToBridgeFactory(y *BridgeFactory) func(string, interface{}) error {
	return func(k string, v interface{}) error {
		var ok bool
		*y, ok = v.(BridgeFactory)
		if !ok {
			return opt_helper.InvalidArgument(k)
		}

		return nil
	}
}

func ToHostnamer(h *Hostnamer) func(string, interface{}) error {
	return func(k string, v interface{}) error {
		var ok bool
		if *h, ok = v.(Hostnamer); !ok {
			return opt_helper.InvalidArgument(k)
		}

		return nil
	}
}
