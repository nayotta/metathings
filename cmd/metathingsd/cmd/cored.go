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
		bind string
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
	srv := service.NewCoreService(
		service.SetLogLevel(V("log_level")),
	)
	pb.RegisterCoreServiceServer(s, srv)

	log.Infof("metathings core service listen on %v", cored_opts.bind)
	return s.Serve(lis)
}

func init() {
	coredCmd.Flags().StringVarP(&cored_opts.bind, "bind", "b", "127.0.0.1:5001", "Metathings Core Service binding address")

	RootCmd.AddCommand(coredCmd)
}
