package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identity"
)

func DecodeValidateTokenResponse(req gorequest.Response, body string) (*pb.ValidateTokenResponse, error) {
	token, err := decodeToken(req, body)
	if err != nil {
		return nil, err
	}

	res := &pb.ValidateTokenResponse{}
	res.Token = token

	return res, nil
}
