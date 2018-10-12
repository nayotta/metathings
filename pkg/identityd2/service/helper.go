package metathings_identityd2_service

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/crypto/bcrypt"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func must_parse_extra(x map[string]*wrappers.StringValue) string {
	var buf []byte
	var err error

	extra_map := pb_helper.ExtractStringMap(x)
	if buf, err = json.Marshal(extra_map); err != nil {
		return `{}`
	}

	return string(buf)
}

func must_parse_password(x string) string {
	buf, _ := bcrypt.GenerateFromPassword([]byte(x), bcrypt.DefaultCost)
	return string(buf)
}

func copy_extra(x string) map[string]string {
	y := map[string]string{}
	if err := json.Unmarshal([]byte(x), &y); err != nil {
		y = map[string]string{}
	}
	return y
}

func copy_domain_children(xs []*storage.Domain) []*pb.Domain {
	ys := []*pb.Domain{}

	for _, x := range xs {
		ys = append(ys, &pb.Domain{
			Id: *x.Id,
		})
	}

	return ys
}

func copy_domain(x *storage.Domain) *pb.Domain {
	y := &pb.Domain{
		Id:    *x.Id,
		Name:  *x.Name,
		Alias: *x.Alias,
		Parent: &pb.Domain{
			Id: *x.Parent.Id,
		},
		Children: copy_domain_children(x.Children),
		Extra:    copy_extra(*x.Extra),
	}

	return y
}

func copy_role_policies(xs []*storage.Policy) []*pb.Policy {
	ys := []*pb.Policy{}

	for _, x := range xs {
		ys = append(ys, &pb.Policy{
			Id: *x.Id,
		})
	}

	return ys
}

func copy_role(x *storage.Role) *pb.Role {
	y := &pb.Role{
		Id: *x.Id,
		Domain: &pb.Domain{
			Id: *x.Domain.Id,
		},
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: *x.Description,
		Policies:    copy_role_policies(x.Policies),
		Extra:       copy_extra(*x.Extra),
	}

	return y
}

func copy_entity(x *storage.Entity) *pb.Entity {
	groups := []*pb.Group{}
	for _, g := range x.Groups {
		groups = append(groups, &pb.Group{
			Id: *g.Id,
		})
	}

	roles := []*pb.Role{}
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	y := &pb.Entity{
		Id: *x.Id,
		Domain: &pb.Domain{
			Id: *x.DomainId,
		},
		Groups: []*pb.Group{},
		Roles:  []*pb.Role{},
		Name:   *x.Name,
		Alias:  *x.Alias,
		Extra:  copy_extra(*x.Extra),
	}

	return y
}

func copy_group(x *storage.Group) *pb.Group {
	roles := []*pb.Role{}
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	entities := []*pb.Entity{}
	for _, e := range x.Entities {
		entities = append(entities, &pb.Entity{
			Id: *e.Id,
		})
	}

	y := &pb.Group{
		Id: *x.Id,
		Domain: &pb.Domain{
			Id: *x.DomainId,
		},
		Roles:       roles,
		Entities:    entities,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: *x.Description,
		Extra:       copy_extra(*x.Extra),
	}

	return y
}

func role_in_entity(ent *storage.Entity, role_id string) bool {
	for _, r := range ent.Roles {
		if *r.Id == role_id {
			return true
		}
	}

	return false
}

func entity_in_group(grp *storage.Group, ent_id string) bool {
	for _, e := range grp.Entities {
		if *e.Id == ent_id {
			return true
		}
	}

	return false
}

func role_in_group(grp *storage.Group, role_id string) bool {
	for _, r := range grp.Roles {
		if *r.Id == role_id {
			return true
		}
	}

	return false
}
