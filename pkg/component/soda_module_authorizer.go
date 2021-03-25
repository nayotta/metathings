package metathings_component

import (
	"sync"

	"github.com/stretchr/objx"
)

type SodaModuleAuthContext struct {
	objx.Map
}

type SodaModuleAuthorizer interface {
	Sign(*SodaModuleAuthContext) (*SodaModuleAuthContext, error)
	Verify(*SodaModuleAuthContext) error
}

type SodaModuleAuthorizerFactory func(*Module) (SodaModuleAuthorizer, error)

var soda_module_authorizer_factories map[string]SodaModuleAuthorizerFactory
var soda_module_authorizer_factories_once sync.Once

func register_soda_module_authorizer_factory(name string, fty SodaModuleAuthorizerFactory) {
	soda_module_authorizer_factories_once.Do(func() {
		soda_module_authorizer_factories = make(map[string]SodaModuleAuthorizerFactory)
	})

	soda_module_authorizer_factories[name] = fty
}

func NewSodaModuleAuthorizer(name string, m *Module) (SodaModuleAuthorizer, error) {
	fty, ok := soda_module_authorizer_factories[name]
	if !ok {
		return nil, ErrUnknownSodaModuleAuthorizerDriver
	}

	return fty(m)
}
