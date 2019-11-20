package cmd_helper

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	net_helper "github.com/nayotta/metathings/pkg/common/net"
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

	stage := GetStageFromEnv(v)
	v = v.Sub(stage)
	return v
}

func GetEndpoint(typ, host, listen string) string {
	switch typ {
	case "auto":
		return getEndpointAuto(typ, host, listen)
	case "manual":
		return getEndpointManual(typ, host, listen)
	default:
		return getEndpointAuto(typ, host, listen)
	}
}

func getEndpointAuto(typ, host, listen string) string {
	port := strings.SplitAfter(listen, ":")[1]
	host = net_helper.GetLocalIP()
	return host + ":" + port
}

func getEndpointManual(typ, host, listen string) string {
	port := strings.SplitAfter(listen, ":")[1]
	return host + ":" + port
}

func LoadConfigFile(opt cmd_contrib.ConfigOptioner, v *viper.Viper) func() {
	return func() {
		if opt.GetConfig() == "" {
			return
		}

		v.SetConfigFile(opt.GetConfig())
		if err := v.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}
}
