package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	descriptor_storage "github.com/nayotta/metathings/pkg/deviced/descriptor_storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateGetDescriptor(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, descriptor_getter) {
				req := in.(*pb.GetDescriptorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_descriptor_sha1,
		},
	)
}

func (self *MetathingsDevicedService) GetDescriptor(ctx context.Context, req *pb.GetDescriptorRequest) (*pb.GetDescriptorResponse, error) {
	sha1 := req.GetDescriptor_().GetSha1().GetValue()

	logger := self.get_logger().WithField("sha1", sha1)

	body, err := self.desc_storage.GetDescriptor(sha1)
	if err != nil {
		if err == descriptor_storage.ErrDescriptorNotFound {
			logger.WithError(err).Errorf("descriptor sha1 not found")
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		logger.WithError(err).Errorf("failed to get descriptor with sha1 in descriptor storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetDescriptorResponse{
		Descriptor_: &pb.Descriptor{
			Sha1: sha1,
			Body: body,
		},
	}

	logger.Debugf("get descriptor")

	return res, nil
}
