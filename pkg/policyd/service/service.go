package metathings_policyd_service

import (
	"context"

	server "github.com/nayotta/metathings/pkg/policyd/casbin-server/server"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

type MetathingsPolicydServiceOption struct {
	AdapterDriver string
	AdapterUri    string
	ModelText     string
}

type MetathingsPolicydService struct {
	*server.Server

	opt *MetathingsPolicydServiceOption
}

func NewMetathingsPolicydService(
	opt *MetathingsPolicydServiceOption,
) (pb.CasbinServer, error) {
	ctx := context.Background()
	srv := &MetathingsPolicydService{
		Server: server.NewServer(),
		opt:    opt,
	}

	new_adapter_res, err := srv.NewAdapter(ctx, &pb.NewAdapterRequest{
		DriverName:    "gorm",
		AdapterName:   opt.AdapterDriver,
		ConnectString: opt.AdapterUri,
	})
	if err != nil {
		return nil, err
	}

	_, err = srv.NewEnforcer(ctx, &pb.NewEnforcerRequest{
		AdapterHandle: new_adapter_res.Handler,
		ModelText:     opt.ModelText,
	})
	if err != nil {
		return nil, err
	}

	return srv, nil
}
