package encode_decode

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type _Role struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	DomainId string `json:"domain_id"`
}

type roleResponseBody struct {
	Role _Role `json:"role"`
}

type rolesResponseBody struct {
	Roles []_Role `json:"roles"`
}

func copyRole(role _Role) *pb.Role {
	return &pb.Role{
		Id:       role.Id,
		Name:     role.Name,
		DomainId: role.DomainId,
	}
}

func copyRoles(roles []_Role) []*pb.Role {
	pb_roles := []*pb.Role{}

	for _, role := range roles {
		pb_roles = append(pb_roles, copyRole(role))
	}

	return pb_roles
}

func decodeRole(res gorequest.Response, body string) (*pb.Role, error) {
	b := roleResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	role := copyRole(b.Role)

	return role, nil
}

func decodeRoles(res gorequest.Response, body string) ([]*pb.Role, error) {
	b := rolesResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	roles := copyRoles(b.Roles)

	return roles, nil
}
