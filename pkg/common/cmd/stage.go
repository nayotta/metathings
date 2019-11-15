package cmd_helper

import "github.com/spf13/viper"

func GetStageFromEnv(vs ...*viper.Viper) string {
	var v *viper.Viper

	if len(vs) > 0 {
		v = vs[0]
	} else {
		v = viper.GetViper()
	}

	stage := v.GetString("stage")
	if stage == "" {
		stage = "dev"
	}
	return stage
}
