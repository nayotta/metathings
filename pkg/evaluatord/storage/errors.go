package metathings_evaluatord_storage

import "errors"

var (
	ErrUnknownTaskStorageDriver = errors.New("unknown task storage driver")
	ErrTaskNotFound             = errors.New("task not found")
)
