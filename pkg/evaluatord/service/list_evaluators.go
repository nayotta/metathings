package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ListEvaluators(ctx context.Context, req *pb.ListEvaluatorsRequest) (res *pb.ListEvaluatorsResponse, err error) {
	evltr := req.GetEvaluator()
	evltr_s := &storage.Evaluator{}

	logger := srv.get_logger()

	if evltr_id := evltr.GetId(); evltr_id != nil {
		evltr_id_str := evltr_id.GetValue()
		evltr_s.Id = &evltr_id_str
	}

	if evltr_alias := evltr.GetAlias(); evltr_alias != nil {
		evltr_alias_str := evltr_alias.GetValue()
		evltr_s.Alias = &evltr_alias_str
	}

	evltrs_s, err := srv.storage.ListEvaluators(ctx, evltr_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to list evaluators in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.ListEvaluatorsResponse{
		Evaluators: copy_evaluators(evltrs_s),
	}

	logger.Debugf("list evaluators")

	return res, nil
}
