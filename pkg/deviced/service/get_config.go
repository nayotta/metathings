package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

func (self *MetathingsDevicedService) ValidateGetConfig(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, config_getter) {
				req := in.(*pb.GetConfigRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_config_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetConfig(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetConfigRequest).GetConfig().GetId().GetValue(), "deviced:get_config")
}

func (self *MetathingsDevicedService) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	var cfg_s *storage.Config
	var err error

	cfg_id_str := req.GetConfig().GetId().GetValue()
	logger := self.logger.WithFields(log.Fields{
		"config": cfg_id_str,
	})

	if cfg_s, err = self.storage.GetConfig(ctx, cfg_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get config in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetConfigResponse{
		Config: copy_config(cfg_s),
	}

	logger.Debugf("get config")

	return res, nil
}
