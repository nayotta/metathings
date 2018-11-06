package cmd

import (
	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	"github.com/spf13/cobra"
)

type DevicedOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
}

func NewDevicedOption() *DevicedOption {
	return &DevicedOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	deviced_opt *DevicedOption
)

var (
	devicedCmd = &cobra.Command{
		Use:   "deviced",
		Short: "Device Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
		}),
		Run: cmd_helper.Run("deviced", runDeviced),
	}
)

func runDeviced() error {
	return nil
}

func init() {
	deviced_opt = NewDevicedOption()

	flags := devicedCmd.Flags()

	flags.StringVarP(deviced_opt.GetListenP(), "listen", "l", "127.0.0.1:5001", "MetaThings Device Service listening address")
	flags.StringVar(deviced_opt.GetStorage().GetDriverP(), "storage-driver", "sqlite3", "MetaThtings Device Service Storage Driver")
	flags.StringVar(deviced_opt.GetStorage().GetUriP(), "storage-uri", "", "MetaThings Deviced Service Storage URI")
	flags.StringVar(deviced_opt.GetCertFileP(), "cert-file", "certs/server.crt", "MetaThings Device Service Credential File")
	flags.StringVar(deviced_opt.GetKeyFileP(), "key-file", "certs/server.key", "MetaThings Device Service Key File")

	RootCmd.AddCommand(devicedCmd)
}
