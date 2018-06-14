package metathings_core_agent_service

import (
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

func (srv *coreAgentService) handleGRPCError(err error, format string, args ...interface{}) error {
	return grpc_helper.HandleGRPCError(srv.logger, err, format, args...)
}
