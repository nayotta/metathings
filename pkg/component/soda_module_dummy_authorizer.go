package metathings_component

import (
	"net/http"

	"github.com/stretchr/objx"
)

type SodaModuleDummyAuthorizer struct{}

func (*SodaModuleDummyAuthorizer) Sign(*SodaModuleAuthContext) (*SodaModuleAuthContext, error) {
	return &SodaModuleAuthContext{
		objx.New(map[string]interface{}{}),
	}, nil
}

func (*SodaModuleDummyAuthorizer) Verify(*SodaModuleAuthContext) error {
	return nil
}

func NewSodaModuleDummyAuthorizer(m *Module) (SodaModuleAuthorizer, error) {
	return &SodaModuleDummyAuthorizer{}, nil
}

func parse_http_dummy_auth_context(r *http.Request) (*SodaModuleAuthContext, error) {
	return &SodaModuleAuthContext{nil}, nil
}

func init() {
	register_soda_module_authorizer_factory("dummy", NewSodaModuleDummyAuthorizer)
	register_http_auth_context_parser("dummy", parse_http_dummy_auth_context)
}
