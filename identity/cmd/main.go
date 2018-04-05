package main

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/identity/service"
)

var (
	bind string
)

var (
	rootCmd = cobra.Command{
		Use:   "metathings-identity-service",
		Short: "MetaThings Identity Service Daemon",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func runGRPC() error {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("[E] failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterIdentityServiceServer(srv, service.NewIdentityService())
	log.Printf("[!] Listen on %v\n", bind)
	return srv.Serve(lis)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&bind, "bind", "b", "127.0.0.1:5000", "Metathings Identity Service binding address")

	rootCmd.Execute()
}
