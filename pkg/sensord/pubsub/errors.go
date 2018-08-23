package pubsub

import "errors"

var (
	ErrExistedPublisher        = errors.New("existed publisher")
	ErrNotFoundPublisher       = errors.New("publisher not found")
	ErrExistedSubscriber       = errors.New("existed subscriber")
	ErrNotFoundSubscriber      = errors.New("subcsriber not found")
	ErrUnregisteredManagerName = errors.New("unregistered pubsub manager")
	ErrUnsubscribable          = errors.New("unsubscribeable")
)
