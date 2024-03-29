package metathings_deviced_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateCreateFlowSet(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() flow_set_getter {
				req := in.(*pb.CreateFlowSetRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			func(x flow_set_getter) error {
				fs := x.GetFlowSet()

				if fs.GetName() == nil {
					return errors.New("flow_set.name is empty")
				}

				return nil
			},
		},
	)
}

func (self *MetathingsDevicedService) CreateFlowSet(ctx context.Context, req *pb.CreateFlowSetRequest) (*pb.CreateFlowSetResponse, error) {
	var err error

	flwst := req.GetFlowSet()

	flwst_id_str := id_helper.NewId()
	if flwst.GetId() != nil {
		flwst_id_str = flwst.GetId().GetValue()
	}
	flwst_name_str := flwst.GetName().GetValue()
	flwst_alias_str := flwst_name_str
	if flwst.GetAlias() != nil {
		flwst_alias_str = flwst.GetAlias().GetValue()
	}

	logger := self.get_logger().WithField("flow_set", flwst_id_str)

	flwst_s := &storage.FlowSet{
		Id:    &flwst_id_str,
		Name:  &flwst_name_str,
		Alias: &flwst_alias_str,
	}

	if flwst_s, err = self.storage.CreateFlowSet(ctx, flwst_s); err != nil {
		logger.WithError(err).Errorf("failed to create flow set in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateFlowSetResponse{
		FlowSet: copy_flow_set(flwst_s),
	}

	logger.Infof("create flow set")

	return res, nil
}
