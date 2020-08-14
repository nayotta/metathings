package metathings_deviced_session_storage

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

func startup_session_key(id string) string {
	return "/metathings/devices/" + id + "/sessions/startup"
}

type RedisSessionStorage struct {
	client redis.Cmdable
}

func (self *RedisSessionStorage) context() context.Context {
	return context.TODO()
}

func (self *RedisSessionStorage) GetStartupSession(id string) (int32, error) {
	ctx := self.context()
	r := self.client.Get(ctx, startup_session_key(id))
	if err := r.Err(); err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	s, err := strconv.ParseInt(r.Val(), 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(s), nil
}

func (self *RedisSessionStorage) SetStartupSessionIfNotExists(id string, sess int32, expire time.Duration) error {
	ctx := self.context()
	r := self.client.SetNX(ctx, startup_session_key(id), sess, expire)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

func (self *RedisSessionStorage) UnsetStartupSession(id string) error {
	ctx := self.context()
	r := self.client.Del(ctx, startup_session_key(id))
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

func (self *RedisSessionStorage) RefreshStartupSession(id string, expire time.Duration) error {
	ctx := self.context()
	r := self.client.Expire(ctx, startup_session_key(id), expire)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

func NewRedisSessionStorage(args ...interface{}) (SessionStorage, error) {
	cli, err := client_helper.NewRedisClient(args...)
	if err != nil {
		return nil, err
	}

	return &RedisSessionStorage{
		client: cli,
	}, nil
}

func init() {
	register_session_storage_factory("redis", NewRedisSessionStorage)
}
