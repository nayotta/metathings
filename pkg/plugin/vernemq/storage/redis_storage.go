package metathings_plugin_vernemq_storage

import (
	"github.com/PeerXu/option-go"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	redis_helper "github.com/nayotta/metathings/pkg/common/redis"
)

type RedisStorage struct {
	client redis.UniversalClient
	logger logrus.FieldLogger
}

func NewRedisStorage(opts ...option.ApplyOption) (Storage, error) {
	o := option.Apply(opts...)

	client, err := redis_helper.GetRedisClient(o)
	if err != nil {
		return nil, err
	}

	logger, err := log_helper.GetLogger(o)
	if err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: client,
		logger: logger,
	}, nil
}

func init() {
	Register("redis", NewRedisStorage)
}
