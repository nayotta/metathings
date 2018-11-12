package metathings_deviced_connection

import (
	"github.com/go-redis/redis"
)

type redisStorage struct {
	client *redis.Client
}

func (self *redisStorage) AddBridgeToDevice(dev_id, br_id string) error {
	tx := self.client.Pipeline()

	tx.SAdd("/devices", dev_id)
	tx.SAdd("/devices/"+dev_id+"/bridges", br_id)

	if _, err := tx.Exec(); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) RemoveBridgeFromDevice(dev_id, br_id string) error {
	if err := self.client.SRem("/devices/"+dev_id+"/bridges", br_id).Err(); err != nil {
		return err
	}

	return nil
}

func (self *redisStorage) ListBridgesFromDevice(dev_id string) ([]string, error) {
	var brs []string
	var err error

	if brs, err = self.client.SMembers("/devices/" + dev_id + "/bridges").Result(); err != nil {
		return nil, err
	}

	return brs, nil
}

func new_redis_storage(args ...interface{}) (Storage, error) {
	var ok bool
	var key string
	var val interface{}

	var opts redis.Options

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	for i := 0; i < len(args); i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return nil, ErrInvalidArgument
		}
		val = args[i+1]

		switch key {
		case "addr":
			opts.Addr, ok = val.(string)
			if !ok {
				return nil, ErrInvalidArgument
			}
		case "db":
			opts.DB, ok = val.(int)
			if !ok {
				return nil, ErrInvalidArgument
			}
		case "password":
			opts.Password, ok = val.(string)
			if !ok {
				return nil, ErrInvalidArgument
			}
		}
	}

	cli := redis.NewClient(&opts)

	return &redisStorage{
		client: cli,
	}, nil
}

func init() {
	register_storage_factory("redis", new_redis_storage)
}
