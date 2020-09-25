package metathings_evaluatord_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateDeleteTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, timer_getter) {
				req := in.(*pb.DeleteTimerRequest)
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

func (srv *MetathingsEvaluatordService) AuthorizeDeleteTimer(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.DeleteTimerRequest).GetTimer().GetId().GetValue(), "evaluatord:delete_timer")
}

func (srv *MetathingsEvaluatordService) DeleteTimer(ctx context.Context, req *pb.DeleteTimerRequest) (_ *empty.Empty, err error) {
	tmr_id_str := req.GetTimer().GetId().GetValue()
	logger := srv.get_logger().WithField("timer", tmr_id_str)

	tmr_api, err := srv.timer_backend.Get(ctx, tmr_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to get timer api in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = tmr_api.Delete(ctx); err != nil {
		logger.WithError(err).Errorf("failed to delete timer in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = srv.timer_storage.DeleteTimer(ctx, tmr_id_str); err != nil {
		logger.WithError(err).Errorf("failed to delete timer in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("delete timer")

	return &empty.Empty{}, nil
}
