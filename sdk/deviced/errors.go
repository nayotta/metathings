package metathings_deviced_sdk

import "errors"

var (
	ErrUnsupportedSimpleStorageFactory = errors.New("unsupported simple storage driver")
	ErrUnsupportedCallerFactory        = errors.New("unsupported caller driver")
	ErrUnsupportedFlowFactory          = errors.New("unsupported flow driver")

	ErrMethodNotFound = errors.New("method not found")
	ErrModuleNotFound = errors.New("module not found")
	ErrConfigNotFound = errors.New("config not found")
)
