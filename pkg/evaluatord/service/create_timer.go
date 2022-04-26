package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	timer "github.com/nayotta/metathings/pkg/evaluatord/timer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateCreateTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() timer_getter {
				req := in.(*pb.CreateTimerRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_timer,
			ensure_timer_id_not_exists(ctx, srv.timer_storage),
			ensure_valid_timer_timezone,
			ensure_valid_timer_schedule,
		},
	)
}

func (srv *MetathingsEvaluatordService) CreateTimer(ctx context.Context, req *pb.CreateTimerRequest) (*pb.CreateTimerResponse, error) {
	var err error

	tmr := req.GetTimer()
	tmr_s := &storage.Timer{}

	tmr_id_str := id_helper.NewId()
	if tmr_id := tmr.GetId(); tmr_id != nil {
		tmr_id_str = tmr_id.GetValue()
	}
	tmr_s.Id = &tmr_id_str

	logger := srv.get_logger().WithField("timer", tmr_id_str)

	tmr_alias_str := tmr_id_str
	if tmr_alias := tmr.GetAlias(); tmr_alias != nil {
		tmr_alias_str = tmr_alias.GetValue()
	}
	tmr_s.Alias = &tmr_alias_str

	tmr_description_str := ""
	if tmr_description := tmr.GetDescription(); tmr_description != nil {
		tmr_description_str = tmr_description.GetValue()
	}
	tmr_s.Description = &tmr_description_str

	tmr_schedule_str := tmr.GetSchedule().GetValue()
	tmr_s.Schedule = &tmr_schedule_str

	tmr_timezone_str := tmr.GetTimezone().GetValue()
	tmr_s.Timezone = &tmr_timezone_str

	tmr_enabled_bool := true
	if tmr_enabled := tmr.GetEnabled(); tmr_enabled != nil {
		tmr_enabled_bool = tmr_enabled.GetValue()
	}
	tmr_s.Enabled = &tmr_enabled_bool

	if tmr_s, err = srv.timer_storage.CreateTimer(ctx, tmr_s); err != nil {
		logger.WithError(err).Errorf("failed to create timer in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if _, err = srv.timer_backend.Create(ctx,
		timer.SetId(tmr_id_str),
		timer.SetSchedule(tmr_schedule_str),
		timer.SetTimezone(tmr_timezone_str),
		timer.SetEnabled(tmr_enabled_bool),
	); err != nil {
		defer srv.timer_storage.DeleteTimer(ctx, tmr_id_str)
		logger.WithError(err).Errorf("failed to create timer in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateTimerResponse{
		Timer: copy_timer(tmr_s),
	}

	logger.Infof("create timer")

	return res, nil
}
