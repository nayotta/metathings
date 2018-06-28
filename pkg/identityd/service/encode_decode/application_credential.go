package encode_decode

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type _applicationCredentialResponseBody_role struct {
	Id       string
	DomainId string `json:"domain_id"`
	Name     string
}

type applicationCredentialResponseBody struct {
	ApplicationCredential struct {
		Id           string
		Name         string
		Secret       string
		Description  string
		ExpiresAt    string `json:"expires_at"`
		ProjectId    string `json:"project_id"`
		Roles        []_applicationCredentialResponseBody_role
		Unrestricted bool
	} `json:"application_credential"`
}

func decodeApplicationCredential(_ gorequest.Response, body string) (*pb.ApplicationCredential, error) {
	b := applicationCredentialResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	ac := b.ApplicationCredential

	t, err := time.Parse(time.RFC3339, ac.ExpiresAt)
	if err != nil {
		return nil, err
	}

	app_cred := &pb.ApplicationCredential{
		Id:           ac.Id,
		Name:         ac.Name,
		Secret:       ac.Secret,
		Description:  ac.Description,
		ExpiresAt:    &timestamp.Timestamp{Seconds: t.Unix()},
		ProjectId:    ac.ProjectId,
		Unrestricted: ac.Unrestricted,
	}

	rs := []*pb.ApplicationCredential__Role{}
	for _, r := range ac.Roles {
		rs = append(rs, &pb.ApplicationCredential__Role{
			Id:       r.Id,
			Name:     r.Name,
			DomainId: r.DomainId,
		})
	}
	app_cred.Roles = rs

	return app_cred, nil
}
