package metathings_component

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	http_helper "github.com/nayotta/metathings/pkg/common/http"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
)

type HttpAuthContextParser func(*http.Request) (*SodaModuleAuthContext, error)

var http_auth_context_parser_map map[string]HttpAuthContextParser
var http_auth_context_parser_map_once sync.Once

func register_http_auth_context_parser(name string, p HttpAuthContextParser) {
	http_auth_context_parser_map_once.Do(func() {
		http_auth_context_parser_map = make(map[string]HttpAuthContextParser)
	})
	http_auth_context_parser_map[name] = p
}

func ParseHttpAuthContext(name string, r *http.Request) (*SodaModuleAuthContext, error) {
	p, ok := http_auth_context_parser_map[name]
	if !ok {
		p = http_auth_context_parser_map["default"]
	}

	return p(r)
}

func parse_http_auth_context(r *http.Request) (*SodaModuleAuthContext, error) {
	auth := r.Header.Get("Authorization")

	tokens := strings.SplitN(auth, " ", 2)
	if len(tokens) != 2 {
		return nil, ErrUnexpectedTokenFormat
	}

	return &SodaModuleAuthContext{objx.New(map[string]interface{}{
		"scheme":     tokens[0],
		"credential": tokens[1],
	})}, nil
}

func convert_protobuf_message_to_interface(x proto.Message) (interface{}, error) {
	var y interface{}

	buf, err := grpc_helper.JSONPBMarshaler.MarshalToString(x)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(buf), &y); err != nil {
		return nil, err
	}

	return y, nil
}

type SodaModuleHttpBackend struct {
	m *Module

	done chan struct{}

	auth  SodaModuleAuthorizer
	httpd *http.Server
}

