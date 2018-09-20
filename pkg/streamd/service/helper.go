package metathings_streamd_service

import (
	"encoding/json"
	"fmt"
	"strconv"

	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	storage "github.com/nayotta/metathings/pkg/streamd/storage"
	stream_manager "github.com/nayotta/metathings/pkg/streamd/stream"
)

type storageStreamToStreamOptionCodec struct{}

func (self *storageStreamToStreamOptionCodec) configToMap(cfg string) map[string]string {
	var m map[string]string
	err := json.Unmarshal([]byte(cfg), &m)
	if err != nil {
		return map[string]string{}
	}

	delete(m, "__metadata__")

	return m
}

func (self *storageStreamToStreamOptionCodec) newUpstream(x storage.Upstream) *stream_manager.UpstreamOption {
	return &stream_manager.UpstreamOption{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: self.configToMap(*x.Config),
	}
}

func (self *storageStreamToStreamOptionCodec) newSource(x storage.Source) *stream_manager.SourceOption {
	return &stream_manager.SourceOption{
		Id:       *x.Id,
		Upstream: self.newUpstream(x.Upstream),
	}

}

func (self *storageStreamToStreamOptionCodec) newInput(x storage.Input) *stream_manager.InputOption {
	return &stream_manager.InputOption{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: self.configToMap(*x.Config),
	}
}

func (self *storageStreamToStreamOptionCodec) newOutput(x storage.Output) *stream_manager.OutputOption {
	return &stream_manager.OutputOption{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: self.configToMap(*x.Config),
	}
}

func (self *storageStreamToStreamOptionCodec) newGroup(x storage.Group) *stream_manager.GroupOption {
	return &stream_manager.GroupOption{
		Id:      *x.Id,
		Inputs:  self.newInputs(x.Inputs),
		Outputs: self.newOutputs(x.Outputs),
	}
}

func (self *storageStreamToStreamOptionCodec) newSources(xs []storage.Source) []*stream_manager.SourceOption {
	var ys []*stream_manager.SourceOption
	for _, x := range xs {
		ys = append(ys, self.newSource(x))
	}
	return ys
}

func (self *storageStreamToStreamOptionCodec) newGroups(xs []storage.Group) []*stream_manager.GroupOption {
	var ys []*stream_manager.GroupOption
	for _, x := range xs {
		ys = append(ys, self.newGroup(x))
	}
	return ys
}

func (self *storageStreamToStreamOptionCodec) newInputs(xs []storage.Input) []*stream_manager.InputOption {
	var ys []*stream_manager.InputOption
	for _, x := range xs {
		ys = append(ys, self.newInput(x))
	}
	return ys
}

func (self *storageStreamToStreamOptionCodec) newOutputs(xs []storage.Output) []*stream_manager.OutputOption {
	var ys []*stream_manager.OutputOption
	for _, x := range xs {
		ys = append(ys, self.newOutput(x))
	}
	return ys
}

func (self *storageStreamToStreamOptionCodec) Encode(x storage.Stream) *stream_manager.StreamOption {
	return &stream_manager.StreamOption{
		Id:      *x.Id,
		Name:    *x.Name,
		Sources: self.newSources(x.Sources),
		Groups:  self.newGroups(x.Groups),
	}
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
