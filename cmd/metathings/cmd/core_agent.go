package cmd

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	service "github.com/bigdatagz/metathings/pkg/core_agent/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

var (
	core_agentd_opts struct {
		bind            string
		core_agent_home string
		core_id         string
	}
)

var (
	coreAgentdCmd = &cobra.Command{
		Use:    "agentd",
		Short:  "Core Agent Service Daemon",
		PreRun: globalPreRunHook,
		Run: func(cmd *cobra.Command, args []string) {
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

	log.WithFields(log.Fields{
		"id":       root_opts.application_credential_id,
		"secret":   root_opts.application_credential_secret,
		"v-id":     V("application_credential_id"),
		"v-secret": V("application_credential_secret"),
	}).
		Debugf("application credential")

	s := grpc.NewServer()
	srv, err := service.NewCoreAgentService(
		service.SetMetathingsAddr(V("addr")),
		service.SetLogLevel(V("log_level")),
		service.SetCoreId(core_agentd_opts.core_id),
		service.SetApplicationCredential(
			V("application_credential_id"),
			V("application_credential_secret"),
		),
	)
	if err != nil {
		return err
	}

	pb.RegisterCoreAgentServiceServer(s, srv)
	log.WithFields(log.Fields{
		"bind":    core_agentd_opts.bind,
		"core_id": core_agentd_opts.core_id,
	}).Infof("metathings core agent service listening")
	return s.Serve(lis)
}

func init() {
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.bind, "bind", "127.0.0.1:5002", "Core Agentd Service binding address")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.core_agent_home, "core-agent-home", "~/.metathings/core_agent/", "Core Agent Home Path")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.core_id, "core-id", "", "Core(Agent) ID")
	coreCmd.AddCommand(coreAgentdCmd)
}
