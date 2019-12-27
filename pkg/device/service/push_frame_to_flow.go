package metathings_device_service

import (
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) PushFrameToFlow(stm pb.DeviceService_PushFrameToFlowServer) error {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	upstm, err := cli.PushFrameToFlow(self.context())
	if err != nil {
		self.logger.WithError(err).Errorf("failed to push frame to flow from deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			cres, err := upstm.Recv()
			if err != nil {
				self.logger.WithError(err).Warningf("failed to receive push frame to flow response from deviced service")
				return
			}

			res := &pb.PushFrameToFlowResponse{
				Id:       cres.GetId(),
				Response: &pb.PushFrameToFlowResponse_Ack_{Ack: &pb.PushFrameToFlowResponse_Ack{}},
			}

			err = stm.Send(res)
			if err != nil {
				self.logger.WithError(err).Warningf("failed to send push frame to flow response to module")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			req, err := stm.Recv()
			if err != nil {
				self.logger.WithError(err).Warningf("failed to receive push frame to flow request from module")
				return
			}

			creq := &deviced_pb.PushFrameToFlowRequest{
				Id: req.GetId(),
			}

			switch req.Request.(type) {
			case *pb.PushFrameToFlowRequest_Config_:
				cfg := req.GetConfig()
				dev := self.pb_device()
				dev.Flows = []*deviced_pb.OpFlow{cfg.GetFlow()}
				creq.Request = &deviced_pb.PushFrameToFlowRequest_Config_{
					Config: &deviced_pb.PushFrameToFlowRequest_Config{
						Device:    dev,
						ConfigAck: cfg.GetConfigAck(),
						PushAck:   cfg.GetPushAck(),
					},
				}

			case *pb.PushFrameToFlowRequest_Frame:
				creq.Request = &deviced_pb.PushFrameToFlowRequest_Frame{
					Frame: req.GetFrame(),
				}
			}

			err = upstm.Send(creq)
			if err != nil {
				self.logger.WithError(err).Warningf("failed to send push frame to flow request to deviced service")
				return
			}
		}
	}()

	wg.Wait()
	return nil
}
