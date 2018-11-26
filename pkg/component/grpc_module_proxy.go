package metathings_component

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/pkg/proto/component"
)

type GrpcModuleProxy struct {
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
}

func (self *GrpcModuleProxy) UnaryCall(ctx context.Context, method string, value *any.Any) (*any.Any, error) {
	cli, cfn, err := self.cli_fty.NewModuleSerivceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	req := &pb.UnaryCallRequest{
		Method: &wrappers.StringValue{Value: method},
		Value:  value,
	}
	res, err := cli.UnaryCall(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (self *GrpcModuleProxy) StreamCall(ctx context.Context, method string) (ModuleProxyStream, error) {
	panic("unimplemented")
}

type GrpcModuleProxyFactory struct{}

func (self *GrpcModuleProxyFactory) NewModuleProxy(args ...interface{}) (ModuleProxy, error) {
	p := &GrpcModuleProxy{}

	if err := opt_helper.Setopt(map[string]func(key string, val interface{}) error{
		"logger": func(key string, val interface{}) error {
			var ok bool
			if p.logger, ok = val.(log.FieldLogger); !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
		"client_factory": func(key string, val interface{}) error {
			var ok bool
			if p.cli_fty, ok = val.(*client_helper.ClientFactory); !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args); err != nil {
		return nil, err
	}

	return p, nil
}

func init() {
	register_module_proxy_factory("grpc", new(GrpcModuleProxyFactory))
}
