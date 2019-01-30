package metathings_toolkit_tag

import "errors"

var (
	ErrInvalidArgument         = errors.New("invalid argument")
	ErrUnknownTagToolkitDriver = errors.New("unknown tag toolkit driver")
	ErrNotFound                = errors.New("not found")
)
