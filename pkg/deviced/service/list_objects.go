package metathings_deviced_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateListObjects(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() object_getter {
				req := in.(*pb.ListObjectsRequest)
				return req
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
	opt := &simple_storage.ListObjectsOption{}

	logger := self.get_logger()

	if recur := req.GetRecursive(); recur != nil {
		opt.Recursive = recur.GetValue()
	}

	if depth := req.GetDepth(); depth != nil {
		opt.Depth = int(depth.GetValue())
	}

	if opt.Recursive && opt.Depth <= 0 {
		opt.Depth = 16
	}

	objs_s, err := self.simple_storage.ListObjects(obj_s, opt)
	if err != nil {
		logger.WithError(err).Errorf("failed to list objects in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger = logger.WithFields(log.Fields{
		"device":    obj_s.Device,
		"object":    obj_s.FullName(),
		"recursive": opt.Recursive,
		"depth":     opt.Depth,
	})

	res := &pb.ListObjectsResponse{
		Objects: copy_objects(objs_s),
	}

	logger.Debugf("list objects")

	return res, nil
}
