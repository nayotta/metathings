package metathings_deviced_service

import (
	"bytes"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidatePutObject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter) {
				req := in.(*pb.PutObjectRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_name,
			ensure_get_object_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizePutObject(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PutObjectRequest).GetObject().GetDevice().GetId().GetValue(), "put_object")
}

func (self *MetathingsDevicedService) PutObject(ctx context.Context, req *pb.PutObjectRequest) (*empty.Empty, error) {
	dev_id := req.GetObject().GetDevice().GetId().GetValue()
	obj := req.GetObject()
	obj_pre_str := obj.GetPrefix().GetValue()
	obj_name_str := obj.GetName().GetValue()
	obj_md := pb_helper.ExtractStringMapToString(obj.GetMetadata())
	content := req.GetContent().GetValue()
	reader := bytes.NewReader(content)
	obj_s := simple_storage.NewObject(obj_pre_str, obj_name_str, obj_md)

	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.simple_storage.PutObject(dev_s, obj_s, reader)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to put object to simple storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"device":      dev_id,
		"object.name": obj_name_str,
	}).Infof("put object")

	return &empty.Empty{}, nil
}
