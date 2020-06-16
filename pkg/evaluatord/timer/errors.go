package metathings_evaluatord_timer

import "errors"

var (
	ErrUnknownTimerBackendDriver = errors.New("unknown timer backend driver")
	ErrTimerIdIsEmpty            = errors.New("timer id is empty")
)
