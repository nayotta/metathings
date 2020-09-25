package metathings_deviced_service

import (
	"bytes"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
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
	return self.authorizer.Authorize(ctx, in.(*pb.PutObjectRequest).GetObject().GetDevice().GetId().GetValue(), "deviced:put_object")
}

func (self *MetathingsDevicedService) PutObject(ctx context.Context, req *pb.PutObjectRequest) (*empty.Empty, error) {
	obj := req.GetObject()
	content := req.GetContent().GetValue()
	reader := bytes.NewReader(content)
	obj_s := parse_object(obj)

	logger := self.get_logger()

	err := self.simple_storage.PutObject(obj_s, reader)
	if err != nil {
		logger.WithError(err).Errorf("failed to put object to simple storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger = logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	})

	logger.Debugf("put object")

	return &empty.Empty{}, nil
}
