package metathings_core_plugin

import (
	"context"
	"errors"
	"plugin"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	agentd_pb "github.com/nayotta/metathings/pkg/proto/core_agent"
)

const METATHINGS_PLUGIN_PREFIX = "mtp"

type CoreService struct {
	opts    opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func MakeCoreService(opts opt_helper.Option, logger log.FieldLogger, cli_fty *client_helper.ClientFactory) CoreService {
	return CoreService{
		opts:    opts,
		logger:  logger,
		cli_fty: cli_fty,
	}
}

func (s CoreService) Init() error {
	req := &agentd_pb.CreateOrGetEntityRequest{
		Name:        &gpb.StringValue{Value: s.opts.GetString("name")},
		ServiceName: &gpb.StringValue{Value: s.opts.GetString("service_name")},
		Endpoint:    &gpb.StringValue{Value: s.opts.GetString("endpoint")},
	}

	ctx := context.Background()
	cli, cfn, err := s.cli_fty.NewCoreAgentServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	res, err := cli.CreateOrGetEntity(ctx, req)
	if err != nil {
		return err
	}

	s.opts.Set("id", res.Entity.Id)

	if !s.opts.GetBool("heartbeat.manual") {
		go s.Heartbeat()
	}

	return nil
}

func (s CoreService) Heartbeat() error {
	for {
		err := s.HeartbeatOnce()
		if err != nil {
			return err
		}
		<-time.After(time.Duration(s.opts.GetInt("heartbeat.interval")) * time.Second)
	}
}

func (s CoreService) HeartbeatOnce() error {
	ctx := context.Background()
	cli, cfn, err := s.cli_fty.NewCoreAgentServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	req := &agentd_pb.HeartbeatRequest{
		EntityId: &gpb.StringValue{Value: s.opts.GetString("id")},
	}

	_, err = cli.Heartbeat(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

type ServicePlugin interface {
	Init(opts opt_helper.Option) error
	Run() error
}

type DispatcherPlugin interface {
	Init(opts opt_helper.Option) error
	UnaryCall(method string, ctx context.Context, req *any.Any) (*any.Any, error)
}

type PluginCommandOptions struct {
	Name        string
	ServiceName string `mapstructure:"service_name"`
}

type PluginDescriptorType int32

const (
	PD_UNKNOWN PluginDescriptorType = iota
	PD_SERVICE
	PD_DISPATCHER
)

type PluginDescriptor struct {
	Type PluginDescriptorType
	Path string
}

type Descriptor struct {
	Name    string
	Plugins map[PluginDescriptorType]PluginDescriptor
}

type ServiceDescriptor struct {
	Version     string
	Descriptors map[string]Descriptor
}

func newServiceDescriptor() *ServiceDescriptor {
	return &ServiceDescriptor{
		Descriptors: make(map[string]Descriptor),
	}
}

func LoadServiceDescriptor(path string) (*ServiceDescriptor, error) {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var sd struct {
		Version  string
		Services []struct {
			Name    string
			Plugins struct {
				Service struct {
					Path string
				}
				Dispatcher struct {
					Path string
				}
			}
		}
	}
	err = v.Unmarshal(&sd)
	if err != nil {
		return nil, err
	}

	serv_desc := newServiceDescriptor()
	serv_desc.Version = sd.Version
	for _, s := range sd.Services {
		d := Descriptor{
			Name:    s.Name,
			Plugins: make(map[PluginDescriptorType]PluginDescriptor),
		}
		d.Plugins[PD_SERVICE] = PluginDescriptor{
			Type: PD_SERVICE,
			Path: s.Plugins.Service.Path,
		}
		d.Plugins[PD_DISPATCHER] = PluginDescriptor{
			Type: PD_DISPATCHER,
			Path: s.Plugins.Dispatcher.Path,
		}
		serv_desc.Descriptors[s.Name] = d
	}

	return serv_desc, nil
}

func (sd *ServiceDescriptor) GetServicePlugin(name string) (ServicePlugin, error) {
	d, ok := sd.Descriptors[name]
	if !ok {
		return nil, ErrNotFound
	}

	sp, ok := d.Plugins[PD_SERVICE]
	if !ok {
		return nil, ErrNotFound
	}

	lib, err := plugin.Open(sp.Path)
	if err != nil {
		return nil, err
	}

	fn, err := lib.Lookup("NewServicePlugin")
	if err != nil {
		return nil, err
	}

	return fn.(func() ServicePlugin)(), nil
}

func (sd *ServiceDescriptor) GetDispatcherPlugin(name string) (DispatcherPlugin, error) {
	d, ok := sd.Descriptors[name]
	if !ok {
		return nil, ErrNotFound
	}

	dp, ok := d.Plugins[PD_DISPATCHER]
	if !ok {
		return nil, ErrNotFound
	}

	lib, err := plugin.Open(dp.Path)
	if err != nil {
		return nil, err
	}

	fn, err := lib.Lookup("NewDispatcherPlugin")
	if err != nil {
		return nil, err
	}

	return fn.(func() DispatcherPlugin)(), nil
}

var (
	ErrNotFound           = errors.New("plugin not found")
	ErrUnknownClient      = errors.New("unknown client")
	ErrUnknownRequestType = errors.New("unknown request type")
)
