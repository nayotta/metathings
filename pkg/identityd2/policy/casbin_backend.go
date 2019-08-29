package metathings_identityd2_policy

import (
	"context"

	casbin_pb "github.com/casbin/casbin-server/proto"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

const (
	CASBIN_BACKEND_DEFAULT_ENFORCER_HANDLER = 0
	CASBIN_BACKEND_POLICY_PTYPE             = "p"
	CASBIN_BACKEND_UNGROUPING_PTYPE         = "g2"
	CASBIN_BACKEND_SUBJECT_PTYPE            = "g2"
	CASBIN_BACKEND_OBJECT_PTYPE             = "g3"
	CASBIN_BACKEND_UNGROUPING               = "ungrouping"
)

type CasbinBackendOption struct {
	EnforcerHandler int32
}

func new_casbin_backend_option() *CasbinBackendOption {
	return &CasbinBackendOption{
		EnforcerHandler: CASBIN_BACKEND_DEFAULT_ENFORCER_HANDLER,
	}
}

type CasbinBackend struct {
	opt     *CasbinBackendOption
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
}

func (cb *CasbinBackend) context() context.Context {
	return context.TODO()
}

func (cb *CasbinBackend) _add_grouping_policy(cli pb.PolicydServiceClient, g, ent, grp, rol string) error {
	var err error

	req := &casbin_pb.PolicyRequest{
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

	req := &casbin_pb.PolicyRequest{
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
	var res *casbin_pb.Array2DReply
	var ys [][]string

	req := &casbin_pb.FilteredPolicyRequest{
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

func (cb *CasbinBackend) _list_policies(cli pb.PolicydServiceClient, p, rol, grp string) ([][]string, error) {
	var err error
	var res *casbin_pb.Array2DReply
	var ys [][]string

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           p,
		FieldIndex:      0,
		FieldValues:     []string{rol, grp},
	}
	if res, err = cli.GetFilteredNamedPolicy(cb.context(), req); err != nil {
		return nil, err
	}

	for _, d2 := range res.GetD2() {
		ys = append(ys, d2.GetD1())
	}

	return ys, nil
}

func (cb *CasbinBackend) _remove_subject_from_group(cli pb.PolicydServiceClient, grp *storage.Group, sub *storage.Entity) error {
	var err error

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_SUBJECT_PTYPE,
		FieldIndex:      0,
		FieldValues:     []string{ConvertSubject(sub), ConvertGroup(grp)},
	}

	if _, err = cli.RemoveFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _add_role_to_group(cli pb.PolicydServiceClient, grp *storage.Group, rol *storage.Role) error {
	var err error

	sub_rol_s := ConvertRole(grp, rol)
	grp_s := ConvertGroup(grp)
	obj_rol_s := ConvertRoleForObject(grp)

	for _, act := range rol.Actions {
		req := &casbin_pb.PolicyRequest{
			EnforcerHandler: cb.opt.EnforcerHandler,
			PType:           CASBIN_BACKEND_POLICY_PTYPE,
			Params:          []string{sub_rol_s, grp_s, obj_rol_s, *act.Name},
		}
		if _, err = cli.AddNamedPolicy(cb.context(), req); err != nil {
			cb._remove_role_from_group(cli, grp, rol)
			return err
		}
	}

	for _, sub := range grp.Subjects {
		if err = cb._add_grouping_policy(cli, CASBIN_BACKEND_SUBJECT_PTYPE, ConvertSubject(sub), ConvertGroup(grp), sub_rol_s); err != nil {
			return err
		}
	}

	return nil
}

func (cb *CasbinBackend) _remove_role_from_group(cli pb.PolicydServiceClient, grp *storage.Group, rol *storage.Role) error {
	var err error

	sub_rol_s := ConvertRole(grp, rol)
	grp_s := ConvertGroup(grp)

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_POLICY_PTYPE,
		FieldIndex:      0,
		FieldValues:     []string{sub_rol_s, grp_s},
	}

	if _, err = cli.RemoveFilteredNamedPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _remove_group_about_subject(cli pb.PolicydServiceClient, grp *storage.Group) error {
	var err error

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_SUBJECT_PTYPE,
		FieldIndex:      1,
		FieldValues:     []string{ConvertGroup(grp)},
	}

	if _, err = cli.RemoveFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _remove_group_about_object(cli pb.PolicydServiceClient, grp *storage.Group) error {
	var err error

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_OBJECT_PTYPE,
		FieldIndex:      1,
		FieldValues:     []string{ConvertGroup(grp)},
	}

	if _, err = cli.RemoveFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _remove_group_about_policy(cli pb.PolicydServiceClient, grp *storage.Group) error {
	var err error

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_POLICY_PTYPE,
		FieldIndex:      1,
		FieldValues:     []string{ConvertGroup(grp)},
	}

	if _, err = cli.RemoveFilteredNamedPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _add_role_to_entity(cli pb.PolicydServiceClient, ent *storage.Entity, rol *storage.Role) error {
	var err error

	req := &casbin_pb.PolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_UNGROUPING_PTYPE,
		Params:          []string{ConvertEntity(ent), CASBIN_BACKEND_UNGROUPING, ConvertUngroupingRole(rol)},
	}
	if _, err = cli.AddNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _remove_role_from_entity(cli pb.PolicydServiceClient, ent *storage.Entity, rol *storage.Role) error {
	var err error

	req := &casbin_pb.FilteredPolicyRequest{
		EnforcerHandler: cb.opt.EnforcerHandler,
		PType:           CASBIN_BACKEND_UNGROUPING_PTYPE,
		FieldIndex:      0,
		FieldValues:     []string{ConvertEntity(ent), CASBIN_BACKEND_UNGROUPING, ConvertUngroupingRole(rol)},
	}
	if _, err = cli.RemoveFilteredNamedGroupingPolicy(cb.context(), req); err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) _enforce(cli pb.PolicydServiceClient, sub, obj *storage.Entity, act *storage.Action) error {
	var err error

	sub_s := ConvertSubject(sub)
	obj_s := ConvertObject(obj)

	reqs := []*casbin_pb.EnforceRequest{
		&casbin_pb.EnforceRequest{
			EnforcerHandler: cb.opt.EnforcerHandler,
			Params:          []string{ConvertEntity(sub), CASBIN_BACKEND_UNGROUPING, obj_s, *act.Name},
		},
	}
	for _, grp := range sub.Groups {
		grp_s := ConvertGroup(grp)

		reqs = append(reqs, &casbin_pb.EnforceRequest{
			EnforcerHandler: cb.opt.EnforcerHandler,
			Params:          []string{sub_s, grp_s, obj_s, *act.Name},
		})
	}

	req := &pb.EnforceBucketRequest{Requests: reqs}
	res, err := cli.EnforceBucket(cb.context(), req)
	if err != nil {
		return err
	}

	if !res.Res {
		return ErrPermissionDenied
	}

	return nil
}

func (cb *CasbinBackend) Enforce(sub, obj *storage.Entity, act *storage.Action) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._enforce(cli, sub, obj, act)
	if err != nil {
		return err
	}

	return nil
}

func (cb *CasbinBackend) CreateGroup(grp *storage.Group) error {
	cb.logger.WithField("group", *grp.Id).Debugf("create group")

	return nil
}

func (cb *CasbinBackend) DeleteGroup(grp *storage.Group) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	if err = cb._remove_group_about_subject(cli, grp); err != nil {
		return err
	}

	if err = cb._remove_group_about_object(cli, grp); err != nil {
		return err
	}

	if err = cb._remove_group_about_policy(cli, grp); err != nil {
		return err
	}

	cb.logger.WithField("group", *grp.Id).Debugf("delete group")

	return nil
}

func (cb *CasbinBackend) AddSubjectToGroup(grp *storage.Group, sub *storage.Entity) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	if err = cb._remove_subject_from_group(cli, grp, sub); err != nil {
		return err
	}

	for _, r := range ConvertRolesForSubject(grp) {
		if err = cb._add_grouping_policy(cli, CASBIN_BACKEND_SUBJECT_PTYPE, ConvertSubject(sub), ConvertGroup(grp), r); err != nil {
			return err
		}
	}

	cb.logger.WithFields(log.Fields{
		"group":   *grp.Id,
		"subject": *sub.Id,
	}).Debugf("add subject to group")

	return nil
}

func (cb *CasbinBackend) RemoveSubjectFromGroup(grp *storage.Group, sub *storage.Entity) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._remove_subject_from_group(cli, grp, sub)
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"group":   *grp.Id,
		"subject": *sub.Id,
	}).Debugf("remove subject from group")

	return nil
}

