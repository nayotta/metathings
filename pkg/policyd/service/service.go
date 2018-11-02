package metathings_policyd_service

import (
	"context"
	"errors"

	server "github.com/nayotta/metathings/pkg/policyd/casbin-server/server"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

type Policy struct {
	Role   string
	Kind   string
	Action string
}

type MetathingsPolicydServiceOption struct {
	AdapterDriver string
	AdapterUri    string
	ModelText     string
	Policies      []Policy
}

type MetathingsPolicydService struct {
	*server.Server

	opt *MetathingsPolicydServiceOption
}

func (self *MetathingsPolicydService) AddPresetPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	for _, p := range self.opt.Policies {
		if len(in.Params) != 2 {
			return nil, errors.New("bad arguments")
		}

		dom := in.Params[0]
		grp := in.Params[1]

		res, err := self.AddPolicy(ctx, &pb.PolicyRequest{
			EnforcerHandler: in.EnforcerHandler,
			Params:          []string{p.Role, dom, grp, p.Kind, p.Action},
		})
		if err != nil {
			return nil, err
		}
		if !res.Res {
			return &pb.BoolReply{Res: false}, nil
		}
	}

	return &pb.BoolReply{Res: true}, nil
}

func (self *MetathingsPolicydService) RemovePresetPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	// TODO(Peer): RemoveFilteredNamedPolicy should better than RemovePolicy?
	for _, p := range self.opt.Policies {
		if len(in.Params) != 2 {
			return nil, errors.New("bad arguments")
		}

		dom := in.Params[0]
		grp := in.Params[1]

		_, err := self.RemovePolicy(ctx, &pb.PolicyRequest{
			EnforcerHandler: in.EnforcerHandler,
			Params:          []string{p.Role, dom, grp, p.Kind, p.Action},
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.BoolReply{Res: true}, nil
}

func NewMetathingsPolicydService(
	opt *MetathingsPolicydServiceOption,
) (pb.PolicydServiceServer, error) {
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
