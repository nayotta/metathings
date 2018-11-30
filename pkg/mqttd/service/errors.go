package metathingsmqttdservice

import "errors"

var (
	// ErrUnsupportMessageType ErrUnsupportMessageType
	ErrUnsupportMessageType = errors.New("unsupport message type")
	//ErrUnsupportPayloadType ErrUnsupportPayloadType
	ErrUnsupportPayloadType = errors.New("unsupport payload type")
	//ErrUnsupportRequestType ErrUnsupportRequestType
	ErrUnsupportRequestType = errors.New("unsupport request type")
	// ErrPluginNotFound ErrPluginNotFound
	ErrPluginNotFound = errors.New("plugin not found")
)
