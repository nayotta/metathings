package metathings_deviced_service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateUploadDescriptor(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() descriptor_getter {
				req := in.(*pb.UploadDescriptorRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_descriptor_body,
		},
	)
}

func (self *MetathingsDevicedService) UploadDescriptor(ctx context.Context, req *pb.UploadDescriptorRequest) (*pb.UploadDescriptorResponse, error) {
	var fds dpb.FileDescriptorSet

	body := req.GetDescriptor_().GetBody().GetValue()
	logger := self.get_logger()

	if err := proto.Unmarshal(body, &fds); err != nil {
		logger.WithError(err).Errorf("failed to unmarshal descriptor body")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	fd, err := desc.CreateFileDescriptorFromSet(&fds)
	if err != nil {
		logger.WithError(err).Errorf("failed to create file descriptor from protoset")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	srvs := fd.GetServices()
	if len(srvs) == 0 {
		err = ErrInvalidProtoset
		logger.WithError(err).Errorf("no service in protoset")
		return nil, status.Errorf(codes.InvalidArgument, ErrInvalidProtoset.Error())
	}

	sha1 := fmt.Sprintf("%x", sha1.Sum(body))
	if err = self.desc_storage.SetDescriptor(sha1, body); err != nil {
		logger.WithError(err).Errorf("failed to set descriptor in descriptor storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger = logger.WithField("sha1", sha1)

	res := &pb.UploadDescriptorResponse{
		Descriptor_: &pb.Descriptor{
			Sha1: sha1,
		},
	}

	logger.Infof("upload descriptor")

	return res, nil
}
