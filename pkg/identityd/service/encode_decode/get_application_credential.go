package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

func DecodeGetApplicationCredential(res gorequest.Response, body string) (*pb.GetApplicationCredentialResponse, error) {
	app_cred, err := decodeApplicationCredential(req, body)
	if err != nil {
		return nil, err
	}

	return &pb.GetApplicationCredentialResponse{ApplicationCredential: app_cred}, nil
}
