package grpc_helper

import (
	"context"
	"reflect"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var cast_srv grpc_auth.ServiceAuthFuncOverride
		var new_ctx context.Context
		var err error
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

		return handler(new_ctx, req)
	}
}
