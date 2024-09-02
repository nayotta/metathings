package redis_helper

import (
	"github.com/PeerXu/option-go"
	"github.com/redis/go-redis/v9"
)

const (
	OPTION_REDIS_CLIENT = "redisClient"
)

var (
	WithRedisClient, GetRedisClient = option.New[redis.UniversalClient](OPTION_REDIS_CLIENT)
)
