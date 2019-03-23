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

func (self *MetathingsDevicedService) ValidateRenameObject(ctx context.Context, in interface{}) error {
	f := func(fn func(*pb.OpObject) error) func(source_getter) error {
		return func(x source_getter) error {
			return fn(x.GetSource())
		}
	}
	g := func(fn func(*pb.OpObject) error) func(destination_getter) error {
		return func(x destination_getter) error {
			return fn(x.GetDestination())
		}
	}

	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, source_getter, destination_getter) {
				req := in.(*pb.RenameObjectRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			f(_ensure_get_object_name),
			f(_ensure_get_object_device_id),
			g(_ensure_get_object_name),
			g(_ensure_get_object_device_id),
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeRenameObject(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.RenameObjectRequest).GetSource().GetDevice().GetId().GetValue(), "rename_object")
}

func (self *MetathingsDevicedService) RenameObject(ctx context.Context, req *pb.RenameObjectRequest) (*empty.Empty, error) {
	src := req.GetSource()
	src_pre_str := src.GetPrefix().GetValue()
	src_name_str := src.GetName().GetValue()
	dst := req.GetDestination()
	dst_pre_str := dst.GetPrefix().GetValue()
	dst_name_str := dst.GetName().GetValue()
	dev_id := src.GetDevice().GetId().GetValue()

	dev_s, err := self.storage.GetDevice(dev_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	src_s := simple_storage.NewObject(src_pre_str, src_name_str, nil)
	dst_s := simple_storage.NewObject(dst_pre_str, dst_name_str, nil)
	err = self.simple_storage.RenameObject(dev_s, src_s, dst_s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"device":      dev_id,
		"source":      path.Join(src_pre_str, src_name_str),
		"destination": path.Join(dst_pre_str, dst_name_str),
	}).Infof("renmae object")

	return &empty.Empty{}, nil
}
