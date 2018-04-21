package cmd_helper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func UnmarshalConfig(dst interface{}) {
	stage := GetStageFromEnv()
	err := viper.Sub(stage).Unmarshal(dst)
	if err != nil {
		log.WithError(err).Fatalf("failed to unmarshal config")
	}
}
