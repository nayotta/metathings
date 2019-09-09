package metathings_deviced_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateListObjects(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter) {
				req := in.(*pb.ListObjectsRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeListObjects(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.ListObjectsRequest).GetObject().GetDevice().GetId().GetValue(), "deviced:list_objects")
}

func (self *MetathingsDevicedService) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.ListObjectsResponse, error) {
	obj := req.GetObject()
	obj_s := parse_object(obj)

	objs_s, err := self.simple_storage.ListObjects(obj_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list objects in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListObjectsResponse{
		Objects: copy_objects(objs_s),
	}

	self.logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	}).Debugf("list objects")

	return res, nil
}
