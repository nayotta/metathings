package driver_helper

import (
	"plugin"

	"github.com/spf13/viper"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Driver interface {
	Init(opt_helper.Option) error
	Close() error
}

type Descriptor struct {
	Name string
	Path string
}

type NewDriverMethod func(opt_helper.Option) (Driver, error)

type DriverFactory struct {
	descriptors map[string]Descriptor
	methods     map[string]NewDriverMethod
}

func loadDescriptor(path string) (map[string]Descriptor, error) {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var ds struct {
		Descriptors []Descriptor
	}
	err = v.Unmarshal(&ds)
	if err != nil {
		return nil, err
	}

	descriptors := map[string]Descriptor{}
	for _, d := range ds.Descriptors {
		descriptors[d.Name] = d
	}

	return descriptors, nil
}

func NewDriverFactory(path string) (*DriverFactory, error) {
	ds, err := loadDescriptor(path)
	if err != nil {
		return nil, err
	}

	return &DriverFactory{
		descriptors: ds,
		methods:     map[string]NewDriverMethod{},
	}, nil
}

func (df *DriverFactory) New(name string, opt opt_helper.Option) (Driver, error) {
	method, err := df.getNewDriverMethod(name)
	if err != nil {
		return nil, err
	}

	drv, err := method(opt)
	if err != nil {
		return nil, err
	}

	return drv, nil
}

func (df *DriverFactory) getNewDriverMethod(name string) (NewDriverMethod, error) {
	method, ok := df.methods[name]
	if ok {
		return method, nil
	}

	ds, ok := df.descriptors[name]
	if !ok {
		return nil, ErrDriverNotFound
	}

	sym, err := plugin.Open(ds.Path)
	if err != nil {
		return nil, err
	}

	fn, err := sym.Lookup("NewDriver")
	if err != nil {
		return nil, err
	}

	method = *fn.(*NewDriverMethod)
	df.methods[name] = method

	return method, nil
}
