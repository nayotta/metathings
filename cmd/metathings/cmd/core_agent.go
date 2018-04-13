package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	service "github.com/bigdatagz/metathings/pkg/core_agent/service"
	core_pb "github.com/bigdatagz/metathings/pkg/proto/core"
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
			token := V("token")
			if core_agentd_opts.core_id == "" {
				ctx := context.Background()
				md := metadata.Pairs("authorization", fmt.Sprintf("mt %v", token))
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

			opts := []service.ServiceOptions{
				service.SetLogLevel(V("log-level")),
				service.SetToken(V("token")),
				service.SetCoreId(core_agentd_opts.core_id),
			}

			_ = service.NewCoreAgentService(opts...)
			log.Debugf("metathings core agentd")
		},
	}
)

func init() {
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.bind, "bind", "127.0.0.1:5002", "Core Agentd Service binding address")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.core_id, "core-id", "", "Core(Agent) ID")
	coreCmd.AddCommand(coreAgentdCmd)
}
