package metathings_device_cloud_profile_storage_redis

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (ps *RedisProfileStorage) UnbindDevices(ctx context.Context, deviceIds []string) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":    "UnbindDevice",
		"devices.id": deviceIds,
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
