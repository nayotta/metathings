package metathings_device_cloud_profile_storage_redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func (ps *RedisProfileStorage) DeleteProfile(ctx context.Context, profileName string) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "DeleteProfile",
		"profile.name": profileName,
	})

	key, err := ps.profileNameToRedisKey(profileName)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal profile name to redis key")
		return err
	}

	_, err = ps.client.Del(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.WithError(err).Debugf("failed to delete profile in redis")
		return err
	}

	logger.Tracef("delete profile in redis")

	return nil
}
