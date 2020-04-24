package metathings_deviced_flow

import "errors"

var (
	ErrUnknownFlowFactory          = errors.New("unknown flow factory")
	ErrUnknownFlowSetFactory       = errors.New("unknown flow set factory")
	ErrGetAliveRedisClientMaxRetry = errors.New("get alive redis client max retry")
)
