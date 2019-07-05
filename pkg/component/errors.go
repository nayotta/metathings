package metathings_component

import "errors"

var (
	ErrUnaryCallTimeout           = errors.New("unary call timeout")
	ErrStreamCallConfigAckTimeout = errors.New("stream call config ack timeout")
	ErrUnexceptedResponse         = errors.New("unexpected response")
	ErrBadScheme                  = errors.New("bad scheme")
	ErrBadServiceEndpoint         = errors.New("bad service endpoint")
	ErrDefaultAddressRequired     = errors.New("default address required")
	ErrDeviceAddressRequired      = errors.New("device address required")
	ErrInvalidArguments           = errors.New("invalid arguments")
	ErrSubscribedSession          = errors.New("subscribed session")
	ErrUnsubscribedSession        = errors.New("unsubscribed session")
	ErrStartTimeout               = errors.New("start timeout")
)
