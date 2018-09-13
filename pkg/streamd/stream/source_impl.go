package stream_manager

import (
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type sourceImplOption struct {
	id       string
	upstream *UpstreamOption

	sym_tbl      SymbolTable
	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
}

type sourceImpl struct {
	logger   log.FieldLogger
	opt      *sourceImplOption
	upstream Upstream
}

func (self *sourceImpl) Upstream() Upstream {
	return self.upstream
}

func (self *sourceImpl) Id() string {
	return self.opt.id
}

type sourceImplFactory struct {
	opt *sourceImplOption
}

func (self *sourceImplFactory) Set(key string, val interface{}) SourceFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "application_credential_manager":
		self.opt.app_cred_mgr = val.(app_cred_mgr.ApplicationCredentialManager)
	case "client_factory":
		self.opt.cli_fty = val.(*client_helper.ClientFactory)
	case "symbol_table":
		self.opt.sym_tbl = val.(SymbolTable)
	case "option":
		opt := val.(*SourceOption)
		self.opt.id = opt.id
		self.opt.upstream = opt.upstream
	}

	return self
}

func (self *sourceImplFactory) New() (Source, error) {
	fty, err := NewUpstreamFactory(self.opt.upstream.name)
	if err != nil {
		return nil, err
	}

	upstream, err := fty.Set("logger", self.opt.logger).
		Set("application_credential_manager", self.opt.app_cred_mgr).
		Set("client_factory", self.opt.cli_fty).
		Set("symbol_table", self.opt.sym_tbl).
		Set("option", self.opt.upstream).
		New()
	if err != nil {
		return nil, err
	}

	src := &sourceImpl{
		opt:      self.opt,
		upstream: upstream,
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "source:default",
		}),
	}

	return src, nil
}

func init() {
	RegisterSourceFactory("default", func() SourceFactory { return &sourceImplFactory{opt: &sourceImplOption{}} })
}
