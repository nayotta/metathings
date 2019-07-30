package metathings_device_cloud_storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
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

func (s *RedisStorage) get_redis_client() (*redis.Client, func(), error) {
	cli, err := s.pool.Get()
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get redis client")
		return nil, nil, err
	}

	return cli.(*redis.Client), func() { s.pool.Put(cli) }, nil
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
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.heartbeat(cli, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to heartbeat in redis")
		return err
	}

	s.get_logger().WithField("module", mdl_id).Debugf("module heartbeat")

	return nil
}

func (s *RedisStorage) heartbeat(cli *redis.Client, mdl_id string) error {
	now := time.Now()
	if err := cli.Set(s.module_heartbeat_key(mdl_id), now.UnixNano(), 0).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) IsDeviceConnectSession(dev_id string, sess string) error {
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	return s.is_device_connect_session(cli, dev_id, sess)
}

func (s *RedisStorage) is_device_connect_session(cli *redis.Client, dev_id string, sess string) error {
	res := cli.Get(s.device_connect_session_key(dev_id))

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
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.set_device_connect_session(cli, dev_id, sess)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to connect device in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"device":  dev_id,
		"session": sess,
	}).Debugf("connect device")

	return nil
}

func (s *RedisStorage) set_device_connect_session(cli *redis.Client, dev_id string, sess string) error {
	if err := cli.Set(s.device_connect_session_key(dev_id), sess, s.opt.Device.Session.Timeout).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) UnsetDeviceConnectSession(dev_id string, sess string) error {
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.unset_device_connect_session(cli, dev_id, sess)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to unconnect device in redis")
		return err
	}

	s.get_logger().WithFields(log.Fields{
		"device":  dev_id,
		"session": sess,
	}).Debugf("unconnect device")

	return nil
}

func (s *RedisStorage) unset_device_connect_session(cli *redis.Client, dev_id string, sess string) error {
	err := s.is_device_connect_session(cli, dev_id, sess)
	if err == nil {
		if err = cli.Del(s.device_connect_session_key(dev_id)).Err(); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (s *RedisStorage) GetDeviceConnectSession(dev_id string) (string, error) {
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return "", err
	}
	defer cfn()

	sess, err := s.get_device_connect_session(cli, dev_id)
	if err != nil {
		if err != ErrNotConnected {
			s.get_logger().WithError(err).Debugf("failed to get device connect session")
		}
		return "", err
	}

	return sess, nil
}

func (s *RedisStorage) get_device_connect_session(cli *redis.Client, dev_id string) (string, error) {
	res := cli.Get(s.device_connect_session_key(dev_id))
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return "", ErrNotConnected
		}
		return "", err
	}

	return res.Val(), nil
}

func (s *RedisStorage) GetHeartbeatAt(mdl_id string) (time.Time, error) {
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return NOTIME, err
	}
	defer cfn()

	t, err := s.get_heartbeat_at(cli, mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get heartbeat time in redis")
		return NOTIME, err
	}

	return t, nil
}

func (s *RedisStorage) get_heartbeat_at(cli *redis.Client, mdl_id string) (time.Time, error) {
	var ts int64
	res := cli.Get(s.module_heartbeat_key(mdl_id))
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
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return err
	}
	defer cfn()

	err = s.set_module_session(cli, mdl_id, sess)
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

func (s *RedisStorage) set_module_session(cli *redis.Client, mdl_id string, sess int64) error {
	if err := cli.Set(s.module_session_key(mdl_id), sess, s.opt.Module.Session.Timeout).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) GetModuleSession(mdl_id string) (int64, error) {
	cli, cfn, err := s.get_redis_client()
	if err != nil {
		return 0, err
	}
	defer cfn()

	sess, err := s.get_module_session(cli, mdl_id)
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

func (s *RedisStorage) get_module_session(cli *redis.Client, mdl_id string) (int64, error) {
	res := cli.Get(s.module_session_key(mdl_id))
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
	opt.Module.Session.Timeout = 127 * time.Second
	opt.Device.Session.Timeout = 131 * time.Second
	opt.Redis.Pool.Init = 1
	opt.Redis.Pool.Max = 5

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
		"addr":   opt_helper.ToString(&opt.Redis.Address),
		"passwd": opt_helper.ToString(&opt.Redis.Password),
		"db":     opt_helper.ToInt(&opt.Redis.Db),
	})(args...); err != nil {
		return nil, err
	}

	if pool, err = pool_helper.NewPool(opt.Redis.Pool.Init, opt.Redis.Pool.Max, func() (pool_helper.Client, error) {
		return client_helper.NewRedisClient(
			"addr", opt.Redis.Address,
			"password", opt.Redis.Password,
			"db", opt.Redis.Db,
		)
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
