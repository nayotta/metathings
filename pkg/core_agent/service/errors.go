package metathings_core_agent_service

import "errors"

var (
	ErrUnsupportMessageType = errors.New("unsupport message type")
	ErrUnsupportPayloadType = errors.New("unsupport payload type")
	ErrPluginNotFound       = errors.New("plugin not found")
)
