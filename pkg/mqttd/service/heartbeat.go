package metathingsmqttdservice

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Hearbeat Hearbeat
// TODO(zh) heart beat make a lot connect, need reduce
func (serv *MetathingsMqttdService) Hearbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	entID := req.GetEntityId().GetValue()
	ctx = context_helper.WithToken(ctx, serv.app_cred_mgr.GetToken())
	cli, closeFn, err := serv.cliFty.NewCoredServiceClient()
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	ents := []*cored_pb.HeartbeatEntity{}
	ts := protobuf_helper.FromTime(time.Now())
	ents = append(ents, &cored_pb.HeartbeatEntity{
		Id:          req.EntityId,
		HeartbeatAt: &ts,
	})

	r := &cored_pb.HeartbeatRequest{
		Session:  &gpb.UInt64Value{Value: serv.heartbeat_session},
		Entities: ents,
	}

	_, err = cli.Heartbeat(ctx, r)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
