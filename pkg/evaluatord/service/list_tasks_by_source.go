package metathings_evaluatord_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateListTasksBySource(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, source_getter) {
				req := in.(*pb.ListTasksBySourceRequest)
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

func (srv *MetathingsEvaluatordService) AuthorizeListTasksBySource(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.ListTasksBySourceRequest).GetSource().GetId().GetValue(), "evaluatord:list_tasks_by_source")
}

func (srv *MetathingsEvaluatordService) ListTasksBySource(ctx context.Context, req *pb.ListTasksBySourceRequest) (*pb.ListTasksBySourceResponse, error) {
	src := req.GetSource()
	src_id := src.GetId().GetValue()
	src_typ := src.GetType().GetValue()
	logger := srv.get_logger().WithFields(logrus.Fields{
		"source":      src_id,
		"source_type": src_typ,
	})
	var opts []storage.ListTasksBySourceOption

	if range_from := req.GetRangeFrom(); range_from != nil {
		start, err := ptypes.Timestamp(range_from)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		logger.WithField("range_from", start)
		opts = append(opts, storage.SetRangeStartOption(start))
	}

	if range_to := req.GetRangeTo(); range_to != nil {
		stop, err := ptypes.Timestamp(range_to)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		logger.WithField("range_to", stop)
		opts = append(opts, storage.SetRangeStopOption(stop))
	}

	tsks, err := srv.task_storage.ListTasksBySource(ctx, &storage.Resource{
		Id:   &src_id,
		Type: &src_typ,
	}, opts...)
	if err != nil {
		logger.WithError(err).Errorf("failed to list tasks by source")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListTasksBySourceResponse{
		Tasks: copy_tasks(tsks),
	}

	logger.Debugf("list tasks by source")

	return res, nil
}
