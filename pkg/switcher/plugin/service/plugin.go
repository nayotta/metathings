package main

import (
	"net"
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mtp "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/switcher"
	service "github.com/nayotta/metathings/pkg/switcher/service"
)

type _driverOptions struct {
	Name       string
	Descriptor string
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
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

func runSwitcherd() error {
	// TODO(Peer): hard code here now.
	port := strings.SplitAfter(root_opts.Listen, ":")[1]
	ep := "localhost" + ":" + port

	opts := mtp.DefaultOptions()
	opts.Set("name", root_opts.Name)
	opts.Set("log.level", root_opts.Log.Level)
	opts.Set("agent.address", root_opts.ServiceConfig.CoreAgentd.Address)
	opts.Set("metathings.address", root_opts.ServiceConfig.Metathingsd.Address)
	opts.Set("endpoint", ep)
	opts.Set("driver.descriptor", root_opts.Driver.Descriptor)
	opts.Set("driver.name", root_opts.Driver.Name)
	opts.Set("driver", cmd_helper.GetFromStage(v).Sub("driver"))

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
	defer func() {
		srv.Close()
		log.Infof("switcher(core) service closed")
	}()

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
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", constant_helper.CONSTANT_CORE_AGENT_ADDRESS, "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", constant_helper.CONSTANT_METATHINGSD_ADDRESS, "Metathings Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.Driver.Descriptor, "driver-descriptor", "~/.metathins/switcher_driver_descriptor.yaml", "Switcher driver descriptor path")
	rootCmd.PersistentFlags().StringVar(&root_opts.Driver.Name, "driver-name", "", "Switcher driver name")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &switcherServicePlugin{}
}
