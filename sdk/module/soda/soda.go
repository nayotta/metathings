package metathings_module_soda_sdk

import (
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/sirupsen/logrus"
)

type SodaClient interface {
	PutObjectStreaming(name string, src io.ReadSeekCloser, len int64, opts PutObjectStreamingOption) error
}

func NewSodaClient() (SodaClient, error) {
	panic("unimplemented")
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
