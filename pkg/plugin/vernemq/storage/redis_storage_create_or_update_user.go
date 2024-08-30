package metathings_plugin_vernemq_storage

import (
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
)

func (s *RedisStorage) CreateOrUpdateUser(user *User) error {
	logger := s.GetLogger().WithField("#method", "CreateOrUpdateUser")

	ctx := s.getContext()
	key := ParseVernemqRedisKey(user.ClientID, user.Username)
	buf, err := json.Marshal(user)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal json to string")
		return err
	}
	val := string(buf)

	if _, err = s.client.Set(ctx, key, val, 0).Result(); err != nil && !errors.Is(err, redis.Nil) {
		logger.WithError(err).Debugf("failed to create or update user to storage")
		return err
	}

	logger.Tracef("create or update to storage")

	return nil
}
