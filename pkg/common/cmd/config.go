package cmd_helper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func UnmarshalConfig(dst interface{}, vs ...*viper.Viper) {
	v := GetFromStage(vs...)
	err := v.Unmarshal(dst)
	if err != nil {
		log.WithError(err).Fatalf("failed to unmarshal config")
	}
}

func GetFromStage(vs ...*viper.Viper) *viper.Viper {
	var v *viper.Viper

	if len(vs) == 0 {
		v = viper.GetViper()
	} else {
		v = vs[0]
	}

	stage := GetStageFromEnv()
	v = v.Sub(stage)
	return v
}
