package metathings_policyd_service

import (
	"context"
	"errors"

	casbin_pb "github.com/casbin/casbin-server/proto"
	"github.com/casbin/casbin-server/server"

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

// Casbin server functions
func (self *MetathingsPolicydService) NewEnforcer(ctx context.Context, in *casbin_pb.NewEnforcerRequest) (*casbin_pb.NewEnforcerReply, error) {
	return self.Server.NewEnforcer(ctx, in)
}

func (self *MetathingsPolicydService) NewAdapter(ctx context.Context, in *casbin_pb.NewAdapterRequest) (*casbin_pb.NewAdapterReply, error) {
	return self.Server.NewAdapter(ctx, in)
}

func (self *MetathingsPolicydService) Enforce(ctx context.Context, in *casbin_pb.EnforceRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.Enforce(ctx, in)
}

func (self *MetathingsPolicydService) LoadPolicy(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.EmptyReply, error) {
	return self.Server.LoadPolicy(ctx, in)
}

func (self *MetathingsPolicydService) SavePolicy(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.EmptyReply, error) {
	return self.Server.SavePolicy(ctx, in)
}

func (self *MetathingsPolicydService) AddPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.AddPolicy(ctx, in)
}

func (self *MetathingsPolicydService) AddNamedPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.AddNamedPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemovePolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemovePolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveNamedPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveNamedPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveFilteredPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveFilteredPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveFilteredNamedPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveFilteredNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetPolicy(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetPolicy(ctx, in)
}
func (self *MetathingsPolicydService) GetNamedPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetNamedPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetFilteredPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetFilteredPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetFilteredNamedPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetFilteredNamedPolicy(ctx, in)
}

func (self *MetathingsPolicydService) AddGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.AddGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) AddNamedGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.AddNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveNamedGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveFilteredGroupingPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveFilteredGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) RemoveFilteredNamedGroupingPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.RemoveFilteredNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetGroupingPolicy(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetNamedGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetFilteredGroupingPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetFilteredGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetFilteredNamedGroupingPolicy(ctx context.Context, in *casbin_pb.FilteredPolicyRequest) (*casbin_pb.Array2DReply, error) {
	return self.Server.GetFilteredNamedGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) GetAllSubjects(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllSubjects(ctx, in)
}

func (self *MetathingsPolicydService) GetAllNamedSubjects(ctx context.Context, in *casbin_pb.SimpleGetRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllNamedSubjects(ctx, in)
}

func (self *MetathingsPolicydService) GetAllObjects(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllObjects(ctx, in)
}

func (self *MetathingsPolicydService) GetAllNamedObjects(ctx context.Context, in *casbin_pb.SimpleGetRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllNamedObjects(ctx, in)
}

func (self *MetathingsPolicydService) GetAllActions(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllActions(ctx, in)
}

func (self *MetathingsPolicydService) GetAllNamedActions(ctx context.Context, in *casbin_pb.SimpleGetRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllNamedActions(ctx, in)
}

func (self *MetathingsPolicydService) GetAllRoles(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllRoles(ctx, in)
}

func (self *MetathingsPolicydService) GetAllNamedRoles(ctx context.Context, in *casbin_pb.SimpleGetRequest) (*casbin_pb.ArrayReply, error) {
	return self.Server.GetAllNamedRoles(ctx, in)
}

func (self *MetathingsPolicydService) HasPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.HasPolicy(ctx, in)
}

func (self *MetathingsPolicydService) HasNamedPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.HasNamedPolicy(ctx, in)
}

func (self *MetathingsPolicydService) HasGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.HasGroupingPolicy(ctx, in)
}

func (self *MetathingsPolicydService) HasNamedGroupingPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	return self.Server.HasNamedGroupingPolicy(ctx, in)
}

// Custom functions
func (self *MetathingsPolicydService) EnforceBucket(ctx context.Context, in *pb.EnforceBucketRequest) (*casbin_pb.BoolReply, error) {
	for _, req := range in.Requests {
		res, err := self.Enforce(ctx, req)
		if err != nil {
			return &casbin_pb.BoolReply{Res: false}, err
		} else if res.Res {
			return &casbin_pb.BoolReply{Res: true}, nil
		}
	}

	return &casbin_pb.BoolReply{Res: false}, nil
}

