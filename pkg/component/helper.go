package metathings_component

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/url"
	"os"
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

func NewPutObjectStreamingOptionFromPath(path string) (*PutObjectStreamingOption, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	sha1 := sha1.New()
	_, err = io.Copy(sha1, f)
	if err != nil {
		return nil, err
	}
	sha1_str := fmt.Sprintf("%x", sha1.Sum(nil))

	opt := &PutObjectStreamingOption{
		Length: stat.Size(),
		Sha1:   sha1_str,
	}

	return opt, nil
}

func chunkSha1sum(p []byte) []byte {
	h := sha1.New()
	h.Write(p)
	return h.Sum(nil)
}
