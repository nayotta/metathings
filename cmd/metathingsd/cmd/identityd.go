package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/identity/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
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
		PreRun: defaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			mode := getModeFromEnv()
			err := viper.Sub(mode).Unmarshal(&identityd_opts)
			if err != nil {
				log.WithError(err).Fatalf("failed to unmarshal config")
			}
			root_opts = &identityd_opts._rootOptions

			identityd_opts.Service = "identityd"
			identityd_opts.Mode = mode
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runIdentityd(); err != nil {
				log.Fatalf("failed to run identityd: %v", err)
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
	srv := service.NewIdentityService(
		service.SetKeystoneBaseURL(identityd_opts.Keystone.Url),
		service.SetLogLevel(identityd_opts.Log.Level),
	)

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
