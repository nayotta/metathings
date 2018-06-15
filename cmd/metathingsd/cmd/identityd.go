package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/identityd/service"
	pb "github.com/nayotta/metathings/pkg/proto/identity"
)

type _keystoneOptions struct {
	Url string
}

type _identitydOptions struct {
	_rootOptions `mapstructure:",squash"`
	Listen       string
	Keystone     _keystoneOptions
}

var (
	identityd_opts *_identitydOptions
)

var (
	identitydCmd = &cobra.Command{
		Use:   "identityd",
		Short: "Identity Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(identityd_opts)
			root_opts = &identityd_opts._rootOptions
			identityd_opts.Service = "identityd"
			identityd_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runIdentityd(); err != nil {
				log.WithError(err).Fatalf("failed to run identityd")
			}
		},
	}
)

func runIdentityd() error {
	lis, err := net.Listen("tcp", identityd_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewIdentityService(
		service.SetKeystoneBaseURL(identityd_opts.Keystone.Url),
		service.SetLogLevel(identityd_opts.Log.Level),
	)
	if err != nil {
		return err
	}
	pb.RegisterIdentityServiceServer(s, srv)
	log.WithFields(log.Fields{
		"bind": identityd_opts.Listen,
	}).Infof("metathings identity service listening")
	return s.Serve(lis)
}

func init() {
	identityd_opts = &_identitydOptions{}

	identitydCmd.Flags().StringVarP(&identityd_opts.Listen, "listen", "l", "127.0.0.1:5000", "Metathings Identity Service listening address")
	identitydCmd.Flags().StringVar(&identityd_opts.Keystone.Url, "keystone-url", "http://localhost:35357", "Backend Keystone Base URL")

	RootCmd.AddCommand(identitydCmd)
}
