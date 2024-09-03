package cmd

import (
	"errors"

	e "github.com/PeerXu/errors-go"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")

	ErrUnsupportedStorageDriver, ErrUnsupportedStorageDriverFn               = e.NewErrorAndErrorFunc[string]("unsupported storage driver")
	ErrUnsupportedProfileStorageDriver, ErrUnsupportedProfileStorageDriverFn = e.NewErrorAndErrorFunc[string]("unsupported profile storage driver")
)
