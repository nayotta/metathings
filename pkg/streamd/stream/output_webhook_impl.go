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

func SetWebhookOutputId(id string) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).id = id
	}
}

func SetWebhookOutputAlias(alias string) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).alias = alias
	}
}

func SetWebhookOutputLogger(logger log.FieldLogger) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).logger = logger
	}
}

func SetWebhookOutputBrokers(brokers []string) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).brokers = brokers
	}
}

func SetWebhookOutputLuanchScript(luanch_script string) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).luanch_script = luanch_script
	}
}

func SetWebhookOutputWebhookBodyTemplate(webhook_body_template string) OutputOption {
	return func(o interface{}) {
		o.(*webhookOutputOption).webhook_body_template = webhook_body_template
	}
}

type webhookOutput struct {
	Emitter
	slck             *sync.Mutex
	logger           log.FieldLogger
	state            OutputState
	opt              webhookOutputOption
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

	eng := NewLuaEngine()
	defer eng.Close()

	eng.SetContext(ctx)
	return eng.Luanch(self.opt.luanch_script, output_data.Metadata().Data(), output_data.Data())
}

func (self *webhookOutput) render_webhook_body(data StreamData) (string, error) {
	return mustache.Render(self.opt.luanch_script, data.Data())
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

	err = processor.Run(ctx)
	if err != nil {
		self.state = OUTPUT_STATE_STOP
		return
	}

	self.goka_group_graph = group_graph
	self.goka_processor = processor

	self.state = OUTPUT_STATE_RUNNING
	self.Emit(START_EVENT, nil)
}

func (self *webhookOutput) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != OUTPUT_STATE_STOP {
		return ErrUnstartable
	}

	self.state = OUTPUT_STATE_STARTING
	go self.start()

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

	return nil
}

func (self *webhookOutput) stop() {
	self.slck.Lock()
	defer self.slck.Unlock()

	self.stop_fn()

	self.state = OUTPUT_STATE_STOP
	self.Emit(STOP_EVENT, nil)
}

func (self *webhookOutput) State() OutputState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *webhookOutput) Close() {
	panic("unimplemented")
}

func newWebhookOutput(os ...OutputOption) (Output, error) {
	opt := webhookOutputOption{}
	for _, o := range os {
		o(&opt)
	}

	wh_op := &webhookOutput{
		Emitter: NewEmitter(),
		slck:    &sync.Mutex{},
		logger:  opt.logger,
		state:   OUTPUT_STATE_STOP,
		opt:     opt,
	}

	return wh_op, nil
}

func init() {
	RegisterOutput("webhook", newWebhookOutput)
}
