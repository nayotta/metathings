package metathings_sensor_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
	driver "github.com/nayotta/metathings/pkg/sensor/driver"
	state_helper "github.com/nayotta/metathings/pkg/sensor/state"
)

type metathingsSensorService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
	snr_mgr *SensorManager

	sensor_st_psr state_helper.SensorStateParser
}

func (srv *metathingsSensorService) copySensorConfig(cfg driver.SensorConfig) map[string]*pb.SensorValue {
	snr_cfg := make(map[string]*pb.SensorValue)
	for _, key := range cfg.Keys() {
		val := cfg.Get(key)
		switch val.(type) {
		case float64:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Double{Double: val.(float64)}}
		case float32:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Float{Float: val.(float32)}}
		case int64:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Int64{Int64: val.(int64)}}
		case uint64:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Uint64{Uint64: val.(uint64)}}
		case int32:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Int32{Int32: val.(int32)}}
		case uint32:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Uint32{Uint32: val.(uint32)}}
		case bool:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_Bool{Bool: val.(bool)}}
		case string:
			snr_cfg[key] = &pb.SensorValue{Value: &pb.SensorValue_String_{String_: val.(string)}}
		}
	}
	return snr_cfg
}

func (srv *metathingsSensorService) copySensorData(dat driver.SensorData) map[string]*pb.SensorValue {
	snr_dat := make(map[string]*pb.SensorValue)
	for _, key := range dat.Keys() {
		val := dat.Get(key)
		switch val.(type) {
		case float64:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Double{Double: val.(float64)}}
		case float32:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Float{Float: val.(float32)}}
		case int64:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Int64{Int64: val.(int64)}}
		case uint64:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Uint64{Uint64: val.(uint64)}}
		case int32:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Int32{Int32: val.(int32)}}
		case uint32:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Uint32{Uint32: val.(uint32)}}
		case bool:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_Bool{Bool: val.(bool)}}
		case string:
			snr_dat[key] = &pb.SensorValue{Value: &pb.SensorValue_String_{String_: val.(string)}}
		}
	}
	return snr_dat
}

func (srv *metathingsSensorService) copySensor(s Sensor) *pb.Sensor {
	snr := s.Driver.Show()
	return &pb.Sensor{
		Name:   s.Name,
		State:  srv.sensor_st_psr.ToValue(snr.State.ToString()),
		Config: srv.copySensorConfig(snr.Config),
	}
}

func (srv *metathingsSensorService) copySensors(ss []Sensor) []*pb.Sensor {
	snrs := make([]*pb.Sensor, 0, len(ss))
	for _, s := range ss {
		snrs = append(snrs, srv.copySensor(s))
	}
	return snrs
}

func (srv *metathingsSensorService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snr, err := srv.snr_mgr.GetSensor(req.Name.Value)
	if err != nil {
		srv.logger.WithError(err).WithField("name", req.Name.Value).Errorf("failed to get sensor")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.GetResponse{
		Sensor: srv.copySensor(snr),
	}

	srv.logger.WithField("sensor", snr).Debugf("get sensor")

	return res, nil
}

func (srv *metathingsSensorService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snrs := srv.copySensors(srv.snr_mgr.ListSensors())
	res := &pb.ListResponse{
		Sensors: snrs,
	}

	srv.logger.WithField("sensors", snrs).Debugf("list sensors")

	return res, nil
}

func (srv *metathingsSensorService) Patch(ctx context.Context, req *pb.PatchRequest) (*pb.PatchResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snr, err := srv.snr_mgr.GetSensor(req.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get sensor")
	}

	cfg := driver.NewSensorConfig()
	for key, val := range req.Config {
		switch val.Value.(type) {
		case *pb.SensorValue_Double:
			cfg.Set(key, val.GetDouble())
		case *pb.SensorValue_Float:
			cfg.Set(key, val.GetFloat())
		case *pb.SensorValue_Int64:
			cfg.Set(key, val.GetInt64())
		case *pb.SensorValue_Uint64:
			cfg.Set(key, val.GetUint64())
		case *pb.SensorValue_Int32:
			cfg.Set(key, val.GetInt32())
		case *pb.SensorValue_Uint32:
			cfg.Set(key, val.GetUint32())
		case *pb.SensorValue_Bool:
			cfg.Set(key, val.GetBool())
		case *pb.SensorValue_String_:
			cfg.Set(key, val.GetString_())
		}
	}
	snr.Driver.Config(cfg)
	res := &pb.PatchResponse{
		Sensor: srv.copySensor(snr),
	}

	fields := make(log.Fields)
	for _, key := range cfg.Keys() {
		fields[key] = cfg.Get(key)
	}
	fields["name"] = req.Name.Value
	log.WithFields(fields).Debugf("patch sensor config")

	return res, nil
}

func (srv *metathingsSensorService) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snr, err := srv.snr_mgr.GetSensor(req.Name.Value)
	if err != nil {
		srv.logger.WithError(err).WithField("name", req.Name.Value).Errorf("failed to get sensor")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snr_dat := srv.copySensorData(snr.Driver.Data())
	res := &pb.GetDataResponse{
		Data: &pb.SensorData{
			Data: snr_dat,
		},
	}

	srv.logger.WithFields(log.Fields{
		"name": req.Name.Value,
		"data": snr_dat,
	}).Debugf("get sensor data")
	return res, nil
}

func (srv *metathingsSensorService) ListData(ctx context.Context, req *pb.ListDataRequest) (*pb.ListDataResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snrs := srv.snr_mgr.ListSensors()
	snr_dats := make(map[string]*pb.SensorData)
	for _, snr := range snrs {
		snr_dats[snr.Name] = &pb.SensorData{Data: srv.copySensorData(snr.Driver.Data())}
	}
	res := &pb.ListDataResponse{
		Datas: snr_dats,
	}

	srv.logger.WithField("datas", snr_dats).Debugf("list sensor data")
	return res, nil
}

func (srv *metathingsSensorService) Close() {

}

func NewSensorService(opt opt_helper.Option) (*metathingsSensorService, error) {
	opt.Set("service_name", "sensor")

	logger, err := log_helper.NewLogger("sensor", opt.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENT_CONFIG] = client_helper.ServiceConfig{opt.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	opt.Set("logger", logger)
	snr_mgr, err := NewSensorManager(opt)
	if err != nil {
		return nil, err
	}
	logger.Debugf("new sensor manager")

	srv := &metathingsSensorService{
		opt:     opt,
		logger:  logger,
		cli_fty: cli_fty,
		snr_mgr: snr_mgr,

		sensor_st_psr: state_helper.SENSOR_STATE_PARSER,
	}

	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	return srv, nil
}
