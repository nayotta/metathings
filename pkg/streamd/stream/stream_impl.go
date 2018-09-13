package stream_manager

import (
	"sync"
	"time"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log "github.com/sirupsen/logrus"
)

type streamImplOption struct {
	id      string
	name    string
	sources []*SourceOption
	groups  []*GroupOption

	sym_tbl      SymbolTable
	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
}

type streamImpl struct {
	Emitter
	opt     *streamImplOption
	state   StreamState
	sources []Source
	groups  []Group
	sym_tbl SymbolTable
	slck    *sync.Mutex
	logger  log.FieldLogger
}

func (self *streamImpl) Id() string {
	return self.opt.id
}

func (self *streamImpl) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != STREAM_STATE_STOP {
		self.logger.WithError(ErrUnstartable).Debugf("stream is unstartable")
		return ErrUnstartable
	}

	self.state = STREAM_STATE_STARTING
	go self.start()
	self.logger.Debugf("stream starting")

	return nil
}

func (self *streamImpl) start() {
	self.slck.Lock()
	defer self.slck.Unlock()

	wg := sync.WaitGroup{}

	for _, source := range self.Sources() {
		upstream := source.Upstream()
		upstream.Once(START_EVENT, func(Event, interface{}) {
			wg.Done()
		})
		wg.Add(1)
		upstream.Start()
	}

	for _, group := range self.Groups() {
		for _, input := range group.Inputs() {
			input.Once(START_EVENT, func(Event, interface{}) {
				wg.Done()
			})
			wg.Add(1)
			input.Start()
		}

		for _, output := range group.Outputs() {
			output.Once(START_EVENT, func(Event, interface{}) {
				wg.Done()
			})
			wg.Add(1)
			output.Start()
		}
	}

	go func() {
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			self.slck.Lock()
			defer self.slck.Unlock()
			self.state = STREAM_STATE_RUNNING
			self.logger.Debugf("stream is running")
		case <-time.After(10 * time.Second):
			self.slck.Lock()
			defer self.slck.Unlock()
			self.logger.Warningf("stream start timeout")
			self.force_stop()
		}
	}()
}

func (self *streamImpl) force_stop() {
	self._internal_stop(func(wg *sync.WaitGroup) {
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			self.logger.Debugf("stream force stop done")
		case <-time.After(5 * time.Second):
			self.logger.Warningf("stream force stop failed")
		}
		self.slck.Lock()
		defer self.slck.Unlock()
		self.state = STREAM_STATE_STOP
	})
}

func (self *streamImpl) _internal_stop(cb func(wg *sync.WaitGroup)) {
	self.slck.Lock()
	defer self.slck.Unlock()

	wg := sync.WaitGroup{}

	for _, source := range self.Sources() {
		upstream := source.Upstream()
		upstream.Once(STOP_EVENT, func(Event, interface{}) {
			wg.Done()
		})
		wg.Add(1)
		upstream.Stop()
	}

	for _, group := range self.Groups() {
		for _, input := range group.Inputs() {
			input.Once(STOP_EVENT, func(Event, interface{}) {
				wg.Done()
			})
			wg.Add(1)
			input.Stop()
		}

		for _, output := range group.Outputs() {
			output.Once(STOP_EVENT, func(Event, interface{}) {
				wg.Done()
			})
			wg.Add(1)
			output.Stop()
		}
	}

	go cb(&wg)
}

func (self *streamImpl) Stop() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != STREAM_STATE_RUNNING {
		self.logger.WithError(ErrUnterminable).Errorf("stream is unterminable")
		return ErrUnterminable
	}

	self.state = STREAM_STATE_TERMINATING
	go self.stop()

	return nil
}

func (self *streamImpl) stop() {
	self._internal_stop(func(wg *sync.WaitGroup) {
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			self.slck.Lock()
			defer self.slck.Unlock()
			self.state = STREAM_STATE_STOP
			self.logger.Debugf("stream is terminated")
		case <-time.After(10 * time.Second):
			self.slck.Lock()
			defer self.slck.Unlock()
			self.logger.Warningf("stream terminate timeout")
			self.force_stop()
		}
	})
}

func (self *streamImpl) State() StreamState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *streamImpl) Close() {
	panic("unimplemented")
}

func (self *streamImpl) Sources() []Source {
	return self.sources
}

func (self *streamImpl) Groups() []Group {
	return self.groups
}

type streamImplFactory struct {
	opt *streamImplOption
}

func (self *streamImplFactory) Set(key string, val interface{}) StreamFactory {
	switch key {
	case "symbol_table":
		self.opt.sym_tbl = val.(SymbolTable)
	case "application_credential_manager":
		self.opt.app_cred_mgr = val.(app_cred_mgr.ApplicationCredentialManager)
	case "client_factory":
		self.opt.cli_fty = val.(*client_helper.ClientFactory)
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "option":
		opt := val.(*StreamOption)
		self.opt.id = opt.id
		self.opt.name = opt.name
		self.opt.sources = opt.sources
		self.opt.groups = opt.groups
	}

	return self
}

func (self *streamImplFactory) New() (Stream, error) {
	sym_tbl := self.make_symbol_table_by_stream_option()

	sources := []Source{}
	for _, src_opt := range self.opt.sources {
		fty := NewDefaultSourceFactory()
		source, err := fty.Set("symbol_table", sym_tbl).
			Set("application_credential_manager", self.opt.app_cred_mgr).
			Set("client_factory", self.opt.cli_fty).
			Set("logger", self.opt.logger).
			Set("option", src_opt).
			New()
		if err != nil {
			return nil, err
		}

		sources = append(sources, source)
	}

	groups := []Group{}
	for _, grp_opt := range self.opt.groups {
		fty := NewDefaultGroupFactory()
		group, err := fty.
			Set("logger", self.opt.logger).
			Set("symbol_table", self.opt.sym_tbl).
			Set("option", grp_opt).
			New()
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	stream := &streamImpl{
		Emitter: NewEmitter(),
		opt:     self.opt,
		state:   STREAM_STATE_STOP,
		sym_tbl: sym_tbl,
		slck:    &sync.Mutex{},
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "stream:default",
		}),
		sources: sources,
		groups:  groups,
	}

	return stream, nil
}

func (self *streamImplFactory) make_symbol_table_by_stream_option() SymbolTable {
	syms := []Symbol{}
	for _, src := range self.opt.sources {
		us := src.upstream
		syms = append(syms, NewSymbol(us.id, COMPONENT_UPSTREAM, us.alias))
	}

	for _, grp := range self.opt.groups {
		for _, in := range grp.inputs {
			syms = append(syms, NewSymbol(in.id, COMPONENT_INPUT, in.alias))
		}

		for _, out := range grp.outputs {
			syms = append(syms, NewSymbol(out.id, COMPONENT_OUTPUT, out.alias))
		}
	}
	sym_tbl := NewSymbolTable(syms)

	return sym_tbl
}

func init() {
	RegisterStreamFactory("default", func() StreamFactory { return &streamImplFactory{&streamImplOption{}} })
}
