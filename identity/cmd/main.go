package main

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/identity/service"
	pb "github.com/bigdatagz/metathings/proto/identity"
)

var (
	bind         string
	ksAdminAddr  string
	ksPublicAddr string
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
		service.SetKeystoneAdminAddress(ksAdminAddr),
		service.SetKeystonePublicAddress(ksPublicAddr),
	)

	pb.RegisterIdentityServiceServer(s, srv)
	log.Printf("[!] Listen on %v\n", bind)
	return s.Serve(lis)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&bind, "bind", "b", "127.0.0.1:5000", "Metathings Identity Service binding address")
	rootCmd.PersistentFlags().StringVar(&ksAdminAddr, "keystone-admin-addr", "", "Backend Keystone Admin Address")
	rootCmd.PersistentFlags().StringVar(&ksPublicAddr, "keystone-public-addr", "", "Backend Keystone Public Address")

	rootCmd.Execute()
}
