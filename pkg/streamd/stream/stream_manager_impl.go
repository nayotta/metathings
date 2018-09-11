package stream_manager

import (
	"fmt"

	pb "github.com/nayotta/metathings/pkg/proto/streamd"
)

type upstreamOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type sourceOption struct {
	id       string
	upstream *upstreamOption
}

type inputOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type outputOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type groupOption struct {
	id      string
	inputs  []*inputOption
	outputs []*outputOption
}

type streamOption struct {
	name    string
	sources []*sourceOption
	groups  []*groupOption
}

func SetNewStreamOption(req *pb.CreateRequest) NewStreamOption {
	configToMap := func(x map[string]*pb.ConfigValue) map[string]string {
		y := map[string]string{}

		for k, v := range x {
			switch v.GetValue().(type) {
			case *pb.ConfigValue_Double:
				y[k] = fmt.Sprintf("%v", v.GetDouble())
			case *pb.ConfigValue_Int64:
				y[k] = fmt.Sprintf("%v", v.GetInt64())
			case *pb.ConfigValue_Uint64:
				y[k] = fmt.Sprintf("%v", v.GetUint64())
			case *pb.ConfigValue_String_:
				y[k] = v.GetString_()
			}
		}

		return y
	}

	newUpstreamOption := func(x *pb.OpUpstream) *upstreamOption {
		return &upstreamOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newSourceOption := func(x *pb.OpSource) *sourceOption {
		return &sourceOption{
			id:       x.GetId().GetValue(),
			upstream: newUpstreamOption(x.GetUpstream()),
		}
	}

	newInputOption := func(x *pb.OpInput) *inputOption {
		return &inputOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newOutputOption := func(x *pb.OpOutput) *outputOption {
		return &outputOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newGroupOption := func(x *pb.OpGroup) *groupOption {
		y := &groupOption{
			id:      x.GetId().GetValue(),
			inputs:  []*inputOption{},
			outputs: []*outputOption{},
		}

		for _, input := range x.GetInputs() {
			y.inputs = append(y.inputs, newInputOption(input))
		}

		for _, output := range x.GetOutputs() {
			y.outputs = append(y.outputs, newOutputOption(output))
		}

		return y
	}

	setSources := func(o *streamOption) {
		o.sources = []*sourceOption{}

		for _, source := range req.GetSources() {
			o.sources = append(o.sources, newSourceOption(source))
		}
	}

	setGroups := func(o *streamOption) {
		o.groups = []*groupOption{}

		for _, group := range req.GetGroups() {
			o.groups = append(o.groups, newGroupOption(group))
		}
	}

	return func(x interface{}) {
		o := x.(*streamOption)
		o.name = req.GetName().GetValue()
		setSources(o)
		setGroups(o)
	}
}

type streamManagerImpl struct{}

func (self *streamManagerImpl) NewStream(opts ...NewStreamOption) (Stream, error) {
	panic("unimplemented")
}

func (self *streamManagerImpl) GetStream(id string) (Stream, error) {
	panic("unimplemented")
}

func newStreamManagerImpl(opts ...StreamManagerOption) (StreamManager, error) {
	panic("unimplemented")
}

func init() {
	RegisterStreamManager("default", newStreamManagerImpl)
}
