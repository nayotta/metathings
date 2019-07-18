package metathings_device_cloud_service

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (s *MetathingsDeviceCloudService) start_push_frame_loop(dev_id string, req *device_pb.PushFrameToFlowRequest, sess string) error {
	req_cfg := req.GetConfig()
	cfg_ack := req_cfg.GetConfigAck().GetValue()
	psh_ack := req_cfg.GetPushAck().GetValue()
	flw_name := req_cfg.GetFlow().GetName().GetValue()
	req_id := req.GetId().GetValue()

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		s.get_logger().WithError(err).Warningf("failed to new deviced service client")
		return err
	}

	logger := s.get_logger().WithFields(log.Fields{
		"device":     dev_id,
		"flow":       flw_name,
		"config_ack": cfg_ack,
		"push_ack":   psh_ack,
	})

	ctx := s.context_with_device(dev_id)
	stm, err := cli.PushFrameToFlow(ctx)
	if err != nil {
		defer cfn()
		logger.WithField("request", req_id).WithError(err).Warningf("failed to create push frame to flow streaming")
		return err
	}
	logger.WithField("request", req_id).Debugf("build push frame to flow streaming")

	cfg := &deviced_pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: req.GetId().GetValue()},
		Request: &deviced_pb.PushFrameToFlowRequest_Config_{
			Config: &deviced_pb.PushFrameToFlowRequest_Config{
				Device: &deviced_pb.OpDevice{
					Id: &wrappers.StringValue{Value: dev_id},
					Flows: []*deviced_pb.OpFlow{
						req_cfg.GetFlow(),
					},
				},
				ConfigAck: &wrappers.BoolValue{Value: cfg_ack},
				PushAck:   &wrappers.BoolValue{Value: psh_ack},
			},
		},
	}

	if err = stm.Send(cfg); err != nil {
		defer cfn()
		logger.WithField("request", req_id).WithError(err).Warningf("failed to config push frame to flow streaming")
		return err
	}
	logger.WithField("request", req_id).Debugf("send push frame to flow config request")

	if cfg_ack {
		logger.WithField("request", req_id).Debugf("waiting for config ack")
		res, err := stm.Recv()
		if err != nil {
			defer cfn()
			logger.WithField("request", req_id).Warningf("failed to recv push frame to flow config response ack")
			return err
		}

		if req_id != res.GetId() {
			defer cfn()
			logger.WithField("request", req_id).Warningf("failed to match push frame to flow config response ack")
			return ErrUnmatchedRequestId
		}
		logger.WithField("request", req_id).Debugf("recv config ack")

	}

	var pffch PushFrameToFlowChannel

	// TODO(Peer): get flow channel driver from request
	flw_ch_drv := "mqtt"
	switch flw_ch_drv {
	case "mqtt":
		pffch, err = NewPushFrameToFlowChannel("mqtt",
			"mqtt_address", s.opt.Connection.Mqtt.Address,
			"mqtt_username", s.opt.Credential.Id,
			"mqtt_password", mosquitto_service.ParseMosquittoPluginPassword(s.opt.Credential.Id, s.opt.Credential.Secret),
			"device_id", dev_id,
			"channel_session", sess,
			"push_ack", psh_ack,
			"logger", s.logger,
		)
	default:
		defer cfn()
		logger.Warningf("unspported flow channel driver")
		return ErrUnsupportedFlowChannelDriver
	}
	if err != nil {
		defer cfn()
		logger.WithError(err).Warningf("failed to new push frame to flow channel")
		return err
	}

	go s.push_frame_loop(stm, pffch, cfn, psh_ack)
	s.get_logger().Debugf("push frame loop started")

	return nil
}

func (s *MetathingsDeviceCloudService) push_frame_loop(stm deviced_pb.DevicedService_PushFrameToFlowClient, pffch PushFrameToFlowChannel, cfn client_helper.CloseFn, push_ack bool) {
	defer func() {
		cfn()
		pffch.Close()
		s.get_logger().Debugf("push frame loop stoped")
	}()
	for {
		frm, ok := <-pffch.Channel()
		if !ok {
			s.get_logger().Debugf("push frame to flow channel closed")
			return
		}

		req := &deviced_pb.PushFrameToFlowRequest{
			Id: &wrappers.StringValue{Value: id_helper.NewId()},
			Request: &deviced_pb.PushFrameToFlowRequest_Frame{
				Frame: frm,
			},
		}

		if err := stm.Send(req); err != nil {
			s.get_logger().WithError(err).Warningf("failed to send frame to streaming")
			return
		}
	}
}
