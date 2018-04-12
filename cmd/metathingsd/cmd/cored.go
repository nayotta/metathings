package cmd

import (
	"net"

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
		Use:   "cored",
		Short: "Cored Service Daemon",
		Run: func(cmd *cobra.Command, args []string) {
			initialize()
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

	s := grpc.NewServer()
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
