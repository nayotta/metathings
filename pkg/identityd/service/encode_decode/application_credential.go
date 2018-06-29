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

type _ApplicationCredential struct {
	Id           string
	Name         string
	Secret       string
	Description  string
	ExpiresAt    string `json:"expires_at"`
	ProjectId    string `json:"project_id"`
	Roles        []_applicationCredentialResponseBody_role
	Unrestricted bool
}

type applicationCredentialResponseBody struct {
	ApplicationCredential _ApplicationCredential `json:"application_credential"`
}

type applicationCredentialsResponseBody struct {
	ApplicationCredentials []_ApplicationCredential `json:"application_credentials"`
}

func copyApplicationCredential(ac _ApplicationCredential) *pb.ApplicationCredential {
	t, err := time.Parse(time.RFC3339, ac.ExpiresAt)
	if err != nil {
		t = time.Unix(0, 0)
	}

	pb_ac := &pb.ApplicationCredential{
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
	pb_ac.Roles = rs

	return pb_ac
}

func decodeApplicationCredential(_ gorequest.Response, body string) (*pb.ApplicationCredential, error) {
	b := applicationCredentialResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	app_cred := copyApplicationCredential(b.ApplicationCredential)

	return app_cred, nil
}

func decodeApplicationCredentials(_ gorequest.Response, body string) ([]*pb.ApplicationCredential, error) {
	b := applicationCredentialsResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	acs := b.ApplicationCredentials
	pb_acs := []*pb.ApplicationCredential{}
	for _, ac := range acs {
		pb_acs = append(pb_acs, copyApplicationCredential(ac))
	}

	return pb_acs, nil
}
