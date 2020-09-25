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

func (srv *MetathingsEvaluatordService) ValidateListEvaluatorsBySource(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, source_getter) {
				req := in.(*pb.ListEvaluatorsBySourceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_source,
			ensure_get_source_id,
			ensure_get_source_type,
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeListEvaluatorsBySource(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.ListEvaluatorsBySourceRequest).GetSource().GetId().GetValue(), "evaluatord:list_evaluators_by_source")
}

func (srv *MetathingsEvaluatordService) ListEvaluatorsBySource(ctx context.Context, req *pb.ListEvaluatorsBySourceRequest) (res *pb.ListEvaluatorsBySourceResponse, err error) {
	src := req.GetSource()
	src_id_str := src.GetId().GetValue()
	src_typ_str := src.GetType().GetValue()

	logger := srv.get_logger().WithField("source", src_id_str)

	src_s := &storage.Resource{
		Id:   &src_id_str,
		Type: &src_typ_str,
	}

	evltrs_s, err := srv.storage.ListEvaluatorsBySource(ctx, src_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to list evaluators by source in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.ListEvaluatorsBySourceResponse{
		Evaluators: copy_evaluators(evltrs_s),
	}

	logger.Debugf("list evaluators by source")

	return res, nil
}