func (self *MetathingsPolicydService) parseEnforceRequest(in *casbin_pb.EnforceRequest) []interface{} {
	params := make([]interface{}, 0, len(in.Params))
	for e := range in.Params {
		params = append(params, in.Params[e])
	}

	return params
}

func (self *MetathingsPolicydService) AddPresetPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	for _, p := range self.opt.Policies {
		if len(in.Params) != 2 {
			return nil, errors.New("bad arguments")
		}

		dom := in.Params[0]
		grp := in.Params[1]

		res, err := self.AddPolicy(ctx, &casbin_pb.PolicyRequest{
			EnforcerHandler: in.EnforcerHandler,
			Params:          []string{p.Role, dom, grp, p.Kind, p.Action},
		})
		if err != nil {
			return nil, err
		}
		if !res.Res {
			return &casbin_pb.BoolReply{Res: false}, nil
		}
	}

	return &casbin_pb.BoolReply{Res: true}, nil
}

func (self *MetathingsPolicydService) RemovePresetPolicy(ctx context.Context, in *casbin_pb.PolicyRequest) (*casbin_pb.BoolReply, error) {
	// TODO(Peer): RemoveFilteredNamedPolicy should better than RemovePolicy?
	for _, p := range self.opt.Policies {
		if len(in.Params) != 2 {
			return nil, errors.New("bad arguments")
		}

		dom := in.Params[0]
		grp := in.Params[1]

		_, err := self.RemovePolicy(ctx, &casbin_pb.PolicyRequest{
			EnforcerHandler: in.EnforcerHandler,
			Params:          []string{p.Role, dom, grp, p.Kind, p.Action},
		})
		if err != nil {
			return nil, err
		}
	}

	return &casbin_pb.BoolReply{Res: true}, nil
}

func (self *MetathingsPolicydService) Initialize(ctx context.Context, in *casbin_pb.EmptyRequest) (*casbin_pb.EmptyReply, error) {
	var err error

	has_policy_req := &casbin_pb.PolicyRequest{
		EnforcerHandler: in.Handler,
		PType:           "p",
		Params:          []string{"rol.sysadmin", "ungrouping", "any", "any"},
	}
	has_policy_res, err := self.HasPolicy(ctx, has_policy_req)
	if err != nil {
		return nil, err
	}

	// WARNING(Peer): `rol.sysadmin` from identityd2.policy.CasbinBackend.convert_ungrouping_role
	// WARNING(Peer): `ungrouping` from identityd2.policy.CASBIN_BACKEND_UNGROUPING_PTYPE
	if !has_policy_res.Res {
		add_policy_req := &casbin_pb.PolicyRequest{
			EnforcerHandler: in.Handler,
			Params:          []string{"rol.sysadmin", "ungrouping", "any", "any"},
		}
		_, err = self.AddPolicy(ctx, add_policy_req)
		if err != nil {
			return nil, err
		}
	}

	return &casbin_pb.EmptyReply{}, nil
}

func NewMetathingsPolicydService(
	opt *MetathingsPolicydServiceOption,
) (pb.PolicydServiceServer, error) {
	ctx := context.Background()
	srv := &MetathingsPolicydService{
		Server: server.NewServer(),
		opt:    opt,
	}

	new_adapter_res, err := srv.NewAdapter(ctx, &casbin_pb.NewAdapterRequest{
		DriverName:    "gorm",
		AdapterName:   opt.AdapterDriver,
		ConnectString: opt.AdapterUri,
	})
	if err != nil {
		return nil, err
	}

	_, err = srv.NewEnforcer(ctx, &casbin_pb.NewEnforcerRequest{
		AdapterHandle: new_adapter_res.Handler,
		ModelText:     opt.ModelText,
	})
	if err != nil {
		return nil, err
	}

	return srv, nil
}
