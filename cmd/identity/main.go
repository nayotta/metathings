package main

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/identity/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

var (
	bind            string
	ksAdminBaseURL  string
	ksPublicBaseURL string
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
		service.SetKeystoneAdminBaseURL(ksAdminBaseURL),
		service.SetKeystonePublicBaseURL(ksPublicBaseURL),
	)

	pb.RegisterIdentityServiceServer(s, srv)
	log.Printf("[!] Listen on %v\n", bind)
	return s.Serve(lis)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&bind, "bind", "b", "127.0.0.1:5000", "Metathings Identity Service binding address")
	rootCmd.PersistentFlags().StringVar(&ksAdminBaseURL, "keystone-admin-url", "http://localhost:35357", "Backend Keystone Admin Base URL")
	rootCmd.PersistentFlags().StringVar(&ksPublicBaseURL, "keystone-public-url", "http://localhost:5000", "Backend Keystone Public Base URL")

	rootCmd.Execute()
}
