package metathings_component

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	anypb "github.com/golang/protobuf/ptypes/any"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/objx"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	pb "github.com/nayotta/metathings/proto/component"
)

const (
	METATHINGS_SODA_MODULE_CLIENT_USERAGENT = "Metathings-Soda-Module-Client"
)

func resolve_http_url(m *Module, meth string, defaults ...string) (string, error) {
	cfg := m.Kernel().Config()

	base := cfg.GetString("backend.target.url")
	base_url, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	path := cfg.GetString(fmt.Sprintf("backend.downstreams.%v.path", meth))
	if path == "" {
		if len(defaults) == 0 {
			return "", ErrDownstreamNotFound
		}

		path = defaults[0]
	}

	uri, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	full_url := base_url.ResolveReference(uri)

	return full_url.String(), nil
}

func unmarshal_any_to_json_string(m *anypb.Any) (string, error) {
	var st stpb.Struct
	if err := ptypes.UnmarshalAny(m, &st); err != nil {
		return "", err
	}

	jsbuf, err := grpc_helper.JSONPBMarshaler.MarshalToString(&st)
	if err != nil {
		return "", err
	}

	return jsbuf, nil
}

func marshal_json_string_to_any(buf string) (*anypb.Any, error) {
	var st stpb.Struct

	err := grpc_helper.JSONPBUnmarshaler.Unmarshal(strings.NewReader(buf), &st)
	if err != nil {
		return nil, err
	}

	m, err := ptypes.MarshalAny(&st)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func WrapHttpAuthContext(r *http.Request, ctx *SodaModuleAuthContext) error {
	scheme := ctx.Get("scheme").String()
	credential := ctx.Get("credential").String()

	if scheme == "" || credential == "" {
		return nil
	}

	r.Header.Set("Authorization", scheme+" "+credential)

	return nil
}

type SodaModuleHttpWrapper struct {
	m *Module

	req_auth SodaModuleAuthorizer
}

func (w *SodaModuleHttpWrapper) http_client() *http.Client {
	cfg := w.m.Kernel().Config()

	timeout := cfg.GetDuration("backend.timeout")
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	return &http.Client{
		Timeout: timeout,
	}
}

func (w *SodaModuleHttpWrapper) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	meth := req.GetMethod().GetValue()

	full_url, err := resolve_http_url(w.m, meth)
	if err != nil {
		return nil, err
	}

	jsbuf, err := unmarshal_any_to_json_string(req.GetValue())
	if err != nil {
		return nil, err
	}

	http_req, err := http.NewRequest(http.MethodPost, full_url, strings.NewReader(jsbuf))
	if err != nil {
		return nil, err
	}

	http_req.Header.Set("Content-Type", "application/json")
	http_req.Header.Set("User-Agent", METATHINGS_SODA_MODULE_CLIENT_USERAGENT)

	// TODO(Peer): pass request body to sign
	signature, err := w.req_auth.Sign(
		&SodaModuleAuthContext{objx.New(map[string]interface{}{})},
	)
	if err != nil {
		return nil, err
	}

	if err = WrapHttpAuthContext(http_req, signature); err != nil {
		return nil, err
	}

	http_res, err := w.http_client().Do(http_req)
	if err != nil {
		return nil, err
	}
	defer http_res.Body.Close()

	// TODO(Peer): handle http error
	buf, err := ioutil.ReadAll(http_res.Body)
	if err != nil {
		return nil, err
	}

	any_val, err := marshal_json_string_to_any(string(buf))
	if err != nil {
		return nil, err
	}

	res := &pb.UnaryCallResponse{
		Method: meth,
		Value:  any_val,
	}

	return res, nil
}

func (w *SodaModuleHttpWrapper) StreamCall(pb.ModuleService_StreamCallServer) error {
	return ErrHandleUnimplemented
}

type SodaModuleHttpWrapperFactory struct{}

func (f *SodaModuleHttpWrapperFactory) NewModuleWrapper(m *Module) (SodaModuleWrapper, error) {
	cfg := m.Kernel().Config()

	req_auth_name := cfg.GetString("backend.request_auth.name")
	if req_auth_name == "" {
		req_auth_name = "dummy"
	}
	req_auth, err := NewSodaModuleAuthorizer(req_auth_name, m)
	if err != nil {
		return nil, err
	}

	return &SodaModuleHttpWrapper{
		m:        m,
		req_auth: req_auth,
	}, nil
}

func init() {
	register_soda_module_wrapper_factory("http", new(SodaModuleHttpWrapperFactory))
}
