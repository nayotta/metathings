package stream_manager

import (
	"context"
	"sync"
	"time"

	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"
)

const (
	FILTER_INPUT_GROUP = "metathings.streamd.input.filter.group"
)

type filterInputOption struct {
	id    string
	alias string

	logger  log.FieldLogger
	brokers []string
	targets []string
	filters map[string]string
	sym_tbl SymbolTable
}

type filterInput struct {
	Emitter
	slck             *sync.Mutex
	logger           log.FieldLogger
	state            InputState
	opt              *filterInputOption
	emitters         map[string]*goka.Emitter
	stop_fn          func()
	goka_group_graph *goka.GroupGraph
	goka_processor   *goka.Processor
}

func (self *filterInput) Id() string {
	return self.opt.id
}

func (self *filterInput) Symbol() string {
	sym := NewSymbol(self.opt.id, COMPONENT_INPUT, self.opt.alias)
	return sym.String()
}

func (self *filterInput) filter(ctx goka.Context, msg interface{}) {
	ip_dat, ok := msg.(*InputData)
	if !ok {
		self.logger.Warningf("failed to convert message to InputData")
		return
	}

	for target, filter := range self.opt.filters {
		ok, err := self.filter_input_data(filter, ip_dat)
		if err != nil {
			self.logger.WithError(err).Warningf("failed to filter input data")
		} else if ok {
			if err = self.emit_input_data(target, ip_dat); err != nil {
				self.logger.WithError(err).Warningf("failed to emit input data")
			}
		}
	}
}

func (self *filterInput) filter_input_data(filter string, input_data *InputData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	eng := NewLuaEngine()
	defer eng.Close()

	eng.SetContext(ctx)
	return eng.Filter(filter, input_data.Metadata().Data(), input_data.Data())
}

func (self *filterInput) emit_input_data(target string, input_data *InputData) error {
	tgr_sym := self.opt.sym_tbl.Lookup(target)

	var codec goka.Codec
	var msg interface{}

	switch tgr_sym.Component() {
	case COMPONENT_INPUT:
		input_data.Metadata().Set("from", self.Symbol())
		msg = input_data
		codec = new(InputDataCodec)
	case COMPONENT_OUTPUT:
		output_data := InputDataToOutputData(input_data)
		output_data.Metadata().Set("from", self.Symbol())
		codec = new(OutputDataCodec)
	}

	var emitter *goka.Emitter
	var ok bool
	var err error

	if emitter, ok = self.emitters[tgr_sym.String()]; !ok {
		emitter, err = goka.NewEmitter(self.opt.brokers, goka.Stream(tgr_sym.String()), codec)
		if err != nil {
			return err
		}

		self.emitters[tgr_sym.String()] = emitter
	}

	err = emitter.EmitSync("", msg)
	if err != nil {
		return err
	}

	return nil
}

func (self *filterInput) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != INPUT_STATE_STOP {
		return ErrUnstartable
	}

	self.state = INPUT_STATE_STARTING
	go self.start()

	return nil
}

func (self *filterInput) start() {
	self.slck.Lock()
	defer self.slck.Unlock()

	group_graph := goka.DefineGroup(FILTER_INPUT_GROUP,
		goka.Input(goka.Stream(self.Symbol()), new(InputDataCodec), self.filter),
	)

	processor, err := goka.NewProcessor(self.opt.brokers, group_graph)
	if err != nil {
		self.state = INPUT_STATE_STOP
		return
	}

	ctx, stop_fn := context.WithCancel(context.Background())
	self.stop_fn = stop_fn

	err = processor.Run(ctx)
	if err != nil {
		self.state = INPUT_STATE_STOP
		return
	}

	self.goka_group_graph = group_graph
	self.goka_processor = processor

	self.state = INPUT_STATE_RUNNING
	self.Emit(START_EVENT, nil)
}

func (self *filterInput) Stop() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != INPUT_STATE_RUNNING {
		return ErrUnterminable
	}

	self.state = INPUT_STATE_TERMINATING

	go self.stop()

	return nil
}

func (self *filterInput) stop() {
	self.slck.Lock()
	defer self.slck.Unlock()

	self.stop_fn()

	self.state = INPUT_STATE_STOP
	self.Emit(STOP_EVENT, nil)
}

func (self *filterInput) State() InputState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *filterInput) Close() {
	panic("unimplemented")
}

type filterInputFactory struct {
	opt *filterInputOption
}

func (self *filterInputFactory) Set(key string, val interface{}) InputFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "symbol_table":
		self.opt.sym_tbl = val.(SymbolTable)
	case "brokers":
		self.opt.brokers = val.([]string)
	case "option":
		opt := val.(*InputOption)
		self.opt.id = opt.Id
		self.opt.alias = opt.Alias
		self.opt.targets = split_and_trim(opt.Config["targets"])
		self.opt.filters = group_by_prefix(opt.Config, "filter.")
	}

	return self
}

func (self *filterInputFactory) New() (Input, error) {
	input := &filterInput{
		Emitter: NewEmitter(),
		slck:    &sync.Mutex{},
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "input:filter",
		}),
		state:    INPUT_STATE_STOP,
		opt:      self.opt,
		emitters: map[string]*goka.Emitter{},
	}

	return input, nil

}

func init() {
	RegisterInputFactory("filter", func() InputFactory {
		return &filterInputFactory{opt: &filterInputOption{}}
	})
}
