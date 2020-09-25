package metathings_deviced_service

import (
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) PullFrameFromFlowSet(req *pb.PullFrameFromFlowSetRequest, stm pb.DevicedService_PullFrameFromFlowSetServer) error {
	cfg_req_id := req.GetId().GetValue()
	cfg := req.GetConfig()
	cfg_flwst := cfg.GetFlowSet()
	cfg_flwst_id := cfg_flwst.GetId().GetValue()
	cfg_ack := cfg.GetConfigAck().GetValue()

	logger := self.get_logger().WithFields(log.Fields{
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
	defer flwst.Close()

	running := true
	packch := make(chan *pb.PullFrameFromFlowSetResponse_Pack_)
	defer close(packch)

	errch := make(chan error, 1)
	defer close(errch)

	go func() {
		pfch := flwst.PullFrame()

		for running {
			select {
			case frm, ok := <-pfch:
				if running {
					if !ok && flwst.Err() != nil {
						errch <- flwst.Err()
						return
					}

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

	dead := make(chan struct{})
	go func() {
		defer close(dead)
		for running {
			if err = stm.Send(&pb.PullFrameFromFlowSetResponse{
				Id: "ffffffffffffffffffffffffffffffff",
				Response: &pb.PullFrameFromFlowSetResponse_Ack_{
					Ack: new(pb.PullFrameFromFlowSetResponse_Ack),
				},
			}); err != nil {
				logger.WithError(err).Errorf("failed to send alive response")
				return
			}

			time.Sleep(time.Duration(self.opt.Methods.PullFrameFromFlowSet.AliveInterval) * time.Second)
		}
		logger.Debugf("alive loop closed")
	}()

	defer func() { running = false }()
	for {
		select {
		case pack := <-packch:
			if err = stm.Send(&pb.PullFrameFromFlowSetResponse{
				Response: pack,
			}); err != nil {
				logger.WithError(err).Errorf("failed to send pack to stream")
				return status.Errorf(codes.Internal, err.Error())
			}
		case err := <-errch:
			logger.WithError(err).Errorf("failed to recv pack response from flowset")
			return status.Errorf(codes.Internal, err.Error())
		case <-dead:
			return status.Errorf(codes.Aborted, "stream closed")
		}
	}
}
