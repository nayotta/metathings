package client_helper

import (
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	gpb "github.com/golang/protobuf/ptypes/wrappers"

	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
)

func NewString(s string) *gpb.StringValue {
	return &gpb.StringValue{Value: s}
}

func MustNewUnaryCallRequest(core_id, entity_name, service_name, method_name string, request proto.Message) *cored_pb.UnaryCallRequest {
	req_val, err := ptypes.MarshalAny(request)
	if err != nil {
		panic(err)
	}

	req := &cored_pb.UnaryCallRequest{
		CoreId: NewString(core_id),
		Payload: &cored_pb.UnaryCallRequestPayload{
			Name:        NewString(entity_name),
			ServiceName: NewString(service_name),
			MethodName:  NewString(method_name),
			Value:       req_val,
		},
	}

	return req
}

func DecodeUnaryCallResponse(call_res *cored_pb.UnaryCallResponse, res proto.Message) error {
	err := ptypes.UnmarshalAny(call_res.Payload.Value, res)
	if err != nil {
		return err
	}

	return nil
}
