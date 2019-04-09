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

func (self *MetathingsDevicedService) ValidateGetObjectContent(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter) {
				req := in.(*pb.GetObjectContentRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_name,
			ensure_get_object_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetObjectContent(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetObjectContentRequest).GetObject().GetDevice().GetId().GetValue(), "get_object_content")
}

func (self *MetathingsDevicedService) GetObjectContent(ctx context.Context, req *pb.GetObjectContentRequest) (*pb.GetObjectContentResponse, error) {
	obj := req.GetObject()
	obj_s := parse_object(obj)

	ch, err := self.simple_storage.GetObjectContent(obj_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get object content in simple storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var contents []byte
	for {
		buf, ok := <-ch
		if !ok {
			break
		}
		contents = append(contents, buf...)
	}

	res := &pb.GetObjectContentResponse{Content: contents}

	self.logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	}).Debugf("get object content")

	return res, nil
}
