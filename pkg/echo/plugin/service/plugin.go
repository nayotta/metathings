package main

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	mtp "github.com/bigdatagz/metathings/pkg/core/plugin"
	service "github.com/bigdatagz/metathings/pkg/echo/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type _coreAgentdOptions struct {
	Address string
}

type _metathingsdOptions struct {
	Address string
}

type _serviceConfigOptions struct {
	CoreAgentd  _coreAgentdOptions  `mapstructure:"core_agentd"`
	Metathingsd _metathingsdOptions `mapstructure:"metathingsd"`
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	ServiceConfig          _serviceConfigOptions `mapstructure:"service_config"`
	Listen                 string
	Name                   string
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

var (
	rootCmd = &cobra.Command{
		Use: "echo",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}
			var _opts _rootOptions
			cmd_helper.UnmarshalConfig(&_opts, v)

			if _opts.ServiceConfig.CoreAgentd.Address == "" {
				_opts.ServiceConfig.CoreAgentd.Address = root_opts.ServiceConfig.CoreAgentd.Address
			}

			root_opts = &_opts
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runEchod(); err != nil {
				log.WithError(err).Fatalf("failed to run echo(core) service")
			}
		},
	}
)

func runEchod() error {
	port := strings.SplitAfter(root_opts.Listen, ":")[1]
	ep := "localhost" + ":" + port

	srv, err := service.NewEchoService(
		service.SetName(root_opts.Name),
		service.SetLogLevel(root_opts.Log.Level),
		service.SetAgentdAddr(root_opts.ServiceConfig.CoreAgentd.Address),
		service.SetMetathingsdAddr(root_opts.ServiceConfig.Metathingsd.Address),
		service.SetEndpoint(ep),
	)
	if err != nil {
		return err
	}

	errs := make(chan error)
	go func() {
		errs <- srv.ConnectToAgent()
	}()
	go func() {
		lis, err := net.Listen("tcp", root_opts.Listen)
		if err != nil {
			errs <- err
			return
		}
		s := grpc.NewServer()
		pb.RegisterEchoServiceServer(s, srv)

		log.WithField("listen", root_opts.Listen).Infof("echo(core) service listening")
		errs <- s.Serve(lis)
	}()
	return <-errs
}

func initConfig() {
	if root_opts.Config != "" {
		v.SetConfigFile(root_opts.Config)
		if err := v.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read plugin config")
		}
	}
}

type echoServicePlugin struct{}

func (p *echoServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *echoServicePlugin) Init(opts mtp.PluginOptions) error {
	args := opts.GetStrings("args")
	rootCmd.SetArgs(args)

	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(mtp.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13401", "Echo(Core Plugin) Service listenting address")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Name, "name", "echod", "Core Service Name")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", "agentd.metathings.local:5002", "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", "api.metathings.ai:80", "Metathings Service Address")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &echoServicePlugin{}
}
