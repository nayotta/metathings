package metathings_component

type SodaModuleSecretAuthorizer struct {
	m *Module

	secret string
}

func (a *SodaModuleSecretAuthorizer) Authorize(ctx *SodaModuleAuthContext) error {
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
		return nil, ErrRequireSodaModuleAuthorizeSecret
	}

	return &SodaModuleSecretAuthorizer{
		m:      m,
		secret: secret,
	}, nil
}

func init() {
	register_soda_module_authorizer_factory("secret", NewSodaModuleSecretAuthorizer)
}
