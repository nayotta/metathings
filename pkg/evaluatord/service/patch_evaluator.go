package metathings_evaluatord_service

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

func (srv *MetathingsEvaluatordService) ValidatePatchEvaluator(ctx context.Context, in interface{}) error {
	return srv.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, evaluator_getter) {
				req := in.(*pb.PatchEvaluatorRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_evaluator,
			ensure_get_evaluator_id,
		},
	)
}

func (srv *MetathingsEvaluatordService) AuthorizePatchEvaluator(ctx context.Context, in interface{}) error {
	return srv.authorizer.Authorize(ctx, in.(*pb.PatchEvaluatorRequest).GetEvaluator().GetId().GetValue(), "evaluatord:patch_evaluator")
}

func (srv *MetathingsEvaluatordService) PatchEvaluator(ctx context.Context, req *pb.PatchEvaluatorRequest) (res *pb.PatchEvaluatorResponse, err error) {
	evltr := req.GetEvaluator()
	evltr_id_str := evltr.GetId().GetValue()
	evltr_s := &storage.Evaluator{}

	logger := srv.get_logger().WithField("evaluator", evltr_id_str)

	if evltr_alias := evltr.GetAlias(); evltr_alias != nil {
		evltr_alias_str := evltr_alias.GetValue()
		evltr_s.Alias = &evltr_alias_str
	}

	if evltr_description := evltr.GetDescription(); evltr_description != nil {
		evltr_description_str := evltr_description.GetValue()
		evltr_s.Description = &evltr_description_str
	}

	if evltr_config := evltr.GetConfig(); evltr_config != nil {
		evltr_config_str, err := new(jsonpb.Marshaler).MarshalToString(evltr_config)
		if err != nil {
			logger.WithError(err).Errorf("failed to marshal config to string")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		evltr_s.Config = &evltr_config_str
	}

	if op := evltr.GetOperator(); op != nil {
		op_s := &storage.Operator{}

		if op_alias := op.GetAlias(); op_alias != nil {
			op_alias_str := op_alias.GetValue()
			op_s.Alias = &op_alias_str
		}

		if op_description := op.GetDescription(); op_description != nil {
			op_description_str := op_description.GetValue()
			op_s.Description = &op_description_str
		}

		if op_driver := op.GetDriver(); op_driver != nil {
			op_driver_str := op_driver.GetValue()

			// SYM:REFACTOR:lua_operator
			switch op_driver_str {
			case "lua":
				fallthrough
			case "default":
				if lua_desc := op.GetLua(); lua_desc != nil {
					op_s.LuaDescriptor = &storage.LuaDescriptor{}
					if desc_code := lua_desc.GetCode(); desc_code != nil {
						desc_code_str := desc_code.GetValue()
						op_s.LuaDescriptor.Code = &desc_code_str
					}
				}
			}
		}

		evltr_s.Operator = op_s
	}

	if evltr_s, err = srv.storage.PatchEvaluator(ctx, evltr_id_str, evltr_s); err != nil {
		logger.WithError(err).Errorf("failed to patch evaluator")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res = &pb.PatchEvaluatorResponse{
		Evaluator: copy_evaluator(evltr_s),
	}

	logger.Infof("patch evaluator")

	return res, nil
}
