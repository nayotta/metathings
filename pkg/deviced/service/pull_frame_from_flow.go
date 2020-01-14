package metathings_deviced_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) PullFrameFromFlow(req *pb.PullFrameFromFlowRequest, stm pb.DevicedService_PullFrameFromFlowServer) error {
	cfg := req.GetConfig()
	dev_r := cfg.GetDevice()
	dev_id := dev_r.GetId().GetValue()
	cfg_ack := cfg.GetConfigAck().GetValue()

	dev_s, err := self.storage.GetDevice(stm.Context(), dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device")
		return status.Errorf(codes.Internal, err.Error())
	}

	var fs []flow.Flow
match_flow_loop:
	for _, flw_r := range dev_r.GetFlows() {
		flw_n_r := flw_r.GetName().GetValue()
		for _, flw_s := range dev_s.Flows {
			if *flw_s.Name != flw_n_r {
				continue
			}

			f, err := self.new_flow(dev_id, *flw_s.Id)
			if err != nil {
				self.logger.WithError(err).Errorf("failed to new flow")
				return status.Errorf(codes.Internal, err.Error())
			}
			defer f.Close()

			fs = append(fs, f)

			break match_flow_loop
		}
	}

	pack_ch := make(chan *pb.PullFrameFromFlowResponse_Pack_)
	for _, f := range fs {
		go func(f flow.Flow) {
			defer f.Close()

			// TODO(Peer): handle quit channel
			// TODO(Peer): send multi-frame by one pack
			pf_ch, _ := f.PullFrame()
			flw := &pb.Flow{Id: f.Id()}

			for {
				select {
				case frm := <-pf_ch:
					pack_ch <- &pb.PullFrameFromFlowResponse_Pack_{
						Pack: &pb.PullFrameFromFlowResponse_Pack{
							Flow:   flw,
							Frames: []*pb.Frame{frm},
						},
					}
				}
			}

		}(f)
	}

	if cfg_ack {
		err = stm.Send(&pb.PullFrameFromFlowResponse{Response: &pb.PullFrameFromFlowResponse_Ack_{
			Ack: &pb.PullFrameFromFlowResponse_Ack{},
		}})
		if err != nil {
			self.logger.WithError(err).Errorf("failed to send config ack message")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	for {
		select {
		case pack := <-pack_ch:
			err = stm.Send(&pb.PullFrameFromFlowResponse{
				Response: pack,
			})
			if err != nil {
				self.logger.WithError(err).Errorf("failed to send pack response")
				return status.Errorf(codes.Internal, err.Error())
			}
		}
	}
}
