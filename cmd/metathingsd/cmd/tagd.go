package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	service "github.com/nayotta/metathings/pkg/tagd/service"
	storage "github.com/nayotta/metathings/pkg/tagd/storage"
	pb "github.com/nayotta/metathings/proto/tagd"
)

type TagdOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Storage                       map[string]interface{}
}

func NewTagdOption() *TagdOption {
	return &TagdOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	tagd_opt *TagdOption
)

func init_tagd_storage(opt *TagdOption) {
	ms := map[string]interface{}{}
	vs := cmd_helper.GetFromStage().Sub("storage")
	for _, key := range vs.AllKeys() {
		if key == "driver" {
			ms[key] = vs.GetString(key)
		} else {
			ms[key] = vs.Get(key)
		}
	}
	opt.Storage = ms
}

var (
	tagdCmd = &cobra.Command{
		Use:   "tagd",
		Short: "Tag Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewTagdOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			init_tagd_storage(opt_t)

			tagd_opt = opt_t
			tagd_opt.SetServiceName("tagd")
			tagd_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("tagd", runTagd),
	}
)

func GetTagdOptions() (
	*TagdOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.OpentracingOptioner,
) {

	return tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt
}

func NewTagdStorage(opt *TagdOption, logger log.FieldLogger) (storage.Storage, error) {
	drv, ok := opt.Storage["driver"]
	if !ok {
		return nil, storage.ErrUnknownDriver
	}

	args := []interface{}{"logger", logger}
	for k, v := range opt.Storage {
		if k == "driver" {
			continue
		}
		args = append(args, k, v)
	}

	stor, err := storage.NewStorage(drv.(string), args...)
	if err != nil {
		return nil, err
	}

	return stor, nil
}

func runTagd() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetTagdOptions,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("tagd"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			token_helper.NewTokenValidator,
			authorizer.NewAuthorizer,
			cmd_contrib.NewValidator,
			NewTagdStorage,
			service.NewMetathingsTagdService,
		),
		fx.Invoke(
			pb.RegisterTagdServiceServer,
		),
	)

	if err := app.Start(context.TODO()); err != nil {
		return err
	}
	defer app.Stop(context.TODO())

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	tagd_opt = NewTagdOption()

	flags := tagdCmd.Flags()

	flags.StringVarP(tagd_opt.GetListenP(), "listen", "l", "127.0.0.1:5002", "Metathings Tag Service listening address")
	flags.StringVar(tagd_opt.GetCertFileP(), "cert-file", "certs/server.crt", "Metathings Tag Service Credential File")
	flags.StringVar(tagd_opt.GetKeyFileP(), "key-file", "certs/server.key", "Metathings Tag Service Key File")

	RootCmd.AddCommand(tagdCmd)
}
