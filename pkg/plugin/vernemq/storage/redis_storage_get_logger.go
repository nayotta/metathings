package metathings_plugin_vernemq_storage

import "github.com/sirupsen/logrus"

func (s *RedisStorage) GetLogger() logrus.FieldLogger {
	return s.logger
}
