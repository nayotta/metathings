package stream_manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

type sensordUpstreamOption struct {
	id    string
	alias string

	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
	snr_id       string
	brokers      []string
	targets      []string
	filters      map[string]string
	sym_tbl      SymbolTable
}

func SetSensordUpstreamId(id string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).id = id
	}
}

func SetSensordUpstreamAlias(alias string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).alias = alias
	}
}

func SetSensordUpstreamLogger(logger log.FieldLogger) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).logger = logger
	}
}

func SetSensordUpstreamApplicationCredentialManager(app_cred_mgr app_cred_mgr.ApplicationCredentialManager) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).app_cred_mgr = app_cred_mgr
	}
}

func SetSensordUpstreamClientFactory(cli_fty *client_helper.ClientFactory) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).cli_fty = cli_fty
	}
}

func SetSensordUpstreamBrokers(brokers []string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).brokers = brokers
	}
}

func SetSensordUpstreamTargets(targets []string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).targets = targets
	}
}

func SetSensordUpstreamFilters(filters map[string]string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).filters = filters
	}
}

func SetSensordUpstreamSymbolTable(sym_tbl SymbolTable) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).sym_tbl = sym_tbl
	}
}

type sensordUpstream struct {
	Emitter
	slck     *sync.Mutex
	logger   log.FieldLogger
	state    UpstreamState
	opt      sensordUpstreamOption
	cfn      client_helper.CloseFn
	emitters map[string]*goka.Emitter
}

func (self *sensordUpstream) Id() string {
	return self.opt.id
}

func (self *sensordUpstream) Symbol() string {
	sym := NewSymbol(self.Id(), COMPONENT_UPSTREAM, self.opt.alias)
	return sym.String()
}

func (self *sensordUpstream) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()
	if self.state != UPSTREAM_STATE_STOP {
		return ErrUnstartable
	}

	self.state = UPSTREAM_STATE_STARTING

	cli, cfn, err := self.opt.cli_fty.NewSensordServiceClient()
	if err != nil {
		self.state = UPSTREAM_STATE_STOP
		self.Emit(ERROR_EVENT, nil)
		return err
	}
	self.cfn = cfn

	go self.start(cli, cfn)

	return nil
}

func (self *sensordUpstream) start(cli sensord_pb.SensordServiceClient, cfn client_helper.CloseFn) {
	defer func() {
		self.slck.Lock()
		defer self.slck.Unlock()

		self.state = UPSTREAM_STATE_STOP
		self.Emit(STOP_EVENT, nil)

		self.logger.WithField("sensor_id", self.opt.snr_id).Infof("upstream terminated")
	}()

	ctx := context.Background()
	tk_ctx := context_helper.WithToken(ctx, self.opt.app_cred_mgr.GetToken())

	stm, err := cli.Subscribe(tk_ctx)
	if err != nil {
		self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to subscribe")
		return
	}

	sub_reqs := &sensord_pb.SubscribeRequests{
		Requests: []*sensord_pb.SubscribeRequest{
			&sensord_pb.SubscribeRequest{
				Payload: &sensord_pb.SubscribeRequest_SubscribeById{
					SubscribeById: &sensord_pb.SubscribeByIdRequest{
						Id: &gpb.StringValue{Value: self.opt.snr_id},
					},
				},
			},
		},
	}

	err = stm.Send(sub_reqs)
	if err != nil {
		self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to subscribe sesnor data")
		return
	}

	self.slck.Lock()
	self.state = UPSTREAM_STATE_RUNNING
	self.Emit(START_EVENT, nil)
	self.slck.Unlock()

	for {
		sub_ress, err := stm.Recv()
		if err != nil {
			self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to recv data from stream")
			return
		}

		for _, sub_res := range sub_ress.Responses {
			upstm_dat := enc_sensord_upstream_data(sub_res.Data)
			for target, filter := range self.opt.filters {
				ok, err := self.filter_upstream_data(filter, upstm_dat)
				if err != nil {
					self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Warningf("failed to filter upstream data")

				} else if ok {
					if err = self.emit_upstream_data(target, upstm_dat); err != nil {
						self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Warningf("failed to emit upstream data")
					}
				}
			}
		}
	}
}

