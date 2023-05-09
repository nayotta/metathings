package metathings_component

import (
	"time"

	option "github.com/PeerXu/option-go"
	option_helper "github.com/nayotta/metathings/pkg/common/option"
	"github.com/stretchr/objx"
)

const (
	OPTION_NAME          = "name"
	OPTION_SHA1SUM       = "sha1sum"
	OPTION_MAX_AGE       = "maxAge"
	OPTION_OFFSET        = "offset"
	OPTION_LENGTH        = "length"
	OPTION_BUFFER_LENGTH = "bufferLength"
	OPTION_WAIT_TIMEOUT  = "waitTimeout"
)

type NewModuleOption func(objx.Map)

func SetVersion(v string) NewModuleOption {
	return func(opt objx.Map) {
		if v != "" {
			opt.Set("version", v)
		}
	}
}

func SetArgs(vs []string) NewModuleOption {
	return func(opt objx.Map) {
		opt.Set("args", vs)
	}
}

func SetTarget(v interface{}) NewModuleOption {
	return func(opt objx.Map) {
		opt.Set("target", v)
	}
}

type NewObjectStreamOption = option.ApplyOption

var (
	WithLogger, GetLogger             = option_helper.WithLogger, option_helper.GetLogger
	WithName, GetName                 = option.New[string](OPTION_NAME)
	WithSha1sum, GetSha1sum           = option.New[string](OPTION_SHA1SUM)
	WithMaxAge, GetMaxAge             = option.New[time.Duration](OPTION_MAX_AGE)
	WithOffset, GetOffset             = option.New[int64](OPTION_OFFSET)
	WithLength, GetLength             = option.New[int64](OPTION_LENGTH)
	WithBufferLength, GetBufferLength = option.New[int](OPTION_BUFFER_LENGTH)
	WithWaitTimeout, GetWaitTimeout   = option.New[time.Duration](OPTION_WAIT_TIMEOUT)
)
