package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateGetEvaluator(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, evaluator_getter) {
				req := in.(*pb.GetEvaluatorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_evaluator,
			ensure_get_evaluator_id,
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeGetEvaluator(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.GetEvaluatorRequest).GetEvaluator().GetId().GetValue(), "evaluatord:get_evaluator")
}

func (srv *MetathingsEvaluatordService) GetEvaluator(ctx context.Context, req *pb.GetEvaluatorRequest) (res *pb.GetEvaluatorResponse, err error) {
	var evltr_s *storage.Evaluator

	evltr_id_str := req.GetEvaluator().GetId().GetValue()
	logger := srv.get_logger().WithField("id", evltr_id_str)

	if evltr_s, err = srv.storage.GetEvaluator(ctx, evltr_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get evaluator in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.GetEvaluatorResponse{
		Evaluator: copy_evaluator(evltr_s),
	}

	logger.Debugf("get evaluator")

	return res, nil
}
