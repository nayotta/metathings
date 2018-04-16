package cmd

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	service "github.com/bigdatagz/metathings/pkg/core_agent/service"
	core_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

var (
	core_agentd_opts struct {
		bind    string
		core_id string
	}
)

var (
	coreAgentdCmd = &cobra.Command{
		Use:    "agentd",
		Short:  "Core Agent Service Daemon",
		PreRun: globalPreRunHook,
		Run: func(cmd *cobra.Command, args []string) {
			if core_agentd_opts.core_id == "" {
				ctx := context.Background()
				md := metadata.Pairs("authorization", fmt.Sprintf("mt %v", V("token")))
				ctx = metadata.NewOutgoingContext(ctx, md)

				opts := []grpc.DialOption{grpc.WithInsecure()}
				conn, err := grpc.Dial(root_opts.addr, opts...)
				if err != nil {
					log.Fatalf("failed to dial to %v: %v", root_opts.addr, err)
				}
				defer conn.Close()
				req := &core_pb.CreateCoreRequest{}
				cli := core_pb.NewCoreServiceClient(conn)
				res, err := cli.CreateCore(ctx, req)
				if err != nil {
					log.Fatalf("failed to create core: %v", err)
				}
				core_agentd_opts.core_id = res.Core.Id
			}

			if err := runCoreAgentd(); err != nil {
				log.Fatalf("%v", err)
			}
		},
	}
)

func runCoreAgentd() error {
	lis, err := net.Listen("tcp", core_agentd_opts.bind)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	srv := service.NewCoreAgentService(
		service.SetToken(V("token")),
		service.SetLogLevel(V("log_level")),
		service.SetCoreId(core_agentd_opts.core_id),
	)

	pb.RegisterCoreAgentServiceServer(s, srv)
	log.WithFields(log.Fields{
		"bind":    core_agentd_opts.bind,
		"core_id": core_agentd_opts.core_id,
	}).Infof("metathings core agent service listening")
	return s.Serve(lis)
}

func init() {
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.bind, "bind", "127.0.0.1:5002", "Core Agentd Service binding address")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.core_id, "core-id", "", "Core(Agent) ID")
	coreCmd.AddCommand(coreAgentdCmd)
}
