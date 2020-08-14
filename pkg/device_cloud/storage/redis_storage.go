package metathings_device_cloud_storage

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
)

type RedisStorageOption struct {
	Module struct {
		Session struct {
			Timeout time.Duration
		}
	}
	Device struct {
		Session struct {
			Timeout time.Duration
		}
	}
	Redis struct {
		Address  string
		Password string
		Db       int
		Pool     struct {
			Init int
			Max  int
		}
	}
}

type RedisStorage struct {
	opt    *RedisStorageOption
	logger log.FieldLogger
	pool   pool_helper.Pool
}

func (s *RedisStorage) get_logger() log.FieldLogger {
	return s.logger
}

func (s *RedisStorage) context() context.Context {
	return context.TODO()
}

func (s *RedisStorage) get_redis_client() (client_helper.RedisClient, func(), error) {
	cli, err := s.pool.Get()
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get redis client")
		return nil, nil, err
	}

	return cli.(client_helper.RedisClient), func() { s.pool.Put(cli) }, nil
}

func (s *RedisStorage) module_heartbeat_key(mdl_id string) string {
	return fmt.Sprintf("mtdc.mdl.%v.hb", mdl_id)
}

func (s *RedisStorage) device_connect_session_key(dev_id string) string {
	return fmt.Sprintf("mtdc.dev.%v.conn_sess", dev_id)
}

func (s *RedisStorage) module_session_key(mdl_id string) string {
	return fmt.Sprintf("mtdc.mdl.%v.sess", mdl_id)
}

func (s *RedisStorage) Heartbeat(mdl_id string) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.heartbeat(cli, ctx, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to heartbeat in redis")
		return err
	}

	s.get_logger().WithField("module", mdl_id).Debugf("module heartbeat")

	return nil
}

func (s *RedisStorage) heartbeat(cli client_helper.RedisClient, ctx context.Context, mdl_id string) error {
	now := time.Now()
	if err := cli.Set(ctx, s.module_heartbeat_key(mdl_id), now.UnixNano(), 0).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) IsDeviceConnectSession(dev_id string, sess string) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	return s.is_device_connect_session(cli, ctx, dev_id, sess)
}

func (s *RedisStorage) is_device_connect_session(cli client_helper.RedisClient, ctx context.Context, dev_id string, sess string) error {
	res := cli.Get(ctx, s.device_connect_session_key(dev_id))

	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return ErrNotConnected
		}
		return err
	}

	if res.Val() != sess {
		return ErrConnectedByOtherDeviceCloud
	}

	return nil
}

func (s *RedisStorage) SetDeviceConnectSession(dev_id string, sess string) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.set_device_connect_session(cli, ctx, dev_id, sess)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to set device connection session in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"device":  dev_id,
		"session": sess,
	}).Debugf("set device connection session")

	return nil
}

func (s *RedisStorage) set_device_connect_session(cli client_helper.RedisClient, ctx context.Context, dev_id string, sess string) error {
	if err := cli.Set(ctx, s.device_connect_session_key(dev_id), sess, s.opt.Device.Session.Timeout).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) UnsetDeviceConnectSession(dev_id string, sess string) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.unset_device_connect_session(cli, ctx, dev_id, sess)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to unset device connect session in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"device":  dev_id,
		"session": sess,
	}).Debugf("unset device connect session")

	return nil
}

func (s *RedisStorage) unset_device_connect_session(cli client_helper.RedisClient, ctx context.Context, dev_id string, sess string) error {
	err := s.is_device_connect_session(cli, ctx, dev_id, sess)
	if err == nil {
		if err = cli.Del(ctx, s.device_connect_session_key(dev_id)).Err(); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (s *RedisStorage) GetDeviceConnectSession(dev_id string) (string, error) {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return "", err
	}
	defer cfn()

	sess, err := s.get_device_connect_session(cli, ctx, dev_id)
	if err != nil {
		if err != ErrNotConnected {
			s.get_logger().WithError(err).Debugf("failed to get device connect session")
		}
		return "", err
	}

	return sess, nil
}

func (s *RedisStorage) get_device_connect_session(cli client_helper.RedisClient, ctx context.Context, dev_id string) (string, error) {
	res := cli.Get(ctx, s.device_connect_session_key(dev_id))
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return "", ErrNotConnected
		}
		return "", err
	}

	return res.Val(), nil
}

