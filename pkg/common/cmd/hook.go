package cmd_helper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type PreRunHookFun func()

var defaultHooks []PreRunHookFun

func PreRunHooks(hooks ...PreRunHookFun) func(*cobra.Command, []string) {
	return func(*cobra.Command, []string) {
		for _, hook := range hooks {
			hook()
		}
	}
}

func DefaultPreRunHooks(hook PreRunHookFun, defaults ...PreRunHookFun) func(*cobra.Command, []string) {
	if len(defaults) == 0 && len(defaultHooks) > 0 {
		defaults = make([]PreRunHookFun, len(defaultHooks))
		copy(defaultHooks, defaults)
	}
	if hook != nil {
		defaults = append([]PreRunHookFun{hook}, defaults...)
	}
	return PreRunHooks(defaults...)
}

func SetDefaultHooks(hooks ...PreRunHookFun) {
	if len(hooks) > 0 {
		defaultHooks = hooks
	}
}

func Run(service string, fn func() error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := fn(); err != nil {
			log.WithError(err).Fatalf("failed to run %v", service)
		}
	}
}

func RunWithArgs(service string, fn func(args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := fn(args); err != nil {
			log.WithError(err).Fatalf("failed to run with args %v", service)
		}
	}
}
