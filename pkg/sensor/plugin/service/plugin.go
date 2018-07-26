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
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
	service "github.com/nayotta/metathings/pkg/sensor/service"
)

type _sensorDriverOption struct {
	Name string
}

type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	Listen                 string
	Endpoint               cmd_helper.EndpointOptions
	Name                   string
	DriverDescriptor       string
}

var (
	root_opts *_rootOptions
	v         *viper.Viper
)

var (
	rootCmd = &cobra.Command{
		Use: "sensor",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			var _opts _rootOptions
			cmd_helper.UnmarshalConfig(&_opts, v)

			if _opts.ServiceConfig.CoreAgentd.Address == "" {
				_opts.ServiceConfig.CoreAgentd.Address = root_opts.ServiceConfig.CoreAgentd.Address
			}

			if _opts.ServiceConfig.Metathingsd.Address == "" {
				_opts.ServiceConfig.Metathingsd.Address = root_opts.ServiceConfig.Metathingsd.Address
			}

			if _opts.DriverDescriptor == "" {
				_opts.DriverDescriptor = root_opts.DriverDescriptor
			}

			root_opts = &_opts
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runSensord(); err != nil {
				log.WithError(err).Fatalf("failed to run sensor(core) service")
			}
		},
	}
)

func runSensord() error {
	ep := cmd_helper.GetEndpoint(root_opts.Endpoint.Type, root_opts.Endpoint.Host, root_opts.Listen)

	opts := mtp.DefaultOptions()
	opts.Set("name", root_opts.Name)
	opts.Set("log.level", root_opts.Log.Level)
	opts.Set("agent.address", root_opts.ServiceConfig.CoreAgentd.Address)
	opts.Set("metathings.address", root_opts.ServiceConfig.Metathingsd.Address)
	opts.Set("endpoint", ep)
	opts.Set("driver.descriptor", root_opts.DriverDescriptor)
	opts.Set("application_credential.id", root_opts.ApplicationCredential.Id)
	opts.Set("application_credential.secret", root_opts.ApplicationCredential.Secret)
	opts.Set("sensors", cmd_helper.GetFromStage(v).Get("sensors"))

	srv, err := service.NewSensorService(opts)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", root_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterSensorServiceServer(s, srv)

	err = srv.Init()
	if err != nil {
		return err
	}
	log.Debugf("sensor(core) service initialized")
	defer func() {
		srv.Close()
		log.Infof("sensor(core) service closed")
	}()

	log.WithField("listen", root_opts.Listen).Infof("sensor(core) service listening")
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

type sensorServicePlugin struct{}

func (p *sensorServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *sensorServicePlugin) Init(opts opt_helper.Option) error {
	args := opts.GetStrings("args")

	rootCmd.SetArgs(args)

	v = viper.New()
	root_opts = &_rootOptions{}
	v.AutomaticEnv()
	v.SetEnvPrefix(mtp.METATHINGS_PLUGIN_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.BindEnv("stage")

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&root_opts.Listen, "listen", "l", "0.0.0.0:13405", "Sensor(Core Plugin) Service listening address")
	rootCmd.PersistentFlags().StringVar(&root_opts.Endpoint.Type, "endpoint-type", "auto", "Get endpoint address type[auto, manual]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Endpoint.Host, "endpoint-host", "", "Endpoint host address (work on endpoint-type is manual)")
	rootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	rootCmd.PersistentFlags().StringVar(&root_opts.Name, "name", "sensord", "Core Service Name")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.CoreAgentd.Address, "agent-addr", constant_helper.CONSTANT_CORE_AGENT_ADDRESS, "Core Agent Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathingsd.Address, "metathings-addr", constant_helper.CONSTANT_METATHINGSD_ADDRESS, "Metathings Service Address")
	rootCmd.PersistentFlags().StringVar(&root_opts.DriverDescriptor, "driver-descriptor", "~/.metathins/sensor_driver_descriptor.yaml", "Motor driver descriptor path")

	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &sensorServicePlugin{}
}
