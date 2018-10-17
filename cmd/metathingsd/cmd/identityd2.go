package cmd

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/identityd2/service"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type Identityd2Option struct {
	// expose detail for viper to unmarshal config file.
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
}

func NewIdentityd2Option() *Identityd2Option {
	return &Identityd2Option{}
}

var (
	identityd2_opt *Identityd2Option
)

var (
	identityd2Cmd = &cobra.Command{
		Use:   "identityd2",
		Short: "Identity-NG Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewIdentityd2Option()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			if opt_t.GetListen() == "" {
				opt_t.SetListen(identityd2_opt.GetListen())
			}

			if opt_t.Storage.Driver == "" {
				opt_t.Storage.Driver = identityd2_opt.Storage.Driver
			}

			if opt_t.Storage.Uri == "" {
				opt_t.Storage.Uri = identityd2_opt.Storage.Uri
			}

			if opt_t.GetCertFile() == "" {
				opt_t.SetCertFile(identityd2_opt.GetCertFile())
			}

			if opt_t.GetKeyFile() == "" {
				opt_t.SetKeyFile(identityd2_opt.GetKeyFile())
			}

			identityd2_opt = opt_t
			identityd2_opt.SetServiceName("identityd2")
			identityd2_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("identityd2", runIdentityd2),
	}
)

func GetIdentityd2Options() (
	*Identityd2Option,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt
}

func NewIdentityd2Storage(opt cmd_contrib.StorageOptioner, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.GetDriver(), opt.GetUri(), "logger", logger)
}

func NewMetathingsIdentitydServiceOption(opt *Identityd2Option) *service.MetathingsIdentitydServiceOption {
	return &service.MetathingsIdentitydServiceOption{
		TokenExpire: 1 * time.Hour,
	}
}

func runIdentityd2() error {
	app := fx.New(
		fx.Provide(
			GetIdentityd2Options,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("identityd2"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			NewIdentityd2Storage,
			NewMetathingsIdentitydServiceOption,
			service.NewMetathingsIdentitydService,
		),
		fx.Invoke(
			pb.RegisterIdentitydServiceServer,
		),
	)

	app.Run()

	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	identityd2_opt = NewIdentityd2Option()

	flags := identityd2Cmd.Flags()

	flags.StringVarP(identityd2_opt.GetListenP(), "listen", "l", "127.0.0.1:5000", "Metathings Identity2 Service listening address")
	flags.StringVar(&identityd2_opt.Storage.Driver, "storage-driver", "sqlite3", "Metathings Identity2 Service Storage Driver")
	flags.StringVar(&identityd2_opt.Storage.Uri, "storage-uri", "", "Metathings Identity2 Service Storage URI")
	flags.StringVar(identityd2_opt.GetCertFileP(), "cert-file", "certs/identityd2-server.crt", "Metathings Identity2 Service Credential File")
	flags.StringVar(identityd2_opt.GetKeyFileP(), "key-file", "certs/identityd2-server.key", "Metathings Identity2 Service Key File")

	RootCmd.AddCommand(identityd2Cmd)
}
