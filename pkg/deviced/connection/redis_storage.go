package metathings_deviced_connection

import (
	"context"
	"fmt"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

func bridge_key(device string, startup int32) string {
	return fmt.Sprintf("/devices/%v/startup/%v/bridges", device, startup)
}

type redisStorage struct {
	client client_helper.RedisClient
}

func (self *redisStorage) context() context.Context {
	return context.TODO()
}

func (self *redisStorage) AddBridgeToDevice(dev_id string, sess int32, br_id string) error {
	tx := self.client.Pipeline()
	ctx := self.context()

	tx.SAdd(ctx, "/devices", dev_id)
	tx.SAdd(ctx, bridge_key(dev_id, sess), br_id)

	if _, err := tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) RemoveBridgeFromDevice(dev_id string, sess int32, br_id string) error {
	ctx := self.context()
	if err := self.client.SRem(ctx, bridge_key(dev_id, sess), br_id).Err(); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) ListBridgesFromDevice(dev_id string, sess int32) ([]string, error) {
	var brs []string
	var err error

	ctx := self.context()
	if brs, err = self.client.SMembers(ctx, bridge_key(dev_id, sess)).Result(); err != nil {
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
