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
	service "github.com/nayotta/metathings/pkg/motor/service"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type _motorDriverOption struct {
	Name string
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	Listen                 string
	Name                   string
	DriverDescriptor       string
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

var (
	rootCmd = &cobra.Command{
		Use: "motor",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}
			var _opts _rootOptions
			cmd_helper.UnmarshalConfig(&_opts, v)

			if _opts.ServiceConfig.CoreAgentd.Address == "" {
				_opts.ServiceConfig.CoreAgentd.Address = root_opts.ServiceConfig.CoreAgentd.Address
			}

			if _opts.DriverDescriptor == "" {
				_opts.DriverDescriptor = root_opts.DriverDescriptor
			}

			root_opts = &_opts
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runMotord(); err != nil {
				log.WithError(err).Fatalf("failed to run motor(core) service")
			}
		},
	}
)

func runMotord() error {
	// TODO(Peer): hard code here now.
	port := strings.SplitAfter(root_opts.Listen, ":")[1]
	ep := "localhost" + ":" + port

	opts := mtp.DefaultOptions()
	opts.Set("name", root_opts.Name)
	opts.Set("log.level", root_opts.Log.Level)
	opts.Set("agent.address", root_opts.ServiceConfig.CoreAgentd.Address)
	opts.Set("metathings.address", root_opts.ServiceConfig.Metathingsd.Address)
	opts.Set("endpoint", ep)
	opts.Set("driver.descriptor", root_opts.DriverDescriptor)
	opts.Set("motors", cmd_helper.GetFromStage(v).Get("motors"))

	srv, err := service.NewMotorService(opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterMotorServiceServer(s, srv)

	err = srv.Init()
	if err != nil {
		return err
	}
	log.Debugf("motor(core) service initialized")
	defer func() {
		srv.Close()
		log.Infof("motor(core) service closed")
	}()

	log.WithField("listen", root_opts.Listen).Infof("motor(core) service listening")
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

type motorServicePlugin struct{}

func (p *motorServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *motorServicePlugin) Init(opts opt_helper.Option) error {
	args := opts.GetStrings("args")
	rootCmd.SetArgs(args)

	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(mtp.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13402", "Motor(Core Plugin) Service listening address")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Name, "name", "motord", "Core Service Name")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", constant_helper.CONSTANT_CORE_AGENTD_ADDRESS, "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", constant_helper.CONSTANT_METATHINGSD_ADDRESS, "Metathings Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.DriverDescriptor, "driver-descriptor", "~/.metathins/motor_driver_descriptor.yaml", "Motor driver descriptor path")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &motorServicePlugin{}
}
