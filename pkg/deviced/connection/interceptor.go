package metathings_deviced_connection

import (
	"errors"

	pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	InterceptorStop = errors.New("interceptor stop")
)

type ConnectResponseMatcher func(*pb.ConnectResponse) bool
type ConnectResponseInterceptor func(*pb.ConnectResponse) error
type ConnectResponseInterceptorChain func(*pb.ConnectResponse) error

func NewConnectResponseInterceptorChain(args ...interface{}) ConnectResponseInterceptorChain {
	matchers := []ConnectResponseMatcher{}
	interceptors := []ConnectResponseInterceptor{}

	for i := 0; i < len(args); i += 2 {
		matchers = append(matchers, args[i].(ConnectResponseMatcher))
		interceptors = append(interceptors, args[i+1].(ConnectResponseInterceptor))
	}

	return func(req *pb.ConnectResponse) error {
		for i := 0; i < len(matchers); i++ {
			matcher := matchers[i]
			interceptor := interceptors[i]
			if matcher(req) {
				if err := interceptor(req); err != nil {
					return err
				}
			}

		}
		return nil
	}
}

func NewConnectResponseUnaryCallMatcher(kind pb.ConnectMessageKind, component, name, method string) ConnectResponseMatcher {
	return func(req *pb.ConnectResponse) bool {
		unary := req.GetUnaryCall()
		if unary == nil {
			return false
		}

		if req.GetKind() != kind {
			return false
		}

		if component != "*" {
			if unary.GetComponent() != component {
				return false
			}
		}

		if name != "*" {
			if unary.GetName() != name {
				return false
			}
		}

		if method != "*" {
			if unary.GetMethod() != method {
				return false
			}
		}

		return true
	}
}
