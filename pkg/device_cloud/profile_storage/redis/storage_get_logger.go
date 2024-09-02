package metathings_device_cloud_profile_storage_redis

import "github.com/sirupsen/logrus"

func (ps *RedisProfileStorage) GetLogger() logrus.FieldLogger {
	return ps.logger.WithField("#instance", "RedisProfileStorage")
}
