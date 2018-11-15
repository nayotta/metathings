package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	service "github.com/nayotta/metathings/pkg/camerad/service"
	storage "github.com/nayotta/metathings/pkg/camerad/storage"
	cli_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/camerad"
)

type RtmpOptioner interface {
	GetRtmpUrlP() *string
	GetRtmpUrl() string
	SetRtmpUrl(string)
}

type RtmpOption struct {
	Rtmp struct {
		Url string
	}
}

func (self *RtmpOption) GetRtmpUrlP() *string {
	return &self.Rtmp.Url
}

func (self *RtmpOption) GetRtmpUrl() string {
	return self.Rtmp.Url
}

func (self *RtmpOption) SetRtmpUrl(url string) {
	self.Rtmp.Url = url
}

type CameradOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	RtmpOption                    `mapstructure:",squash"`
}

func NewCameradOption() *CameradOption {
	return &CameradOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	camerad_opt *CameradOption
)

var (
	cameradCmd = &cobra.Command{
		Use:   "camerad",
		Short: "Camera Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.GetConfig() == "" {
				return
			}

			opt_t := NewCameradOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			if opt_t.GetListen() == "" {
				opt_t.SetListen(camerad_opt.GetListen())
			}

			if opt_t.Storage.Driver == "" {
				opt_t.Storage.Driver = camerad_opt.Storage.Driver
			}

			if opt_t.Storage.Uri == "" {
				opt_t.Storage.Uri = camerad_opt.Storage.Uri
			}

			if opt_t.GetCertFile() == "" {
				opt_t.SetCertFile(camerad_opt.GetCertFile())
			}

			if opt_t.GetKeyFile() == "" {
				opt_t.SetKeyFile(camerad_opt.GetKeyFile())
			}

			// BUG(Peer): chould not get address from config?
			if opt_t.GetServiceEndpoint(cli_helper.DEFAULT_CONFIG).GetAddress() == "" {
				opt_t.GetServiceEndpoint(cli_helper.DEFAULT_CONFIG).SetAddress(camerad_opt.GetServiceEndpoint(cli_helper.DEFAULT_CONFIG).GetAddress())
			}

			camerad_opt.SetServiceName("camerad")
			camerad_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("camerad", runCamerad),
	}
)

func GetCameradOptions() (
	*CameradOption,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return camerad_opt,
		camerad_opt,
		camerad_opt,
		camerad_opt,
		camerad_opt
}

func NewCameradStorage(opt cmd_contrib.StorageOptioner, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.GetDriver(), opt.GetUri(), logger)
}

func NewMetathingsCameradServiceOption(opt *CameradOption) *service.MetathingsCameradServiceOption {
	return &service.MetathingsCameradServiceOption{
		RtmpUrl: opt.Rtmp.Url,
	}
}

func runCamerad() error {
	app := fx.New(
		fx.Provide(
			GetCameradOptions,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("camerad"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewCredentialManager,
			state_helper.NewCameraStateParser,
			token_helper.NewTokenValidator,
			NewCameradStorage,
			NewMetathingsCameradServiceOption,
			service.NewMetathingsCameradService,
		),
		fx.Invoke(
			pb.RegisterCameradServiceServer,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	// camerad_opts = &_cameradOptions{}
	camerad_opt = NewCameradOption()

	cameradCmd.Flags().StringVarP(camerad_opt.GetListenP(), "listen", "l", "127.0.0.1:5002", "Metathings Camera Service listening address")
	cameradCmd.Flags().StringVar(camerad_opt.GetServiceEndpoint(cli_helper.DEFAULT_CONFIG).GetAddressP(), "host", "mt-api.nayotta.com", "MetaThings Service Address")
	cameradCmd.Flags().StringVar(camerad_opt.GetDriverP(), "storage-driver", "sqlite3", "Storage Driver [sqlite3]")
	cameradCmd.Flags().StringVar(camerad_opt.GetUriP(), "storage-uri", "", "Storage URI")
	cameradCmd.Flags().StringVar(camerad_opt.GetRtmpUrlP(), "rtmp-url", "", "RTMP Server URL")

	RootCmd.AddCommand(cameradCmd)
}
