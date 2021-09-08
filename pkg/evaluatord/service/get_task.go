package metathings_evaluatord_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateGetTask(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() task_getter {
				req := in.(*pb.GetTaskRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_task,
			ensure_get_task_id,
		},
	)
}

func (srv *MetathingsEvaluatordService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	tsk_id := req.GetTask().GetId().GetValue()
	logger := srv.get_logger().WithFields(logrus.Fields{
		"id": tsk_id,
	})
	tsk, err := srv.task_storage.GetTask(ctx, tsk_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get task")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetTaskResponse{
		Task: copy_task(tsk),
	}

	logger.Debugf("get task")

	return res, nil
}
