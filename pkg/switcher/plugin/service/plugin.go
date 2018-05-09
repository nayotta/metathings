package main

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	opt_helper "github.com/bigdatagz/metathings/pkg/common/option"
	mtp "github.com/bigdatagz/metathings/pkg/core/plugin"
	pb "github.com/bigdatagz/metathings/pkg/proto/switcher"
	service "github.com/bigdatagz/metathings/pkg/switcher/service"
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

type _driverOptions struct {
	Name       string
	Descriptor string
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	ServiceConfig          _serviceConfigOptions `mapstructure:"service_config"`
	Listen                 string
	Name                   string
	Driver                 _driverOptions
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

var (
	rootCmd = &cobra.Command{
		Use: "switcher",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}
			var _opts _rootOptions
			cmd_helper.UnmarshalConfig(&_opts, v)

			if _opts.ServiceConfig.CoreAgentd.Address == "" {
				_opts.ServiceConfig.CoreAgentd.Address = root_opts.ServiceConfig.CoreAgentd.Address
			}

			if _opts.Driver.Descriptor == "" {
				_opts.Driver.Descriptor = root_opts.Driver.Descriptor
			}

			root_opts = &_opts
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runSwitcherd(); err != nil {
				log.WithError(err).Fatalf("failed to run switcher(core) service")
			}
		},
	}
)

func defaultOptions() opt_helper.Option {
	return opt_helper.Option{
		"heartbeat.interval": 15,
	}
}

func runSwitcherd() error {
	port := strings.SplitAfter(root_opts.Listen, ":")[1]
	ep := "localhost" + ":" + port

	opts := defaultOptions()
	opts.Set("name", root_opts.Name)
	opts.Set("log.level", root_opts.Log.Level)
	opts.Set("agent.address", root_opts.ServiceConfig.CoreAgentd.Address)
	opts.Set("metathings.address", root_opts.ServiceConfig.Metathingsd.Address)
	opts.Set("endpoint", ep)
	opts.Set("driver.descriptor", root_opts.Driver.Descriptor)
	opts.Set("driver.name", root_opts.Driver.Name)
	opts.Set("driver", v.Sub("driver"))

	srv, err := service.NewSwitcherService(opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterSwitcherServiceServer(s, srv)

	err = srv.Init()
	if err != nil {
		return err
	}
	log.Debugf("switcher(core) service initialized")

	log.WithField("listen", root_opts.Listen).Infof("switcher(core) service listening")
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

type switcherServicePlugin struct{}

func (p *switcherServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *switcherServicePlugin) Init(opts opt_helper.Option) error {
	args := opts.GetStrings("args")
	rootCmd.SetArgs(args)

	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(mtp.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13401", "Switcher(Core Plugin) Service listenting address")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Name, "name", "switcherd", "Core Service Name")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", "agentd.metathings.local:5002", "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", "api.metathings.ai:80", "Metathings Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.Driver.Descriptor, "driver-descriptor", "~/.metathins/switcher_driver_descriptor.yaml", "Switcher driver descriptor path")
	rootCmd.PersistentFlags().StringVar(&root_opts.Driver.Name, "driver-name", "", "Switcher driver name")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &switcherServicePlugin{}
}
