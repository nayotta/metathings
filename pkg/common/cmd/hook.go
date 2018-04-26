package cmd_helper

import (
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
