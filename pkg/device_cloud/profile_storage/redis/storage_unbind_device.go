package metathings_device_cloud_profile_storage_redis

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (ps *RedisProfileStorage) UnbindDevice(ctx context.Context, deviceNames []string) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "UnbindDevice",
		"devices.name": deviceNames,
	})

	var keys []string
	for _, deviceName := range deviceNames {
		key, err := ps.deviceNameToBindingProfileRedisKey(deviceName)
		if err != nil {
			logger.WithError(err).Debugf("failed to parse device name to redis key")
			return err
		}
		keys = append(keys, key)
	}

	pipe := ps.client.Pipeline()
	for _, key := range keys {
		pipe.Del(ctx, key)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.WithError(err).Debugf("failed to unbind device in redis")
		return err
	}

	logger.Tracef("unbind profile in redis")

	return nil
}
