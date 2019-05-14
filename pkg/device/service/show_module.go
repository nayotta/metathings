package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) ShowModule(ctx context.Context, req *empty.Empty) (*pb.ShowModuleResponse, error) {
	tkn := context_helper.ExtractToken(ctx)

	mdl_id := tkn.GetEntity().GetId()

	for _, m := range self.info.Modules {
		if m.Id == mdl_id {
			res := &pb.ShowModuleResponse{
				Module: m,
			}

			self.logger.WithField("module", mdl_id).Debugf("show module")

			return res, nil
		}
	}

	err := ErrModuleNotFound
	self.logger.WithError(err).Errorf("failed to found module")

	return nil, status.Errorf(codes.NotFound, err.Error())
}
