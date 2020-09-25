package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ListTimers(ctx context.Context, req *pb.ListTimersRequest) (*pb.ListTimersResponse, error) {
	tmr := req.GetTimer()
	tmr_s := &storage.Timer{}

	logger := srv.get_logger()

	if tmr_id := tmr.GetId(); tmr_id != nil {
		tmr_id_str := tmr_id.GetValue()
		tmr_s.Id = &tmr_id_str
	}

	if tmr_alias := tmr.GetAlias(); tmr_alias != nil {
		tmr_alias_str := tmr_alias.GetValue()
		tmr_s.Alias = &tmr_alias_str
	}

	tmrs_s, err := srv.timer_storage.ListTimers(ctx, tmr_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to list timers in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListTimersResponse{
		Timers: copy_timers(tmrs_s),
	}

	logger.Debugf("list timers")

	return res, nil
}
