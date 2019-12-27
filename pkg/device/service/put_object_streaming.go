package metathings_device_service

import (
	"io"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
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
		obj := yreq.GetObject()
		obj.Device = self.pb_device()

		y.Request = &deviced_pb.PutObjectStreamingRequest_Metadata_{
			Metadata: &deviced_pb.PutObjectStreamingRequest_Metadata{
				Object: obj,
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
	logger := self.logger.WithField("#method", "PutObjectStreaming")

	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to connect deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	upstm, err := cli.PutObjectStreaming(self.context())
	if err != nil {
		logger.WithError(err).Errorf("failed to put object streaming from deviced service")
		return status.Errorf(codes.Internal, err.Error())
	}

	errs := make(chan error, 1)
	defer func() {
		close(errs)
		errs = nil
	}()

	n2s_quit := make(chan struct{})
	s2n_quit := make(chan struct{})

	go func() {
		defer close(n2s_quit)
		for {
			cres, err := upstm.Recv()
			if err != nil {
				if err != io.EOF && grpc_helper.ExpectCodes(err, codes.Internal) != nil {
					logger.WithError(err).Warningf("failed to receive put object streaming response from deviced service")
					if errs != nil {
						errs <- err
					}
				}
				return
			}

			res := self.parse_put_object_streaming_response(cres)
			if err = stm.Send(res); err != nil {
				logger.WithError(err).Warningf("failed to send put object streaming response to module")
				if errs != nil {
					errs <- err
				}
				return
			}

		}
	}()

	go func() {
		defer close(s2n_quit)
		for {
			req, err := stm.Recv()
			if err != nil {
				if grpc_helper.ExpectCodes(err, codes.Canceled) != nil {
					logger.WithError(err).Warningf("failed to receive put object streaming request from module")
				}
				if errs != nil {
					errs <- err
				}
				return
			}

			creq := self.parse_put_object_streaming_request(req)
			if err = upstm.Send(creq); err != nil {
				logger.WithError(err).Warningf("failed to send put object streaming request to deviced service")
				if errs != nil {
					errs <- err
				}
				return
			}
		}
	}()

	select {
	case <-n2s_quit:
		logger.WithFields(log.Fields{
			"from": "north",
			"to":   "south",
		}).Debugf("stream closed")
	case <-s2n_quit:
		logger.WithFields(log.Fields{
			"from": "south",
			"to":   "north",
		}).Debugf("stream closed")
	case err = <-errs:
		logger.WithError(err).Errorf("failed to file sync")
		return err
	}

	return nil
}
