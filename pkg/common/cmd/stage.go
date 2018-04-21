package cmd_helper

import "github.com/spf13/viper"

func GetStageFromEnv() string {
	stage := viper.GetString("stage")
	if stage == "" {
		stage = "dev"
	}
	return stage
}
