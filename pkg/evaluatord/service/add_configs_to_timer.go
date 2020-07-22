package metathings_evaluatord_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateAddConfigsToTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, timer_getter) {
				req := in.(*pb.AddConfigsToTimerRequest)
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

func (srv *MetathingsEvaluatordService) AuthorizeAddConfigsToTimer(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.AddConfigsToTimerRequest).GetTimer().GetId().GetValue(), "evaluatord:add_configs_to_timer")
}

func (srv *MetathingsEvaluatordService) AddConfigsToTimer(ctx context.Context, req *pb.AddConfigsToTimerRequest) (*empty.Empty, error) {
	timer_id_str := req.GetTimer().GetId().GetValue()
	logger := srv.get_logger().WithField("timer", timer_id_str)

	var cfgs []string
	for _, cfg := range req.GetConfigs() {
		cfgs = append(cfgs, cfg.GetId().GetValue())
	}

	if err := srv.timer_storage.AddConfigsToTimer(ctx, timer_id_str, cfgs); err != nil {
		logger.WithError(err).Errorf("failed to add configs to timer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("add configs to timer")

	return &empty.Empty{}, nil
}
