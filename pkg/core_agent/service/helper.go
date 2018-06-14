package metathings_core_agent_service

import (
	"time"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	UNAVAILABLE_TIME = time.Unix(0, 0)
)

func (srv *coreAgentService) handleGRPCError(err error, format string, args ...interface{}) error {
	return grpc_helper.HandleGRPCError(srv.logger, err, format, args...)
}
