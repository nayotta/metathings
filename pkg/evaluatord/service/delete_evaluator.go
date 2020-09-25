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

func (srv *MetathingsEvaluatordService) ValidateDeleteEvaluator(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, evaluator_getter) {
				req := in.(*pb.DeleteEvaluatorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_evaluator,
			ensure_get_evaluator_id,
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeDeleteEvaluator(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.DeleteEvaluatorRequest).GetEvaluator().GetId().GetValue(), "evaluatord:delete_evaluator")
}

func (srv *MetathingsEvaluatordService) DeleteEvaluator(ctx context.Context, req *pb.DeleteEvaluatorRequest) (_ *empty.Empty, err error) {
	evltr_id_str := req.GetEvaluator().GetId().GetValue()
	logger := srv.get_logger().WithField("id", evltr_id_str)

	if err = srv.storage.DeleteEvaluator(ctx, evltr_id_str); err != nil {
		logger.WithError(err).Errorf("failed to delete evaluator in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("delete evaluator")

	return &empty.Empty{}, nil
}
