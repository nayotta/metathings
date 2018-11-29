package cmd

import (
	"context"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	connection "github.com/nayotta/metathings/pkg/mqttd/connection"
	service "github.com/nayotta/metathings/pkg/mqttd/service"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// MqttdOption MqttdOption
type MqttdOption struct {
	// expose detail for viper to unmarshal config file.
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	MqttBridge                    struct {
		Bridge map[string]interface{}
	}
}

// NewMqttdOption NewMqttdOption
func NewMqttdOption() *MqttdOption {
	return &MqttdOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	mqttdOpts *MqttdOption
)

func initMqttBridgeConfig(opt *MqttdOption) {
	mmbc := map[string]interface{}{}
	vmbc := cmd_helper.GetFromStage().Sub("mqtt_bridge").Sub("bridge")
	for _, key := range vmbc.AllKeys() {
		mmbc[key] = vmbc.Get(key)
	}
	opt.MqttBridge.Bridge = mmbc
}

func parseMqttdBridgeOption(x map[string]interface{}) (string, []interface{}, error) {
	var key string
	var val interface{}
	var name string
	var ok bool

	y := []interface{}{}

	if val, ok = x["name"]; !ok {
		return "", nil, ErrInvalidArgument
	}

	if name, ok = val.(string); !ok {
		return "", nil, ErrInvalidArgument
	}

	for key, val = range x {
		if key == "name" {
			continue
		}

		y = append(y, key, val)
	}

	return name, y, nil
}

// NewMqttBridge NewMqttBridge
func NewMqttBridge(opt *MqttdOption, logger log.FieldLogger) (connection.MqttBridge, error) {
	var args []interface{}
	var err error
	var br connection.MqttBridge

	if _, args, err = parseMqttdBridgeOption(opt.MqttBridge.Bridge); err != nil {
		return nil, err
	}
	args = append(args, "logger", logger)

	if br, err = connection.NewMqttBridge(args); err != nil {
		return nil, err
	}
	if err = br.InitMqttBridge(); err != nil {
		return nil, err
	}

	return br, nil
}

var (
	mqttdCmd = &cobra.Command{
		Use:   "mqttd",
		Short: "Mqttd Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {

			if base_opt.Config == "" {
				return
			}

			optT := NewMqttdOption()
			cmd_helper.UnmarshalConfig(&optT)
			base_opt = &optT.BaseOption

			init_service_cmd_option(optT, mqttdOpts)
			initMqttBridgeConfig(optT)

			mqttdOpts = optT
			mqttdOpts.SetServiceName("mqttd")
			mqttdOpts.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runMqttd(); err != nil {
				log.WithError(err).Fatalf("failed to run mqttd")
			}
		},
	}
)

// GetMqttdOptions GetMqttdOptions
func GetMqttdOptions() (
	*MqttdOption,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
) {
	return mqttdOpts,
		mqttdOpts,
		mqttdOpts,
		mqttdOpts,
		mqttdOpts,
		mqttdOpts,
		mqttdOpts
}

// NewMqttdStorage NewMqttdStorage
func NewMqttdStorage(opt cmd_contrib.StorageOptioner, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.GetDriver(), opt.GetUri(), "logger", logger)
}

// NewMetathingsMqttdServiceOption NewMetathingsMqttdServiceOption
func NewMetathingsMqttdServiceOption(opt *MqttdOption) *service.MetathingsMqttdServiceOption {
	return &service.MetathingsMqttdServiceOption{}
}

func runMqttd() error {
	app := fx.New(
		fx.Provide(
			GetMqttdOptions,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("mqttd"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewTokener,
			token_helper.NewTokenValidator,
			NewMetathingsMqttdServiceOption,
			policy.NewEnforcer,
			NewMqttdStorage,
			NewMqttBridge,
			service.NewMetathingsMqttdService,
		),
		fx.Invoke(
			pb.RegisterMqttdServiceServer,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}
	defer app.Stop(context.Background())

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	mqttdOpts = NewMqttdOption()

	flags := mqttdCmd.Flags()

	flags.StringVarP(mqttdOpts.GetListenP(), "listen", "l", "127.0.0.1:5001", "MetaThings Mqttd Service listening address")
	flags.StringVar(mqttdOpts.GetStorage().GetDriverP(), "storage-driver", "sqlite3", "MetaThtings Mqttd Service Storage Driver")
	flags.StringVar(mqttdOpts.GetStorage().GetUriP(), "storage-uri", "", "MetaThings Mqttd Service Storage URI")
	flags.StringVar(mqttdOpts.GetCertFileP(), "cert-file", "certs/server.crt", "MetaThings Mqttd Service Credential File")
	flags.StringVar(mqttdOpts.GetKeyFileP(), "key-file", "certs/server.key", "MetaThings Mqttd Service Key File")

	RootCmd.AddCommand(mqttdCmd)
}
