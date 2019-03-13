package metathings_identityd2_policy

import (
	"context"
	"fmt"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

const (
	CASBIN_BACKEND_SUBJECT_PTYPE = "g2"
	CASBIN_BACKEND_OBJECT_PTYPE  = "g3"
)

type CasbinBackendOption struct {
	EnforcerHandler int32
}

type CasbinBackend struct {
	opt     *CasbinBackendOption
	cli_fty *client_helper.ClientFactory
}

func (cb *CasbinBackend) context() context.Context {
	return context.TODO()
}

func (cb *CasbinBackend) Enforce(sub, obj *storage.Entity, act *storage.Action) (bool, error) {
	panic("unimplemented")
}

func (cb *CasbinBackend) CreateGroup(grp *storage.Group) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) DeleteGroup(grp *storage.Group) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) _add_grouping_policy(cli pb.PolicydServiceClient, g, ent, grp, rol string) error {
	var err error

	req := &pb.PolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           g,
		Params:          []string{ent, grp, rol},
	}
	if _, err = cli.AddNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _remove_grouping_policy(cli pb.PolicydServiceClient, g, ent, grp, rol string) error {
	var err error

	req := &pb.PolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           g,
		Params:          []string{ent, grp, rol},
	}
	if _, err = cli.RemoveNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _list_grouping_policies(cli pb.PolicydServiceClient, g, ent, grp string) ([][]string, error) {
	var err error
	var res *pb.Array2DReply
	var ys [][]string

	req := &pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           g,
		FieldIndex:      0,
		FieldValues:     []string{ent, grp},
	}

	if res, err = cli.GetFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return nil, err
	}

	for _, d2 := range res.GetD2() {
		ys = append(ys, d2.GetD1())
	}

	return ys, nil
}

func (cb *CasbinBackend) _remove_subject_from_group(cli pb.PolicydServiceClient, grp *storage.Group, sub *storage.Entity) error {
	var err error

	req := &pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_SUBJECT_PTYPE,
		FieldIndex:      0,
		FieldValues:     []string{cb.convert_subject(sub), cb.convert_group(grp)},
	}

	if _, err = cli.RemoveFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) convert_group(grp *storage.Group) string {
	return fmt.Sprintf("dom.%s.grp.%s", *grp.DomainId, *grp.Id)
}

func (cb *CasbinBackend) convert_subject(sub *storage.Entity) string {
	return fmt.Sprintf("sub.%s", *sub.Id)
}

func (cb *CasbinBackend) convert_object(obj *storage.Entity) string {
	return fmt.Sprintf("obj.%s", *obj.Id)
}

func (cb *CasbinBackend) convert_role_for_object(grp *storage.Group) string {
	return fmt.Sprintf("dom.%s.grp.%s.rol.%s", *grp.DomainId, *grp.Id, "data")
}

func (cb *CasbinBackend) convert_roles_for_subject(grp *storage.Group) []string {
	var ys []string

	for _, r := range grp.Roles {
		ys = append(ys, fmt.Sprintf("dom.%s.grp.%s.rol.%s", *grp.DomainId, *grp.Id, *r.Id))
	}

	return ys
}

func (cb *CasbinBackend) AddSubjectToGroup(grp *storage.Group, sub *storage.Entity) error {
	var err error

	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	if err = cb._remove_subject_from_group(cli, grp, sub); err != nil {
		return err
	}

	for _, r := range cb.convert_roles_for_subject(grp) {
		if err = cb._add_grouping_policy(cli, CASBIN_BACKEND_SUBJECT_PTYPE, cb.convert_subject(sub), cb.convert_group(grp), r); err != nil {
			return err
		}
	}

	return nil
}

func (cb *CasbinBackend) RemoveSubjectFromGroup(grp *storage.Group, sub *storage.Entity) error {
	var err error

	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	return cb._remove_subject_from_group(cli, grp, sub)
}

func (cb *CasbinBackend) AddObjectToGroup(grp *storage.Group, obj *storage.Entity) error {
	var err error

	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	return cb._add_grouping_policy(cli, CASBIN_BACKEND_OBJECT_PTYPE, cb.convert_object(obj), cb.convert_group(grp), cb.convert_role_for_object(grp))
}

func (cb *CasbinBackend) RemoveObjectFromGroup(grp *storage.Group, obj *storage.Entity) error {
	var err error

	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	return cb._remove_grouping_policy(cli, CASBIN_BACKEND_OBJECT_PTYPE, cb.convert_object(obj), cb.convert_group(grp), cb.convert_role_for_object(grp))
}

func (cb *CasbinBackend) AddRoleToGroup(grp *storage.Group, rol *storage.Role) error {
	panic("unimplemented")
}

func (cb *CasbinBackend) RemoveRoleFromGroup(grp *storage.Group, rol *storage.Role) error {
	panic("unimplemented")
}

func casbin_backend_factory(args ...interface{}) (Backend, error) {
	panic("unimplemented")
}

func init() {
	register_backend_factory("casbin", casbin_backend_factory)
}
