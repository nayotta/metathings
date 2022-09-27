package metathings_component

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	anypb "github.com/golang/protobuf/ptypes/any"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/objx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nhooyr.io/websocket"

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

func unmarshal_any_to_bytes(m *anypb.Any) ([]byte, error) {
	var bs wrappers.BytesValue
	if err := ptypes.UnmarshalAny(m, &bs); err != nil {
		return nil, err
	}
	return bs.GetValue(), nil
}

func marshal_bytes_to_any(buf []byte) (*anypb.Any, error) {
	bs := &wrappers.BytesValue{Value: buf}
	m, err := ptypes.MarshalAny(bs)
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

func (w *SodaModuleHttpWrapper) StreamCall(upstm pb.ModuleService_StreamCallServer) error {
	var req *pb.StreamCallRequest
	var err error
	sess := rand.Int63()

	logger := w.Logger().WithFields(log.Fields{
		"#method": "StreamCall",
		"session": sess,
	})

	if req, err = upstm.Recv(); err != nil {
		logger.WithError(err).Errorf("failed to receive config message")
		return status.Errorf(codes.Internal, err.Error())
	}

	cfg := req.GetConfig()
	if cfg == nil {
		err = ErrInvalidArguments
		logger.WithError(err).Errorf("failed to get config from message")
		return status.Errorf(codes.FailedPrecondition, err.Error())
	}

	meth := cfg.GetMethod().GetValue()
	full_url, err := resolve_http_url(w.m, meth)
	switch err {
	case ErrDownstreamNotFound:
		logger.WithError(err).Errorf("down stream not found")
		return status.Errorf(codes.NotFound, err.Error())
	case nil:
		break
	default:
		logger.WithError(err).Errorf("failed to resolve method")
		return status.Errorf(codes.Internal, err.Error())
	}
	logger = logger.WithFields(log.Fields{
		"method": meth,
		"url":    full_url,
	})

	wsConn, _, err := websocket.Dial(context.Background(), full_url, nil)
	if err != nil {
		logger.WithError(err).Errorf("failed to dial websocket")
		return status.Errorf(codes.Internal, err.Error())
	}

	cctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	up2down_wait := make(chan struct{})
	down2up_wait := make(chan struct{})
	go w.stm_up2down(upstm, wsConn, sess, up2down_wait)
	go w.stm_down2up(upstm, cctx, wsConn, sess, down2up_wait)

	logger.Debugf("stream call started")
	select {
	case <-up2down_wait:
	case <-down2up_wait:
	}

	logger.Debugf("stream call done")

	return nil
}

func (w *SodaModuleHttpWrapper) Logger() *log.Entry {
	return w.m.Logger().WithField("#instance", "SodaModuleHttpWrapper")
}

func (w *SodaModuleHttpWrapper) stm_up2down(upstm pb.ModuleService_StreamCallServer, downstm *websocket.Conn, sess int64, wait chan struct{}) {
	logger := w.Logger().WithFields(log.Fields{
		"dir":     "up->down",
		"session": sess,
	})

	defer close(wait)
	for epoch := uint64(0); ; epoch++ {
		logger := logger.WithFields(log.Fields{
			"epoch": epoch,
		})

		req, err := upstm.Recv()
		if err != nil {
			logger.WithError(err).Debugf("failed to recv msg from upstm")
			return
		}
		logger.Tracef("recv msg from upstm")

		buf, err := unmarshal_any_to_bytes(req.GetData().GetValue())
		if err != nil {
			logger.WithError(err).Debugf("failed to unmarshal any to json string")
			return
		}

		if err = downstm.Write(context.Background(), websocket.MessageBinary, buf); err != nil {
			logger.WithError(err).Debugf("failed to write json buffer to downstm")
			return
		}
		logger.Tracef("send json buffer to downstm")
	}
}

func (w *SodaModuleHttpWrapper) stm_down2up(upstm pb.ModuleService_StreamCallServer, dsCtx context.Context, downstm *websocket.Conn, sess int64, wait chan struct{}) {
	logger := w.Logger().WithFields(log.Fields{
		"dir":     "down->up",
		"session": sess,
	})

	defer close(wait)
	for epoch := uint64(0); ; epoch++ {
		logger := logger.WithFields(log.Fields{
			"epoch": epoch,
		})

		_, buf, err := downstm.Read(dsCtx)
		if err != nil {
			logger.WithError(err).Debugf("failed to recv msg from downstm")
			return
		}
		logger.Tracef("recv msg from downstm")

		anyVal, err := marshal_bytes_to_any(buf)
		if err != nil {
			logger.WithError(err).Debugf("failed to marshal json string to any")
			return
		}

		res := &pb.StreamCallResponse{
			Response: &pb.StreamCallResponse_Data{
				Data: &pb.StreamCallDataResponse{
					Value: anyVal,
				},
			},
		}

		if err = upstm.Send(res); err != nil {
			logger.WithError(err).Debugf("failed to send msg to upstm")
			return
		}
		logger.Tracef("send msg to upstm")
	}
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
