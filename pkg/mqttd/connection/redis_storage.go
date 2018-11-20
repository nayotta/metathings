package metathingsmqttdconnection

import (
	"github.com/go-redis/redis"
)

type redisStorage struct {
	client *redis.Client
}

// AddBridgeToDevice AddBridgeToDevice
func (that *redisStorage) AddBridgeToDevice(devID, brID string) error {
	tx := that.client.Pipeline()

	tx.SAdd("/devices", devID)
	tx.SAdd("/devices/"+devID+"/bridges", brID)

	if _, err := tx.Exec(); err != nil {
		return err
	}

	return nil
}

// RemoveBridgeFromDevice RemoveBridgeFromDevice
func (that *redisStorage) RemoveBridgeFromDevice(devID, brID string) error {
	if err := that.client.SRem("/devices/"+devID+"/bridges", brID).Err(); err != nil {
		return err
	}

	return nil
}

// ListBridgesFromDevice ListBridgesFromDevice
func (that *redisStorage) ListBridgesFromDevice(devID string) ([]string, error) {
	var brs []string
	var err error

	if brs, err = that.client.SMembers("/devices/" + devID + "/bridges").Result(); err != nil {
		return nil, err
	}

	return brs, nil
}

func newRedisStorage(args ...interface{}) (Storage, error) {
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
	registerStorageFactory("redis", newRedisStorage)
}
