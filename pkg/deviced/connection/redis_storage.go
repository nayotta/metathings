package metathings_deviced_connection

import (
	"fmt"

	"github.com/go-redis/redis"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

func bridge_key(device string, startup int32) string {
	return fmt.Sprintf("/devices/%v/startup/%v/bridges", device, startup)
}

type redisStorage struct {
	client *redis.Client
}

func (self *redisStorage) AddBridgeToDevice(dev_id string, sess int32, br_id string) error {
	tx := self.client.Pipeline()

	tx.SAdd("/devices", dev_id)
	tx.SAdd(bridge_key(dev_id, sess), br_id)

	if _, err := tx.Exec(); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) RemoveBridgeFromDevice(dev_id string, sess int32, br_id string) error {
	if err := self.client.SRem(bridge_key(dev_id, sess), br_id).Err(); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) ListBridgesFromDevice(dev_id string, sess int32) ([]string, error) {
	var brs []string
	var err error

	if brs, err = self.client.SMembers(bridge_key(dev_id, sess)).Result(); err != nil {
		return nil, err
	}

	return brs, nil
}

func new_redis_storage(args ...interface{}) (Storage, error) {
	cli, err := client_helper.NewRedisClient(args...)
	if err != nil {
		return nil, err
	}

	return &redisStorage{
		client: cli,
	}, nil
}

func init() {
	register_storage_factory("redis", new_redis_storage)
}
