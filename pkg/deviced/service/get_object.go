package metathings_deviced_service

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateGetObject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() object_getter {
				req := in.(*pb.GetObjectRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_name,
			ensure_get_object_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetObject(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetObjectRequest).GetObject().GetDevice().GetId().GetValue(), "deviced:get_object")
}

func (self *MetathingsDevicedService) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	obj := req.GetObject()
	obj_s := parse_object(obj)

	logger := self.get_logger()

	obj_s, err := self.simple_storage.GetObject(obj_s)
	if err != nil {
		switch err {
		case os.ErrNotExist:
			logger.WithError(err).Warningf("object not found")
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			logger.WithError(err).Errorf("failed to get object in simple storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	logger = logger.WithFields(log.Fields{
		"device": obj_s.Device,
		"object": obj_s.FullName(),
	})

	res := &pb.GetObjectResponse{
		Object: copy_object(obj_s),
	}

	logger.Debugf("get object")

	return res, nil
}