func (s *RedisStorage) GetHeartbeatAt(mdl_id string) (time.Time, error) {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return NOTIME, err
	}
	defer cfn()

	t, err := s.get_heartbeat_at(cli, ctx, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get heartbeat time in redis")
		return NOTIME, err
	}

	return t, nil
}

func (s *RedisStorage) get_heartbeat_at(cli client_helper.RedisClient, ctx context.Context, mdl_id string) (time.Time, error) {
	var ts int64
	res := cli.Get(ctx, s.module_heartbeat_key(mdl_id))
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			ts = 0
		} else {
			return NOTIME, err
		}
	}

	ts, err := strconv.ParseInt(res.Val(), 10, 64)
	if err != nil {
		ts = 0
	}

	return time.Unix(0, ts), nil
}

func (s *RedisStorage) SetModuleSession(mdl_id string, sess int64) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.set_module_session(cli, ctx, mdl_id, sess)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to set module session in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"module":  mdl_id,
		"session": sess,
	}).Debugf("set module session")

	return nil
}

func (s *RedisStorage) set_module_session(cli client_helper.RedisClient, ctx context.Context, mdl_id string, sess int64) error {
	if err := cli.Set(ctx, s.module_session_key(mdl_id), sess, s.opt.Module.Session.Timeout).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) UnsetModuleSession(mdl_id string) error {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.unset_module_session(cli, ctx, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to unset module session in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"module": mdl_id,
	}).Debugf("unset module session")

	return nil
}

func (s *RedisStorage) unset_module_session(cli client_helper.RedisClient, ctx context.Context, mdl_id string) error {
	if err := cli.Del(ctx, s.module_session_key(mdl_id)).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) GetModuleSession(mdl_id string) (int64, error) {
	ctx := s.context()
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return 0, err
	}
	defer cfn()

	sess, err := s.get_module_session(cli, ctx, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get module session in redis")
		return 0, err
	}

	s.get_logger().WithFields(log.Fields{
		"module":  mdl_id,
		"session": sess,
	}).Debugf("get module session")

	return sess, nil
}

func (s *RedisStorage) get_module_session(cli client_helper.RedisClient, ctx context.Context, mdl_id string) (int64, error) {
	res := cli.Get(ctx, s.module_session_key(mdl_id))
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	sess, err := strconv.ParseInt(res.Val(), 10, 64)
	if err != nil {
		return 0, err
	}

	return sess, nil
}

type RedisStorageFactory struct{}

func (f *RedisStorageFactory) New(args ...interface{}) (Storage, error) {
	var pool pool_helper.Pool
	var logger log.FieldLogger
	var err error

	opt := &RedisStorageOption{}
	opt.Module.Session.Timeout = 601 * time.Second
	opt.Device.Session.Timeout = 607 * time.Second
	opt.Redis.Pool.Init = 1
	opt.Redis.Pool.Max = 5

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":                 opt_helper.ToLogger(&logger),
		"module_session_timeout": opt_helper.ToDuration(&opt.Module.Session.Timeout),
		"device_session_timeout": opt_helper.ToDuration(&opt.Device.Session.Timeout),
		"pool_initial":           opt_helper.ToInt(&opt.Redis.Pool.Init),
		"pool_max":               opt_helper.ToInt(&opt.Redis.Pool.Max),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	if pool, err = pool_helper.NewPool(opt.Redis.Pool.Init, opt.Redis.Pool.Max, func() (pool_helper.Client, error) {
		return client_helper.NewRedisClient(args...)
	}); err != nil {
		logger.WithError(err).Debugf("failed to new redis client pool")
		return nil, err
	}

	s := &RedisStorage{
		opt:    opt,
		logger: logger,
		pool:   pool,
	}

	return s, nil
}

func init() {
	register_storage_factory("redis", new(RedisStorageFactory))
}
