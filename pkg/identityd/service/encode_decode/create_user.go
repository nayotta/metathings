package encode_decode

import (
	"context"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
	"github.com/parnurzeal/gorequest"
)

type createUserRequestBody struct {
	User struct {
		Name             string `json:"name"`
		Password         string `json:"password"`
		Enabled          bool   `json:"enabled"`
		DomainId         string `json:"domain_id"`
		DefaultProjectId string `json:"default_project_id"`
	} `json:"user"`
}

func EncodeCreateUser(ctx context.Context, req *pb.CreateUserRequest) (res interface{}, err error) {
	body := &createUserRequestBody{}
	body.User.Name = req.GetName().GetValue()
	body.User.Password = req.GetPassword().GetValue()
	body.User.DomainId = req.GetDomainId().GetValue()

	default_project_id := req.GetDefaultProjectId()
	if default_project_id != nil {
		body.User.DefaultProjectId = default_project_id.GetValue()
	}

	enabled := req.GetEnabled()
	if enabled != nil {
		body.User.Enabled = enabled.GetValue()
	} else {
		body.User.Enabled = true
	}

	return body, nil
}

func DecodeCreateUser(res gorequest.Response, body string) (*pb.CreateUserResponse, error) {
	usr, err := decodeUser(res, body)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: usr}, nil
}
