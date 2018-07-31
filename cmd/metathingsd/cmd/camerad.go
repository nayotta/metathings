package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/nayotta/metathings/pkg/camerad/service"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	pb "github.com/nayotta/metathings/pkg/proto/camerad"
)

type _cameradOptions struct {
	_rootOptions  `mapstructure:",squash"`
	Listen        string
	Storage       cmd_helper.StorageOptions
	ServiceConfig cmd_helper.ServiceConfigOptions `mapstructure:"service_config"`
	RtmpUrl       string                          `mapstructure:"rtmp_url"`
}

var (
	camerad_opts *_cameradOptions
)

var (
	cameradCmd = &cobra.Command{
		Use:   "camerad",
		Short: "Camera Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(camerad_opts)
			root_opts = &camerad_opts._rootOptions
			camerad_opts.Service = "camerad"
			camerad_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCamerad(); err != nil {
				log.WithError(err).Fatalf("failed to run camerad")
			}
		},
	}
)

func runCamerad() error {
	lis, err := net.Listen("tcp", camerad_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewCameradService(
		service.SetLogLevel(camerad_opts.Log.Level),
		service.SetIdentitydAddr(camerad_opts.ServiceConfig.Identityd.Address),
		service.SetCoredAddr(camerad_opts.ServiceConfig.Cored.Address),
		service.SetStorage(camerad_opts.Storage.Driver, camerad_opts.Storage.Uri),
		service.SetApplicationCredential(
			camerad_opts.ApplicationCredential.Id,
			camerad_opts.ApplicationCredential.Secret,
		),
		service.SetRtmpUrl(camerad_opts.RtmpUrl),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new camera service")
		return err
	}

	pb.RegisterCameradServiceServer(s, srv)

	log.WithField("listen", camerad_opts.Listen).Infof("metathings camera service listening")
	return s.Serve(lis)
}

func init() {
	camerad_opts = &_cameradOptions{}

	cameradCmd.Flags().StringVarP(&camerad_opts.Listen, "listen", "l", "127.0.0.1:5002", "Metathings Camera Service listening address")
	cameradCmd.Flags().StringVar(&camerad_opts.ServiceConfig.Identityd.Address, "identityd-addr", "mt-api.nayotta.com", "MetaThings Identity Service Address")
	cameradCmd.Flags().StringVar(&camerad_opts.ServiceConfig.Cored.Address, "cored-addr", "mt-api.nayotta.com", "MetaThings Core Service Address")
	cameradCmd.Flags().StringVar(&camerad_opts.Storage.Driver, "storage-driver", "sqlite3", "Storage Driver [sqlite3]")
	cameradCmd.Flags().StringVar(&camerad_opts.Storage.Uri, "storage-uri", "", "Storage URI")
	cameradCmd.Flags().StringVar(&camerad_opts.RtmpUrl, "rtmp-url", "", "RTMP Server URL")

	RootCmd.AddCommand(cameradCmd)
}