func (cb *CasbinBackend) AddObjectToGroup(grp *storage.Group, obj *storage.Entity) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._add_grouping_policy(cli, CASBIN_BACKEND_OBJECT_PTYPE, ConvertObject(obj), ConvertGroup(grp), ConvertRoleForObject(grp))
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"group":  *grp.Id,
		"object": *obj.Id,
	}).Debugf("add object to group")

	return nil
}

func (cb *CasbinBackend) RemoveObjectFromGroup(grp *storage.Group, obj *storage.Entity) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._remove_grouping_policy(cli, CASBIN_BACKEND_OBJECT_PTYPE, ConvertObject(obj), ConvertGroup(grp), ConvertRoleForObject(grp))
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"group":  *grp.Id,
		"object": *obj.Id,
	}).Debugf("remove object from group")

	return nil
}

func (cb *CasbinBackend) AddRoleToGroup(grp *storage.Group, rol *storage.Role) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._add_role_to_group(cli, grp, rol)
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"group": *grp.Id,
		"role":  *grp.Id,
	}).Debugf("add role to group")

	return nil
}

func (cb *CasbinBackend) RemoveRoleFromGroup(grp *storage.Group, rol *storage.Role) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._remove_role_from_group(cli, grp, rol)
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"group": *grp.Id,
		"role":  *rol.Name,
	}).Debugf("remove role from group")

	return nil
}

