package file_helper

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrHashNotMatch    = errors.New("hash not match")
	ErrBadChunkIndex   = errors.New("bad chunk index")
	ErrEmptyCacheFile  = errors.New("empty cache file")
)

var (
	DONE = errors.New("done")
)
