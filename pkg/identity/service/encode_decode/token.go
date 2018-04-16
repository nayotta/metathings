package encode_decode

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"

	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type _domain struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type tokenResponseBody struct {
	Token struct {
		Project struct {
			Domain _domain `json:"domain"`
			Id     string  `json:"id"`
			Name   string  `json:"name"`
		} `json:"project"`
		User struct {
			Domain _domain `json:"domain"`
			Id     string  `json:"id"`
			Name   string  `json:"name"`
		} `json:"user"`
		Methods []string `json:"methods"`
		Roles   []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"roles"`
		IsDomain  bool   `json:"is_domain"`
		ExipresAt string `json:"exipres_at"`
		IssuedAt  string `json:"issued_at"`

		ApplicationCredentialRestricted bool `json:"application_credential_restricted"`
	} `json:"token"`
}

func decodeToken(_ gorequest.Response, body string) (*pb.Token, error) {
	b := tokenResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	t := b.Token
	token := &pb.Token{
		Methods:  t.Methods,
		IsDomain: t.IsDomain,
		User: &pb.Token__User{
			Id:   t.User.Id,
			Name: t.User.Name,
			Domain: &pb.Token__Domain{
				Id:   t.User.Domain.Id,
				Name: t.User.Domain.Name,
			},
		},
		Roles: []*pb.Token__Role{},
		Project: &pb.Token__Project{
			Id:   t.Project.Id,
			Name: t.Project.Name,
			Domain: &pb.Token__Domain{
				Id:   t.Project.Domain.Id,
				Name: t.Project.Domain.Name,
			},
		},
	}
	for _, r := range t.Roles {
		token.Roles = append(token.Roles, &pb.Token__Role{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	return token, nil
}
