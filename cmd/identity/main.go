package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/identity/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

var (
	bind      string
	ksBaseURL string
)

var (
	rootCmd = cobra.Command{
		Use:   "metathings-identity-service",
		Short: "MetaThings Identity Service Daemon",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runGRPC(); err != nil {
				log.Fatalf("[E] failed to runGRPC: %v\n", err)
			}
		},
	}
)

func runGRPC() error {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("[E] failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	srv := service.NewIdentityService(
		service.SetKeystoneBaseURL(ksBaseURL),
	)

	pb.RegisterIdentityServiceServer(s, srv)
	log.Printf("Listen on %v\n", bind)
	return s.Serve(lis)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&bind, "bind", "b", "127.0.0.1:5000", "Metathings Identity Service binding address")
	rootCmd.PersistentFlags().StringVar(&ksBaseURL, "keystone-base-url", "http://localhost:35357", "Backend Keystone Base URL")

	rootCmd.Execute()
}
