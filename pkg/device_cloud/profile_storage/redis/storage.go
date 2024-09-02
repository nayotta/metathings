package metathings_device_cloud_profile_storage_redis

import (
	"github.com/PeerXu/option-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	redis_helper "github.com/nayotta/metathings/pkg/common/redis"
	core "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/core"
	intf "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
)

const NAME = "redis"

type RedisProfileStorage struct {
	defaultProfile string

	client redis.UniversalClient
	logger logrus.FieldLogger
}

func NewRedisProfileStorage(opts ...option.ApplyOption) (intf.ProfileStorage, error) {
	o := option.ApplyWithDefault(core.DefaultNewStorageOptions(), opts...)

	logger, err := log_helper.GetLogger(o)
	if err != nil {
		return nil, err
	}

	client, err := redis_helper.GetRedisClient(o)
	if err != nil {
		return nil, err
	}

	defaultProfile, err := core.GetDefaultProfile(o)
	if err != nil {
		return nil, err
	}

	return &RedisProfileStorage{
		defaultProfile: defaultProfile,

		client: client,
		logger: logger,
	}, nil
}

func init() {
	core.Register(NAME, NewRedisProfileStorage)
}
