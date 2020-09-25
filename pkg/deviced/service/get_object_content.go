package metathings_deviced_service

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
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
	return self.authorizer.Authorize(ctx, in.(*pb.GetObjectContentRequest).GetObject().GetDevice().GetId().GetValue(), "deviced:get_object_content")
}

func (self *MetathingsDevicedService) GetObjectContent(ctx context.Context, req *pb.GetObjectContentRequest) (*pb.GetObjectContentResponse, error) {
	obj := req.GetObject()
	obj_s := parse_object(obj)

	logger := self.get_logger()

	ch, err := self.simple_storage.GetObjectContent(obj_s)
	if err != nil {
		switch err {
		case os.ErrNotExist:
			logger.WithError(err).Warningf("object not found")
		default:
			logger.WithError(err).Errorf("failed to get object content in simple storage")
		}
		return nil, self.ParseError(err)
	}

	var contents []byte
	for {
		buf, ok := <-ch
		if !ok {
			break
		}
		contents = append(contents, buf...)
	}

	logger = logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	})

	res := &pb.GetObjectContentResponse{Content: contents}

	logger.Debugf("get object content")

	return res, nil
}
