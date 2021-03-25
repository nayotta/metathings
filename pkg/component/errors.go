package metathings_component

import (
	"errors"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	ErrUnknownSodaModuleWrapperDriver      = errors.New("unknown soda module wrapper driver")
	ErrUnknownSodaModuleBackendDriver      = errors.New("unknown soda module backend driver")
	ErrUnknownSodaModuleAuthorizerDriver   = errors.New("unknown soda module authorizer driver")
	ErrRequireSodaModuleAuthorizerSecret   = errors.New("require soda module authorizer secret")
	ErrRequireSodaModuleAuthorizerUsername = errors.New("require soda module authorizer username")
	ErrRequireSodaModuleAuthorizerPassword = errors.New("require soda module authorizer password")
	ErrUnauthorized                        = errors.New("unauthorized")
	ErrUnexpectedTokenFormat               = errors.New("unexpected token format")
	ErrUnaryCallTimeout                    = errors.New("unary call timeout")
	ErrStreamCallConfigAckTimeout          = errors.New("stream call config ack timeout")
	ErrStreamCallConfig                    = errors.New("stream call config error")
	ErrUnexceptedResponse                  = errors.New("unexpected response")
	ErrBadScheme                           = errors.New("bad scheme")
	ErrBadServiceEndpoint                  = errors.New("bad service endpoint")
	ErrDefaultAddressRequired              = errors.New("default address required")
	ErrDeviceAddressRequired               = errors.New("device address required")
	ErrInvalidArguments                    = errors.New("invalid arguments")
	ErrSubscribedSession                   = errors.New("subscribed session")
	ErrUnsubscribedSession                 = errors.New("unsubscribed session")
	ErrStartTimeout                        = errors.New("start timeout")
	ErrDownstreamNotFound                  = errors.New("downstream not found")
)

var sodaModuleErrorMapping = grpc_helper.ErrorMapping{}
