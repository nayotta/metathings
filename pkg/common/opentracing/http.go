package opentracing_helper

import (
	"context"
	"io"
	"net/http"
	"reflect"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func GetSpanFromRequest(r *http.Request, name string) opentracing.Span {
	tracer := opentracing.GlobalTracer()
	span_ctx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	sp := tracer.StartSpan(name, ext.RPCServerOption(span_ctx))
	ext.HTTPMethod.Set(sp, r.Method)
	ext.HTTPUrl.Set(sp, r.URL.String())
	ext.Component.Set(sp, "net/http")
	r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))
	return sp
}

func InjectSpanToRequest(span opentracing.Span, r *http.Request) {
	tracer := opentracing.GlobalTracer()

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, r.URL.String())
	ext.HTTPMethod.Set(span, r.Method)

	tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
}

type IsTracedGetter interface {
	IsTraced(*http.Request) bool
}

type ComponentNameGetter interface {
	ComponentName() string
}

func Middleware(srv interface{}, name string, options ...nethttp.MWOption) http.HandlerFunc {
	srv_val := reflect.ValueOf(srv)
	meth_val := srv_val.MethodByName(name)
	h := meth_val.Interface().(func(http.ResponseWriter, *http.Request))

	var isTraced func(*http.Request) bool
	isTracedGetter, ok := srv.(IsTracedGetter)
	if !ok {
		isTraced = func(*http.Request) bool { return false }
	} else {
		isTraced = isTracedGetter.IsTraced
	}

	var componentName string
	componentNameGetter, ok := srv.(ComponentNameGetter)
	if !ok {
		componentName = "net/http"
	} else {
		componentName = componentNameGetter.ComponentName()
	}
	options = append(options,
		nethttp.OperationNameFunc(func(r *http.Request) string {
			return r.Method + " " + r.URL.Path
		}),
		nethttp.MWSpanFilter(isTraced),
		nethttp.MWComponentName(componentName),
	)

	return nethttp.MiddlewareFunc(opentracing.GlobalTracer(), h, options...)
}

func NewRequester(srv IsTracedGetter) func(...nethttp.ClientOption) func(context.Context, string, string, io.Reader) (*http.Request, *nethttp.Tracer, error) {
	return func(options ...nethttp.ClientOption) func(context.Context, string, string, io.Reader) (*http.Request, *nethttp.Tracer, error) {
		return func(ctx context.Context, method, url string, body io.Reader) (*http.Request, *nethttp.Tracer, error) {
			req, err := http.NewRequest(method, url, body)
			if err != nil {
				return nil, nil, err
			}

			if !srv.IsTraced(nil) {
				return req, nil, nil
			}

			options = append(options,
				nethttp.OperationName(req.Method+" "+req.URL.Path),
			)
			req = req.WithContext(ctx)
			req, hr := nethttp.TraceRequest(opentracing.GlobalTracer(), req, options...)

			return req, hr, nil
		}
	}
}

func GetHTTPClient(srv IsTracedGetter) *http.Client {
	if !srv.IsTraced(nil) {
		return http.DefaultClient
	}

	return &http.Client{Transport: &nethttp.Transport{}}
}
