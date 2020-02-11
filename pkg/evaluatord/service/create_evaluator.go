package metathings_evaluatord_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidateCreateEvaluator(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, evaluator_getter) {
				req := in.(*pb.CreateEvaluatorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_evaluator,
			ensure_evaluator_id_not_exists(ctx, srv.storage),
			ensure_operator_id_not_exists(ctx, srv.storage),
			ensure_valid_operator_driver,
		},
	)
}

func (srv *MetathingsEvaluatordService) CreateEvaluator(ctx context.Context, req *pb.CreateEvaluatorRequest) (res *pb.CreateEvaluatorResponse, err error) {
	evltr := req.GetEvaluator()
	evltr_s := &storage.Evaluator{}

	evltr_id_str := id_helper.NewId()
	if evltr_id := evltr.GetId(); evltr_id != nil {
		evltr_id_str = evltr_id.GetValue()
	}
	evltr_s.Id = &evltr_id_str

	logger := srv.get_logger().WithField("id", evltr_id_str)

	evltr_alias_str := evltr_id_str
	if evltr_alias := evltr.GetAlias(); evltr_alias != nil {
		evltr_alias_str = evltr_alias.GetValue()
	}
	evltr_s.Alias = &evltr_alias_str

	evltr_description_str := ""
	if evltr_description := evltr.GetDescription(); evltr_description != nil {
		evltr_description_str = evltr_description.GetValue()
	}
	evltr_s.Description = &evltr_description_str

	evltr_sources := []*storage.Resource{}
	for _, evltr_src := range evltr.GetSources() {
		src_id_str := evltr_src.GetId().GetValue()
		src_type_str := evltr_src.GetType().GetValue()
		evltr_sources = append(evltr_sources, &storage.Resource{
			Id:   &src_id_str,
			Type: &src_type_str,
		})
	}
	evltr_s.Sources = evltr_sources

	op := evltr.GetOperator()
	op_s := &storage.Operator{}

	op_id_str := id_helper.NewId()
	if op_id := op.GetId(); op_id != nil {
		op_id_str = op_id.GetValue()
	}
	op_s.Id = &op_id_str

	op_alias_str := op_id_str
	if op_alias := op.GetAlias(); op_alias != nil {
		op_alias_str = op_alias.GetValue()
	}
	op_s.Alias = &op_alias_str

	op_description_str := ""
	if op_description := op.GetDescription(); op_description != nil {
		op_description_str = op_description.GetValue()
	}
	op_s.Description = &op_description_str

	op_driver_str := "default"
	if op_driver := op.GetDriver(); op_driver != nil {
		op_driver_str = op_driver.GetValue()
	}
	op_s.Driver = &op_driver_str

	switch op_driver_str {
	case "lua":
		fallthrough
	case "default":
		desc := op.GetLua()
		desc_code_str := desc.GetCode().GetValue()
		op_s.LuaDescriptor = &storage.LuaDescriptor{
			Code: &desc_code_str,
		}
	}
	evltr_s.Operator = op_s

	evltr_s, err = srv.storage.CreateEvaluator(ctx, evltr_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to create evaluator in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.CreateEvaluatorResponse{
		Evaluator: copy_evaluator(evltr_s),
	}

	logger.Infof("create evaluator")

	return res, nil
}
