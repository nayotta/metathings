package stream_manager

import (
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type UpstreamOption struct {
	Id     string
	Name   string
	Alias  string
	Config map[string]string
}

type SourceOption struct {
	Id       string
	Upstream *UpstreamOption
}

type InputOption struct {
	Id     string
	Name   string
	Alias  string
	Config map[string]string
}

type OutputOption struct {
	Id     string
	Name   string
	Alias  string
	Config map[string]string
}

type GroupOption struct {
	Id      string
	Inputs  []*InputOption
	Outputs []*OutputOption
}

type StreamOption struct {
	Id      string
	Name    string
	Sources []*SourceOption
	Groups  []*GroupOption
}

type streamManagerImplOption struct {
	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
	brokers      []string
}

type streamManagerImpl struct {
	logger  log.FieldLogger
	opt     *streamManagerImplOption
	streams map[string]Stream
}

func (self *streamManagerImpl) NewStream(opt *StreamOption, extra map[string]interface{}) (Stream, error) {
	fty := NewDefaultStreamFactory()
	stm, err := fty.Set("option", opt).
		Set("application_credential", extra["application_credential"]).
		Set("client_factory", extra["client_factory"]).
		Set("logger", extra["logger"]).
		Set("brokers", extra["brokers"]).
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
	case "brokers":
		self.opt.brokers = val.([]string)
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
