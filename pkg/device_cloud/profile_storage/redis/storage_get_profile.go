package metathings_device_cloud_profile_storage_redis

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	core "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/core"
	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

func (ps *RedisProfileStorage) GetProfile(ctx context.Context, profileName string) (p intf.Profile, err error) {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":      "GetProfile",
		"profile.name": profileName,
	})

	key, err := ps.profileNameToRedisKey(profileName)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal profile to redis key")
		return
	}

	str, err := ps.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Debugf("profile not found")
			err = core.ErrProfileNotFoundFn(profileName)
			return
		}
		logger.WithError(err).Debugf("failed to get profile from redis")
		return
	}

	p, err = ps.redisValToProfile(str)
	if err != nil {
		logger.WithError(err).WithField("profile.string_b64", base64.StdEncoding.EncodeToString([]byte(str))).Debugf("failed to parse profile string to profile")
		return
	}

	logger.Tracef("get profile from redis")

	return
}
