package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateGetTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, timer_getter) {
				req := in.(*pb.GetTimerRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_timer,
			ensure_get_timer_id,
			ensure_timer_id_exists(ctx, srv.timer_storage),
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeGetTimer(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.GetTimerRequest).GetTimer().GetId().GetValue(), "evaluatord:get_timer")
}

func (srv *MetathingsEvaluatordService) GetTimer(ctx context.Context, req *pb.GetTimerRequest) (res *pb.GetTimerResponse, err error) {
	var tmr_s *storage.Timer

	tmr_id_str := req.GetTimer().GetId().GetValue()
	logger := srv.get_logger().WithField("timer", tmr_id_str)

	if tmr_s, err = srv.timer_storage.GetTimer(ctx, tmr_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get timer in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.GetTimerResponse{
		Timer: copy_timer(tmr_s),
	}

	logger.Debugf("get timer")

	return res, nil
}
