package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/ptypes/wrappers"
	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/proto/deviced"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type UploadDescriptorOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	Protoset                     string
}

func NewUploadDescriptorOption() *UploadDescriptorOption {
	return &UploadDescriptorOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	upload_descriptor_opt *UploadDescriptorOption
	uploadDescriptorCmd   = &cobra.Command{
		Use:   "upload",
		Short: "Upload descriptor",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				upload_descriptor_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(upload_descriptor_opt)
			base_opt = &upload_descriptor_opt.BaseOption

			upload_descriptor_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("upload descriptor", upload_descriptor),
	}
)

func GetUploadDescriptorOptions() (
	*UploadDescriptorOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return upload_descriptor_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(upload_descriptor_opt, upload_descriptor_opt),
		upload_descriptor_opt
}

func upload_descriptor() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetUploadDescriptorOptions,
			cmd_contrib.NewLogger("upload_descriptor"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *UploadDescriptorOption, cli_fty *client_helper.ClientFactory) {
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

						return _upload_descriptor(opt, cli)
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

func _upload_descriptor(opt *UploadDescriptorOption, cli pb.DevicedServiceClient) error {
	ctx := context_helper.WithToken(context.TODO(), opt.GetToken())

	buf, err := ioutil.ReadFile(opt.Protoset)
	if err != nil {
		return err
	}

	req := &pb.UploadDescriptorRequest{
		Descriptor_: &pb.OpDescriptor{
			Body: &wrappers.BytesValue{
				Value: buf,
			},
		},
	}

	res, err := cli.UploadDescriptor(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println(res.Descriptor_.Sha1)

	return nil
}

func init() {
	upload_descriptor_opt = NewUploadDescriptorOption()

	flags := uploadDescriptorCmd.Flags()

	flags.StringVar(&upload_descriptor_opt.Protoset, "protoset", "", "Protoset file")

	descriptorCmd.AddCommand(uploadDescriptorCmd)
}
