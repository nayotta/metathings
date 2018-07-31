package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/cored/service"
	pb "github.com/nayotta/metathings/pkg/proto/cored"
)

type _coredOptions struct {
	_rootOptions  `mapstructure:",squash"`
	Listen        string
	Storage       cmd_helper.StorageOptions
	ServiceConfig cmd_helper.ServiceConfigOptions `mapstructure:"service_config"`
	Heartbeat     _heartbeatOptions
}

var (
	cored_opts *_coredOptions
)

var (
	coredCmd = &cobra.Command{
		Use:   "cored",
		Short: "Cored Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			var opts _coredOptions
			cmd_helper.UnmarshalConfig(&opts)
			if opts.Heartbeat.CoreAliveTimeout == 0 {
				opts.Heartbeat.CoreAliveTimeout = cored_opts.Heartbeat.CoreAliveTimeout
			}
			if opts.Heartbeat.EntityAliveTimeout == 0 {
				opts.Heartbeat.EntityAliveTimeout = cored_opts.Heartbeat.EntityAliveTimeout
			}
			cored_opts = &opts
			root_opts = &cored_opts._rootOptions
			cored_opts.Service = "cored"
			cored_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCored(); err != nil {
				log.WithError(err).Fatalf("failed to run cored")
			}
		},
	}
)

func runCored() error {
	lis, err := net.Listen("tcp", cored_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewCoredService(
		service.SetLogLevel(cored_opts.Log.Level),
		service.SetIdentitydAddr(cored_opts.ServiceConfig.Identityd.Address),
		service.SetStorage(cored_opts.Storage.Driver, cored_opts.Storage.Uri),
		service.SetApplicationCredential(
			cored_opts.ApplicationCredential.Id,
			cored_opts.ApplicationCredential.Secret,
		),
		service.SetCoreAliveTimeout(cored_opts.Heartbeat.CoreAliveTimeout),
		service.SetEntityAliveTimeout(cored_opts.Heartbeat.EntityAliveTimeout),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new core service")
		return err
	}

	pb.RegisterCoredServiceServer(s, srv)

	log.WithField("listen", cored_opts.Listen).Infof("metathings core service listening")
	return s.Serve(lis)
}

func init() {
	cored_opts = &_coredOptions{}

	coredCmd.Flags().StringVarP(&cored_opts.Listen, "listen", "l", "127.0.0.1:5001", "MetaThings Core Service listening address")
	coredCmd.Flags().StringVar(&cored_opts.ServiceConfig.Identityd.Address, "identityd-addr", "", "MetaThings Identity Service address")
	coredCmd.Flags().StringVar(&cored_opts.Storage.Driver, "storage-driver", "sqlite3", "Storage Driver [sqlite3]")
	coredCmd.Flags().StringVar(&cored_opts.Storage.Uri, "storage-uri", "", "Storage URI")
	coredCmd.Flags().IntVar(&cored_opts.Heartbeat.CoreAliveTimeout, "core-alive-timeout", 30, "Core Agent alive timeout")
	coredCmd.Flags().IntVar(&cored_opts.Heartbeat.EntityAliveTimeout, "entity-alive-timeout", 30, "Entity alive timeout")

	RootCmd.AddCommand(coredCmd)
}
