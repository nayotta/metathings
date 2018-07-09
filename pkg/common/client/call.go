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

func NewUInt64(x uint64) *gpb.UInt64Value {
	return &gpb.UInt64Value{Value: x}
}

func NewFloat32(x float32) *gpb.FloatValue {
	return &gpb.FloatValue{Value: x}
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

func MustNewStreamCallConfigRequest(core_id, entity_name, service_name, method_name string) *cored_pb.StreamCallRequest {
	cfg_req := &cored_pb.StreamCallRequest{
		CoreId: NewString(core_id),
		Payload: &cored_pb.StreamCallRequestPayload{
			Payload: &cored_pb.StreamCallRequestPayload_Config{
				Config: &cored_pb.StreamCallConfigRequest{
					Name:        NewString(entity_name),
					ServiceName: NewString(service_name),
					MethodName:  NewString(method_name),
				},
			},
		},
	}
	return cfg_req
}

func MustNewStreamCallDataRequest(core_id string, pb proto.Message) *cored_pb.StreamCallRequest {
	req_val, err := ptypes.MarshalAny(pb)
	if err != nil {
		panic(err)
	}
	dat_req := &cored_pb.StreamCallRequest{
		CoreId: NewString(core_id),
		Payload: &cored_pb.StreamCallRequestPayload{
			Payload: &cored_pb.StreamCallRequestPayload_Data{
				Data: &cored_pb.StreamCallDataRequest{Value: req_val},
			},
		},
	}
	return dat_req
}