func (b *SodaModuleHttpBackend) authorize_middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cfg := b.m.Kernel().Config()
			auth_name := cfg.GetString("backend.auth.name")

			logger := b.m.Logger()

			jw := http_helper.WrapJSONResponseWriter(w)

			ctx, err := ParseHttpAuthContext(auth_name, r)
			if err != nil {
				logger.WithError(err).Errorf("failed to parse http request to auth context")
				jw.WriteHeader(http.StatusUnauthorized)
				jw.WriteJSON(http_helper.ConvertError(err))
				return
			}

			err = b.auth.Verify(ctx)
			if err != nil {
				logger.WithError(err).Errorf("failed to authorize with auth context")
				jw.WriteHeader(http.StatusUnauthorized)
				jw.WriteJSON(http_helper.ConvertError(err))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (b *SodaModuleHttpBackend) handle_show(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "show")

	jw := http_helper.WrapJSONResponseWriter(w)

	mdl, err := b.m.Kernel().Show()
	if err != nil {
		logger.WithError(err).Errorf("failed to show in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	m, err := convert_protobuf_message_to_interface(mdl)
	if err != nil {
		logger.WithError(err).Errorf("failed to convert module to string map")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusOK)
	jw.WriteJSON(map[string]interface{}{
		"module": m,
	})

	logger.Debugf("show")
}

func (b *SodaModuleHttpBackend) handle_push_frame_to_flow_once(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "push_frame_to_flow_once")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	id := bodyx.Get("id").String()
	if id == "" {
		id = id_helper.NewId()
	}

	frm_ts, err := cast.ToTimeE(bodyx.Get("frame.ts").String())
	if err != nil {
		frm_ts = time.Now()
	}

	opt := &PushFrameToFlowOnceOption{
		Id: &id,
		Ts: &frm_ts,
	}

	flow_name := bodyx.Get("flow.name").String()
	frm_data := bodyx.Get("frame.data").Inter()

	logger = logger.WithFields(log.Fields{
		"id":   id,
		"name": flow_name,
		"ts":   frm_ts,
	})

	if err = b.m.Kernel().PushFrameToFlowOnce(flow_name, frm_data, opt); err != nil {
		logger.WithError(err).Errorf("failed to push frame to flow once in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusNoContent)

	logger.Debugf("push frame to flow once")
}

func (b *SodaModuleHttpBackend) handle_heartbeat(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "heartbeat")

	jw := http_helper.WrapJSONResponseWriter(w)

	err := b.m.Kernel().Heartbeat()
	if err != nil {
		logger.WithError(err).Errorf("failed to heartbeat in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusNoContent)

	logger.Debugf("heartbeat")
}

func (b *SodaModuleHttpBackend) handle_put_object(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "put_object")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	name := bodyx.Get("object.name").String()
	content := bodyx.Get("object.content").String()

	logger = logger.WithField("name", name)

	if err = b.m.Kernel().PutObject(name, strings.NewReader(content)); err != nil {
		logger.WithError(err).Errorf("failed to put object to device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusNoContent)

	logger.Debugf("put object")
}

func (b *SodaModuleHttpBackend) handle_get_object(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "get_object")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	name := bodyx.Get("object.name").String()

	logger = logger.WithField("name", name)

	obj, err := b.m.Kernel().GetObject(name)
	if err != nil {
		logger.WithError(err).Errorf("failed to get object in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	o, err := convert_protobuf_message_to_interface(obj)
	if err != nil {
		logger.WithError(err).Errorf("failed to convert object to string map")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusOK)
	jw.WriteJSON(map[string]interface{}{
		"object": o,
	})

	logger.Debugf("get object")
}

func (b *SodaModuleHttpBackend) handle_get_object_content(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "get_object_content")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	name := bodyx.Get("object.name").String()

	logger = logger.WithField("name", name)

	content, err := b.m.Kernel().GetObjectContent(name)
	if err != nil {
		logger.WithError(err).Errorf("failed to get object content in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	jw.WriteHeader(http.StatusOK)
	jw.WriteJSON(map[string]interface{}{
		"content": content,
	})

	logger.Debugf("get object content")
}

func (b *SodaModuleHttpBackend) handle_remove_object(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "remove_object")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	name := bodyx.Get("object.name").String()

	logger = logger.WithField("name", name)

	err = b.m.Kernel().RemoveObject(name)
	if err != nil {
		logger.WithError(err).Errorf("failed to remove object in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)

	logger.Debugf("remove object")
}

func (b *SodaModuleHttpBackend) handle_rename_object(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "rename_object")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	source := bodyx.Get("source.name").String()
	destination := bodyx.Get("destination.name").String()

	logger = logger.WithFields(log.Fields{
		"source":      source,
		"destination": destination,
	})

	err = b.m.Kernel().RenameObject(source, destination)
	if err != nil {
		logger.WithError(err).Errorf("failed to rename object in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)

	logger.Debugf("rename object")
}

func (b *SodaModuleHttpBackend) handle_list_objects(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "list_objects")

	jr := http_helper.WrapJSONRequest(r)
	jw := http_helper.WrapJSONResponseWriter(w)

	bodyx, err := jr.JSON()
	if err != nil {
		logger.WithError(err).Errorf("failed to parse request body")
		jw.WriteHeader(http.StatusBadRequest)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	name := bodyx.Get("object.name").String()
	recursive := bodyx.Get("recursive").Bool()
	depth := cast.ToInt32(bodyx.Get("depth").Inter())

	logger = logger.WithFields(log.Fields{
		"name":      name,
		"recursive": recursive,
		"depth":     depth,
	})

	objs, err := b.m.Kernel().ListObjects(name, &ListObjectsOption{
		Recursive: recursive,
		Depth:     depth,
	})
	if err != nil {
		logger.WithError(err).Errorf("failed to list objects in device")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
	}

	var os []interface{}

	for _, obj := range objs {
		// HACK: rewrite error LastModified
		if obj.LastModified.Seconds < 0 || obj.LastModified.Nanos < 0 {
			obj.LastModified = &timestamppb.Timestamp{}
		}

		o, err := convert_protobuf_message_to_interface(obj)
		if err != nil {
			logger.WithError(err).Errorf("failed to convert object to string map")
			jw.WriteHeader(http.StatusInternalServerError)
			jw.WriteJSON(http_helper.ConvertError(err))
			return
		}

		os = append(os, o)
	}

	jw.WriteHeader(http.StatusOK)
	jw.WriteJSON(map[string]interface{}{
		"objects": os,
	})

	logger.Debugf("list objects")
}

func (b *SodaModuleHttpBackend) health_http_client() *http.Client {
	cfg := b.m.Kernel().Config()
	timeout := cfg.GetDuration("heartbeat.timeout")

	return &http.Client{
		Timeout: timeout,
	}
}

func (b *SodaModuleHttpBackend) rerverse_heartbeat_loop() {
	for b.is_running() {
		go b.reverse_heartbeat_once()
		time.Sleep(time.Duration(b.m.Kernel().Config().GetInt("heartbeat.interval")) * time.Second)
	}
}

func (b *SodaModuleHttpBackend) reverse_heartbeat_once() {
	logger := b.m.Logger()

	url, err := resolve_http_url(b.m, "$health", "/healthz")
	if err != nil {
		logger.WithError(err).Errorf("failed to resolve health url")
		return
	}

	http_req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.WithError(err).Errorf("failed to new health request")
		return
	}

	http_req.Header.Set("User-Agent", METATHINGS_SODA_MODULE_CLIENT_USERAGENT)

	http_res, err := b.health_http_client().Do(http_req)
	if err != nil {
		logger.WithError(err).Errorf("failed to send health request to downstream")
		return
	}
	defer http_res.Body.Close()

	if http_res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(http_res.Body)
		if err != nil {
			logger.WithError(err).Errorf("failed to read response body")
		}

		logger.WithFields(log.Fields{
			"status_code": http_res.StatusCode,
			"body":        string(body),
		})
		return
	}

	if err = b.m.Kernel().Heartbeat(); err != nil {
		logger.WithError(err).Errorf("failed to heartbeat to module")
		return
	}

	logger.Debugf("module heartbeat")
}

func (b *SodaModuleHttpBackend) is_running() bool {
	select {
	case _, running := <-b.done:
		return running
	default:
		return true
	}
}

func (b *SodaModuleHttpBackend) Start() error {
	logger := b.m.Logger()

	cfg := b.m.Kernel().Config()
	host := cfg.GetString("backend.host")
	port := cfg.GetInt("backend.port")
	addr := fmt.Sprintf("%s:%d", host, port)

	router := mux.NewRouter()

	heartbeat_strategy := cfg.GetString("heartbeat.strategy")
	if heartbeat_strategy == "manual" {
		router.HandleFunc("/v1/actions/heartbeat", b.handle_heartbeat).Methods("POST")
	}

	sr := router.PathPrefix("/v1/actions").Subrouter()
	sr.HandleFunc("/show", b.handle_show).Methods("POST")
	sr.HandleFunc("/push_frame_to_flow_once", b.handle_push_frame_to_flow_once).Methods("POST")
	sr.HandleFunc("/put_object", b.handle_put_object).Methods("POST")
	sr.HandleFunc("/get_object", b.handle_get_object).Methods("POST")
	sr.HandleFunc("/get_object_content", b.handle_get_object_content).Methods("POST")
	sr.HandleFunc("/remove_object", b.handle_remove_object).Methods("POST")
	sr.HandleFunc("/rename_object", b.handle_rename_object).Methods("POST")
	sr.HandleFunc("/list_objects", b.handle_list_objects).Methods("POST")
	router.Use(b.authorize_middleware())

	b.httpd = &http.Server{
		Handler: router,
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		if err = b.httpd.Serve(lis); err != nil {
			logger.WithError(err).Warningf("soda module http server exit")
		}

		close(b.done)
	}()

	switch heartbeat_strategy {
	case "reverse":
		go b.rerverse_heartbeat_loop()
	default:
	}

	return nil
}

func (b *SodaModuleHttpBackend) Stop() error {
	return b.httpd.Shutdown(context.TODO())
}

func (b *SodaModuleHttpBackend) Done() <-chan struct{} {
	return b.done
}

func (b *SodaModuleHttpBackend) Health() error {
	return nil
}

func NewSodaModuleHttpBackend(m *Module) (SodaModuleBackend, error) {
	cfg := m.Kernel().Config()

	auth_name := cfg.GetString("backend.auth.name")
	if auth_name == "" {
		auth_name = "dummy"
	}

	auth, err := NewSodaModuleAuthorizer(auth_name, m)
	if err != nil {
		return nil, err
	}

	return &SodaModuleHttpBackend{
		m:    m,
		auth: auth,
		done: make(chan struct{}),
	}, nil
}

func init() {
	register_soda_module_backend_factory("http", NewSodaModuleHttpBackend)
	register_http_auth_context_parser("default", parse_http_auth_context)
}
