package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	timer_backend "github.com/nayotta/metathings/pkg/evaluatord/timer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidatePatchTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() timer_getter {
				req := in.(*pb.PatchTimerRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_timer,
			ensure_get_timer_id,
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizePatchTimer(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.PatchTimerRequest).GetTimer().GetId().GetValue(), "evaluatord:patch_timer")
}

func (srv *MetathingsEvaluatordService) PatchTimer(ctx context.Context, req *pb.PatchTimerRequest) (*pb.PatchTimerResponse, error) {
	var err error
	tmr := req.GetTimer()
	tmr_id_str := tmr.GetId().GetValue()
	tmr_s := &storage.Timer{}
	opts := []timer_backend.TimerOption{}

	logger := srv.get_logger().WithField("timer", tmr_id_str)

	if tmr_alias := tmr.GetAlias(); tmr_alias != nil {
		tmr_alias_str := tmr_alias.GetValue()
		tmr_s.Alias = &tmr_alias_str
	}

	if tmr_description := tmr.GetDescription(); tmr_description != nil {
		tmr_description_str := tmr_description.GetValue()
		tmr_s.Description = &tmr_description_str
	}

	if tmr_schedule := tmr.GetSchedule(); tmr_schedule != nil {
		tmr_schedule_str := tmr_schedule.GetValue()
		tmr_s.Schedule = &tmr_schedule_str
		opts = append(opts, timer_backend.SetSchedule(tmr_schedule_str))
	}

	if tmr_timezone := tmr.GetTimezone(); tmr_timezone != nil {
		tmr_timezone_str := tmr_timezone.GetValue()
		tmr_s.Timezone = &tmr_timezone_str
		opts = append(opts, timer_backend.SetTimezone(tmr_timezone_str))
	}

	if tmr_enabled := tmr.GetEnabled(); tmr_enabled != nil {
		tmr_enabled_bool := tmr_enabled.GetValue()
		tmr_s.Enabled = &tmr_enabled_bool
		opts = append(opts, timer_backend.SetEnabled(tmr_enabled_bool))
	}

	if len(opts) != 0 {
		timer_api, err := srv.timer_backend.Get(ctx, tmr_id_str)
		if err != nil {
			logger.WithError(err).Errorf("failed to get timer in backend")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if err = timer_api.Set(ctx, opts...); err != nil {
			logger.WithError(err).Errorf("failed to set timer in backend")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if tmr_s, err = srv.timer_storage.PatchTimer(ctx, tmr_id_str, tmr_s); err != nil {
		logger.WithError(err).Errorf("failed to patch timer in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchTimerResponse{
		Timer: copy_timer(tmr_s),
	}

	logger.Infof("patch timer")

	return res, nil
}
