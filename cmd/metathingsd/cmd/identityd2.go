package cmd

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
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

			init_cmd_option(opt_t, identityd2_opt)
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
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
) {
	return identityd2_opt,
		identityd2_opt,
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

func initIdentityd2() error {
	app := fx.New(
		fx.Provide(
			GetIdentityd2Options,
			cmd_contrib.NewLogger("identityd2"),
			cmd_contrib.NewClientFactory,
			policy.NewEnforcer,
			NewIdentityd2Storage,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, stor storage.Storage, enf policy.Enforcer, logger log.FieldLogger) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						var err error

						if err = enf.Initialize(); err != nil {
							return err
						}

						dom_id_str := "default"
						dom_name_str := "default"
						dom_alias_str := "default"
						dom_parent_id_str := ""

						dom := &storage.Domain{
							Id:       &dom_id_str,
							Name:     &dom_name_str,
							Alias:    &dom_alias_str,
							ParentId: &dom_parent_id_str,
						}

						if _, err = stor.CreateDomain(dom); err != nil {
							return err
						}

						if err = enf.AddGroup(dom_id_str, policy.UNGROUPED); err != nil {
							return err
						}

						if err = enf.AddObjectToKind(dom_id_str, service.KIND_DOMAIN); err != nil {
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

						if _, err = stor.CreateRole(sysadmin); err != nil {
							return err
						}

						if err = enf.AddObjectToKind(sysadmin_id_str, service.KIND_ROLE); err != nil {
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

						if _, err = stor.CreateEntity(admin); err != nil {
							return err
						}

						if err = enf.AddObjectToKind(admin_id_str, service.KIND_ENTITY); err != nil {
							return err
						}

						if err = stor.AddRoleToEntity(admin_id_str, sysadmin_id_str); err != nil {
							return err
						}

						if err = enf.AddSubjectToRole(admin_id_str, sysadmin_name_str); err != nil {
							return err
						}

						if err = stor.AddEntityToDomain(dom_id_str, admin_id_str); err != nil {
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
		fx.Provide(
			GetIdentityd2Options,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("identityd2"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			NewIdentityd2Storage,
			NewMetathingsIdentitydServiceOption,
			policy.NewEnforcer,
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
