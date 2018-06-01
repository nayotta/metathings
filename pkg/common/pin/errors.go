package pin_helper

import "errors"

var (
	ErrUnknownModel = errors.New("unknown model")
	ErrUnknownPin   = errors.New("unknown pin")
)
