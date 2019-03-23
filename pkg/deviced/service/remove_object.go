package metathings_deviced_service

import (
	"context"
	"path"

	"github.com/golang/protobuf/ptypes/empty"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return self.authorizer.Authorize(ctx, in.(*pb.RemoveObjectRequest).GetObject().GetDevice().GetId().GetValue(), "remove_object")
}

func (self *MetathingsDevicedService) RemoveObject(ctx context.Context, req *pb.RemoveObjectRequest) (*empty.Empty, error) {
	obj := req.GetObject()
	dev_id := obj.GetDevice().GetId().GetValue()
	obj_pre_str := obj.GetPrefix().GetValue()
	obj_name_str := obj.GetName().GetValue()
	obj_s := simple_storage.NewObject(obj_pre_str, obj_name_str, nil)

	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.simple_storage.RemoveObject(dev_s, obj_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to remove object from simple storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"device": dev_id,
		"object": path.Join(obj_pre_str, obj_name_str),
	}).Infof("remove object")

	return &empty.Empty{}, nil
}
