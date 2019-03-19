package metathings_deviced_connection

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	rand_helper "github.com/nayotta/metathings/pkg/common/rand"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func generate_session() int32 {
	return rand_helper.Int31()
}

func parse_bridge_id(device string, session int32) string {
	return id_helper.NewNamedId(fmt.Sprintf("device.%v.session.%v", device, session))
}

func new_config_ack_request_message(sess int32) *pb.ConnectRequest {
	return &pb.ConnectRequest{
		SessionId: &wrappers.Int32Value{Value: sess},
		Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
		Union: &pb.ConnectRequest_StreamCall{
			StreamCall: &pb.OpStreamCallValue{
				Union: &pb.OpStreamCallValue_ConfigAck{
					ConfigAck: &pb.OpStreamCallConfigAck{},
				},
			},
		},
	}
}

func new_exit_request_message(sess int32) *pb.ConnectRequest {
	return &pb.ConnectRequest{
		SessionId: &wrappers.Int32Value{Value: sess},
		Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
		Union: &pb.ConnectRequest_StreamCall{
			StreamCall: &pb.OpStreamCallValue{
				Union: &pb.OpStreamCallValue_Exit{
					Exit: &pb.OpStreamCallExit{},
				},
			},
		},
	}
}

func new_exit_response_message(sess int32) *pb.ConnectResponse {
	return &pb.ConnectResponse{
		SessionId: sess,
		Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
		Union: &pb.ConnectResponse_StreamCall{
			StreamCall: &pb.StreamCallValue{
				Union: &pb.StreamCallValue_Exit{
					Exit: &pb.StreamCallExit{},
				},
			},
		},
	}
}

func must_marshal_message(msg proto.Message) []byte {
	buf, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return buf
}
