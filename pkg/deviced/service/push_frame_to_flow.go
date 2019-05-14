package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) PushFrameToFlow(stm pb.DevicedService_PushFrameToFlowServer) error {
	req, err := stm.Recv()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to receive config request")
		return status.Errorf(codes.Internal, err.Error())
	}

	var dev_r *pb.OpDevice
	var dev_id string
	var cfg_ack, push_ack bool

	cfg := req.GetConfig()
	dev_r = cfg.GetDevice()
	dev_id = dev_r.GetId().GetValue()
	cfg_ack = cfg.GetConfigAck().GetValue()
	push_ack = cfg.GetPushAck().GetValue()

	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device")
		return status.Errorf(codes.Internal, err.Error())
	}

	var f flow.Flow
match_flow_loop:
	for _, flw_r := range dev_r.GetFlows() {
		flw_n_r := flw_r.GetName().GetValue()
		for _, flw_s := range dev_s.Flows {
			if *flw_s.Name != flw_n_r {
				continue
			}

			f, err = self.new_flow(dev_id, *flw_s.Id)
			if err != nil {
				self.logger.WithError(err).Errorf("failed to new flow")
				return status.Errorf(codes.Internal, err.Error())
			}
			defer f.Close()

			break match_flow_loop
		}
	}

	if f == nil {
		err = ErrFlowNotFound
		self.logger.WithError(err).Errorf("failed to get flow")
		return status.Errorf(codes.NotFound, err.Error())
	}

	if cfg_ack {
		err = stm.Send(&pb.PushFrameToFlowResponse{
			Id:       req.GetId().GetValue(),
			Response: &pb.PushFrameToFlowResponse_Ack_{Ack: &pb.PushFrameToFlowResponse_Ack{}},
		})
		if err != nil {
			self.logger.WithError(err).Errorf("failed to send config ack message")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	for {
		req, err = stm.Recv()
		if err != nil {
			self.logger.WithError(err).Errorf("failed to receive frame request")
			return status.Errorf(codes.Internal, err.Error())
		}

		opfrm := req.GetFrame()
		err = f.PushFrame(&pb.Frame{Data: opfrm.GetData()})
		if err != nil {
			self.logger.WithError(err).Errorf("failed to push frame to flow")
			return status.Errorf(codes.Internal, err.Error())
		}

		if push_ack {
			err = stm.Send(&pb.PushFrameToFlowResponse{
				Id:       req.GetId().GetValue(),
				Response: &pb.PushFrameToFlowResponse_Ack_{Ack: &pb.PushFrameToFlowResponse_Ack{}},
			})
			if err != nil {
				self.logger.WithError(err).Errorf("failed to send push ack message")
				return status.Errorf(codes.Internal, err.Error())
			}
		}
	}
}
