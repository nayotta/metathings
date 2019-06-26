package metathings_device_cloud_storage

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
)

type RedisStorageOption struct {
	ModuleHeartbeatExpireTime time.Duration
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
	if err := cli.Set(fmt.Sprintf("mtdc.mdl.%v.hb", mdl_id), &now, s.opt.ModuleHeartbeatExpireTime).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) IsConnected(sess string, dev_id string) error {
	panic("unimplemented")
}

func (s *RedisStorage) ConnectDevice(sess string, dev_id string) error {
	panic("unimplemented")
}

func (s *RedisStorage) UnconnectDevice(sess string, dev_id string) error {
	panic("unimplemented")
}

func (s *RedisStorage) GetHeartbeatAt(mdl_id string) (time.Time, error) {
	panic("unimplemented")
}

type RedisStorageFactory struct{}

func (f *RedisStorageFactory) New(args ...interface{}) (Storage, error) {
	return &RedisStorage{}, nil
}

func init() {
	register_storage_factory("redis", new(RedisStorageFactory))
}
