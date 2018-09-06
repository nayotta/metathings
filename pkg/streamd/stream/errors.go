package stream_manager

import "errors"

var (
	ErrUnstartable  = errors.New("unstartable")
	ErrUnterminable = errors.New("unterminable")
)
