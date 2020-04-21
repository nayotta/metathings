package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ListConfigs(ctx context.Context, req *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	var cfgs_s []*storage.Config
	var err error

	cfg := req.GetConfig()
	cfg_s := &storage.Config{}
	logger := self.logger

	if cfg.GetId().GetValue() != "" {
		cfg_s.Id = &cfg.Id.Value
	}

	if cfg.GetAlias().GetValue() != "" {
		cfg_s.Alias = &cfg.Alias.Value
	}

	if cfgs_s, err = self.storage.ListConfigs(ctx, cfg_s); err != nil {
		logger.WithError(err).Errorf("failed to list configs in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListConfigsResponse{
		Configs: copy_configs(cfgs_s),
	}

	logger.Debugf("list configs")

	return res, nil
}
