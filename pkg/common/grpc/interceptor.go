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

		var startat time.Time
		inner_logger := logger.WithFields(log.Fields{
			"method": info.FullMethod,
		})
		defer func(callat time.Time) {
			inner_logger.WithField("elapsed", time.Since(callat)).Debugf("call tracing")
		}(time.Now())

		if cast_srv, ok = info.Server.(grpc_auth.ServiceAuthFuncOverride); !ok {
			startat = time.Now()
			res, err = handler(new_ctx, req)
			inner_logger.WithField("handle_elapsed", time.Since(startat))
			return
		}

		startat = time.Now()
		if new_ctx, err = cast_srv.AuthFuncOverride(ctx, info.FullMethod); err != nil {
			return nil, err
		}
		inner_logger.WithField("validate_token_elapsed", time.Since(startat))

		md, _ := ParseMethodDescription(info.FullMethod)
		srv_val := reflect.ValueOf(info.Server)

		vlt_func := srv_val.MethodByName("Validate" + md.Method)
		if vlt_func.Kind() == reflect.Func {
			startat = time.Now()
			if err = vlt_func.Interface().(func(context.Context, interface{}) error)(new_ctx, req); err != nil {
				return nil, err
			}
			inner_logger.WithField("validate_request_elapsed", time.Since(startat))
		}

		auth_func := srv_val.MethodByName("Authorize" + md.Method)
		if auth_func.Kind() == reflect.Func {
			startat = time.Now()
			if err = auth_func.Interface().(func(context.Context, interface{}) error)(new_ctx, req); err != nil {
				return nil, err
			}
			inner_logger.WithField("authorize_request_elapsed", time.Since(startat))
		}

		startat = time.Now()
		res, err = handler(new_ctx, req)
		inner_logger.WithField("handle_elapsed", time.Since(startat))
		return
	}
}

func StreamServerInterceptor(logger log.FieldLogger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		var cast_srv grpc_auth.ServiceAuthFuncOverride
		var new_ctx context.Context
		var ok bool

		var startat time.Time
		inner_logger := logger.WithFields(log.Fields{
			"method": info.FullMethod,
		})
		defer func(callat time.Time) {
			inner_logger.WithField("elapsed", time.Since(callat)).Debugf("call tracing")
		}(time.Now())

		if cast_srv, ok = srv.(grpc_auth.ServiceAuthFuncOverride); !ok {
			startat = time.Now()
			err = handler(srv, stream)
			inner_logger.WithField("handle_elapsed", time.Since(startat))
			return
		}

		startat = time.Now()
		if new_ctx, err = cast_srv.AuthFuncOverride(stream.Context(), info.FullMethod); err != nil {
			return err
		}
		inner_logger.WithField("validate_token_elapsed", time.Since(startat))

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = new_ctx

		md, _ := ParseMethodDescription(info.FullMethod)
		srv_val := reflect.ValueOf(srv)

		vlt_func := srv_val.MethodByName("Validate" + md.Method)
		if vlt_func.Kind() == reflect.Func {
			startat = time.Now()
			if err = vlt_func.Interface().(func(grpc.ServerStream) error)(wrapped); err != nil {
				return err
			}
			inner_logger.WithField("validate_request_elapsed", time.Since(startat))
		}

		auth_func := srv_val.MethodByName("Authorize" + md.Method)
		if auth_func.Kind() == reflect.Func {
			startat = time.Now()
			if err = auth_func.Interface().(func(grpc.ServerStream) error)(wrapped); err != nil {
				return err
			}
			inner_logger.WithField("authorize_request_elapsed", time.Since(startat))
		}

		startat = time.Now()
		err = handler(srv, wrapped)
		inner_logger.WithField("handle_elapsed", time.Since(startat))
		return
	}
}
