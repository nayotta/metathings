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
	_rootOptions          `mapstructure:",squash"`
	Listen                string
	AgentdConfig          _agentdConfigOptions `mapstructure:"agentd_config"`
	ServiceDescriptorPath string               `mapstructure:"service_descriptor"`
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

			var opt _coreAgentdOptions
			cmd_helper.UnmarshalConfig(&opt)

			if opt.ServiceDescriptorPath == "" {
				opt.ServiceDescriptorPath = core_agentd_opts.ServiceDescriptorPath
			}

			core_agentd_opts = &opt
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
	srv, err := service.NewCoreAgentService(
		service.SetMetathingsAddr(root_opts.ServiceConfig.Metathings.Address),
		service.SetLogLevel(root_opts.Log.Level),
		service.SetCoreAgentHome(core_agentd_opts.AgentdConfig.Home),
		service.SetCoreId(core_agentd_opts.AgentdConfig.Id),
		service.SetApplicationCredential(
			root_opts.ApplicationCredential.Id,
			root_opts.ApplicationCredential.Secret,
		),
		service.SetServiceDescriptorPath(core_agentd_opts.ServiceDescriptorPath),
	)
	if err != nil {
		return err
	}

	errs := make(chan error)
	go func() {
		log.Infof("serve on stram")
		errs <- srv.ServeOnStream()
	}()
	go func() {
		lis, err := net.Listen("tcp", core_agentd_opts.Listen)
		if err != nil {
			errs <- err
			return
		}
		s := grpc.NewServer()
		pb.RegisterCoreAgentServiceServer(s, srv)

		log.WithField("listen", core_agentd_opts.Listen).Infof("metathings core agent service listening")
		errs <- s.Serve(lis)
	}()
	return <-errs
}

func init() {
	core_agentd_opts = &_coreAgentdOptions{}

	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.Listen, "bind", "127.0.0.1:5002", "Core Agentd Service binding address")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.AgentdConfig.Home, "core-agent-home", "~/.metathings/core-agent/", "Core Agent Home Path")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.ServiceDescriptorPath, "service-descriptor-path", "~/.metathings/service_descriptor.yaml", "Core Service Plugin Descriptor")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.AgentdConfig.Id, "core-id", "", "Core(Agent) ID")
	coreCmd.AddCommand(coreAgentdCmd)
}
