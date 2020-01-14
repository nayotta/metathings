package cmd

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	service "github.com/nayotta/metathings/pkg/identityd2/service"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type Identityd2Option struct {
	// expose detail for viper to unmarshal config file.
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Init                          int
}

func NewIdentityd2Option() *Identityd2Option {
	return &Identityd2Option{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
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

			init_service_cmd_option(opt_t, identityd2_opt)
			opt_t.Init = identityd2_opt.Init

			identityd2_opt = opt_t
			identityd2_opt.SetServiceName("identityd2")
			identityd2_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if identityd2_opt.Init > 0 {
				if err = initIdentityd2(); err != nil {
					log.WithError(err).Fatalf("failed to init identityd2 service")
				}
			} else {
				if err = runIdentityd2(); err != nil {
					log.WithError(err).Fatalf("failed to run identityd2 service")
				}
			}
		},
	}
)

func GetIdentityd2Options() (
	*Identityd2Option,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.WebhookServiceOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt,
		identityd2_opt
}

type NewIdentityd2StorageParams struct {
	fx.In

	Option cmd_contrib.StorageOptioner
	Logger log.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewIdentityd2Storage(p NewIdentityd2StorageParams) (storage.Storage, error) {
	return storage.NewStorage(p.Option.GetDriver(), p.Option.GetUri(), "logger", p.Logger, "tracer", p.Tracer)
}

func NewMetathingsIdentitydServiceOption(opt *Identityd2Option) *service.MetathingsIdentitydServiceOption {
	return &service.MetathingsIdentitydServiceOption{
		TokenExpire:      1 * time.Hour,
		CredentialExpire: 100 * 365 * 24 * time.Hour, // 100 years.
	}
}

func NewIdentityd2Backend(cli_fty *client_helper.ClientFactory, logger log.FieldLogger) (policy.Backend, error) {
	base_backend, err := policy.NewBackend(
		"casbin",
		"logger", logger,
		"client_factory", cli_fty,
		"casbin_enforcer_handler", int32(0),
	)
	if err != nil {
		return nil, err
	}
	logger.Debugf("new casbin backend")

	vc := cmd_helper.GetFromStage().Sub("cache")
	if vc == nil {
		return base_backend, nil
	}

	var cache_backend policy.Backend
	drv := vc.GetString("driver")
	switch drv {
	case "mongo":
		cache_backend, err = policy.NewBackend(
			"cache",
			"logger", logger,
			"mongo_uri", vc.GetString("uri"),
			"mongo_database", vc.GetString("database"),
			"mongo_collection", vc.GetString("collection"),
			"backend", base_backend,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, opt_helper.InvalidArgument("driver")
	}
	logger.Debugf("new cache backend")

	return cache_backend, nil
}

func initIdentityd2() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetIdentityd2Options,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("identityd2"),
			cmd_contrib.NewClientFactory,
			NewIdentityd2Backend,
			NewIdentityd2Storage,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, stor storage.Storage, bck policy.Backend, logger log.FieldLogger) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						var ok bool
						var err error
						ctx := context.TODO()

						if ok, err = stor.IsInitialized(ctx); err != nil {
							return err
						} else if ok {
							return nil
						}

						dom_id_str := const_helper.DEFAULT_DOMAIN
						dom_name_str := const_helper.DEFAULT_DOMAIN
						dom_alias_str := const_helper.DEFAULT_DOMAIN
						dom_parent_id_str := ""

						dom := &storage.Domain{
							Id:       &dom_id_str,
							Name:     &dom_name_str,
							Alias:    &dom_alias_str,
							ParentId: &dom_parent_id_str,
						}

						if _, err = stor.CreateDomain(ctx, dom); err != nil {
							return err
						}

						sysadmin_id_str := id_helper.NewId()
						sysadmin_name_str := "sysadmin"
						sysadmin_alias_str := "sysadmin"
						sysadmin := &storage.Role{
							Id:    &sysadmin_id_str,
							Name:  &sysadmin_name_str,
							Alias: &sysadmin_alias_str,
						}

						if sysadmin, err = stor.CreateRole(ctx, sysadmin); err != nil {
							return err
						}

						admin_id_str := id_helper.NewId()
						admin_name_str := "admin"
						admin_alias_str := "admin"
						admin_passwd_str := passwd_helper.MustParsePassword("admin")

						admin := &storage.Entity{
							Id:       &admin_id_str,
							Name:     &admin_name_str,
							Alias:    &admin_alias_str,
							Password: &admin_passwd_str,
						}

						if admin, err = stor.CreateEntity(ctx, admin); err != nil {
							return err
						}

						if err = stor.AddEntityToDomain(ctx, dom_id_str, *admin.Id); err != nil {
							return err
						}

						if err = stor.AddRoleToEntity(ctx, *admin.Id, *sysadmin.Id); err != nil {
							return err
						}

						if err = bck.AddRoleToEntity(ctx, admin, sysadmin); err != nil {
							return err
						}

						if err = stor.Initialize(ctx); err != nil {
							return err
						}

						return nil
					},
				})
			},
		),
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		return err
	}

	return nil
}

func runIdentityd2() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetIdentityd2Options,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("identityd2"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewValidator,
			cmd_contrib.NewWebhookService,
			NewIdentityd2Storage,
			NewIdentityd2Backend,
			NewMetathingsIdentitydServiceOption,
			service.NewMetathingsIdentitydService,
		),
		fx.Invoke(
			pb.RegisterIdentitydServiceServer,
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
	identityd2_opt = NewIdentityd2Option()

	flags := identityd2Cmd.Flags()

	flags.StringVarP(identityd2_opt.GetListenP(), "listen", "l", "127.0.0.1:5000", "Metathings Identity2 Service listening address")
	flags.StringVar(identityd2_opt.GetStorage().GetDriverP(), "storage-driver", "sqlite3", "Metathings Identity2 Service Storage Driver")
	flags.StringVar(identityd2_opt.GetStorage().GetUriP(), "storage-uri", "", "Metathings Identity2 Service Storage URI")
	flags.StringVar(identityd2_opt.GetCertFileP(), "cert-file", "certs/server.crt", "Metathings Identity2 Service Credential File")
	flags.StringVar(identityd2_opt.GetKeyFileP(), "key-file", "certs/server.key", "Metathings Identity2 Service Key File")

	flags.CountVar(&identityd2_opt.Init, "init", "Initial Metathings Identity2 Service")

	RootCmd.AddCommand(identityd2Cmd)
}
