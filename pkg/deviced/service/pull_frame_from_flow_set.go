package metathings_deviced_service

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) PullFrameFromFlowSet(req *pb.PullFrameFromFlowSetRequest, stm pb.DevicedService_PullFrameFromFlowSetServer) error {
	cfg_req_id := req.GetId().GetValue()
	cfg := req.GetConfig()
	cfg_flwst := cfg.GetFlowSet()
	cfg_flwst_id := cfg_flwst.GetId().GetValue()
	cfg_ack := cfg.GetConfigAck().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"config": cfg_req_id,
	})

	_, err := self.storage.GetFlowSet(stm.Context(), cfg_flwst_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get flow set in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	flwst, err := self.flwst_fty.New(&flow.FlowSetOption{FlowSetId: cfg_flwst_id})
	if err != nil {
		logger.WithError(err).Errorf("failed to new flow set")
		return status.Errorf(codes.Internal, err.Error())
	}

	packch := make(chan *pb.PullFrameFromFlowSetResponse_Pack_)
	pfch, quit := flwst.PullFrame()
	defer close(quit)

	go func() {
		defer flwst.Close()
		defer close(packch)

		for {
			select {
			case frm := <-pfch:
				packch <- &pb.PullFrameFromFlowSetResponse_Pack_{
					Pack: &pb.PullFrameFromFlowSetResponse_Pack{
						Device: frm.Device,
						Frames: []*pb.Frame{
							frm.Frame,
						},
					},
				}
			}
		}
	}()

	if cfg_ack {
		if err = stm.Send(&pb.PullFrameFromFlowSetResponse{
			Response: &pb.PullFrameFromFlowSetResponse_Ack_{
				Ack: &pb.PullFrameFromFlowSetResponse_Ack{},
			},
		}); err != nil {
			logger.WithError(err).Errorf("failed to send config ack to stream")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	for {
		select {
		case pack := <-packch:
			if err = stm.Send(&pb.PullFrameFromFlowSetResponse{
				Response: pack,
			}); err != nil {
				logger.WithError(err).Errorf("failed to send pack to stream")
				return status.Errorf(codes.Internal, err.Error())
			}
		}
	}
}
