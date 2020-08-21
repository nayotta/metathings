package client_helper

import (
	"time"

	"github.com/go-redis/redis/v8"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type RedisClient interface {
	redis.Cmdable
	Close() error
}

func NewRedisClient(args ...interface{}) (RedisClient, error) {
	// redis client
	var addr string
	var db int
	var username string
	var password string
	var max_retries int
	var min_retry_backoff time.Duration
	var max_retry_backoff time.Duration
	var dial_timeout time.Duration
	var read_timeout time.Duration
	var write_timeout time.Duration

	// redis cluster client
	var is_cluster bool
	var addrs []string
	var max_redirects int

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"addr":              opt_helper.ToString(&addr),
		"db":                opt_helper.ToInt(&db),
		"username":          opt_helper.ToString(&username),
		"password":          opt_helper.ToString(&password),
		"max_retries":       opt_helper.ToInt(&max_retries),
		"min_retry_backoff": opt_helper.ToDuration(&min_retry_backoff),
		"max_retry_backoff": opt_helper.ToDuration(&max_retry_backoff),
		"dial_timeout":      opt_helper.ToDuration(&dial_timeout),
		"read_timeout":      opt_helper.ToDuration(&read_timeout),
		"write_timeout":     opt_helper.ToDuration(&write_timeout),
		"is_cluster":        opt_helper.ToBool(&is_cluster),
		"addrs":             opt_helper.ToStringSlice(&addrs),
		"max_redirects":     opt_helper.ToInt(&max_redirects),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	if is_cluster {
		opt := &redis.ClusterOptions{
			Addrs:           addrs,
			Username:        username,
			Password:        password,
			MaxRedirects:    max_redirects,
			MaxRetries:      max_redirects,
			MinRetryBackoff: min_retry_backoff,
			MaxRetryBackoff: max_retry_backoff,
			DialTimeout:     dial_timeout,
			ReadTimeout:     read_timeout,
			WriteTimeout:    write_timeout,
		}
		cli := redis.NewClusterClient(opt)
		return cli, nil
	} else {
		opt := &redis.Options{
			Addr:            addr,
			DB:              db,
			Username:        username,
			Password:        password,
			MaxRetries:      max_retries,
			MinRetryBackoff: min_retry_backoff,
			MaxRetryBackoff: max_retry_backoff,
			DialTimeout:     dial_timeout,
			ReadTimeout:     read_timeout,
			WriteTimeout:    write_timeout,
		}
		cli := redis.NewClient(opt)
		return cli, nil
	}
}
