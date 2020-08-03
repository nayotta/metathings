package metathings_evaluatord_sdk

import (
	"context"

	"github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type DummyDataLauncher struct {
	logger logrus.FieldLogger
}

func (ddl *DummyDataLauncher) Launch(ctx context.Context, src Resource, dat Data) error {
	ddl.logger.WithFields(logrus.Fields{
		"source":      src.GetId(),
		"source_type": src.GetType(),
		"device":      ExtractDevice(ctx),
	}).Debugf("launch data")

	return nil
}

func NewDummyDataLauncher(args ...interface{}) (DataLauncher, error) {
	var logger logrus.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
	})(args); err != nil {
		return nil, err
	}

	return &DummyDataLauncher{
		logger: logger,
	}, nil
}

func init() {
	registry_data_launcher("dummy", NewDummyDataLauncher)
}
