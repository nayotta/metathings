package metathings_deviced_connection

import (
	"encoding/base64"
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrChannelClosed   = errors.New("channel closed")
	// ErrUnexpectedResponse       = errors.New("unexpected response")
	ErrBridgeClosed             = errors.New("bridge closed")
	ErrReceiveTimeout           = errors.New("receive timeout")
	ErrDuplicatedDeviceInstance = errors.New("duplicated device instance")
	ErrDeviceOffline            = errors.New("device offline")
)

func ErrUnexpectedResponse(res []byte) error {
	return errors.New(fmt.Sprintf("unexpected response: %v", base64.StdEncoding.EncodeToString(res)))
}
