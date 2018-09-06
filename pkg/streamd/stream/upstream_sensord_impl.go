package stream_manager

import (
	"context"
	"sync"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

type sensordUpstreamOption struct {
	id string

	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
	snr_id       string
	brokers      []string
	targets      []string
	sym_tbl      SymbolTable
}

func SetSensordUpstreamId(id string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).id = id
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

func SetBrokers(brokers []string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).brokers = brokers
	}
}

func SetTargets(targets []string) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).targets = targets
	}
}

func SetSymbolTable(sym_tbl SymbolTable) UpstreamOption {
	return func(o interface{}) {
		o.(*sensordUpstreamOption).sym_tbl = sym_tbl
	}
}

type sensordUpstream struct {
	Emitter
	slck   *sync.Mutex
	logger log.FieldLogger
	state  UpstreamState
	opt    sensordUpstreamOption
	cfn    client_helper.CloseFn
}

func (self *sensordUpstream) Id() string {
	return self.opt.id
}

func (self *sensordUpstream) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()
	if self.state != UPSTREAM_STATE_STOP {
		return ErrUnstartable
	}

	self.state = UPSTREAM_STATE_STARTING
	self.Emit(START_EVENT, nil)

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
	self.slck.Unlock()

	for {
		sub_ress, err := stm.Recv()
		if err != nil {
			self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to recv data from stream")
			return
		}

		for _, sub_res := range sub_ress.Responses {
			usd := enc_sensord_upstream_data(sub_res.Data)
			log.WithField("data", usd).Infof("recv data")
		}
	}
}

func enc_sensord_upstream_data(snr_dat *sensord_pb.SensorData) *UpstreamData {
	md := map[string]interface{}{}
	md["sensor_id"] = snr_dat.SensorId
	md["sensor_name"] = snr_dat.Data["$sensor.name"].GetString_()
	md["created_at"] = protobuf_helper.ToTime(*snr_dat.CreatedAt)
	md["arrvied_at"] = protobuf_helper.ToTime(*snr_dat.ArrivedAt)

	d := map[string]interface{}{}
	for k, v := range snr_dat.Data {
		if len(k) > 0 && k[0] != '$' {
			d[k] = dec_sensordData_value(v)
		}
	}

	return NewUpstreamData(d, md)
}

func dec_sensordData_value(v *sensor_pb.SensorValue) interface{} {
	switch v.Value.(type) {
	case *sensor_pb.SensorValue_Double:
		return v.GetDouble()
	case *sensor_pb.SensorValue_Float:
		return v.GetFloat()
	case *sensor_pb.SensorValue_Int64:
		return v.GetInt64()
	case *sensor_pb.SensorValue_Uint64:
		return v.GetUint64()
	case *sensor_pb.SensorValue_Int32:
		return v.GetInt32()
	case *sensor_pb.SensorValue_Uint32:
		return v.GetUint32()
	case *sensor_pb.SensorValue_Bool:
		return v.GetBool()
	case *sensor_pb.SensorValue_String_:
		return v.GetString_()
	default:
		panic("unimplemented")
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
		Emitter: NewEmitter(),
		slck:    &sync.Mutex{},
		state:   UPSTREAM_STATE_STOP,
		opt:     opt,
	}
	return snrd_upstm, nil
}

func init() {
	RegisterUpstream("sensord", newSensordUpstream)
}
