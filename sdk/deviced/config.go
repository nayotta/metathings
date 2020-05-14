package metathings_deviced_sdk

import (
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/objx"

	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Config struct {
	m objx.Map
	*pb.Config
}

func (c *Config) Get(selector string) *objx.Value {
	return c.m.Get(selector)
}

func (c *Config) Set(selector string, value interface{}) {
	c.m.Set(selector, value)
}

func FromConfig(x *pb.Config) (*Config, error) {
	js_str, err := new(jsonpb.Marshaler).MarshalToString(x.GetBody())
	if err != nil {
		return nil, err
	}

	sm := map[string]interface{}{}
	err = json.Unmarshal([]byte(js_str), &sm)
	if err != nil {
		return nil, err
	}

	m := objx.New(sm)

	return &Config{
		m:      m,
		Config: x,
	}, nil
}

func ToProtobuf(x *Config) (*pb.OpConfig, error) {
	js_str, err := x.m.JSON()
	if err != nil {
		return nil, err
	}

	var body stpb.Struct
	err = jsonpb.UnmarshalString(js_str, &body)
	if err != nil {
		return nil, err
	}

	return &pb.OpConfig{
		Id:    &wrappers.StringValue{Value: x.GetId()},
		Alias: &wrappers.StringValue{Value: x.GetAlias()},
		Body:  &body,
	}, nil
}

func LookupConfig(cfgs []*pb.Config, alias string) (*Config, error) {
	for _, cfg := range cfgs {
		if cfg.GetAlias() == alias {
			return FromConfig(cfg)
		}
	}

	return nil, ErrConfigNotFound
}
