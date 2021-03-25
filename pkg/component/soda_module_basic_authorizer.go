package metathings_component

import (
	"encoding/base64"
	"strings"

	"github.com/stretchr/objx"
)

type SodaModuleBasicAuthorizer struct {
	m *Module

	username string
	password string
}

func (a *SodaModuleBasicAuthorizer) Sign(ctx *SodaModuleAuthContext) (*SodaModuleAuthContext, error) {
	var sb strings.Builder
	sb.Write([]byte(a.username))
	sb.Write([]byte(":"))
	sb.Write([]byte(a.password))
	cred := base64.StdEncoding.EncodeToString([]byte(sb.String()))

	return &SodaModuleAuthContext{
		objx.New(map[string]interface{}{
			"scheme":     "Basic",
			"credential": cred,
		}),
	}, nil
}

func (a *SodaModuleBasicAuthorizer) Verify(ctx *SodaModuleAuthContext) error {
	scheme := ctx.Get("scheme").String()
	credential := ctx.Get("credential").String()

	if scheme != "Basic" {
		return ErrUnauthorized
	}

	buf, err := base64.StdEncoding.DecodeString(credential)
	if err != nil {
		return ErrUnauthorized
	}

	ss := strings.SplitN(string(buf), ":", 2)
	if len(ss) != 2 {
		return ErrUnauthorized
	}

	if a.username != ss[0] || a.password != ss[1] {
		return ErrUnauthorized
	}

	return nil
}

func NewSodaModuleBasicAuthorizer(m *Module) (SodaModuleAuthorizer, error) {
	cfg := m.Kernel().Config()

	username := cfg.GetString("backend.request_auth.username")
	if username == "" {
		return nil, ErrRequireSodaModuleAuthorizerUsername
	}

	password := cfg.GetString("backend.request_auth.password")
	if password == "" {
		return nil, ErrRequireSodaModuleAuthorizerPassword
	}

	return &SodaModuleBasicAuthorizer{
		m:        m,
		username: username,
		password: password,
	}, nil
}

func init() {
	register_soda_module_authorizer_factory("basic", NewSodaModuleBasicAuthorizer)
}
