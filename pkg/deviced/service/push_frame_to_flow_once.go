package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	evaluatord_sdk "github.com/nayotta/metathings/sdk/evaluatord"
)

func (self *MetathingsDevicedService) ValidatePushFrameToFlowOnce(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.PushFrameToFlowOnceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizePushFrameToFlowOnce(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PushFrameToFlowOnceRequest).GetDevice().GetId().GetValue(), "deviced:push_frame_to_flow_once")
}

func (self *MetathingsDevicedService) PushFrameToFlowOnce(ctx context.Context, req *pb.PushFrameToFlowOnceRequest) (*empty.Empty, error) {
	tkn_txt, _ := grpc_helper.GetTokenFromContext(ctx)
	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	req_id := req.GetId().GetValue()
	if req_id == "" {
		req_id = id_helper.NewId()
	}

	logger := self.get_logger().WithFields(logrus.Fields{
		"request_id": req_id,
		"device":     dev_id_str,
	})

	dev_s, err := self.storage.GetDevice(ctx, dev_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to get device")
		return nil, self.ParseError(err)
	}

	var flw_id_str string
	var flw_name_str string
	var f flow.Flow
match_flow_loop:
	for _, flw := range dev.GetFlows() {
		flw_name_str = flw.GetName().GetValue()
		for _, flw_s := range dev_s.Flows {
			if *flw_s.Name != flw_name_str {
				continue
			}

			flw_id_str = *flw_s.Id
		}
		break match_flow_loop
	}

	logger = logger.WithField("flow", flw_id_str)

	if flw_id_str == "" {
		err = ErrFlowNotFound
		logger.WithError(err).Errorf("failed to find flow by name")
		return nil, self.ParseError(err)
	}

	if f, err = self.new_flow(dev_id_str, flw_id_str); err != nil {
		logger.WithError(err).Errorf("failed to new flow")
		return nil, self.ParseError(err)
	}
	defer f.Close()

	flwsts_s, err := self.storage.ListViewFlowSetsByFlowId(ctx, flw_id_str)
	if err != nil {
		logger.WithError(err).Errorf("failed to list view flow set by flow id in storage")
		return nil, self.ParseError(err)
	}

	fss, err := self.new_flow_sets(flwsts_s)
	if err != nil {
		logger.WithError(err).Errorf("failed to new flow sets")
		return nil, self.ParseError(err)
	}
	for _, fs := range fss {
		defer fs.Close()
	}

	req_frm := req.GetFrame()
	req_dat := req_frm.GetData()
	frm := &pb.Frame{Data: req_dat}
	req_dat_str, err := new(jsonpb.Marshaler).MarshalToString(req_dat)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal data to json string")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	evltrsdk_dat, err := evaluatord_sdk.DataFromBytes([]byte(req_dat_str))
	if err != nil {
		logger.WithError(err).Errorf("failed to transfer json string to evaluatord sdk data")
		return nil, self.ParseError(err)
	}

	err = f.PushFrame(frm)
	if err != nil {
		logger.WithError(err).Errorf("failed to push frame to flow")
		return nil, self.ParseError(err)
	}

	nctx := evaluatord_sdk.WithToken(ctx, tkn_txt)
	nctx = evaluatord_sdk.WithDevice(nctx, dev_id_str)

	if err = self.data_launcher.Launch(nctx, evaluatord_sdk.NewResource(f.Id(), RESOURCE_TYPE_FLOW), evltrsdk_dat); err != nil {
		logger.WithError(err).Warningf("failed to launch data")
	}

	flwst_frm_dev := &pb.Device{
		Id: dev_id_str,
		Flows: []*pb.Flow{
			{Name: flw_name_str},
		},
	}

	for _, fs := range fss {
		if err = fs.PushFrame(&flow.FlowSetFrame{
			Device: flwst_frm_dev,
			Frame:  frm,
		}); err != nil {
			logger.WithError(err).WithField("flow_set", fs.Id()).Errorf("failed to push frame to flow set")
			return nil, self.ParseError(err)
		}

		if err = self.data_launcher.Launch(nctx, evaluatord_sdk.NewResource(fs.Id(), RESOURCE_TYPE_FLOWSET), evltrsdk_dat); err != nil {
			logger.WithError(err).Warningf("failed to launch data")
		}
	}

	logger.Infof("push frame to flow once")

	return &empty.Empty{}, nil
}
