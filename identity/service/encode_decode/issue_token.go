package encode_decode

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/bigdatagz/metathings/proto/identity"
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

func DecodeIssueTokenResponse(res http.Response, body string) (*pb.IssueTokenResponse, error) {
	b := issueTokenResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}
	issueTokenRes := &pb.IssueTokenResponse{}
	token := &pb.Token{
		Methods:  b.Methods,
		IsDomain: b.IsDomain,
		User: &pb.Token__User{
			Id:   b.User.Id,
			Name: b.User.Name,
			Domain: &pb.Token__Domain{
				Id:   b.User.Domain.Id,
				Name: b.User.Domain.Name,
			},
		},
		Roles: []*pb.Token__Role{},
		Project: &pb.Token__Project{
			Id:   b.Project.Id,
			Name: b.Project.Name,
			Domain: &pb.Token__Domain{
				Id:   b.Project.Domain.Id,
				Name: b.Project.Domain.Name,
			},
		},
	}
	for _, r := range b.Roles {
		token.Roles = append(token.Roles, &pb.Token__Role{
			Id:   r.Id,
			Name: r.Name,
		})
	}
	issueTokenRes.Token = token

	return issueTokenRes, nil
}
