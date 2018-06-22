package client_helper

import (
	"fmt"
	"strings"

	"google.golang.org/grpc"

	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	camera_pb "github.com/nayotta/metathings/pkg/proto/camera"
	camerad_pb "github.com/nayotta/metathings/pkg/proto/camerad"
	agent_pb "github.com/nayotta/metathings/pkg/proto/core_agent"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
	echo_pb "github.com/nayotta/metathings/pkg/proto/echo"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd"
	motor_pb "github.com/nayotta/metathings/pkg/proto/motor"
	switcher_pb "github.com/nayotta/metathings/pkg/proto/switcher"
)

const (
	DEFAULT_CONFIG = iota
	IDENTITYD_CONFIG
	CORED_CONFIG
	CAMERAD_CONFIG
	AGENT_CONFIG
	ECHO_CONFIG
	SWITCHER_CONFIG
	MOTOR_CONFIG
	CAMERA_CONFIG
)

func parseAddress(addr string) string {
	if !strings.Contains(addr, ":") {
		addr = fmt.Sprintf("%v:%v", addr, constant_helper.CONSTANT_METATHINGSD_DEFAULT_PORT)
	}
	return addr
}

type CloseFn func()
type ServiceConfigs map[int]ServiceConfig
type DialOptionFn func() []grpc.DialOption

type ServiceConfig struct {
	Address string
}

type ClientFactory struct {
	defaultDialOptionFn DialOptionFn
	configs             ServiceConfigs
}

func (f *ClientFactory) NewConnection(cfg_val int, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if f.defaultDialOptionFn != nil {
		opts = append(opts, f.defaultDialOptionFn()...)
	}

	cfg, ok := f.configs[cfg_val]
	if !ok {
		cfg = f.configs[DEFAULT_CONFIG]
	}

	conn, err := grpc.Dial(parseAddress(cfg.Address), opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (f *ClientFactory) NewCoredServiceClient(opts ...grpc.DialOption) (cored_pb.CoredServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CORED_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return cored_pb.NewCoredServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewIdentitydServiceClient(opts ...grpc.DialOption) (identityd_pb.IdentitydServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(IDENTITYD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return identityd_pb.NewIdentitydServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewCoreAgentServiceClient(opts ...grpc.DialOption) (agent_pb.CoreAgentServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(AGENT_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return agent_pb.NewCoreAgentServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewEchoServiceClient(opts ...grpc.DialOption) (echo_pb.EchoServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(ECHO_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return echo_pb.NewEchoServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewSwitcherServiceClient(opts ...grpc.DialOption) (switcher_pb.SwitcherServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(SWITCHER_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return switcher_pb.NewSwitcherServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewMotorServiceClient(opts ...grpc.DialOption) (motor_pb.MotorServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(MOTOR_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return motor_pb.NewMotorServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewCameraServiceClient(opts ...grpc.DialOption) (camera_pb.CameraServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CAMERA_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return camera_pb.NewCameraServiceClient(conn), closeFn, nil
}

func (f *ClientFactory) NewCameradServiceClient(opts ...grpc.DialOption) (camerad_pb.CameradServiceClient, CloseFn, error) {
	conn, err := f.NewConnection(CAMERAD_CONFIG, opts...)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		conn.Close()
	}

	return camerad_pb.NewCameradServiceClient(conn), closeFn, nil
}

func NewClientFactory(configs ServiceConfigs, optFn DialOptionFn) (*ClientFactory, error) {
	if _, ok := configs[DEFAULT_CONFIG]; !ok {
		return nil, ErrMissingDefaultConfig
	}

	return &ClientFactory{
		configs:             configs,
		defaultDialOptionFn: optFn,
	}, nil
}

func NewDefaultServiceConfigs(addr string) ServiceConfigs {
	return ServiceConfigs{
		DEFAULT_CONFIG: ServiceConfig{addr},
	}
}

func WithInsecureOptionFunc() DialOptionFn {
	return func() []grpc.DialOption {
		return []grpc.DialOption{grpc.WithInsecure()}
	}
}
