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
	pb "github.com/nayotta/metathings/pkg/proto/servo"
	service "github.com/nayotta/metathings/pkg/servo/service"
)

type _servoDriverOption struct {
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
		Use: "servo",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			var _opts _rootOptions
			cmd_helper.UnmarshalConfig(&_opts, v)
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runServod(); err != nil {
				log.WithError(err).Fatalf("failed to run servo(core) service")
			}
		},
	}
)

func runServod() error {
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
	opts.Set("servos", cmd_helper.GetFromStage(v).Get("servos"))

	srv, err := service.NewServoService(opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterServoServiceServer(s, srv)

	err = srv.Init()
	if err != nil {
		return err
	}
	log.Debugf("servo(core) service initialized")
	defer func() {
		srv.Close()
		log.Infof("servo(core) service closed")
	}()

	log.WithField("listen", root_opts.Listen).Infof("servo(core) service listening")
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

type servoServicePlugin struct{}

func (p *servoServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *servoServicePlugin) Init(opts opt_helper.Option) error {
	args := opts.GetStrings("args")
	rootCmd.SetArgs(args)

	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(mtp.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13404", "Servo(Core Plugin) Service listening address")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Name, "name", "servod", "Core Service Name")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", constant_helper.CONSTANT_CORE_AGENT_ADDRESS, "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", constant_helper.CONSTANT_METATHINGSD_ADDRESS, "Metathings Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.DriverDescriptor, "driver-descriptor", "~/.metathins/servo_driver_descriptor.yaml", "Motor driver descriptor path")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &servoServicePlugin{}
}