// TODO(Peer): dont compile filter every time
func (self *sensordUpstream) filter_upstream_data(filter string, upstream_data *UpstreamData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	eng := NewLuaEngine()
	defer eng.Close()

	eng.SetContext(ctx)
	return eng.Filter(filter, upstream_data.Metadata().Data(), upstream_data.Data())

}

func (self *sensordUpstream) emit_upstream_data(target string, upstream_data *UpstreamData) error {
	sym := self.opt.sym_tbl.Lookup(target)

	var codec goka.Codec
	var msg interface{}

	switch sym.Component() {
	case COMPONENT_INPUT:
		input_data := UpstreamDataToInputData(upstream_data)
		input_data.Metadata().Set("from", self.Symbol())
		msg = input_data
		codec = new(InputDataCodec)
	case COMPONENT_OUTPUT:
		output_data := UpstreamDataToOutputData(upstream_data)
		output_data.Metadata().Set("from", self.Symbol())
		msg = output_data
		codec = new(OutputDataCodec)
	}

	var emitter *goka.Emitter
	var ok bool
	var err error

	if emitter, ok = self.emitters[sym.String()]; !ok {
		emitter, err = goka.NewEmitter(self.opt.brokers, goka.Stream(sym.String()), codec)
		if err != nil {
			return err
		}

		self.emitters[sym.String()] = emitter
	}

	err = emitter.EmitSync("", msg)
	if err != nil {
		return err
	}

	return nil
}

func enc_sensord_upstream_data(snr_dat *sensord_pb.SensorData) *UpstreamData {
	md := map[string]interface{}{}
	md["sensor_id"] = snr_dat.SensorId
	md["sensor_name"] = snr_dat.Data["$sensor.name"].GetString_()
	md["created_at"] = fmt.Sprint(protobuf_helper.ToTime(*snr_dat.CreatedAt).UnixNano())
	md["arrvied_at"] = fmt.Sprint(protobuf_helper.ToTime(*snr_dat.ArrivedAt).UnixNano())

	d := map[string]interface{}{}
	for k, v := range snr_dat.Data {
		if len(k) > 0 && k[0] != '$' {
			d[k] = dec_sensord_data_value(v)
		}
	}

	return NewUpstreamData(d, md)
}

func dec_sensord_data_value(v *sensor_pb.SensorValue) string {
	switch v.Value.(type) {
	case *sensor_pb.SensorValue_Double:
		return fmt.Sprint(v.GetDouble())
	case *sensor_pb.SensorValue_Float:
		return fmt.Sprint(v.GetFloat())
	case *sensor_pb.SensorValue_Int64:
		return fmt.Sprint(v.GetInt64())
	case *sensor_pb.SensorValue_Uint64:
		return fmt.Sprint(v.GetUint64())
	case *sensor_pb.SensorValue_Int32:
		return fmt.Sprint(v.GetInt32())
	case *sensor_pb.SensorValue_Uint32:
		return fmt.Sprint(v.GetUint32())
	case *sensor_pb.SensorValue_String_:
		return v.GetString_()
	default:
		return ""
	}
}

func (self *sensordUpstream) Stop() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != UPSTREAM_STATE_RUNNING {
		return ErrUnterminable
	}

	self.state = UPSTREAM_STATE_TERMINATING
	self.cfn()

	for _, emitter := range self.emitters {
		emitter.Finish()
	}

	self.logger.WithField("sensor_id", self.opt.snr_id).Debugf("upstream terminating")
	return nil
}

func (self *sensordUpstream) State() UpstreamState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *sensordUpstream) Close() {
	panic("unimplemented")
}

func newSensordUpstream(os ...UpstreamOption) (Upstream, error) {
	opt := sensordUpstreamOption{}
	for _, o := range os {
		o(&opt)
	}

	snrd_upstm := &sensordUpstream{
		Emitter:  NewEmitter(),
		slck:     &sync.Mutex{},
		logger:   opt.logger,
		state:    UPSTREAM_STATE_STOP,
		opt:      opt,
		emitters: map[string]*goka.Emitter{},
	}
	return snrd_upstm, nil
}

func init() {
	RegisterUpstream("sensord", newSensordUpstream)
}
