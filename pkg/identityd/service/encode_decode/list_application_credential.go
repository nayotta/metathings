package encode_decode

import (
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

func DecodeListApplicationCredential(req gorequest.Response, body string) (*pb.ListApplicationCredentialsResponse, error) {
	app_creds, err := decodeApplicationCredentials(req, body)
	if err != nil {
		return nil, err
	}

	res := &pb.ListApplicationCredentialsResponse{
		ApplicationCredentials: app_creds,
	}

	return res, nil
}
