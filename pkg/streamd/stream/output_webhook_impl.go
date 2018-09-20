package stream_manager

import (
	"context"
	"sync"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/lovoo/goka"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

const (
	WEBHOOK_OUTPUT_GROUP = "metathings.streamd.output.webhook.group"
)

type webhookOutputOption struct {
	id    string
	alias string

	logger                log.FieldLogger
	brokers               []string
	luanch_script         string
	webhook_body_template string
	webhook_url           string
}

type webhookOutput struct {
	Emitter
	slck             *sync.Mutex
	logger           log.FieldLogger
	state            OutputState
	opt              *webhookOutputOption
	stop_fn          func()
	goka_group_graph *goka.GroupGraph
	goka_processor   *goka.Processor
}

func (self *webhookOutput) Id() string {
	return self.opt.id
}

func (self *webhookOutput) Symbol() string {
	sym := NewSymbol(self.opt.id, COMPONENT_OUTPUT, self.opt.alias)
	return sym.String()
}

func (self *webhookOutput) luancher(ctx goka.Context, msg interface{}) {
	op_dat, ok := msg.(*OutputData)
	if !ok {
		self.logger.Warningf("failed to convert message to OutputData")
		return
	}
	self.logger.WithField("from", op_dat.Metadata().AsString("from")).Debugf("receive data")

	dat, err := self.luanch_output_data(op_dat)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to luanch output data")
		return
	}

	webhook_body, err := self.render_webhook_body(dat)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to render webhook body")
		return
	}

	_, _, errs := gorequest.New().Post(self.opt.webhook_url).Send(webhook_body).End()
	if len(errs) > 0 {
		self.logger.WithError(errs[0]).Errorf("failed to post webhook body to webhook url")
		return
	}

	self.logger.Debugf("send webhook body to webhook url")
}

func (self *webhookOutput) luanch_output_data(output_data *OutputData) (StreamData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	eng := NewLuaEngine(self.opt.logger)
	defer eng.Close()

	eng.SetContext(ctx)
	return eng.Luanch(self.opt.luanch_script, output_data.Metadata().Data(), output_data.Data())
}

func (self *webhookOutput) render_webhook_body(data StreamData) (string, error) {
	return mustache.Render(self.opt.webhook_body_template, data.Data())
}

func (self *webhookOutput) start() {
	self.slck.Lock()
	defer self.slck.Unlock()

	group_graph := goka.DefineGroup(WEBHOOK_OUTPUT_GROUP,
		goka.Input(goka.Stream(self.Symbol()), new(OutputDataCodec), self.luancher),
	)

	processor, err := goka.NewProcessor(self.opt.brokers, group_graph)
	if err != nil {
		self.state = OUTPUT_STATE_STOP
		return
	}

	ctx, stop_fn := context.WithCancel(context.Background())
	self.stop_fn = stop_fn

	self.goka_group_graph = group_graph
	self.goka_processor = processor

	go func() {
		self.state = OUTPUT_STATE_RUNNING
		self.Emit(START_EVENT, nil)
		self.logger.Debugf("output.webhook started")

		err = processor.Run(ctx)
		if err != nil {
			self.logger.WithError(err).Warningf("output.webhook failed, force stop")
			go self.stop()
			return
		}
	}()
}

func (self *webhookOutput) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != OUTPUT_STATE_STOP {
		return ErrUnstartable
	}

	self.state = OUTPUT_STATE_STARTING
	go self.start()
	self.logger.Debugf("output.webhook starting")

	return nil
}

func (self *webhookOutput) Stop() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != OUTPUT_STATE_RUNNING {
		return ErrUnterminable
	}

	self.state = OUTPUT_STATE_TERMINATING
	go self.stop()
	self.logger.Debugf("output.webhook terminating")

	return nil
}

func (self *webhookOutput) stop() {
	self.slck.Lock()
	defer self.slck.Unlock()

	self.stop_fn()

	self.state = OUTPUT_STATE_STOP
	self.Emit(STOP_EVENT, nil)
	self.logger.Debugf("output.webhook terminated")
}

func (self *webhookOutput) State() OutputState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *webhookOutput) Close() {
	panic("unimplemented")
}

type webhookOutputFactory struct {
	opt *webhookOutputOption
}

func (self *webhookOutputFactory) Set(key string, val interface{}) OutputFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "brokers":
		self.opt.brokers = val.([]string)
	case "option":
		opt := val.(*OutputOption)
		self.opt.id = opt.Id
		self.opt.alias = opt.Alias
		self.opt.luanch_script = opt.Config["luanch_script"]
		self.opt.webhook_body_template = opt.Config["webhook_body_template"]
		self.opt.webhook_url = opt.Config["webhook_url"]
	}

	return self
}

func (self *webhookOutputFactory) New() (Output, error) {
	output := &webhookOutput{
		Emitter: NewEmitter(),
		slck:    &sync.Mutex{},
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "output:webhook",
		}),
		state: OUTPUT_STATE_STOP,
		opt:   self.opt,
	}

	return output, nil
}

func init() {
	RegisterOutputFactory("webhook", func() OutputFactory {
		return &webhookOutputFactory{
			opt: &webhookOutputOption{},
		}
	})
}
