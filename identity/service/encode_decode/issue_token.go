package encode_decode

import (
	"context"

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
