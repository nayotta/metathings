package metathings_deviced_connection

import (
	"errors"
)

var (
	ErrUnknownBridgeDriver = errors.New("unknown bridge driver")
)

type Bridge interface {
	Id() string
	Send([]byte) error
	Recv() ([]byte, error)
}

type BridgeFactory interface {
	BuildBridge(device_id string, session int32) (Bridge, error)
	GetBridge(br_id string) (Bridge, error)
}

var bridge_factory_factries map[string]func(...interface{}) (BridgeFactory, error)

func register_bridge_factory_factory(name string, fn func(args ...interface{}) (BridgeFactory, error)) {
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

func init() {
	bridge_factory_factries = make(map[string]func(...interface{}) (BridgeFactory, error))
}