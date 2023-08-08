package metathings_module_soda_sdk

import (
	"github.com/PeerXu/option-go"

	option_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	OPTION_URL = "url"
)

type PutObjectStreamingOption struct {
	Sha1sum string
}

type NewSodaClientOption = option.ApplyOption

var (
	WithLogger, GetLogger         = option_helper.WithLogger, option_helper.GetLogger
	WithURL, GetURL               = option.New[string](OPTION_URL)
	WithHTTPClient, GetHTTPClient = option_helper.WithHTTPClient, option_helper.GetHTTPClient
)
