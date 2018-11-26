package metathingsmqttdservice

import "errors"

var (
	// ErrUnsupportMessageType ErrUnsupportMessageType
	ErrUnsupportMessageType = errors.New("unsupport message type")
	//ErrUnsupportPayloadType ErrUnsupportPayloadType
	ErrUnsupportPayloadType = errors.New("unsupport payload type")
	// ErrPluginNotFound ErrPluginNotFound
	ErrPluginNotFound = errors.New("plugin not found")
)
