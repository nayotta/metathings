package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidatePatchConfig(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, config_getter) {
				req := in.(*pb.PatchConfigRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_config_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizePatchConfig(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PatchConfigRequest).GetConfig().GetId().GetValue(), "deviced:patch_config")
}

func (self *MetathingsDevicedService) PatchConfig(ctx context.Context, req *pb.PatchConfigRequest) (*pb.PatchConfigResponse, error) {
	cfg_s := &storage.Config{}
	var err error

	cfg := req.GetConfig()
	cfg_id_str := cfg.GetId().GetValue()
	logger := self.get_logger().WithField("config", cfg_id_str)

	if alias := cfg.GetAlias(); alias != nil {
		cfg_s.Alias = &alias.Value
	}

	if body := cfg.GetBody(); body != nil {
		body_str, err := grpc_helper.JSONPBMarshaler.MarshalToString(body)
		if err != nil {
			logger.WithError(err).Errorf("invalid body")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		cfg_s.Body = &body_str
	}

	if cfg_s, err = self.storage.PatchConfig(ctx, cfg_id_str, cfg_s); err != nil {
		logger.WithError(err).Errorf("failed to patch config in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchConfigResponse{
		Config: copy_config(cfg_s),
	}

	logger.Infof("patch config")

	return res, nil
}
