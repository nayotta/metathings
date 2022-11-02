package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type StreamCallDeviceOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:", squash"`

	Device      string
	Module      string
	Method      string
	Protobufset string
	Soda        bool
	Data        string
	File        string
}

func NewStreamCallDeviceOption() *StreamCallDeviceOption {
	return &StreamCallDeviceOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	stream_call_device_opt *StreamCallDeviceOption
	streamCallDeviceCmd    = &cobra.Command{
		Use:   "stream-call",
		Short: "Stream Call Device Service",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				stream_call_device_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(stream_call_device_opt)
			base_opt = &stream_call_device_opt.BaseOption

			stream_call_device_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("stream call device", stream_call_device),
	}
)

func GetStreamCallDeviceOptions() (
	*StreamCallDeviceOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return stream_call_device_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(stream_call_device_opt, stream_call_device_opt),
		stream_call_device_opt
}

func stream_call_device() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetStreamCallDeviceOptions,
			cmd_contrib.NewLogger("stream_call_device"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *StreamCallDeviceOption, cli_fty *client_helper.ClientFactory) {
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

						return _stream_call_device(opt, cli)
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

func _stream_call_device(opt *StreamCallDeviceOption, cli pb.DevicedServiceClient) error {
	var req_buf []byte
	var err error
	var any_req *any.Any

	var rd *bufio.Reader
	if opt.Data == "" {
		if opt.File == "" {
			opt.File = "-"
		}

		if opt.File == "-" {
			rd = bufio.NewReader(os.Stdin)
		} else {
			buf, err := ioutil.ReadFile(opt.File)
			if err != nil {
				return err
			}
			rd = bufio.NewReader(strings.NewReader(string(buf)))
		}
	} else {
		rd = bufio.NewReader(strings.NewReader(opt.Data))
	}

	if opt.Soda {
		var bs wrappers.BytesValue
		bs.Value = req_buf

		if any_req, err = ptypes.MarshalAny(&bs); err != nil {
			return err
		}
	} else {
		panic("not support non-soda")
	}

	ctx := context_helper.WithToken(context.TODO(), opt.GetToken())

	stm, err := cli.StreamCall(ctx)
	if err != nil {
		return err
	}

	cfg_req := &pb.StreamCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: opt.Device},
		},
		Value: &pb.OpStreamCallValue{
			Union: &pb.OpStreamCallValue_Config{
				Config: &pb.OpStreamCallConfig{
					Name:   &wrappers.StringValue{Value: opt.Module},
					Method: &wrappers.StringValue{Value: opt.Method},
				},
			},
		},
	}

	if err = stm.Send(cfg_req); err != nil {
		return err
	}

	cfg_res, err := stm.Recv()
	if err != nil {
		return err
	}

	if cfg_res.GetValue().GetConfigAck() == nil {
		panic("expect config ack")
	}

	go func() {
		for {
			buf, err := rd.ReadString('\n')
			if err != nil {
				panic(err)
			}
			buf = strings.ReplaceAll(buf, "\n", "")

			bs := &wrappers.StringValue{Value: buf}
			any_req, err = ptypes.MarshalAny(bs)
			if err != nil {
				panic(err)
			}

			dat_req := &pb.StreamCallRequest{
				Value: &pb.OpStreamCallValue{
					Union: &pb.OpStreamCallValue_Value{
						Value: any_req,
					},
				},
			}

			if err = stm.Send(dat_req); err != nil {
				panic(err)
			}
		}
	}()

	for {
		dat_res, err := stm.Recv()
		if err != nil {
			return err
		}

		var msg_res wrappers.BytesValue
		if err = ptypes.UnmarshalAny(dat_res.GetValue().GetValue(), &msg_res); err != nil {
			return err
		}

		out := string(msg_res.GetValue())

		fmt.Println(out)
	}
}

func init() {
	stream_call_device_opt = NewStreamCallDeviceOption()

	flags := streamCallDeviceCmd.Flags()

	flags.StringVar(stream_call_device_opt.GetTokenP(), "token", "", "Token")
	flags.StringVar(&stream_call_device_opt.Device, "device", "", "Device ID")
	flags.StringVar(&stream_call_device_opt.Module, "module", "", "Module Name")
	flags.StringVar(&stream_call_device_opt.Method, "method", "", "Device Method")
	flags.StringVar(&stream_call_device_opt.Data, "data", "", "Data")
	flags.BoolVar(&stream_call_device_opt.Soda, "soda", false, "Unary call with soda mode")
	flags.StringVar(&stream_call_device_opt.Protobufset, "protobufset", "", "Protobuf set file")
	flags.StringVar(&stream_call_device_opt.File, "file", "", "Data file")

	deviceCmd.AddCommand(streamCallDeviceCmd)
}
