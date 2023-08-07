package metathings_component

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	http_helper "github.com/nayotta/metathings/pkg/common/http"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
)

const (
	HTTP_SODA_OBJECT_STREAM_NAME          = "Metathings-Soda-Object-Stream-Name"
	HTTP_SODA_OBJECT_SHA1SUM              = "Metathings-Soda-Object-Sha1sum"
	HTTP_SODA_OBJECT_LENGTH               = "Metathings-Soda-Object-Length"
	HTTP_SODA_OBJECT_UPLOADED_LENGTH      = "Metathings-Soda-Object-Uploaded-Length"
	HTTP_SODA_OBJECT_STREAM_MAX_AGE       = "Metathings-Soda-Object-Stream-Max-Age"
	HTTP_SODA_OBJECT_STREAM_REMAINED      = "Metathings-Soda-Object-Stream-Remained"
	HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET  = "Metathings-Soda-Object-Stream-Chunk-Offset"
	HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH  = "Metathings-Soda-Object-Stream-Chunk-Length"
	HTTP_SODA_OBJECT_STREAM_CHUNK_SHA1SUM = "Metathings-Soda-Object-Stream-Chunk-Sha1sum"
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

	oss   sync.Map
	auth  SodaModuleAuthorizer
	httpd *http.Server

	objectStreams sync.Map
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
		logger.WithError(err).Errorf("failed to parse request body")
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

func (b *SodaModuleHttpBackend) handle_put_object_streaming(w http.ResponseWriter, r *http.Request) {
	logger := b.m.Logger().WithField("action", "put_object_streaming")

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
	length := bodyx.Get("object.length").Int()
	sha1sum := bodyx.Get("object.sha1sum").String()

	logger = logger.WithFields(logging.Fields{
		"name":    name,
		"length":  length,
		"sha1sum": sha1sum,
	})

	osName := b.parse_object_stream_name(name, sha1sum)
	os, err := b.get_object_stream(osName)
	if err != nil {
		os, err = NewObjectStream(
			WithLogger(b.m.RawLogger()),
			WithName(osName),
			WithSha1sum(sha1sum),
			WithLength(int64(length)),
		)
		if err != nil {
			logger.WithError(err).Debugf("failed to new object stream")
			jw.WriteHeader(http.StatusInternalServerError)
			jw.WriteJSON(http_helper.ConvertError(err))
			return
		}

		if err = b.add_object_stream(osName, os); err != nil {
			logger.WithError(err).Debugf("failed to add object stream")
			jw.WriteHeader(http.StatusInternalServerError)
			jw.WriteJSON(http_helper.ConvertError(err))
			return
		}

		go func() {
			defer b.remove_object_stream(osName)

			if err = b.m.Kernel().PutObjectStreaming(name, os, &PutObjectStreamingOption{
				Sha1:   sha1sum,
				Length: int64(length),
			}); err != nil {
				logger.WithError(err).Errorf("failed to put object streaming")
			}
		}()
	} else {
		logger.Warningf("put object streaming is in progressing")
	}

	jw.Header().Add(HTTP_SODA_OBJECT_STREAM_NAME, osName)
	jw.Header().Add(HTTP_SODA_OBJECT_STREAM_MAX_AGE, cast.ToString(os.MaxAge()))
	jw.Header().Add(HTTP_SODA_OBJECT_STREAM_REMAINED, cast.ToString(os.Remained()))
	jw.WriteHeader(http.StatusNoContent)

	return
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
		// HACK: rewrite not-unmarshalable fields
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

func (b *SodaModuleHttpBackend) handle_object_stream_write_chunk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["object_stream"]
	logger := b.m.Logger().WithFields(logging.Fields{
		"action":        "write_object_chunk",
		"object_stream": name,
	})
	jw := http_helper.WrapJSONResponseWriter(w)

	os, err := b.get_object_stream(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stream")
		jw.WriteJSONError(http.StatusNotFound, err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		logger.WithError(err).Debugf("failed to get upload file")
		jw.WriteJSONError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	opts, err := b.parse_write_object_chunk_options(r)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse write object chunk options")
		jw.WriteJSONError(http.StatusBadRequest, err)
		return
	}
	logger = logger.WithFields(logging.Fields{
		"chunk-length":  opts.Length,
		"chunk-offset":  opts.Offset,
		"chunk-sha1sum": opts.Sha1sum,
	})

	var buf bytes.Buffer

	n, err := io.CopyN(&buf, file, opts.Length)
	if err != nil {
		logger.WithError(err).Debugf("failed to read data from upload file")
		jw.WriteJSONError(http.StatusInternalServerError, err)
		return
	}

	// check length
	if opts.Length != n {
		err = fmt.Errorf("invalid length")
		logger.WithError(err).WithField("acutal-length", n).Debugf("unmatched upload file length")
		jw.WriteJSONError(http.StatusBadRequest, err)
		return
	}

	// check sha1sum
	actualSha1sum := sha1sum(buf.Bytes())
	if opts.Sha1sum != actualSha1sum {
		err = fmt.Errorf("invalid sha1sum")
		logger.WithError(err).WithField("actual-sha1sum", actualSha1sum).Debugf("unmatched upload file sha1sum")
		jw.WriteJSONError(http.StatusBadRequest, err)
		return
	}

	if _, err = os.Write(buf.Bytes()); err != nil {
		logger.WithError(err).Debugf("failed to write buffer to object stream")
		jw.WriteJSONError(http.StatusInternalServerError, err)
		return
	}

	jw.WriteHeader(http.StatusNoContent)
	logger.Tracef("write object chunk")

	return
}

func (b *SodaModuleHttpBackend) handle_object_stream_next_chunk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["object_stream"]

	logger := b.m.Logger().WithFields(logging.Fields{
		"action":        "next_object_stream",
		"object_stream": name,
	})

	jw := http_helper.WrapJSONResponseWriter(w)

	os, err := b.get_object_stream(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stream")
		jw.WriteHeader(http.StatusNotFound)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	offset, err := os.Seek(0, io.SeekCurrent)
	if err != nil {
		logger.WithError(err).Debugf("failed to seek")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	h := jw.Header()
	h.Add(HTTP_SODA_OBJECT_STREAM_REMAINED, cast.ToString(os.Remained()))
	h.Add(HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET, cast.ToString(offset))
	h.Add(HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH, cast.ToString(os.BufferLength()))
	jw.WriteHeader(http.StatusNoContent)

	logger.Tracef("show object stream")

	return
}

func (b *SodaModuleHttpBackend) handle_object_stream_show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["object_stream"]

	logger := b.m.Logger().WithFields(logging.Fields{
		"action":        "show_object_stream",
		"object_stream": name,
	})

	jw := http_helper.WrapJSONResponseWriter(w)

	os, err := b.get_object_stream(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stream")
		jw.WriteHeader(http.StatusNotFound)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	h := jw.Header()
	h.Add(HTTP_SODA_OBJECT_STREAM_NAME, os.Name())
	h.Add(HTTP_SODA_OBJECT_SHA1SUM, os.Sha1sum())
	h.Add(HTTP_SODA_OBJECT_LENGTH, cast.ToString(os.Length()))
	h.Add(HTTP_SODA_OBJECT_UPLOADED_LENGTH, cast.ToString(os.Uploaded()))
	h.Add(HTTP_SODA_OBJECT_STREAM_MAX_AGE, cast.ToString(os.MaxAge()))
	h.Add(HTTP_SODA_OBJECT_STREAM_REMAINED, cast.ToString(os.Remained()))
	h.Add(HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET, cast.ToString(os.Offset()))
	h.Add(HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH, cast.ToString(os.BufferLength()))
	jw.WriteHeader(http.StatusNoContent)

	logger.Tracef("show object stream")

	return
}

func (b *SodaModuleHttpBackend) handle_object_stream_cancel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["object_stream"]

	logger := b.m.Logger().WithFields(logging.Fields{
		"action":        "cancel",
		"object_stream": name,
	})

	jw := http_helper.WrapJSONResponseWriter(w)

	os, err := b.get_object_stream(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stream")
		jw.WriteHeader(http.StatusNotFound)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	if err = os.Close(); err != nil {
		logger.WithError(err).Debugf("failed to cancel object stream")
		jw.WriteHeader(http.StatusInternalServerError)
		jw.WriteJSON(http_helper.ConvertError(err))
		return
	}

	logger.Tracef("cancel object stream")

	return
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

func (b *SodaModuleHttpBackend) add_object_stream(name string, os ObjectStream) error {
	_, loaded := b.objectStreams.LoadOrStore(name, os)
	if loaded {
		return ErrObjectStreamFound
	}

	return nil
}

func (b *SodaModuleHttpBackend) remove_object_stream(name string) error {
	b.objectStreams.Delete(name)
	return nil
}

func (b *SodaModuleHttpBackend) get_object_stream(name string) (ObjectStream, error) {
	v, ok := b.objectStreams.Load(name)
	if !ok {
		return nil, ErrObjectStreamNotFound
	}

	return v.(ObjectStream), nil
}

func (b *SodaModuleHttpBackend) parse_write_object_chunk_options(r *http.Request) (opts WriteObjectChunkOptions, err error) {
	opts.Sha1sum = r.Form.Get(HTTP_SODA_OBJECT_STREAM_CHUNK_SHA1SUM)
	opts.Length = cast.ToInt64(r.Form.Get(HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH))
	opts.Offset = cast.ToInt64(r.Form.Get(HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET))
	return
}

func (b *SodaModuleHttpBackend) parse_object_stream_name(name, sha1sum string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s$%s", name, sha1sum)))
	return hex.EncodeToString(h.Sum(nil))
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
		router.HandleFunc("/v1/actions/heartbeat", b.handle_heartbeat).Methods(http.MethodPost)
	}

	sr := router.PathPrefix("/v1/actions").Subrouter()
	sr.HandleFunc("/show", b.handle_show).Methods(http.MethodPost)
	sr.HandleFunc("/push_frame_to_flow_once", b.handle_push_frame_to_flow_once).Methods(http.MethodPost)
	sr.HandleFunc("/put_object", b.handle_put_object).Methods(http.MethodPost)
	sr.HandleFunc("/put_object_streaming", b.handle_put_object_streaming).Methods(http.MethodPost)
	sr.HandleFunc("/get_object", b.handle_get_object).Methods(http.MethodPost)
	sr.HandleFunc("/get_object_content", b.handle_get_object_content).Methods(http.MethodPost)
	sr.HandleFunc("/remove_object", b.handle_remove_object).Methods(http.MethodPost)
	sr.HandleFunc("/rename_object", b.handle_rename_object).Methods(http.MethodPost)
	sr.HandleFunc("/list_objects", b.handle_list_objects).Methods(http.MethodPost)
	router.Use(b.authorize_middleware())

	// object stream
	ossr := router.PathPrefix("/v1/object_streams/{object_stream}/actions").Subrouter()
	ossr.HandleFunc("/write_chunk", b.handle_object_stream_write_chunk).Methods(http.MethodPost)
	ossr.HandleFunc("/next_chunk", b.handle_object_stream_next_chunk).Methods(http.MethodPost)
	ossr.HandleFunc("/show", b.handle_object_stream_show).Methods(http.MethodPost)
	ossr.HandleFunc("/cancel", b.handle_object_stream_cancel).Methods(http.MethodPost)

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
