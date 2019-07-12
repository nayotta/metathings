package metathings_device_cloud_service

import (
	"github.com/golang/protobuf/ptypes/wrappers"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (s *MetathingsDeviceCloudService) start_push_frame_loop(dev_id string, req *PushFrameToFlowRequest, sess string) {
	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		s.get_logger().WithError(err).Warningf("failed to new deviced service client")
		return
	}

	ctx := s.context_with_device(dev_id)
	stm, err := cli.PushFrameToFlow(ctx)
	if err != nil {
		defer cfn()
		s.get_logger().WithError(err).Warningf("failed to create push frame to flow streaming")
		return
	}

	cfg := &pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: req.Id},
		Request: &pb.PushFrameToFlowRequest_Config_{
			Config: &pb.PushFrameToFlowRequest_Config{
				Device: &pb.OpDevice{
					Id: &wrappers.StringValue{Value: dev_id},
				},
				ConfigAck: &wrappers.BoolValue{Value: req.ConfigAck},
				PushAck:   &wrappers.BoolValue{Value: req.PushAck},
			},
		},
	}

	if err = stm.Send(cfg); err != nil {
		defer cfn()
		s.get_logger().WithError(err).Warningf("failed to config push frame to flow streaming")
		return
	}

	if req.ConfigAck {
		res, err := stm.Recv()
		if err != nil {
			defer cfn()
			s.get_logger().WithError(err).Warningf("failed to recv push frame to flow config response ack")
			return
		}

		if req.Id != res.GetId() {
			defer cfn()
			s.get_logger().Warningf("failed to match push frame to flow config response ack")
			return
		}
	}

	// TODO(Peer): more arguments
	pffch, err := NewPushFrameToFlowChannel("mqtt")
	if err != nil {
		s.get_logger().WithError(err).Warningf("failed to new push frame to flow channel")
		return
	}

	go s.push_frame_loop(stm, pffch, cfn, false)
}

func (s *MetathingsDeviceCloudService) push_frame_loop(stm pb.DevicedService_PushFrameToFlowClient, pffch PushFrameToFlowChannel, cfn client_helper.CloseFn, push_ack bool) {
	defer func() {
		cfn()
		pffch.Close()
	}()
	ch := pffch.Channel()
	for {
		frm, ok := <-ch
		if !ok {
			s.get_logger().Debugf("push frame to flow channel closed")
			return
		}

		req := &pb.PushFrameToFlowRequest{
			Id: &wrappers.StringValue{Value: id_helper.NewId()},
			Request: &pb.PushFrameToFlowRequest_Frame{
				Frame: frm,
			},
		}

		if err := stm.Send(req); err != nil {
			s.get_logger().WithError(err).Warningf("failed to send frame to streaming")
			return
		}
	}
}
