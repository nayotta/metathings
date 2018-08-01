package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

func DecodeGetUser(res gorequest.Response, body string) (*pb.GetUserResponse, error) {
	usr, err := decodeUser(res, body)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{User: usr}, nil
}
