package metathings_deviced_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateCreateConfig(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, config_getter) {
				req := in.(*pb.CreateConfigRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			func(x config_getter) error {
				cfg := x.GetConfig()

				if cfg.GetAlias() == nil {
					return errors.New("config.alias is empty")
				}

				if cfg.GetBody() == nil {
					return errors.New("config.body is empty")
				}

				return nil
			},
		},
	)
}

func (self *MetathingsDevicedService) CreateConfig(ctx context.Context, req *pb.CreateConfigRequest) (*pb.CreateConfigResponse, error) {
	var err error

	cfg := req.GetConfig()

	cfg_id_str := id_helper.NewId()
	if cfg.GetId() != nil {
		cfg_id_str = cfg.GetId().GetValue()
	}
	cfg_alias_str := cfg.GetAlias().GetValue()
	cfg_body := cfg.GetBody()

	cfg_body_str, err := new(jsonpb.Marshaler).MarshalToString(cfg_body)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to marshal config body to string")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cfg_s := &storage.Config{
		Id:    &cfg_id_str,
		Alias: &cfg_alias_str,
		Body:  &cfg_body_str,
	}

	if cfg_s, err = self.storage.CreateConfig(ctx, cfg_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create config in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateConfigResponse{
		Config: copy_config(cfg_s),
	}

	self.logger.WithField("id", cfg_id_str).Infof("create config")

	return res, nil
}
