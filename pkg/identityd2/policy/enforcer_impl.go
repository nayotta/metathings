package metathings_identityd2_policy

import (
	"context"
	"errors"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

type CasbinEnforcer struct {
	cli_fty *client_helper.ClientFactory
}

func to_strings(x interface{}) ([]string, error) {
	switch y := x.(type) {
	case string:
		return []string{y}, nil
	case []string:
		return y, nil
	default:
		return nil, errors.New("unexpected argument")
	}
}

func (self *CasbinEnforcer) Enforce(domain, group, subject, object, action interface{}) error {
	var doms, grps, subs, objs, acts []string
	var err error

	if doms, err = to_strings(domain); err != nil {
		return err
	}
	if grps, err = to_strings(group); err != nil {
		return err
	}
	if subs, err = to_strings(subject); err != nil {
		return err
	}
	if objs, err = to_strings(object); err != nil {
		return err
	}
	if acts, err = to_strings(action); err != nil {
		return err
	}

	grps = append(grps, UNGROUPED)

	cli, cfn, err := self.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	var reqs []*pb.EnforceRequest
	for _, dom := range doms {
		for _, grp := range grps {
			for _, sub := range subs {
				for _, obj := range objs {
					for _, act := range acts {
						reqs = append(reqs, &pb.EnforceRequest{
							EnforcerHandler: 0,
							Params:          []string{sub, dom, grp, obj, act},
						})
					}
				}
			}
		}
	}

	enforce_bucket_req := &pb.EnforceBucketRequest{Requests: reqs}
	enforce_bucket_res, err := cli.EnforceBucket(context.Background(), enforce_bucket_req)
	if err != nil {
		return err
	}

	if !enforce_bucket_res.Res {
		return ErrPermissionDenied
	}

	return nil
}

func (self *CasbinEnforcer) AddGroup(domain, group string) error {
	cli, cfn, err := self.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	req := &pb.PolicyRequest{
		EnforcerHandler: 0,
		Params:          []string{domain, group},
	}
	if _, err = cli.AddPresetPolicy(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func (self *CasbinEnforcer) RemoveGroup(domain, group string) error {
	cli, cfn, err := self.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	req := &pb.PolicyRequest{
		EnforcerHandler: 0,
		Params:          []string{domain, group},
	}
	if _, err = cli.RemovePresetPolicy(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func (self *CasbinEnforcer) add_grouping_policy(n, x, y string) error {
	cli, cfn, err := self.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	req := &pb.PolicyRequest{
		EnforcerHandler: 0,
		PType:           n,
		Params:          []string{x, y},
	}
	if _, err = cli.AddNamedGroupingPolicy(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func (self *CasbinEnforcer) remove_grouping_policy(n, x, y string) error {
	cli, cfn, err := self.cli_fty.NewPolicydServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	req := &pb.PolicyRequest{
		EnforcerHandler: 0,
		PType:           n,
		Params:          []string{x, y},
	}
	if _, err = cli.RemoveNamedGroupingPolicy(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func (self *CasbinEnforcer) AddSubjectToRole(subject, role string) error {
	return self.add_grouping_policy("g", subject, role)
}

func (self *CasbinEnforcer) RemoveSubjectFromRole(subject, role string) error {
	return self.remove_grouping_policy("g", subject, role)
}

func (self *CasbinEnforcer) AddObjectToKind(object, kind string) error {
	return self.add_grouping_policy("g2", object, kind)
}

func (self *CasbinEnforcer) RemoveObjectFromKind(object, kind string) error {
	return self.remove_grouping_policy("g2", object, kind)
}

func NewEnforcer(cli_fty *client_helper.ClientFactory) Enforcer {
	return &CasbinEnforcer{
		cli_fty: cli_fty,
	}
}
