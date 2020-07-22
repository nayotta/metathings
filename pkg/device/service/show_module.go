package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) ShowModule(ctx context.Context, req *empty.Empty) (*pb.ShowModuleResponse, error) {
	tkn := context_helper.ExtractToken(ctx)
	mdl_id := tkn.GetEntity().GetId()
	logger := self.logger.WithField("module", mdl_id)

	mdl, err := self.get_module_info(mdl_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get module info")
		return nil, self.ParseError(err)

	}
	logger = logger.WithField("module.name", mdl.GetName())

	logger.Debugf("show module")

	return &pb.ShowModuleResponse{Module: mdl}, nil
}
