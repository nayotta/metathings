package metathings_device_cloud_profile_storage_redis

import (
	"context"
	"errors"

	"github.com/PeerXu/option-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	core "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/core"
	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

func (ps *RedisProfileStorage) GetProfileByDevice(ctx context.Context, deviceId string, opts ...option.ApplyOption) (p intf.Profile, err error) {
	logger := ps.GetLogger().WithFields(logrus.Fields{
		"#method":   "GetProfileByDevice",
		"device.id": deviceId,
	})

	o := option.ApplyWithDefault(core.DefaultGetProfileByDeviceOptions(), opts...)
	disableDefaultProfile, _ := core.GetDisableDefaultProfile(o)

	key, err := ps.deviceIdToBindingProfileRedisKey(deviceId)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse device id to redis key")
		return
	}

	profileName, err := ps.client.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.WithError(err).Debugf("failed to get profile from redis")
			return
		}

		if disableDefaultProfile {
			err = core.ErrDeviceNotFoundFn(deviceId)
			logger.WithError(err).Debugf("device not bind for any profile")
			return
		}

		profileName = ps.defaultProfile
	}

	p, err = ps.GetProfile(ctx, profileName)
	if err != nil {
		logger.WithError(err).Debugf("failed to get profile from storage")
		return
	}

	logger.Tracef("get profile by device")

	return
}
