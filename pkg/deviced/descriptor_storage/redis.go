package metathings_deviced_descriptor_storage

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/go-redis/redis/v8"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type RedisDescriptorStorage struct {
	client client_helper.RedisClient
}

func (rds *RedisDescriptorStorage) context() context.Context {
	return context.TODO()
}

func (rds *RedisDescriptorStorage) sha1_key(sha1 string) string {
	return fmt.Sprintf("/desc/%v", sha1)
}

func (rds *RedisDescriptorStorage) SetDescriptor(sha1 string, body []byte) error {
	ctx := rds.context()
	return rds.client.Set(ctx, rds.sha1_key(sha1), base64.StdEncoding.EncodeToString(body), 0).Err()
}

func (rds *RedisDescriptorStorage) GetDescriptor(sha1 string) ([]byte, error) {
	ctx := rds.context()
	buf, err := rds.client.Get(ctx, rds.sha1_key(sha1)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrDescriptorNotFound
		}

		return nil, err
	}

	return base64.StdEncoding.DecodeString(buf)
}

func NewRedisDescriptorStorage(args ...interface{}) (DescriptorStorage, error) {
	cli, err := client_helper.NewRedisClient(args...)
	if err != nil {
		return nil, err
	}

	return &RedisDescriptorStorage{
		client: cli,
	}, nil
}

func init() {
	register_descriptor_storage_factory("redis", NewRedisDescriptorStorage)
}
