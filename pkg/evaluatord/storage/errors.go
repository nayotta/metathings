package metathings_evaluatord_storage

import "errors"

var (
	ErrUnknownTaskStorageDriver  = errors.New("unknown task storage driver")
	ErrUnknownTimerStorageDriver = errors.New("unknown timer storage driver")
	ErrTaskNotFound              = errors.New("task not found")
)
