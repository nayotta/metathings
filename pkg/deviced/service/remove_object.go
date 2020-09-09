package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateRemoveObject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter) {
				req := in.(*pb.RemoveObjectRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_name,
			ensure_get_object_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeRemoveObject(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveObjectRequest).GetObject().GetDevice().GetId().GetValue(), "deviced:remove_object")
}

func (self *MetathingsDevicedService) RemoveObject(ctx context.Context, req *pb.RemoveObjectRequest) (*empty.Empty, error) {
	obj := req.GetObject()
	obj_s := parse_object(obj)

	logger := self.get_logger()

	err := self.simple_storage.RemoveObject(obj_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to remove object from simple storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger = logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	})

	logger.Infof("remove object")

	return &empty.Empty{}, nil
}
