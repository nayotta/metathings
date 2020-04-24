package pool_helper

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Close() error {
	c.Called()
	return nil
}

type PoolTestSuite struct {
	suite.Suite
	pool Pool
}

func (s *PoolTestSuite) SetupTest() {
	var err error

	s.pool, err = NewPool(3, 5, func() (Client, error) {
		return new(MockClient), nil
	})

	s.Require().Nil(err)
}

func (s *PoolTestSuite) BeforeTest(suiteName, testName string) {
	map[string]func(){
		"TestPool":         s.setupTestPool,
		"TestPutWithFull":  s.setupTestPutWithFull,
		"TestGetWithEmpty": s.setupTestGetWithEmpty,
	}[testName]()
}

func (s *PoolTestSuite) setupTestPool() {}

func (s *PoolTestSuite) TestPool() {
	s.Equal(3, s.pool.Size())
	cli, err := s.pool.Get()
	s.Require().Nil(err)
	s.Equal(2, s.pool.Size())
	s.pool.Put(cli)
	s.Equal(3, s.pool.Size())
}

func (s *PoolTestSuite) setupTestPutWithFull() {
	s.pool.Put(new(MockClient))
	s.pool.Put(new(MockClient))
}

func (s *PoolTestSuite) TestPutWithFull() {
	cli := new(MockClient)
	cli.On("Close")

	s.Equal(5, s.pool.Size())
	s.pool.Put(cli)
	s.Equal(5, s.pool.Size())
}

func (s *PoolTestSuite) setupTestGetWithEmpty() {
	s.pool.Get()
	s.pool.Get()
	s.pool.Get()
}

func (s *PoolTestSuite) TestGetWithEmpty() {
	s.Equal(0, s.pool.Size())
	_, err := s.pool.Get()
	s.Require().Nil(err)
	s.Equal(0, s.pool.Size())
}

func TestPoolTestSuite(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}
