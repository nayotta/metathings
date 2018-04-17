package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/core/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/core"
)

var (
	cored_opts struct {
		bind           string
		identityd_addr string
	}
)

var (
	coredCmd = &cobra.Command{
		Use:    "cored",
		Short:  "Cored Service Daemon",
		PreRun: globalPreRunHook,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCored(); err != nil {
				log.Fatalf("failed to run cored: %v", err)
			}
		},
	}
)

func runCored() error {
	lis, err := net.Listen("tcp", cored_opts.bind)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewCoreService(
		service.SetIdentitydAddr(cored_opts.identityd_addr),
		service.SetLogLevel(V("log_level")),
		service.SetApplicationCredential(
			V("application_credential_id"),
			V("application_credential_secret"),
		),
	)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new core service")
		return err
	}

	pb.RegisterCoreServiceServer(s, srv)

	log.WithFields(log.Fields{
		"bind": cored_opts.bind,
	}).Infof("metathings core service listening")
	return s.Serve(lis)
}

func init() {
	coredCmd.Flags().StringVarP(&cored_opts.bind, "bind", "b", "127.0.0.1:5001", "MetaThings Core Service binding address")
	coredCmd.Flags().StringVar(&cored_opts.identityd_addr, "identityd-addr", "", "MetaThings Identity Service address")
	coredCmd.MarkFlagRequired("identityd-addr")

	RootCmd.AddCommand(coredCmd)
}
