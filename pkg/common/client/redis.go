package client_helper

import (
	"time"

	"github.com/redis/go-redis/v9"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type RedisClient = redis.UniversalClient

func NewRedisClient(args ...interface{}) (RedisClient, error) {
	// redis client
	var addr string
	var db int
	var username string
	var password string
	var protocol int
	var sentinel_username string
	var sentinel_password string
	var client_name string
	var max_retries int
	var min_retry_backoff time.Duration
	var max_retry_backoff time.Duration
	var dial_timeout time.Duration
	var read_timeout time.Duration
	var write_timeout time.Duration

	var pool_fifo bool
	var pool_size int
	var pool_timeout time.Duration
	var pool_min_idle_conns int
	var pool_max_idle_conns int
	var pool_max_active_conns int
	var pool_conn_max_idle_time time.Duration
	var pool_conn_max_life_time time.Duration

	// redis cluster client
	var addrs []string
	var max_redirects int
	var read_only bool
	var route_by_latency bool
	var route_randomly bool

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"addr":              opt_helper.ToString(&addr),
		"db":                opt_helper.ToInt(&db),
		"username":          opt_helper.ToString(&username),
		"password":          opt_helper.ToString(&password),
		"protocol":          opt_helper.ToInt(&protocol),
		"sentinel_username": opt_helper.ToString(&sentinel_username),
		"sentinel_password": opt_helper.ToString(&sentinel_password),
		"client_name":       opt_helper.ToString(&client_name),
		"max_retries":       opt_helper.ToInt(&max_retries),
		"min_retry_backoff": opt_helper.ToDuration(&min_retry_backoff),
		"max_retry_backoff": opt_helper.ToDuration(&max_retry_backoff),
		"dial_timeout":      opt_helper.ToDuration(&dial_timeout),
		"read_timeout":      opt_helper.ToDuration(&read_timeout),
		"write_timeout":     opt_helper.ToDuration(&write_timeout),

		// cluster
		"addrs":            opt_helper.ToStringSlice(&addrs),
		"max_redirects":    opt_helper.ToInt(&max_redirects),
		"read_only":        opt_helper.ToBool(&read_only),
		"route_by_latency": opt_helper.ToBool(&route_by_latency),
		"route_randomly":   opt_helper.ToBool(&route_randomly),

		// pool
		"pool_fifo":               opt_helper.ToBool(&pool_fifo),
		"pool_size":               opt_helper.ToInt(&pool_size),
		"pool_timeout":            opt_helper.ToDuration(&pool_timeout),
		"pool_min_idle_conns":     opt_helper.ToInt(&pool_min_idle_conns),
		"pool_max_idle_conns":     opt_helper.ToInt(&pool_max_idle_conns),
		"pool_max_active_conns":   opt_helper.ToInt(&pool_max_active_conns),
		"pool_conn_max_idle_time": opt_helper.ToDuration(&pool_conn_max_idle_time),
		"pool_conn_max_life_time": opt_helper.ToDuration(&pool_conn_max_life_time),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	if addr != "" {
		addrs = append(addrs, addr)
	}

	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:            addrs,
		DB:               db,
		Username:         username,
		Password:         password,
		ClientName:       client_name,
		Protocol:         protocol,
		SentinelUsername: sentinel_username,
		SentinelPassword: sentinel_password,
		MaxRedirects:     max_redirects,
		MaxRetries:       max_redirects,
		MinRetryBackoff:  min_retry_backoff,
		MaxRetryBackoff:  max_retry_backoff,
		DialTimeout:      dial_timeout,
		ReadTimeout:      read_timeout,
		WriteTimeout:     write_timeout,
		PoolFIFO:         pool_fifo,
		PoolSize:         pool_size,
		PoolTimeout:      pool_timeout,
		MinIdleConns:     pool_min_idle_conns,
		MaxIdleConns:     pool_max_idle_conns,
		MaxActiveConns:   pool_max_active_conns,
		ConnMaxIdleTime:  pool_conn_max_idle_time,
		ConnMaxLifetime:  pool_conn_max_life_time,
	}), nil
}
