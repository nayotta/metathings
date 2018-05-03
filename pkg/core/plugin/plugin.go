package metathings_core_plugin

import (
	"context"
	"errors"
	"plugin"

	"github.com/spf13/viper"
)

const METATHINGS_PLUGIN_PREFIX = "mtp"

type Option struct {
	Args []string
}

type ServicePlugin interface {
	Init(opt Option) error
	Run() error
}

type DispatcherPlugin interface {
	Dispatch(context.Context, interface{}) (interface{}, error)
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
	ErrNotFound = errors.New("plugin not found")
)
