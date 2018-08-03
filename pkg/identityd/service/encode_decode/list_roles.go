package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

func DecodeListRoles(res gorequest.Response, body string) (*pb.ListRolesResponse, error) {
	roles, err := decodeRoles(res, body)
	if err != nil {
		return nil, err
	}
	return &pb.ListRolesResponse{Roles: roles}, nil
}
