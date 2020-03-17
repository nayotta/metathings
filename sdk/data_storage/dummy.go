package metathings_data_storage_sdk

import (
	"context"

	log "github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
)

type DummyDataStorage struct {
	logger log.FieldLogger
}

func (s *DummyDataStorage) get_logger() log.FieldLogger {
	return s.logger
}

func (s *DummyDataStorage) Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error {
	s.get_logger().WithFields(log.Fields{
		"measurement": measurement,
		"tags":        tags,
		"data":        data,
	}).Debugf("write")

	return nil
}

func NewDummyDataStorage(args ...interface{}) (DataStorage, error) {
	logger, err := log_helper.NewLogger("dummy-data-storage", "debug")
	if err != nil {
		return nil, err
	}

	return &DummyDataStorage{
		logger: logger,
	}, nil
}

func init() {
	registry_data_storage_factory("dummy", NewDummyDataStorage)
}
