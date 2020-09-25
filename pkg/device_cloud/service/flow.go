package metathings_device_cloud_service

import (
	"math"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	config_helper "github.com/nayotta/metathings/pkg/common/config"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	device_pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (s *MetathingsDeviceCloudService) start_push_frame_loop(dev_id string, req *device_pb.PushFrameToFlowRequest, sess string) error {
	req_cfg := req.GetConfig()
	cfg_ack := req_cfg.GetConfigAck().GetValue()
	psh_ack := req_cfg.GetPushAck().GetValue()
	flw_name := req_cfg.GetFlow().GetName().GetValue()
	req_id := req.GetId().GetValue()

	logger := s.get_logger().WithFields(log.Fields{
		"device":  dev_id,
		"flow":    flw_name,
		"request": req_id,
	})

	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Warningf("failed to new deviced service client")
		return err
	}

	ctx := s.context_with_device(dev_id)
	stm, err := cli.PushFrameToFlow(ctx)
	if err != nil {
		defer cfn()
		logger.WithError(err).Warningf("failed to create push frame to flow streaming")
		return err
	}
	logger.Debugf("build push frame to flow streaming")

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
		logger.WithError(err).Warningf("failed to config push frame to flow streaming")
		return err
	}
	logger.Debugf("send push frame to flow config request")

	if cfg_ack {
		logger.Debugf("waiting for config ack")
		res, err := stm.Recv()
		if err != nil {
			defer cfn()
			logger.Warningf("failed to recv push frame to flow config response ack")
			return err
		}

		if req_id != res.GetId() {
			defer cfn()
			logger.Warningf("failed to match push frame to flow config response ack")
			return ErrUnmatchedRequestId
		}
		logger.Debugf("recv config ack")

	}

	var pffch PushFrameToFlowChannel

	drv, args, err := config_helper.ParseConfigOption("driver", s.opt.Connection)
	if err != nil {
		defer cfn()
		logger.WithError(err).Warningf("failed to parse config option")
		return err
	}

	switch drv {
	case "mqtt":
		args = append(
			args,
			"mqtt_username", s.opt.Credential.Id,
			"mqtt_password", mosquitto_service.ParseMosquittoPluginPassword(s.opt.Credential.Id, s.opt.Credential.Secret),
			"device_id", dev_id,
			"channel_session", sess,
			"push_ack", psh_ack,
			"logger", s.logger,
		)

		pffch, err = NewPushFrameToFlowChannel("mqtt", args...)
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

	go s.push_frame_loop(stm, pffch, cfn, psh_ack, logger)
	logger.Debugf("push frame loop started")

	return nil
}

func (s *MetathingsDeviceCloudService) push_frame_loop(stm deviced_pb.DevicedService_PushFrameToFlowClient, pffch PushFrameToFlowChannel, cfn client_helper.CloseFn, push_ack bool, logger log.FieldLogger) {
	defer func() {
		cfn()
		pffch.Close()
		logger.Debugf("push frame loop stoped")
	}()

	var acks sync.Map
	var ack_timeout_cnt int
	var push_ack_failed_limit int

	if push_ack {
		push_ack_failed_limit = s.opt.Methods.PushFrameToFlow.PushAckFailedLimit
	} else {
		push_ack_failed_limit = math.MaxInt64
	}

	if push_ack {
		go func() {
			defer logger.Debugf("push ack loop closed")
			for ack_timeout_cnt < push_ack_failed_limit {
				res, err := stm.Recv()
				if err != nil {
					logger.WithError(err).Debugf("failed to recv message from stream")
					return
				}

				req_id := res.GetId()
				go func(req_id string) {
					inner_logger := logger.WithFields(log.Fields{
						"current_request": req_id,
					})

					defer func() {
						if recover() != nil {
							inner_logger.Debugf("ack channel already closed")
						}
					}()

					chi, ok := acks.Load(req_id)
					if !ok {
						inner_logger.Debugf("failed to get ack channel in ack map")
					}
					chi.(chan struct{}) <- struct{}{}
				}(req_id)
			}
		}()
	}

	for ack_timeout_cnt < push_ack_failed_limit {
		frm, ok := <-pffch.Channel()
		if !ok {
			logger.Debugf("push frame to flow channel closed")
			return
		}

		req_id := id_helper.NewId()

		req := &deviced_pb.PushFrameToFlowRequest{
			Id: &wrappers.StringValue{Value: req_id},
			Request: &deviced_pb.PushFrameToFlowRequest_Frame{
				Frame: frm,
			},
		}

		if push_ack {
			ch := make(chan struct{})
			acks.Store(req_id, ch)
			go func() {
				inner_logger := logger.WithField("current_request", req_id)
				defer func() {
					close(ch)
					acks.Delete(req_id)
				}()
				select {
				case <-time.After(s.opt.Methods.PushFrameToFlow.PushAckTimeout):
					ack_timeout_cnt++
					inner_logger.WithField("ack_timeout_count", ack_timeout_cnt).Debugf("failed to get ack response")
				case _, ok := <-ch:
					if ok {
						// recv ack response
						ack_timeout_cnt = 0
					} else {
						ack_timeout_cnt++
					}
				}
			}()
		}

		if err := stm.Send(req); err != nil {
			logger.WithError(err).Warningf("failed to send frame to streaming")
			return
		}
	}
}
