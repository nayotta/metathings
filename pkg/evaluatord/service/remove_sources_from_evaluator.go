package metathings_evaluatord_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateRemoveSourcesFromEvaluator(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, evaluator_getter) {
				req := in.(*pb.RemoveSourcesFromEvaluatorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_evaluator,
			ensure_get_evaluator_id,
			ensure_evaluator_id_exists(ctx, srv.storage),
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizeRemoveSourcesFromEvaluator(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.RemoveSourcesFromEvaluatorRequest).GetEvaluator().GetId().GetValue(), "evaluatord:remove_sources_from_evaluator")
}

func (srv *MetathingsEvaluatordService) RemoveSourcesFromEvaluator(ctx context.Context, req *pb.RemoveSourcesFromEvaluatorRequest) (_ *empty.Empty, err error) {
	evltr_id_str := req.GetEvaluator().GetId().GetValue()
	logger := srv.get_logger().WithField("evaluator", evltr_id_str)

	srcs_s := []*storage.Resource{}
	for _, src := range req.GetSources() {
		src_id_str := src.GetId().GetValue()
		src_typ_str := src.GetType().GetValue()
		srcs_s = append(srcs_s, &storage.Resource{
			Id:   &src_id_str,
			Type: &src_typ_str,
		})
	}

	if err = srv.storage.RemoveSourcesFromEvaluator(ctx, evltr_id_str, srcs_s); err != nil {
		logger.WithError(err).Errorf("failed to remove sources from evaluator")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("remove sources from evaluator")

	return &empty.Empty{}, nil
}