func (cb *CasbinBackend) AddRoleToEntity(ent *storage.Entity, rol *storage.Role) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._add_role_to_entity(cli, ent, rol)
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"entity": *ent.Id,
		"role":   *rol.Name,
	}).Debugf("add role to entity")

	return nil
}

func (cb *CasbinBackend) RemoveRoleFromEntity(ent *storage.Entity, rol *storage.Role) error {
	cli, cfn, err := cb.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	err = cb._remove_role_from_entity(cli, ent, rol)
	if err != nil {
		return err
	}

	cb.logger.WithFields(log.Fields{
		"entity": *ent.Id,
		"role":   *rol.Name,
	}).Debugf("remove role from entity")

	return nil
}

func casbin_backend_factory(args ...interface{}) (Backend, error) {
	var ok bool
	var logger log.FieldLogger
	var cli_fty *client_helper.ClientFactory
	opt := new_casbin_backend_option()

	err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
		"client_factory": func(key string, val interface{}) error {
			cli_fty, ok = val.(*client_helper.ClientFactory)
			if !ok {
				return opt_helper.InvalidArgument("client_factory")
			}
			return nil
		},
		"casbin_enforcer_handler": opt_helper.ToInt32(&opt.EnforcerHandler),
	})(args...)
	if err != nil {
		return nil, err
	}

	cli, cfn, err := cli_fty.NewPolicydServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	b := &CasbinBackend{
		opt:     opt,
		cli_fty: cli_fty,
		logger:  logger,
	}

	if _, err = cli.Initialize(b.context(), &casbin_pb.EmptyRequest{Handler: opt.EnforcerHandler}); err != nil {
		return nil, err
	}

	return b, nil
}

func init() {
	register_backend_factory("casbin", casbin_backend_factory)
}
