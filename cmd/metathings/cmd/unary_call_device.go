package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	pb "github.com/nayotta/metathings/proto/deviced"
)

type UnaryCallDeviceOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	Device                       string
	Module                       string
	Method                       string
	Protobufset                  string
	Soda                         bool
	Data                         string
	File                         string
}

func NewUnaryCallDeviceOption() *UnaryCallDeviceOption {
	return &UnaryCallDeviceOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	unary_call_device_opt *UnaryCallDeviceOption
)

var (
	unaryCallDeviceCmd = &cobra.Command{
		Use:   "unary-call",
		Short: "Unary Call Device Service",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				unary_call_device_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(unary_call_device_opt)
			base_opt = &unary_call_device_opt.BaseOption

			unary_call_device_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("unary call device", unary_call_device),
	}
)

func GetUnaryCallDeviceOptions() (
	*UnaryCallDeviceOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return unary_call_device_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(unary_call_device_opt, unary_call_device_opt),
		unary_call_device_opt
}

func unary_call_device() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetUnaryCallDeviceOptions,
			cmd_contrib.NewLogger("unary_call_device"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *UnaryCallDeviceOption, cli_fty *client_helper.ClientFactory) {
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

						return _unary_call_device(opt, cli)
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

func _unary_call_device(opt *UnaryCallDeviceOption, cli pb.DevicedServiceClient) error {
	var buf []byte
	var err error
	var any_req *any.Any
	var md *desc.MethodDescriptor

	if opt.Data == "" {
		if opt.File == "" {
			return errors.New("require data or data file")
		}

		if opt.File == "-" {
			buf, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
		} else {
			buf, err = ioutil.ReadFile(opt.File)
			if err != nil {
				return err
			}
		}
	} else {
		buf = []byte(opt.Data)
	}

	if opt.Soda {
		var st stpb.Struct

		if err = grpc_helper.JSONPBUnmarshaler.Unmarshal(bytes.NewReader(buf), &st); err != nil {
			return err
		}

		if any_req, err = ptypes.MarshalAny(&st); err != nil {
			return err
		}
	} else {
		var fds dpb.FileDescriptorSet

		req_buf, err := ioutil.ReadFile(opt.Protobufset)
		if err != nil {
			panic(err)
		}

		if err = proto.Unmarshal(req_buf, &fds); err != nil {
			panic(err)
		}

		fd, err := desc.CreateFileDescriptorFromSet(&fds)
		if err != nil {
			panic(err)
		}

		srvs := fd.GetServices()
		if len(srvs) == 0 {
			panic("unexpected protobufset")
		}

		md = srvs[0].FindMethodByName(opt.Method)
		msg_req := dynamic.NewMessage(md.GetInputType())

		if err = msg_req.UnmarshalJSON(req_buf); err != nil {
			return err
		}

		if any_req, err = ptypes.MarshalAny(msg_req); err != nil {
			return err
		}
	}

	req := &pb.UnaryCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: opt.Device},
		},
		Value: &pb.OpUnaryCallValue{
			Name:   &wrappers.StringValue{Value: opt.Module},
			Method: &wrappers.StringValue{Value: opt.Method},
			Value:  any_req,
		},
	}

	ctx := context_helper.WithToken(context.TODO(), opt.GetToken())

	res, err := cli.UnaryCall(ctx, req)
	if err != nil {
		return err
	}

	var msg_res proto.Message
	if opt.Soda {
		msg_res = new(stpb.Struct)
	} else {
		msg_res = dynamic.NewMessage(md.GetOutputType())
	}

	if err = ptypes.UnmarshalAny(res.GetValue().GetValue(), msg_res); err != nil {
		return err
	}

	out, err := grpc_helper.JSONPBMarshaler.MarshalToString(msg_res)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func init() {
	unary_call_device_opt = NewUnaryCallDeviceOption()

	flags := unaryCallDeviceCmd.Flags()

	flags.StringVar(unary_call_device_opt.GetTokenP(), "token", "", "Token")
	flags.StringVar(&unary_call_device_opt.Device, "device", "", "Device ID")
	flags.StringVar(&unary_call_device_opt.Module, "module", "", "Module Name")
	flags.StringVar(&unary_call_device_opt.Method, "method", "", "Device Method")
	flags.StringVar(&unary_call_device_opt.Data, "data", "", "Data")
	flags.BoolVar(&unary_call_device_opt.Soda, "soda", false, "Unary call with soda mode")
	flags.StringVar(&unary_call_device_opt.Protobufset, "protobufset", "", "Protobuf set file")
	flags.StringVar(&unary_call_device_opt.File, "file", "", "Data file")

	deviceCmd.AddCommand(unaryCallDeviceCmd)
}
