package pubsub

import "errors"

var (
	ErrExistedPublisher      = errors.New("existed publisher")
	ErrExistedSubscriber     = errors.New("existed subscriber")
	ErrUnregisterManagerName = errors.New("unregistered pubsub manager")
	ErrUnsubscribable        = errors.New("unsubscribeable")
)
