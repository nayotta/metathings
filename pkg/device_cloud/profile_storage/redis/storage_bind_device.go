package metathings_device_cloud_profile_storage_redis

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (ps *RedisProfileStorage) BindDevice(ctx context.Context, deviceNames []string, profileName string) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "BindDevice",
		"devices.name": deviceNames,
		"profile.name": profileName,
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
		pipe.Set(ctx, key, profileName, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.WithError(err).Debugf("failed to bind profile to devices in redis")
		return err
	}

	logger.Tracef("bind profile to devices in redis")

	return nil
}
