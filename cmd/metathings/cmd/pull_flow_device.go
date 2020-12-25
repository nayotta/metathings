package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type PullFlowDeviceOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	Device                       string
	Flow                         string
}

func NewPullFlowDeviceOption() *PullFlowDeviceOption {
	return &PullFlowDeviceOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	pull_flow_device_opt *PullFlowDeviceOption
)

var (
	pullFlowDeviceCmd = &cobra.Command{
		Use:   "pull-flow",
		Short: "Pull flow from device",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				pull_flow_device_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(pull_flow_device_opt)
			base_opt = &pull_flow_device_opt.BaseOption

			pull_flow_device_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("pull flow device", pull_flow_device),
	}
)

func GetPullFlowDeviceOptions() (
	*PullFlowDeviceOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return pull_flow_device_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(pull_flow_device_opt, pull_flow_device_opt),
		pull_flow_device_opt
}

func pull_flow_device() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetPullFlowDeviceOptions,
			cmd_contrib.NewLogger("pull_flow_device"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *PullFlowDeviceOption, cli_fty *client_helper.ClientFactory) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						if opt.GetToken() == "" {
							opt.SetToken(os.Getenv("MT_TOKEN"))
						}

						cli, cfn, err := cli_fty.NewDevicedServiceClient()
						if err != nil {
							return err
						}
						defer cfn()

						return _pull_flow_device(opt, cli)
					},
				})
			},
		),
	)

	if err := app.Start(context.TODO()); err != nil {
		return err
	}

	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func _pull_flow_device(opt *PullFlowDeviceOption, cli pb.DevicedServiceClient) error {
	ctx := context_helper.WithToken(context.TODO(), opt.GetToken())

	req_id := id_helper.NewId()
	req := &pb.PullFrameFromFlowRequest{
		Id: &wrappers.StringValue{Value: req_id},
		Request: &pb.PullFrameFromFlowRequest_Config_{
			Config: &pb.PullFrameFromFlowRequest_Config{
				Device: &pb.OpDevice{
					Id: &wrappers.StringValue{Value: opt.Device},
					Flows: []*pb.OpFlow{
						{Name: &wrappers.StringValue{Value: opt.Flow}},
					},
				},
				ConfigAck: &wrappers.BoolValue{Value: true},
			},
		},
	}

	stm, err := cli.PullFrameFromFlow(ctx, req)
	if err != nil {
		return err
	}

	res, err := stm.Recv()
	if err != nil {
		return err
	}

	if res.GetAck() == nil {
		return errors.New("unexpected response")
	}

	for {
		res, err := stm.Recv()
		if err != nil {
			return err
		}

		pack := res.GetPack()
		if pack == nil {
			fmt.Printf("%v#%s: pingpong\n", res.GetId(), time.Now())
			continue
		}

		frms := pack.GetFrames()

		for _, frm := range frms {
			ts, err := ptypes.Timestamp(frm.GetTs())
			if err != nil {
				return err
			}

			buf, err := grpc_helper.JSONPBMarshaler.MarshalToString(frm.GetData())
			if err != nil {
				return err
			}
			fmt.Printf("%v#%s: %v\n", res.GetId(), ts, buf)
		}
	}
}

func init() {
	pull_flow_device_opt = NewPullFlowDeviceOption()

	flags := pullFlowDeviceCmd.Flags()

	flags.StringVar(pull_flow_device_opt.GetTokenP(), "token", "", "Token")
	flags.StringVar(&pull_flow_device_opt.Device, "device", "", "Device ID")
	flags.StringVar(&pull_flow_device_opt.Flow, "flow", "", "Flow name")

	deviceCmd.AddCommand(pullFlowDeviceCmd)
}
