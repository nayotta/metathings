package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

func DecodeListUsers(res gorequest.Response, body string) (*pb.ListUsersResponse, error) {
	usrs, err := decodeUsers(res, body)
	if err != nil {
		return nil, err
	}
	return &pb.ListUsersResponse{Users: usrs}, nil
}
