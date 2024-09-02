package metathings_plugin_vernemq_storage

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func (s *RedisStorage) RemoveUser(clientid, username string) error {
	logger := s.GetLogger().WithFields(logrus.Fields{
		"#method":   "RemoveUser",
		"client-id": clientid,
		"username":  username,
	})
	ctx := s.getContext()
	key := ParseVernemqRedisKey(clientid, username)

	_, err := s.client.Del(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.WithError(err).Debugf("failed to delete user from storage")
		return err
	}

	logger.Tracef("delete user from storage")

	return nil
}
