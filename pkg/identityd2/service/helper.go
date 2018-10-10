package metathings_identityd2_service

import (
	"encoding/json"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

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
