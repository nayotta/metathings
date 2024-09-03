package metathings_device_cloud_profile_storage_redis

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (ps *RedisProfileStorage) BindDevices(ctx context.Context, deviceIds []string, profileName string) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "BindDevice",
		"devices.id":   deviceIds,
		"profile.name": profileName,
	})

	var keys []string
	for _, deviceId := range deviceIds {
		key, err := ps.deviceIdToBindingProfileRedisKey(deviceId)
		if err != nil {
			logger.WithError(err).Debugf("failed to parse device id to redis key")
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
