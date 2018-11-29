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

// Heartbeat Heartbeat
// TODO(zh) heart beat make a lot connect, need reduce
func (serv *MetathingsMqttdService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	ctx = context_helper.WithToken(ctx, serv.appCredMgr.GetToken())
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
		Session:  &gpb.UInt64Value{Value: serv.heartbeatSession},
		Entities: ents,
	}

	_, err = cli.Heartbeat(ctx, r)
	if err != nil {
		serv.logger.WithError(err).Errorf("heart beat failed")
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
