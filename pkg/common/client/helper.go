package client_helper

import (
	"crypto/tls"
	"fmt"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
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

type DoneFn = func() error

type ServiceConfigs map[ClientType]ServiceConfig

func (self ServiceConfigs) SetServiceConfig(typ ClientType, cfg ServiceConfig) {
	self[typ] = cfg
}

type ServiceConfig struct {
	Address              string
	TransportCredentials credentials.TransportCredentials
}

type ClientFactory struct {
	opts              *newClientFactoryOption
	defaultDialOption []grpc.DialOption
	configs           ServiceConfigs
	pools_mtx         sync.Mutex
	pools             map[ClientType]pool_helper.Pool
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

func (f *ClientFactory) GetConnection(cfg_val ClientType, opts ...grpc.DialOption) (*grpc.ClientConn, DoneFn, error) {
	var err error

	if f.opts.DialPoolSize <= 0 || len(opts) != 0 {
		conn, err := f.NewConnection(cfg_val, opts...)
		if err != nil {
			return nil, nil, err
		}

		return conn, conn.Close, nil
	}

	f.pools_mtx.Lock()
	p, ok := f.pools[cfg_val]
	if !ok {
		if p, err = pool_helper.NewPool(1, int(f.opts.DialPoolSize), func() (pool_helper.Client, error) {
			return f.NewConnection(cfg_val)
		}); err != nil {
			return nil, nil, err
		}
		f.pools[cfg_val] = p
	}
	f.pools_mtx.Unlock()

	conn, err := p.Get()
	if err != nil {
		return nil, nil, err
	}
	grpcConn := conn.(*grpc.ClientConn)
	for grpcConn.GetState() == connectivity.TransientFailure ||
		grpcConn.GetState() == connectivity.Shutdown {
		conn, err = p.Get()
		if err != nil {
			return nil, nil, err
		}
		grpcConn = conn.(*grpc.ClientConn)
	}

	return grpcConn, func() error { return p.Put(conn) }, nil
}

func (f *ClientFactory) NewPolicydServiceClient(opts ...grpc.DialOption) (policyd_pb.PolicydServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(POLICYD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return policyd_pb.NewPolicydServiceClient(conn), done, nil
}

func (f *ClientFactory) NewIdentityd2ServiceClient(opts ...grpc.DialOption) (identityd2_pb.IdentitydServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(IDENTITYD2_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return identityd2_pb.NewIdentitydServiceClient(conn), done, nil
}

func (f *ClientFactory) NewDevicedServiceClient(opts ...grpc.DialOption) (deviced_pb.DevicedServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(DEVICED_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return deviced_pb.NewDevicedServiceClient(conn), done, nil
}

func (f *ClientFactory) NewEvaluatordServiceClient(opts ...grpc.DialOption) (evaluatord_pb.EvaluatordServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(EVALUATORD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return evaluatord_pb.NewEvaluatordServiceClient(conn), done, nil
}

func (f *ClientFactory) NewDeviceServiceClient(opts ...grpc.DialOption) (device_pb.DeviceServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(DEVICE_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return device_pb.NewDeviceServiceClient(conn), done, nil
}

func (f *ClientFactory) NewModuleServiceClient(opts ...grpc.DialOption) (component_pb.ModuleServiceClient, DoneFn, error) {
	conn, done, err := f.GetConnection(MODULE_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	return component_pb.NewModuleServiceClient(conn), done, nil
}

type newClientFactoryOption struct {
	DialPoolSize int32
}

func newNewClientFactoryOption() *newClientFactoryOption {
	return &newClientFactoryOption{
		DialPoolSize: 0,
	}
}

type NewClientFactoryOption func(*newClientFactoryOption)

func SetDialPoolSize(siz int32) NewClientFactoryOption {
	return func(o *newClientFactoryOption) {
		o.DialPoolSize = siz
	}
}

func NewClientFactory(configs ServiceConfigs, dial_opts []grpc.DialOption, opts ...NewClientFactoryOption) (*ClientFactory, error) {
	o := newNewClientFactoryOption()

	for _, opt := range opts {
		opt(o)
	}

	if _, ok := configs[DEFAULT_CONFIG]; !ok {
		return nil, ErrMissingDefaultConfig
	}

	return &ClientFactory{
		opts:              o,
		configs:           configs,
		defaultDialOption: dial_opts,
		pools:             make(map[ClientType]pool_helper.Pool),
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
