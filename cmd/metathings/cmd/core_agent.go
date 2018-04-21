package cmd

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	service "github.com/bigdatagz/metathings/pkg/core_agent/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

type _agentdConfigOptions struct {
	Id   string
	Home string
}

type _coreAgentdOptions struct {
	_rootOptions
	Listen       string
	AgentdConfig _agentdConfigOptions
}

var (
	core_agentd_opts *_coreAgentdOptions
)

var (
	coreAgentdCmd = &cobra.Command{
		Use:   "agentd",
		Short: "Core Agent Service Daemon",
		PreRun: defaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(core_agentd_opts)
			root_opts = &core_agentd_opts._rootOptions
			core_agentd_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCoreAgentd(); err != nil {
				log.WithError(err).Fatalf("failed to start core agentd")
			}
		},
	}
)

func runCoreAgentd() error {
	lis, err := net.Listen("tcp", core_agentd_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	srv, err := service.NewCoreAgentService(
		service.SetMetathingsAddr(root_opts.ServiceConfig.Metathings.Address),
		service.SetLogLevel(root_opts.Log.Level),
		service.SetCoreId(core_agentd_opts.AgentdConfig.Id),
		service.SetApplicationCredential(
			root_opts.ApplicationCredential.Id,
			root_opts.ApplicationCredential.Secret,
		),
	)
	if err != nil {
		return err
	}

	pb.RegisterCoreAgentServiceServer(s, srv)
	log.WithFields(log.Fields{
		"bind":    core_agentd_opts.Listen,
		"core_id": core_agentd_opts.AgentdConfig.Id,
	}).Infof("metathings core agent service listening")
	return s.Serve(lis)
}

func init() {
	core_agentd_opts = &_coreAgentdOptions{}

	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.Listen, "bind", "127.0.0.1:5002", "Core Agentd Service binding address")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.AgentdConfig.Home, "core-agent-home", "~/.metathings/core_agent/", "Core Agent Home Path")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.AgentdConfig.Id, "core-id", "", "Core(Agent) ID")
	coreCmd.AddCommand(coreAgentdCmd)
}
