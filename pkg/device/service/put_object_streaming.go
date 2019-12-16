package metathings_device_service

import (
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) parse_put_object_streaming_response(x *deviced_pb.PutObjectStreamingResponse) *pb.PutObjectStreamingResponse {
	y := &pb.PutObjectStreamingResponse{
		Id: x.GetId(),
		Response: &pb.PutObjectStreamingResponse_Chunks{
			Chunks: x.GetChunks(),
		},
	}

	return y
}

func (self *MetathingsDeviceServiceImpl) parse_put_object_streaming_request(x *pb.PutObjectStreamingRequest) *deviced_pb.PutObjectStreamingRequest {
	y := &deviced_pb.PutObjectStreamingRequest{
		Id: x.GetId(),
	}

	switch x.Request.(type) {
	case *pb.PutObjectStreamingRequest_Metadata_:
		yreq := x.GetMetadata()
		y.Request = &deviced_pb.PutObjectStreamingRequest_Metadata_{
			Metadata: &deviced_pb.PutObjectStreamingRequest_Metadata{
				Object: yreq.GetObject(),
				Sha1:   yreq.GetSha1(),
			},
		}
	case *pb.PutObjectStreamingRequest_Chunks:
		yreq := x.GetChunks()
		y.Request = &deviced_pb.PutObjectStreamingRequest_Chunks{
			Chunks: yreq,
		}
	}

	return y
}

func (self *MetathingsDeviceServiceImpl) PutObjectStreaming(stm pb.DeviceService_PutObjectStreamingServer) error {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	upstm, err := cli.PutObjectStreaming(self.context())
	if err != nil {
		self.logger.WithError(err).Errorf("failed to put object streaming from deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			cres, err := upstm.Recv()
			if err != nil {
				self.logger.WithError(err).Warningf("failed to receive put object streaming response from deviced service")
				return
			}

			res := self.parse_put_object_streaming_response(cres)
			if err = stm.Send(res); err != nil {
				self.logger.WithError(err).Warningf("failed to send put object streaming response to module")
				return
			}

		}
	}()

	go func() {
		defer wg.Done()
		for {
			req, err := stm.Recv()
			if err != nil {
				self.logger.WithError(err).Warningf("failed to receive put object streaming request from module")
				return
			}

			creq := self.parse_put_object_streaming_request(req)
			if err = upstm.Send(creq); err != nil {
				self.logger.WithError(err).Warningf("failed to send put object streaming request to deviced service")
			}
		}
	}()

	wg.Wait()
	return nil
}
