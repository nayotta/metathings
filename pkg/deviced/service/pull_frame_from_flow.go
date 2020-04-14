package metathings_deviced_service

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsDevicedService) PullFrameFromFlow(req *pb.PullFrameFromFlowRequest, stm pb.DevicedService_PullFrameFromFlowServer) error {
	cfg := req.GetConfig()
	dev_r := cfg.GetDevice()
	dev_id := dev_r.GetId().GetValue()
	cfg_ack := cfg.GetConfigAck().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"device":  dev_id,
		"#method": "PullFrameFromFlow",
	})

	dev_s, err := self.storage.GetDevice(stm.Context(), dev_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get device")
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
				logger.WithError(err).Errorf("failed to new flow")
				return status.Errorf(codes.Internal, err.Error())
			}
			defer f.Close()

			fs = append(fs, f)

			break match_flow_loop
		}
	}

	running := true
	pack_ch := make(chan *pb.PullFrameFromFlowResponse_Pack_)
	defer close(pack_ch)

	err_ch := make(chan error, 1)
	defer close(err_ch)

	for _, f := range fs {
		go func(f flow.Flow) {
			// TODO(Peer): send multi-frame by one pack
			pf_ch := f.PullFrame()
			flw := &pb.Flow{Id: f.Id()}

			for running {
				select {
				case frm, ok := <-pf_ch:
					if running {
						if !ok && f.Err() != nil {
							err_ch <- f.Err()
							return
						}

						pack_ch <- &pb.PullFrameFromFlowResponse_Pack_{
							Pack: &pb.PullFrameFromFlowResponse_Pack{
								Flow:   flw,
								Frames: []*pb.Frame{frm},
							},
						}
					}
				}
			}
		}(f)
	}

	if cfg_ack {
		err = stm.Send(&pb.PullFrameFromFlowResponse{
			Id: req.GetId().GetValue(),
			Response: &pb.PullFrameFromFlowResponse_Ack_{
				Ack: &pb.PullFrameFromFlowResponse_Ack{},
			}})
		if err != nil {
<<<<<<< HEAD
			self.logger.WithError(err).Errorf("failed to send config ack to stream")
=======
			logger.WithError(err).Errorf("failed to send config ack message")
>>>>>>> v1.1.21.1
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	dead := make(chan struct{})
	go func() {
		defer close(dead)
		for running {
			if err = stm.Send(&pb.PullFrameFromFlowResponse{
				Id: "ffffffffffffffffffffffffffffffff",
				Response: &pb.PullFrameFromFlowResponse_Ack_{
					Ack: &pb.PullFrameFromFlowResponse_Ack{},
				},
			}); err != nil {
				logger.WithError(err).Errorf("failed to send alive response")
				return
			}

			time.Sleep(time.Duration(self.opt.Methods.PullFrameFromFlow.AliveInterval) * time.Second)
		}
		logger.Debugf("alive loop closed")
	}()

	defer func() { running = false }()
	for {
		select {
		case pack := <-pack_ch:
			err = stm.Send(&pb.PullFrameFromFlowResponse{
				Response: pack,
			})
			if err != nil {
<<<<<<< HEAD
				self.logger.WithError(err).Errorf("failed to send pack to stream")
=======
				logger.WithError(err).Errorf("failed to send pack response")
>>>>>>> v1.1.21.1
				return status.Errorf(codes.Internal, err.Error())
			}
		case err := <-err_ch:
			logger.WithError(err).Errorf("failed to recv pack response from flow")
			return status.Errorf(codes.Internal, err.Error())
		case <-dead:
			return status.Errorf(codes.Aborted, "stream closed")
		}
	}
}
