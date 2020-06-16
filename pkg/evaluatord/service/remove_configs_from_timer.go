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

func (srv *MetathingsEvaluatordService) ValidateRemoveConfigsFromTimer(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, timer_getter) {
				req := in.(*pb.RemoveConfigsFromTimerRequest)
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

func (srv *MetathingsEvaluatordService) AuthorizeRemoveConfigsFromTimer(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.RemoveConfigsFromTimerRequest).GetTimer().GetId().GetValue(), "evaluatord:remove_configs_from_timer")
}

func (srv *MetathingsEvaluatordService) RemoveConfigsFromTimer(ctx context.Context, req *pb.RemoveConfigsFromTimerRequest) (*empty.Empty, error) {
	timer_id_str := req.GetTimer().GetId().GetValue()
	logger := srv.get_logger().WithField("timer", timer_id_str)

	var cfgs []string
	for _, cfg := range req.GetConfigs() {
		cfgs = append(cfgs, cfg.GetId().GetValue())
	}

	if err := srv.timer_storage.RemoveConfigsFromTimer(ctx, timer_id_str, cfgs); err != nil {
		logger.WithError(err).Errorf("failed to remove configs from timer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("remove configs from timer")

	return &empty.Empty{}, nil
}
