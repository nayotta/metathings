package encode_decode

import (
	"context"
	"encoding/json"

	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
	"github.com/parnurzeal/gorequest"
)

type issueTokenScopedViaPasswordRequestBody struct {
	Auth struct {
		Identity struct {
			Methods  []string `json:"methods"`
			Password struct {
				User struct {
					Id       string `json:"id,omitempty"`
					Name     string `json:"name,omitempty"`
					Password string `json:"password"`
					Domain   struct {
						Id   string `json:"id,omitempty"`
						Name string `json:"name,omitempty"`
					} `json:"domain,omitempty"`
				} `json:"user"`
			} `json:"password"`
		} `json:"identity"`
		Scope struct {
			Project struct {
				Id string `json:"id"`
			} `json:"project"`
		} `json:"scope"`
	} `json:"auth"`
}

type _domain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type issueTokenResponseBody struct {
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
	} `json:"token"`
}

func EncodeIssueTokenRequest(ctx context.Context, req *pb.IssueTokenRequest) (interface{}, error) {
	if req.GetMethod() != pb.AUTH_METHOD_PASSWORD {
		return nil, Unimplemented
	}

	payload := req.GetPayload().(*pb.IssueTokenRequest_Password).Password
	user_id := payload.GetId()
	username := payload.GetUsername()
	password := payload.GetPassword()
	domain_id := payload.GetDomainId()
	domain_name := payload.GetDomainName()
	scope := payload.GetScope()
	body := issueTokenScopedViaPasswordRequestBody{}

	body.Auth.Identity.Methods = []string{"password"}

	if user_id != nil {
		body.Auth.Identity.Password.User.Id = user_id.Value
	} else if username != nil {
		body.Auth.Identity.Password.User.Name = username.Value
		if domain_id != nil {
			body.Auth.Identity.Password.User.Domain.Id = domain_id.Value
		} else if domain_name != nil {
			body.Auth.Identity.Password.User.Domain.Name = domain_name.Value
		}
	}
	if password != nil {
		body.Auth.Identity.Password.User.Password = password.Value
	}

	if scope != nil {
		project_id := scope.GetProjectId()
		if project_id != nil {
			body.Auth.Scope.Project.Id = project_id.Value
		}
	}

	return &body, nil
}

func DecodeIssueTokenResponse(_ gorequest.Response, body string) (*pb.IssueTokenResponse, error) {
	b := issueTokenResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}
	t := b.Token
	res := &pb.IssueTokenResponse{}
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
	res.Token = token

	return res, nil
}
