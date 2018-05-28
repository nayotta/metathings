package cmd_helper

import "github.com/nayotta/viper"

func GetStageFromEnv() string {
	stage := viper.GetString("stage")
	if stage == "" {
		stage = "dev"
	}
	return stage
}
