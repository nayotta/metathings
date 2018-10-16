package cmd

import (
	"context"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	service "github.com/nayotta/metathings/pkg/identityd2/service"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type Identityd2Option struct {
	_rootOption `mapstructure:",squash"`
	Listen      string
	CertFile    string
	KeyFile     string
	Storage     cmd_helper.StorageOption
}

var (
	identityd2_opt *Identityd2Option
)

var (
	identityd2Cmd = &cobra.Command{
		Use:   "identityd2",
		Short: "Identity-NG Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			opt_t := &Identityd2Option{}
			cmd_helper.UnmarshalConfig(opt_t)
			root_opt = &opt_t._rootOption

			if opt_t.Listen == "" {
				opt_t.Listen = identityd2_opt.Listen
			}

			if opt_t.Storage.Driver == "" {
				opt_t.Storage.Driver = identityd2_opt.Storage.Driver
			}

			if opt_t.Storage.Uri == "" {
				opt_t.Storage.Uri = identityd2_opt.Storage.Uri
			}

			if opt_t.CertFile == "" {
				opt_t.CertFile = identityd2_opt.CertFile
			}

			if opt_t.KeyFile == "" {
				opt_t.KeyFile = identityd2_opt.KeyFile
			}

			identityd2_opt = opt_t

			identityd2_opt.Service = "identityd2"
			identityd2_opt.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: cmd_helper.Run("identityd2", runIdentityd2),
	}
)

func GetIdentityd2Option() *Identityd2Option {
	return identityd2_opt
}

func NewListenter(opt *Identityd2Option) (net.Listener, error) {
	return net.Listen("tcp", opt.Listen)
}

func NewCredentials(opt *Identityd2Option) (credentials.TransportCredentials, error) {
	if opt.CertFile != "" && opt.KeyFile != "" {
		return credentials.NewServerTLSFromFile(opt.CertFile, opt.KeyFile)
	}
	return nil, nil
}

func NewLogger(opt *Identityd2Option) (log.FieldLogger, error) {
	return log_helper.NewLogger("identityd2", opt.Log.Level)
}

type NewGrpcServerParams struct {
	fx.In

	Opt   *Identityd2Option
	Lis   net.Listener
	Creds credentials.TransportCredentials
}

func NewGrpcServer(params NewGrpcServerParams, lc fx.Lifecycle, logger log.FieldLogger) *grpc.Server {

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	}

	if params.Creds != nil {
		opts = append(opts, grpc.Creds(params.Creds))
	}

	s := grpc.NewServer(opts...)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go s.Serve(params.Lis)
			logger.Infof("metathings identityd2 service start")
			return nil
		},
		OnStop: func(context.Context) error {
			s.Stop()
			logger.Infof("metathings identityd2 service stop")
			return nil
		},
	})

	return s
}

func NewStorage(opt *Identityd2Option, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.Storage.Driver, opt.Storage.Uri, "logger", logger)
}

func Register(server *grpc.Server, service *service.MetathingsIdentitydService) {
	pb.RegisterIdentitydServiceServer(server, service)
}

func NewMetathingsIdentitydServiceOption(opt *Identityd2Option) *service.MetathingsIdentitydServiceOption {
	return &service.MetathingsIdentitydServiceOption{
		TokenExpire: 1 * time.Hour,
	}
}

func runIdentityd2() error {
	app := fx.New(
		fx.Provide(
			GetIdentityd2Option,
			NewCredentials,
			NewLogger,
			NewListenter,
			NewGrpcServer,
			NewStorage,
			NewMetathingsIdentitydServiceOption,
			service.NewMetathingsIdentitydService,
		),
		fx.Invoke(
			Register,
		),
	)

	app.Run()

	err := app.Err()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	identityd2_opt = &Identityd2Option{}

	flags := identityd2Cmd.Flags()

	flags.StringVarP(&identityd2_opt.Listen, "listen", "l", "127.0.0.1:5000", "Metathings Identity2 Service listening address")
	flags.StringVar(&identityd2_opt.Storage.Driver, "storage-driver", "sqlite3", "Metathings Identity2 Service Storage Driver")
	flags.StringVar(&identityd2_opt.Storage.Uri, "storage-uri", "", "Metathings Identity2 Service Storage URI")
	flags.StringVar(&identityd2_opt.CertFile, "cert-file", "certs/identityd2-server.crt", "Metathings Identity2 Service Credential File")
	flags.StringVar(&identityd2_opt.KeyFile, "key-file", "certs/identityd2-server.key", "Metathings Identity2 Service Key File")

	RootCmd.AddCommand(identityd2Cmd)
}
