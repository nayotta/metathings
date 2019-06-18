package metathings_device_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	op_mdl := req.GetModule()
	name := op_mdl.GetName().GetValue()

	mdl, err := self.mdl_db.Lookup(name)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to lookup module")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	mdl.Heartbeat()

	self.logger.WithFields(log.Fields{
		"name": name,
	}).Debugf("module heartbeat")

	return &empty.Empty{}, nil
}

func (self *MetathingsDeviceServiceImpl) heartbeat_loop() {
	self.conn_stm_wg.Wait()
	for {
		go self.heartbeat_once()
		time.Sleep(self.opt.HeartbeatInterval)
	}
}

func (self *MetathingsDeviceServiceImpl) heartbeat_once() {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to connect to deviced service")
		return
	}
	defer cfn()

	now := time.Now()
	pb_now := protobuf_helper.FromTime(now)
	pb_mdls := []*deviced_pb.OpModule{}
	for _, mdl := range self.mdl_db.All() {
		var stat state_pb.ModuleState

		hbt := mdl.HeartbeatAt()
		pb_hbt := protobuf_helper.FromTime(hbt)
		if mdl.IsAlive() {
			stat = state_pb.ModuleState_MODULE_STATE_ONLINE
		} else {
			stat = state_pb.ModuleState_MODULE_STATE_OFFLINE
		}

		pb_mdl := &deviced_pb.OpModule{
			Id:          &wrappers.StringValue{Value: mdl.Id()},
			HeartbeatAt: &pb_hbt,
			State:       stat,
		}

		pb_mdls = append(pb_mdls, pb_mdl)
	}

	req := &deviced_pb.HeartbeatRequest{
		Device: &deviced_pb.OpDevice{
			Id:          &wrappers.StringValue{Value: self.info.GetId()},
			HeartbeatAt: &pb_now,
			Modules:     pb_mdls,
		},
		StartupSession: &wrappers.Int32Value{Value: self.startup_session},
	}

	_, err = cli.Heartbeat(self.context(), req)
	if err != nil {
		// TODO(Peer): reconncect streaming, not stop device and restart
		defer self.Stop()

		self.logger.WithError(err).Warningf("failed to heartbeat")
		return
	}

	self.logger.WithFields(log.Fields{
		"heartbeat_at": now,
	}).Debugf("heartbeat")
}
