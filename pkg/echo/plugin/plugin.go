package main

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	service "github.com/bigdatagz/metathings/pkg/echo/service"
	pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type _coreAgentdOptions struct {
	Address string
}

type _serviceConfigOptions struct {
	CoreAgentd _coreAgentdOptions `mapstructure:"core_agentd"`
}

type _rootOptions struct {
	cmd_helper.RootOptions
	ServiceConfig _serviceConfigOptions `mapstructure:"service_config"`
	Service       string
	Listen        string
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

var (
	rootCmd = &cobra.Command{
		Use: "echo",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runEchod(); err != nil {
				log.WithError(err).Fatalf("failed to run echo(core) service")
			}
		},
	}
)

func runEchod() error {
	logger, err := log_helper.NewLogger("echo", root_opts.Log.Level)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	srv := service.NewEchoService(logger)

	pb.RegisterEchoServiceServer(s, srv)
	logger.WithField("listen", root_opts.Listen).Infof("echo(core) service listening")
	return s.Serve(lis)
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

func (p *echoServicePlugin) Init(opt plugin.Option) error {
	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(plugin.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13401", "Echo(Core Plugin) Service listenting address")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")

	return nil
}

func NewPlugin() plugin.CorePlugin {
	return &echoServicePlugin{}
}
