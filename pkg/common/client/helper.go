package client_helper

import (
	"crypto/tls"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	component_pb "github.com/nayotta/metathings/proto/component"
	device_pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
	evaluatord_pb "github.com/nayotta/metathings/proto/evaluatord"
	identityd2_pb "github.com/nayotta/metathings/proto/identityd2"
	policyd_pb "github.com/nayotta/metathings/proto/policyd"
)

type ClientType int32

const (
	DEFAULT_CONFIG ClientType = iota
	POLICYD_CONFIG
	IDENTITYD2_CONFIG
	DEVICED_CONFIG
	EVALUATORD_CONFIG
	DEVICE_CONFIG
	MODULE_CONFIG
	OVERFLOW_CONFIG
)

var (
	client_type_names = []string{
		"default",
		"policyd",
		"identityd2",
		"deviced",
		"evaluatord",
		"device",
		"module",
		"overflow",
	}
)

func (self ClientType) String() string {
	return client_type_names[self]
}

func parseAddress(addr string) string {
	if !strings.Contains(addr, ":") {
		addr = fmt.Sprintf("%v:%v", addr, constant_helper.CONSTANT_METATHINGSD_DEFAULT_PORT)
	}
	return addr
}

type CloseFn func() error

type ServiceConfigs map[ClientType]ServiceConfig

func (self ServiceConfigs) SetServiceConfig(typ ClientType, cfg ServiceConfig) {
	self[typ] = cfg
}

type ServiceConfig struct {
	Address              string
	TransportCredentials credentials.TransportCredentials
}

type ClientFactory struct {
	defaultDialOption []grpc.DialOption
	configs           ServiceConfigs
}

func (f *ClientFactory) NewConnection(cfg_val ClientType, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, f.defaultDialOption...)

	cfg, ok := f.configs[cfg_val]
	if !ok {
		cfg = f.configs[DEFAULT_CONFIG]
	}

	if cfg.TransportCredentials != nil {
		opts = append(opts, grpc.WithTransportCredentials(cfg.TransportCredentials))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(parseAddress(cfg.Address), opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (f *ClientFactory) NewPolicydServiceClient(opts ...grpc.DialOption) (policyd_pb.PolicydServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(POLICYD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return policyd_pb.NewPolicydServiceClient(conn), conn.Close, nil
}

func (f *ClientFactory) NewIdentityd2ServiceClient(opts ...grpc.DialOption) (identityd2_pb.IdentitydServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(IDENTITYD2_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return identityd2_pb.NewIdentitydServiceClient(conn), conn.Close, nil
}

func (f *ClientFactory) NewDevicedServiceClient(opts ...grpc.DialOption) (deviced_pb.DevicedServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(DEVICED_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return deviced_pb.NewDevicedServiceClient(conn), conn.Close, nil
}

func (f *ClientFactory) NewEvaluatordServiceClient(opts ...grpc.DialOption) (evaluatord_pb.EvaluatordServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(EVALUATORD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return evaluatord_pb.NewEvaluatordServiceClient(conn), conn.Close, nil
}

func (f *ClientFactory) NewDeviceServiceClient(opts ...grpc.DialOption) (device_pb.DeviceServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(DEVICE_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return device_pb.NewDeviceServiceClient(conn), conn.Close, nil
}

func (f *ClientFactory) NewModuleServiceClient(opts ...grpc.DialOption) (component_pb.ModuleServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(MODULE_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return component_pb.NewModuleServiceClient(conn), conn.Close, nil
}

func NewClientFactory(configs ServiceConfigs, opts []grpc.DialOption) (*ClientFactory, error) {
	if _, ok := configs[DEFAULT_CONFIG]; !ok {
		return nil, ErrMissingDefaultConfig
	}

	return &ClientFactory{
		configs:           configs,
		defaultDialOption: opts,
	}, nil
}

func NewDefaultServiceConfigs(addr string, cred credentials.TransportCredentials) ServiceConfigs {
	return ServiceConfigs{
		DEFAULT_CONFIG: ServiceConfig{
			Address:              addr,
			TransportCredentials: cred,
		},
	}
}

func DefaultDialOption() []grpc.DialOption {
	return []grpc.DialOption{}
}

func NewClientTransportCredentials(cert_file, key_file string, plain_text, insecure bool) (credentials.TransportCredentials, error) {
	if cert_file != "" && key_file != "" {
		return credentials.NewServerTLSFromFile(cert_file, key_file)
	} else if insecure {
		return credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		}), nil
	} else if plain_text {
		return nil, nil
	}
	return credentials.NewTLS(nil), nil
}

func NewServerTransportCredentials(cert_file, key_file string) (credentials.TransportCredentials, error) {
	if cert_file != "" && key_file != "" {
		return credentials.NewServerTLSFromFile(cert_file, key_file)
	} else {
		return nil, nil
	}
}

func ToClientFactory(v **ClientFactory) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		if *v, ok = val.(*ClientFactory); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
