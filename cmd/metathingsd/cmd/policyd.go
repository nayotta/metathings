package cmd

import (
	"bufio"
	"context"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/policyd/service"
	pb "github.com/nayotta/metathings/proto/policyd"
)

type PolicydOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	ModelFile                     string `mapstructure:"model_file"`
	PolicyFile                    string `mapstructure:"policy_file"`
}

func NewPolicydOption() *PolicydOption {
	return &PolicydOption{}
}

var (
	policyd_opt *PolicydOption
)

var (
	policydCmd = &cobra.Command{
		Use:   "policyd",
		Short: "Policy Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewPolicydOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			init_service_cmd_option(opt_t, policyd_opt)

			if opt_t.ModelFile == "" {
				opt_t.ModelFile = policyd_opt.ModelFile
			}
			if opt_t.PolicyFile == "" {
				opt_t.PolicyFile = policyd_opt.PolicyFile
			}

			policyd_opt = opt_t
			policyd_opt.SetServiceName("policyd")
			policyd_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if err = runPolicyd(); err != nil {
				log.WithError(err).Fatalf("failed to run policyd service")
			}
		},
	}
)

func GetPolicydOptions() (
	*PolicydOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return policyd_opt,
		policyd_opt,
		policyd_opt,
		policyd_opt,
		policyd_opt,
		policyd_opt,
		policyd_opt
}

func NewMetathingsPolicydServiceOption(opt *PolicydOption) (*service.MetathingsPolicydServiceOption, error) {
	var buf []byte
	var err error

	if buf, err = ioutil.ReadFile(opt.ModelFile); err != nil {
		return nil, err
	}

	var polices []service.Policy
	f, err := os.Open(opt.PolicyFile)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		tokens := strings.Split(line, ",")
		polices = append(polices, service.Policy{
			Role:   strings.TrimSpace(tokens[0]),
			Kind:   strings.TrimSpace(tokens[1]),
			Action: strings.TrimSpace(tokens[2]),
		})
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return &service.MetathingsPolicydServiceOption{
		AdapterDriver: opt.GetStorage().GetDriver(),
		AdapterUri:    opt.GetStorage().GetUri(),
		ModelText:     string(buf),
		Policies:      polices,
	}, nil
}

func runPolicyd() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetPolicydOptions,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("policyd"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewGrpcServer,
			NewMetathingsPolicydServiceOption,
			service.NewMetathingsPolicydService,
		),
		fx.Invoke(
			pb.RegisterPolicydServiceServer,
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
	policyd_opt = NewPolicydOption()

	flags := policydCmd.Flags()

	flags.StringVarP(policyd_opt.GetListenP(), "listen", "l", "127.0.0.1:4001", "Metathings Policy Service listening address")
	flags.StringVar(policyd_opt.GetStorage().GetDriverP(), "storage-driver", "sqlite3", "Metathings Policy Service Storage Driver")
	flags.StringVar(policyd_opt.GetStorage().GetUriP(), "storage-uri", "", "Metathings Policy Service Storage URI")
	flags.StringVar(policyd_opt.GetCertFileP(), "cert-file", "certs/policyd-server.crt", "Metathings Policy Service Credential File")
	flags.StringVar(policyd_opt.GetKeyFileP(), "key-file", "certs/policyd-server.key", "Metathings Policy Service Key File")

	flags.StringVar(&policyd_opt.ModelFile, "model-file", "", "Metathings Policy Service Model File")
	flags.StringVar(&policyd_opt.PolicyFile, "policy-file", "", "Metathings Policy Service Policy File")

	RootCmd.AddCommand(policydCmd)
}
