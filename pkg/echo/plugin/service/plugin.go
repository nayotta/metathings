package main

import (
	"net"
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mtp "github.com/nayotta/metathings/pkg/cored/plugin"
	service "github.com/nayotta/metathings/pkg/echo/service"
	pb "github.com/nayotta/metathings/pkg/proto/echo"
)

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
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
	// TODO(Peer): hard code here now.
	port := strings.SplitAfter(root_opts.Listen, ":")[1]
	ep := "localhost" + ":" + port

	opts := mtp.DefaultOptions()
	opts.Set("name", root_opts.Name)
	opts.Set("log.level", root_opts.Log.Level)
	opts.Set("agent.address", root_opts.ServiceConfig.CoreAgentd.Address)
	opts.Set("metathings.address", root_opts.ServiceConfig.Metathingsd.Address)
	opts.Set("endpoint", ep)

	srv, err := service.NewEchoService(opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, srv)

	err = srv.Init()
	if err != nil {
		return err
	}
	log.Debugf("echo(core) service initialized")
	go func() {
		log.Infof("echo(core) service closed")
	}()

	log.WithField("listen", root_opts.Listen).Infof("echo(core) service listening")
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

func (p *echoServicePlugin) Init(opts opt_helper.Option) error {
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
