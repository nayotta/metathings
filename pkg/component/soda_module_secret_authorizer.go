package metathings_component

import "github.com/stretchr/objx"

type SodaModuleSecretAuthorizer struct {
	m *Module

	secret string
}

func (a *SodaModuleSecretAuthorizer) Sign(ctx *SodaModuleAuthContext) (*SodaModuleAuthContext, error) {
	return &SodaModuleAuthContext{
		objx.New(map[string]interface{}{
			"scheme":     "Bearer",
			"credential": a.secret,
		}),
	}, nil
}

func (a *SodaModuleSecretAuthorizer) Verify(ctx *SodaModuleAuthContext) error {
	scheme := ctx.Get("scheme").String()
	credential := ctx.Get("credential").String()

	if scheme != "Bearer" || credential != a.secret {
		return ErrUnauthorized
	}

	return nil
}

func NewSodaModuleSecretAuthorizer(m *Module) (SodaModuleAuthorizer, error) {
	cfg := m.Kernel().Config()

	secret := cfg.GetString("backend.auth.secret")
	if secret == "" {
		return nil, ErrRequireSodaModuleAuthorizerSecret
	}

	return &SodaModuleSecretAuthorizer{
		m:      m,
		secret: secret,
	}, nil
}

func init() {
	register_soda_module_authorizer_factory("secret", NewSodaModuleSecretAuthorizer)
}
