package metathings_device_cloud_profile_storage_redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

func (ps *RedisProfileStorage) CreateOrUpdateProfile(ctx context.Context, profile intf.Profile) error {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "CreateOrUpdateProfile",
		"profile.name": profile.Name,
	})

	key, err := ps.profileToRedisKey(profile)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal profile to redis key")
		return err
	}

	val, err := ps.profileToRedisVal(profile)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal profile to redis val")
		return err
	}

	_, err = ps.client.Set(ctx, key, val, 0).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.WithError(err).Debugf("failed to create or update profile to redis")
		return err
	}

	logger.Tracef("create or update profile in redis")

	return nil
}
