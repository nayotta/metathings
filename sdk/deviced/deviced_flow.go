package metathings_deviced_sdk

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type DevicedFlow struct {
	logger  logrus.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (f *DevicedFlow) PushFrame(ctx context.Context, device, flow string, data interface{}) error {
	logger := f.logger.WithFields(logrus.Fields{
		"device": device,
		"flow":   flow,
	})

	cli, cfn, err := f.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	defer cfn()

	var payload stpb.Struct
	buf, err := json.Marshal(data)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal data to json string")
		return err
	}

	err = jsonpb.UnmarshalString(string(buf), &payload)
	if err != nil {
		logger.WithError(err).Debugf("failed to unmarshal string to payload")
		return err
	}

	frm := &pb.OpFrame{
		Ts:   ptypes.TimestampNow(),
		Data: &payload,
	}

	req := &pb.PushFrameToFlowOnceRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device},
			Flows: []*pb.OpFlow{
				{Name: &wrappers.StringValue{Value: flow}},
			},
		},
		Frame: frm,
	}

	_, err = cli.PushFrameToFlowOnce(ctx, req)
	if err != nil {
		logger.WithError(err).Debugf("failed to push frame to flow once")
		return err
	}

	return nil
}

func NewDevicedFlow(args ...interface{}) (Flow, error) {
	var logger logrus.FieldLogger
	var cli_fty *client_helper.ClientFactory

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":         opt_helper.ToLogger(&logger),
		"client_factory": client_helper.ToClientFactory(&cli_fty),
	})(args...); err != nil {
		return nil, err
	}

	return &DevicedFlow{
		logger:  logger,
		cli_fty: cli_fty,
	}, nil
}

func init() {
	register_flow_factory("default", NewDevicedFlow)
}
