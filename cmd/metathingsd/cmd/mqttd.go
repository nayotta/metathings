package cmd

import (
	"context"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	service "github.com/nayotta/metathings/pkg/mqttd/service"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// MqttdOption MqttdOption
type MqttdOption struct {
	// expose detail for viper to unmarshal config file.
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
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
			baseOpt = &optT.BaseOption

			init_service_cmd_option(optT, mqttdOpts)
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
	return mqttdOpt,
		mqttdOpt,
		mqttdOpt,
		mqttdOpt,
		mqttdOpt,
		mqttdOpt,
		mqttdOpt
}

// NewMqttdStorage NewMqttdStorage
func NewMqttdStorage(opt cmd_contrib.StorageOptioner, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.GetDriver(), opt.GetUri(), "logger", logger)
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
			policy.NewEnforcer,
			NewMqttdStorage,
			service.NewMetathingsMqttdService,
		),
		fx.Invoke(
			pb.RegisterDevicedServiceServer,
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
	flags.StringVar(&mqttdOpts.AgentdConfig.Id, "core-id", "", "Core(mqttd) ID")

	RootCmd.AddCommand(mqttdCmd)
}
