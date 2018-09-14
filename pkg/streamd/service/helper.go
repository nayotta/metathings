package metathings_streamd_service

import (
	"encoding/json"
	"fmt"
	"strconv"

	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	stream_manager "github.com/nayotta/metathings/pkg/streamd/stream"
)

func encode_create_request_to_stream_option(req *pb.CreateRequest) *stream_manager.StreamOption {
	configToMap := func(x map[string]*pb.ConfigValue) map[string]string {
		y := map[string]string{}

		for k, v := range x {
			switch v.GetValue().(type) {
			case *pb.ConfigValue_Double:
				y[k] = fmt.Sprintf("%v", v.GetDouble())
			case *pb.ConfigValue_String_:
				y[k] = v.GetString_()
			}
		}

		return y
	}

	newUpstreamOption := func(x *pb.OpUpstream) *stream_manager.UpstreamOption {
		return &stream_manager.UpstreamOption{
			Id:     x.GetId().GetValue(),
			Name:   x.GetName().GetValue(),
			Alias:  x.GetAlias().GetValue(),
			Config: configToMap(x.GetConfig()),
		}
	}

	newSourceOption := func(x *pb.OpSource) *stream_manager.SourceOption {
		return &stream_manager.SourceOption{
			Id:       x.GetId().GetValue(),
			Upstream: newUpstreamOption(x.GetUpstream()),
		}
	}

	newInputOption := func(x *pb.OpInput) *stream_manager.InputOption {
		return &stream_manager.InputOption{
			Id:     x.GetId().GetValue(),
			Name:   x.GetName().GetValue(),
			Alias:  x.GetAlias().GetValue(),
			Config: configToMap(x.GetConfig()),
		}
	}

	newOutputOption := func(x *pb.OpOutput) *stream_manager.OutputOption {
		return &stream_manager.OutputOption{
			Id:     x.GetId().GetValue(),
			Name:   x.GetName().GetValue(),
			Alias:  x.GetAlias().GetValue(),
			Config: configToMap(x.GetConfig()),
		}
	}

	newGroupOption := func(x *pb.OpGroup) *stream_manager.GroupOption {
		y := &stream_manager.GroupOption{
			Id:      x.GetId().GetValue(),
			Inputs:  []*stream_manager.InputOption{},
			Outputs: []*stream_manager.OutputOption{},
		}

		for _, input := range x.GetInputs() {
			y.Inputs = append(y.Inputs, newInputOption(input))
		}

		for _, output := range x.GetOutputs() {
			y.Outputs = append(y.Outputs, newOutputOption(output))
		}

		return y
	}

	newSources := func(x []*pb.OpSource) []*stream_manager.SourceOption {
		sources := []*stream_manager.SourceOption{}

		for _, source := range req.GetSources() {
			sources = append(sources, newSourceOption(source))
		}

		return sources
	}

	newGroups := func(x []*pb.OpGroup) []*stream_manager.GroupOption {
		groups := []*stream_manager.GroupOption{}

		for _, group := range req.GetGroups() {
			groups = append(groups, newGroupOption(group))
		}

		return groups
	}

	opt := &stream_manager.StreamOption{
		Id:      req.GetId().GetValue(),
		Name:    req.GetName().GetValue(),
		Sources: newSources(req.GetSources()),
		Groups:  newGroups(req.GetGroups()),
	}

	return opt
}

func encode_config_to_json_string(x map[string]*pb.ConfigValue) (string, error) {
	y := map[string]string{}
	y_meta := map[string]string{}

	for k, v := range x {
		switch v.GetValue().(type) {
		case *pb.ConfigValue_Double:
			y[k] = fmt.Sprintf("%v", v.GetDouble())
			y_meta[k] = "double"
		case *pb.ConfigValue_String_:
			y[k] = fmt.Sprintf("%v", v.GetString_())
			y_meta[k] = "string"
		}
	}

	buf, err := json.Marshal(y_meta)
	if err != nil {
		return "", err
	}
	y["__metadata__"] = string(buf)

	buf, err = json.Marshal(y)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func decode_json_string_to_config(x string) (map[string]*pb.ConfigValue, error) {
	y := map[string]*pb.ConfigValue{}
	m := map[string]string{}
	err := json.Unmarshal([]byte(x), &m)
	if err != nil {
		return nil, err
	}

	m_meta_str, ok := m["__metadata__"]
	m_meta := map[string]string{}
	err = json.Unmarshal([]byte(m_meta_str), &m_meta)
	if err != nil {
		return nil, err
	}
	delete(m, "__metadata__")
	for k, v := range m {
		if !ok || m_meta[k] == "string" {
			y[k] = &pb.ConfigValue{Value: &pb.ConfigValue_String_{String_: v}}
		} else {
			d, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			y[k] = &pb.ConfigValue{Value: &pb.ConfigValue_Double{Double: d}}
		}
	}

	return y, nil
}

func must_decode_json_string_to_config(x string) map[string]*pb.ConfigValue {
	y, err := decode_json_string_to_config(x)
	if err != nil {
		return map[string]*pb.ConfigValue{}
	}
	return y
}
