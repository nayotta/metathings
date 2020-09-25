package metathings_deviced_service

import (
	"github.com/golang/protobuf/jsonpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	pb "github.com/nayotta/metathings/proto/deviced"
	evaluatord_sdk "github.com/nayotta/metathings/sdk/evaluatord"
)

// TODO(Peer): merge PushFrameToFlow and PushFrameToFlowOnce to one function
func (self *MetathingsDevicedService) PushFrameToFlow(stm pb.DevicedService_PushFrameToFlowServer) error {

	logger := self.get_logger().WithField("#method", "PushFrameToFlow")

	req, err := stm.Recv()
	if err != nil {
		logger.WithError(err).Errorf("failed to receive config request")
		return status.Errorf(codes.Internal, err.Error())
	}

	var dev_r *pb.OpDevice
	var dev_id string
	var cfg_ack, push_ack bool
	var flw_id_r string
	var flw_name_r string
	var fss []flow.FlowSet

	req_id := req.GetId().GetValue()
	cfg := req.GetConfig()
	dev_r = cfg.GetDevice()
	dev_id = dev_r.GetId().GetValue()
	cfg_ack = cfg.GetConfigAck().GetValue()
	push_ack = cfg.GetPushAck().GetValue()
	ctx := stm.Context()

	logger = logger.WithFields(log.Fields{
		"request": req_id,
		"device":  dev_id,
	})

	tkn_txt, err := grpc_helper.GetTokenFromContext(ctx)
	if err != nil {
		logger.WithError(err).Errorf("failed to get token from context")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	dev_s, err := self.storage.GetDevice(ctx, dev_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get device")
		return status.Errorf(codes.Internal, err.Error())
	}

	var f flow.Flow
match_flow_loop:
	for _, flw_r := range dev_r.GetFlows() {
		flw_name_r = flw_r.GetName().GetValue()
		for _, flw_s := range dev_s.Flows {
			if *flw_s.Name != flw_name_r {
				continue
			}

			flw_id_r = *flw_s.Id
			break match_flow_loop
		}
	}

	logger = logger.WithField("flow", flw_name_r)

	if flw_id_r == "" {
		err = ErrFlowNotFound
		logger.WithError(err).Errorf("failed to find flow by name")
		return status.Errorf(codes.InvalidArgument, ErrFlowNotFound.Error())
	}

	f, err = self.new_flow(dev_id, flw_id_r)
	if err != nil {
		logger.WithError(err).Errorf("failed to new flow")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer f.Close()

	flwsts_s, err := self.storage.ListViewFlowSetsByFlowId(stm.Context(), flw_id_r)
	if err != nil {
		logger.WithError(err).Errorf("failed to list view flow set by flow id in storage")
		return status.Errorf(codes.Internal, err.Error())
	}

	fss, err = self.new_flow_sets(flwsts_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to new flow sets")
		return status.Errorf(codes.Internal, err.Error())
	}
	for _, fs := range fss {
		defer fs.Close()
	}

	logger.WithFields(log.Fields{
		"cfg_ack":  cfg_ack,
		"push_ack": push_ack,
	}).Debugf("recv flow config request")

	if cfg_ack {
		err = stm.Send(&pb.PushFrameToFlowResponse{
			Id:       req_id,
			Response: &pb.PushFrameToFlowResponse_Ack_{Ack: &pb.PushFrameToFlowResponse_Ack{}},
		})
		if err != nil {
			logger.WithError(err).Errorf("failed to send config ack message")
			return status.Errorf(codes.Internal, err.Error())
		}
		logger.Debugf("send flow config ack response")
	}

	flwst_frm_dev := &pb.Device{
		Id: dev_id,
		Flows: []*pb.Flow{
			{Name: flw_name_r},
		},
	}

	for {
		req, err = stm.Recv()
		req_id = req.GetId().GetValue()
		if err != nil {
			logger.WithError(err).Errorf("failed to receive frame request")
			return status.Errorf(codes.Internal, err.Error())
		}
		logger.Debugf("recv data request")

		// TODO(Peer): async recv and send frame
		opfrm := req.GetFrame()
		opdat := opfrm.GetData()
		frm := &pb.Frame{Data: opdat}
		opdat_str, err := new(jsonpb.Marshaler).MarshalToString(opdat)
		if err != nil {
			logger.WithError(err).Errorf("failed to marshal data to json string")
			return status.Errorf(codes.InvalidArgument, err.Error())
		}

		evltrsdk_dat, err := evaluatord_sdk.DataFromBytes([]byte(opdat_str))
		if err != nil {
			logger.WithError(err).Errorf("failed to transfer json string to evaluatord sdk data")
			return status.Errorf(codes.Internal, err.Error())
		}

		err = f.PushFrame(frm)
		if err != nil {
			logger.WithError(err).Errorf("failed to push frame to flow")
			return status.Errorf(codes.Internal, err.Error())
		}
		logger.Debugf("push frame to flow")

		nctx := evaluatord_sdk.WithToken(ctx, tkn_txt)
		nctx = evaluatord_sdk.WithDevice(nctx, dev_id)
		err = self.data_launcher.Launch(
			nctx,
			evaluatord_sdk.NewResource(f.Id(), RESOURCE_TYPE_FLOW),
			evltrsdk_dat)
		if err != nil {
			logger.WithError(err).Warningf("failed to launch data")
		}

		for _, fs := range fss {
			if err = fs.PushFrame(&flow.FlowSetFrame{
				Device: flwst_frm_dev,
				Frame:  frm,
			}); err != nil {
				logger.WithError(err).WithField("flow_set_id", fs.Id()).Errorf("failed to push frame to flow set")
				return status.Errorf(codes.Internal, err.Error())
			}

			if err = self.data_launcher.Launch(
				nctx,
				evaluatord_sdk.NewResource(fs.Id(), RESOURCE_TYPE_FLOWSET),
				evltrsdk_dat); err != nil {
				logger.WithError(err).Warningf("failed to launch data")
			}
		}

		if push_ack {
			err = stm.Send(&pb.PushFrameToFlowResponse{
				Id:       req.GetId().GetValue(),
				Response: &pb.PushFrameToFlowResponse_Ack_{Ack: &pb.PushFrameToFlowResponse_Ack{}},
			})
			if err != nil {
				logger.WithError(err).Errorf("failed to send push ack message")
				return status.Errorf(codes.Internal, err.Error())
			}
			logger.Debugf("send flow data ack response")
		}
	}
}
