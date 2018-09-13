package stream_manager

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
)

type UpstreamOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type SourceOption struct {
	id       string
	upstream *UpstreamOption
}

type InputOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type OutputOption struct {
	id     string
	name   string
	alias  string
	config map[string]string
}

type GroupOption struct {
	id      string
	inputs  []*InputOption
	outputs []*OutputOption
}

type StreamOption struct {
	id      string
	name    string
	sources []*SourceOption
	groups  []*GroupOption
}

func PbCreateRequestToStreamOption(req *pb.CreateRequest) *StreamOption {
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

	newUpstreamOption := func(x *pb.OpUpstream) *UpstreamOption {
		return &UpstreamOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newSourceOption := func(x *pb.OpSource) *SourceOption {
		return &SourceOption{
			id:       x.GetId().GetValue(),
			upstream: newUpstreamOption(x.GetUpstream()),
		}
	}

	newInputOption := func(x *pb.OpInput) *InputOption {
		return &InputOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newOutputOption := func(x *pb.OpOutput) *OutputOption {
		return &OutputOption{
			id:     x.GetId().GetValue(),
			name:   x.GetName().GetValue(),
			alias:  x.GetAlias().GetValue(),
			config: configToMap(x.GetConfig()),
		}
	}

	newGroupOption := func(x *pb.OpGroup) *GroupOption {
		y := &GroupOption{
			id:      x.GetId().GetValue(),
			inputs:  []*InputOption{},
			outputs: []*OutputOption{},
		}

		for _, input := range x.GetInputs() {
			y.inputs = append(y.inputs, newInputOption(input))
		}

		for _, output := range x.GetOutputs() {
			y.outputs = append(y.outputs, newOutputOption(output))
		}

		return y
	}

	newSources := func(x []*pb.OpSource) []*SourceOption {
		sources := []*SourceOption{}

		for _, source := range req.GetSources() {
			sources = append(sources, newSourceOption(source))
		}

		return sources
	}

	newGroups := func(x []*pb.OpGroup) []*GroupOption {
		groups := []*GroupOption{}

		for _, group := range req.GetGroups() {
			groups = append(groups, newGroupOption(group))
		}

		return groups
	}

	opt := &StreamOption{
		id:      req.GetId().GetValue(),
		name:    req.GetName().GetValue(),
		sources: newSources(req.GetSources()),
		groups:  newGroups(req.GetGroups()),
	}

	return opt
}

type streamManagerImplOption struct {
	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
}

type streamManagerImpl struct {
	logger  log.FieldLogger
	opt     *streamManagerImplOption
	streams map[string]Stream
}

func (self *streamManagerImpl) NewStream(opt StreamOption, extra map[string]interface{}) (Stream, error) {
	fty := NewDefaultStreamFactory()
	stm, err := fty.Set("option", opt).
		Set("application_credential", extra["application_credential"]).
		Set("client_factory", extra["client_factory"]).
		Set("logger", extra["logger"]).
		New()
	if err != nil {
		self.logger.WithError(err).Debugf("failed to new stream")
		return nil, err
	}

	self.streams[stm.Id()] = stm

	return stm, nil
}

func (self *streamManagerImpl) GetStream(id string) (Stream, error) {
	stm, ok := self.streams[id]
	if !ok {
		return nil, ErrStreamNotFound
	}
	return stm, nil
}

type streamManagerImplFactory struct {
	opt *streamManagerImplOption
}

func (self *streamManagerImplFactory) Set(key string, val interface{}) StreamManagerFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "application_credential_manager":
		self.opt.app_cred_mgr = val.(app_cred_mgr.ApplicationCredentialManager)
	case "client_factory":
		self.opt.cli_fty = val.(*client_helper.ClientFactory)
	}

	return self
}

func (self *streamManagerImplFactory) New() (StreamManager, error) {
	return &streamManagerImpl{
		opt: self.opt,
		logger: self.opt.logger.WithFields(log.Fields{
			"#component": "stream_manager:default",
		}),
		streams: map[string]Stream{},
	}, nil
}

func init() {
	RegisterStreamManagerFactory("default", func() StreamManagerFactory { return &streamManagerImplFactory{opt: &streamManagerImplOption{}} })
}
