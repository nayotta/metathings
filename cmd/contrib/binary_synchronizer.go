package cmd_contrib

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/nayotta/metathings/pkg/common/binary_synchronizer"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
)

type NewBinarySynchronizerParams struct {
	fx.In

	Logger log.FieldLogger
	Option map[string]interface{} `name:"binary_synchronizer_option"`
}

type BinarySynchronizerOption struct {
	fx.Out
	Option map[string]interface{} `name:"binary_synchronizer_option"`
}

func NewBinarySynchronizer(p NewBinarySynchronizerParams) (binary_synchronizer.BinarySynchronizer, error) {
	args := cfg_helper.FlattenConfigOption(p.Option, "logger", p.Logger)
	return binary_synchronizer.NewBinarySynchronizer(args...)
}
