package grpc_helper

import (
	"context"
	"reflect"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(logger log.FieldLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
		var cast_srv grpc_auth.ServiceAuthFuncOverride
		var new_ctx context.Context
		var ok bool

		if cast_srv, ok = info.Server.(grpc_auth.ServiceAuthFuncOverride); !ok {
			return handler(new_ctx, req)
		}

		if new_ctx, err = cast_srv.AuthFuncOverride(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		md, _ := ParseMethodDescription(info.FullMethod)
		srv_val := reflect.ValueOf(info.Server)

		vlt_func := srv_val.MethodByName("Validate" + md.Method)
		if vlt_func.Kind() == reflect.Func {
			if err = vlt_func.Interface().(func(context.Context, interface{}) error)(new_ctx, req); err != nil {
				return nil, err
			}
		}

		auth_func := srv_val.MethodByName("Authorize" + md.Method)
		if auth_func.Kind() == reflect.Func {
			if err = auth_func.Interface().(func(context.Context, interface{}) error)(new_ctx, req); err != nil {
				return nil, err
			}
		}

		defer func(callat time.Time) {
			logger.WithFields(log.Fields{
				"#method":  info.FullMethod,
				"#elapsed": time.Since(callat),
			}).Debugf("call tracing")
		}(time.Now())
		return handler(new_ctx, req)
	}
}

func StreamServerInterceptor(logger log.FieldLogger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		var cast_srv grpc_auth.ServiceAuthFuncOverride
		var new_ctx context.Context
		var ok bool

		if cast_srv, ok = srv.(grpc_auth.ServiceAuthFuncOverride); !ok {
			return handler(srv, stream)
		}

		if new_ctx, err = cast_srv.AuthFuncOverride(stream.Context(), info.FullMethod); err != nil {
			return err
		}

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = new_ctx

		md, _ := ParseMethodDescription(info.FullMethod)
		srv_val := reflect.ValueOf(srv)

		vlt_func := srv_val.MethodByName("Validate" + md.Method)
		if vlt_func.Kind() == reflect.Func {
			if err = vlt_func.Interface().(func(grpc.ServerStream) error)(wrapped); err != nil {
				return err
			}
		}

		auth_func := srv_val.MethodByName("Authorize" + md.Method)
		if auth_func.Kind() == reflect.Func {
			if err = auth_func.Interface().(func(grpc.ServerStream) error)(wrapped); err != nil {
				return err
			}
		}

		defer func(callat time.Time) {
			logger.WithFields(log.Fields{
				"#method":  info.FullMethod,
				"#elapsed": time.Since(callat),
			}).Debugf("call tracing")
		}(time.Now())
		return handler(srv, wrapped)
	}
}
