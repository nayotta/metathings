package client_helper

import "github.com/go-redis/redis"

func NewRedisClient(args ...interface{}) (*redis.Client, error) {
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

	return cli, nil
}
