package cmd

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/core_agent/service"
	pb "github.com/nayotta/metathings/pkg/proto/core_agent"
)

type _agentdConfigOptions struct {
	Id   string
	Home string
}

type _heartbeatOptions struct {
	Interval int
}

type _coreAgentdOptions struct {
	_rootOptions      `mapstructure:",squash"`
	Listen            string
	AgentdConfig      _agentdConfigOptions `mapstructure:"agentd_config"`
	ServiceDescriptor string               `mapstructure:"service_descriptor"`
	Heartbeat         _heartbeatOptions
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

			if opt.ServiceDescriptor == "" {
				opt.ServiceDescriptor = core_agentd_opts.ServiceDescriptor
			}
			if opt.Heartbeat.Interval == 0 {
				opt.Heartbeat.Interval = core_agentd_opts.Heartbeat.Interval
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
		service.SetServiceDescriptor(core_agentd_opts.ServiceDescriptor),
		service.SetHeartbeatInterval(core_agentd_opts.Heartbeat.Interval),
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
		log.Infof("start heartbeat loop")
		srv.HeartbeatLoop()
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
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.ServiceDescriptor, "service-descriptor", "~/.metathings/service_descriptor.yaml", "Core Service Plugin Descriptor")
	coreAgentdCmd.Flags().StringVar(&core_agentd_opts.AgentdConfig.Id, "core-id", "", "Core(Agent) ID")
	coreAgentdCmd.Flags().IntVar(&core_agentd_opts.Heartbeat.Interval, "heartbeat-interval", 5, "Core(Agent) heartbeat interval")
	coreCmd.AddCommand(coreAgentdCmd)
}
