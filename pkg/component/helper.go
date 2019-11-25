package metathings_component

import (
	"net/url"
	"strings"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type TransportCredential struct {
	Insecure  bool
	PlainText bool   `mapstructure:"plain_text"`
	KeyFile   string `mapstructure:"key_file"`
	CertFile  string `mapstructure:"cert_file"`
}

type ServiceEndpoint struct {
	TransportCredential `mapstructure:",squash"`
	Address             string
}

func ToModule(v **Module) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(*Module); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}

type Endpoint struct {
	*url.URL
}

func (ep *Endpoint) IsMetathingsProtocol() bool {
	return strings.HasPrefix(strings.ToLower(ep.Scheme), "mtp")
}

func (ep *Endpoint) GetTransportProtocol(defaults ...string) string {
	tp := "grpc"
	if len(defaults) > 0 {
		tp = defaults[0]
	}
	scheme := strings.ToLower(ep.Scheme)
	if strings.HasPrefix(scheme, "mtp+") {
		tp = strings.TrimPrefix(scheme, "mtp+")
	}
	return tp
}

func ParseEndpoint(ep string) (*Endpoint, error) {
	url, err := url.Parse(ep)
	if err != nil {
		return nil, err
	}

	return &Endpoint{url}, nil
}
