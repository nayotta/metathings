package metathings_device_cloud_storage

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
)

type redisStorageTestSuite struct {
	suite.Suite
	s *RedisStorage
}

var (
	testModuleID                    = "test-module-id"
	testDeviceID                    = "test-device-id"
	testNotExistsDeviceID           = "test-not-exists-device-id"
	testModuleSession               = int64(0)
	testAnotherModuleSession        = int64(1)
	testDeviceConnectSession        = "test-device-connect-session"
	testAnotherDeviceConnectSession = "test-another-device-connect-session"
)

func (suite *redisStorageTestSuite) SetupTest() {
	logger, err := log_helper.NewLogger("test", "debug")
	suite.Nil(err)

	redis_addr := test_helper.GetTestRedisAddr()
	redis_passwd := test_helper.GetTestRedisPassword()
	redis_db, _ := strconv.Atoi(test_helper.GetTestRedisDB())

	stor, err := NewStorage("redis",
		"logger", logger,
		"addr", redis_addr,
		"passwd", redis_passwd,
		"db", redis_db,
	)
	suite.Nil(err)

	suite.s = stor.(*RedisStorage)
}

func (suite *redisStorageTestSuite) TestHeartbeat() {
	suite.s.opt.Module.Heartbeat.Timeout = 100 * time.Millisecond

	err := suite.s.Heartbeat(testModuleID)
	suite.Nil(err)

	t, err := suite.s.GetHeartbeatAt(testModuleID)
	suite.Nil(err)
	suite.True(int64(time.Now().Sub(t)) < int64(100*time.Millisecond))

	time.Sleep(200 * time.Millisecond)

	t, err = suite.s.GetHeartbeatAt(testModuleID)
	suite.Nil(err)
	suite.True(t.Equal(NOTIME))
}

func (suite *redisStorageTestSuite) TestModuleSession() {
	suite.s.opt.Module.Session.Timeout = 100 * time.Millisecond

	sess, err := suite.s.GetModuleSession(testModuleID)
	suite.Nil(err)
	suite.Equal(sess, int64(0))

	err = suite.s.SetModuleSession(testModuleID, testModuleSession)
	suite.Nil(err)

	sess, err = suite.s.GetModuleSession(testModuleID)
	suite.Nil(err)
	suite.Equal(sess, testModuleSession)

	time.Sleep(200 * time.Millisecond)
	sess, err = suite.s.GetModuleSession(testModuleID)
	suite.Nil(err)
	suite.Equal(sess, int64(0))
}

func (suite *redisStorageTestSuite) TestDeviceConnectSession() {
	suite.s.opt.Device.Session.Timeout = 100 * time.Millisecond

	_, err := suite.s.GetDeviceConnectSession(testDeviceID)
	suite.Equal(err, ErrNotConnected)

	err = suite.s.SetDeviceConnectSession(testDeviceID, testDeviceConnectSession)
	suite.Nil(err)

	sess, err := suite.s.GetDeviceConnectSession(testDeviceID)
	suite.Nil(err)
	suite.Equal(sess, testDeviceConnectSession)

	time.Sleep(200 * time.Millisecond)
	_, err = suite.s.GetDeviceConnectSession(testDeviceID)
	suite.Equal(err, ErrNotConnected)
}

func TestRedisStorageTestSuite(t *testing.T) {
	suite.Run(t, new(redisStorageTestSuite))
}
