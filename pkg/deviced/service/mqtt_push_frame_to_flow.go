package metathings_deviced_service

import (
	"context"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MqttPushFrameToFlow MqttPushFrameToFlow
func (self *MetathingsDevicedService) MqttPushFrameToFlow(ctx context.Context, req *pb.MqttPushFrameToFlowRequest) (*pb.MqttPushFrameToFlowResponse, error) {
	var dev_r *pb.OpDevice
	var dev_id string
	var push_ack bool

	dev_id = req.GetFlowId().GetValue()
	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var f flow.Flow
match_flow_loop:
	for _, flw_s := range dev_s.Flows {
		if *flw_s.Id != req.GetFlowId().GetValue() {
			continue
		}

		f, err = self.new_flow(dev_id, *flw_s.Id)
		if err != nil {
			self.logger.WithError(err).Errorf("failed to new flow")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		defer f.Close()

		break match_flow_loop
	}

	if f == nil {
		err = ErrFlowNotFound
		self.logger.WithError(err).Errorf("failed to get flow")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	opfrm := req.GetFrame()
	err = f.PushFrame(&pb.Frame{Data: opfrm.GetData()})
	if err != nil {
		self.logger.WithError(err).Errorf("failed to push frame to flow")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.MqttPushFrameToFlowResponse{
		Id: req.GetId().GetValue(),
		Response: &pb.MqttPushFrameToFlowResponse_Ack_{},
	}

	return res, nil
}
