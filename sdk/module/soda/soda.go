package metathings_module_soda_sdk

import (
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/PeerXu/option-go"
	"github.com/sirupsen/logrus"

	option_helper "github.com/nayotta/metathings/pkg/common/option"
)

func NewDefaultSodaClientOption() option.Option {
	return option.NewOption(map[string]any{
		option_helper.OPTION_HTTP_CLIENT: http.DefaultClient,
	})
}

type SodaClient interface {
	PutObjectStreaming(name string, src io.ReadSeekCloser, len int64, opts PutObjectStreamingOption) error
}

func NewSodaClient(opts ...NewSodaClientOption) (SodaClient, error) {
	o := option.ApplyWithDefault(NewDefaultSodaClientOption(), opts...)

	logger, err := GetLogger(o)
	if err != nil {
		return nil, err
	}

	httpClient, err := GetHTTPClient(o)
	if err != nil {
		return nil, err
	}

	url, err := GetURL(o)
	if err != nil {
		return nil, err
	}

	return &sodaClient{
		logger:     logger,
		httpClient: httpClient,
		baseURL:    url,
	}, nil
}

type sodaClient struct {
	logger     *logrus.Entry
	httpClient *http.Client
	baseURL    string
}

func (c *sodaClient) joinPath(s string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = path.Join(u.Path, s)
	return u.String()
}
