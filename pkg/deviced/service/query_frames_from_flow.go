package metathings_deviced_service

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) QueryFramesFromFlow(ctx context.Context, req *pb.QueryFramesFromFlowRequest) (*pb.QueryFramesFromFlowResponse, error) {
	res := &pb.QueryFramesFromFlowResponse{}

	dev_r := req.GetDevice()
	dev_id := dev_r.GetId().GetValue()
	flws_r := dev_r.GetFlows()
	now := time.Now()
	begin_at_r := now.Add(-24 * time.Hour)
	end_at_r := now
	if req.GetFrom() != nil {
		begin_at_r = pb_helper.ToTime(*req.GetFrom())
	}
	if req.GetTo() != nil {
		end_at_r = pb_helper.ToTime(*req.GetTo())
	}

	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, flw_r := range flws_r {
		flw_n_r := flw_r.GetName().GetValue()

		for _, flw_s := range dev_s.Flows {
			if *flw_s.Name != flw_n_r {
				continue
			}

			f, err := self.new_flow(dev_id, *flw_s.Id)
			defer f.Close()
			if err != nil {
				self.logger.WithError(err).Errorf("failed to new flow")
				return nil, status.Errorf(codes.Internal, err.Error())
			}

			frms, err := f.QueryFrame(&flow.FlowFilter{
				BeginAt: begin_at_r,
				EndAt:   end_at_r,
			})

			if err != nil {
				self.logger.WithError(err).Errorf("failed to query frame")
				return nil, status.Errorf(codes.Internal, err.Error())
			}

			pack := &pb.QueryFramesFromFlowResponse_Pack{
				Flow:   &pb.Flow{Id: *flw_s.Id},
				Frames: frms,
			}

			res.Packs = append(res.Packs, pack)
			break
		}
	}

	self.logger.Debugf("query frame from flow")

	return res, nil
}
