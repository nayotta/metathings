package state_helper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type stateParserTestSuite struct {
	suite.Suite
	sp StateParser
}

func (suite *stateParserTestSuite) SetupTest() {
	suite.sp = NewStateParser("test", map[int32]string{
		0: "TEST_UNKNOWN",
		1: "TEST_ONLINE",
		2: "TEST_OFFLINE",
	}, map[string]int32{
		"TEST_UNKNOWN": 0,
		"TEST_ONLINE":  1,
		"TEST_OFFLINE": 2,
	})
}

func (suite *stateParserTestSuite) TestToString() {
	suite.Equal("unknown", suite.sp.ToString(0))
	suite.Equal("online", suite.sp.ToString(1))
	suite.Equal("offline", suite.sp.ToString(2))
}

func (suite *stateParserTestSuite) TestToStringUnknown() {
	suite.Equal("unknown", suite.sp.ToString(-1))
}

func (suite *stateParserTestSuite) TestToValue() {
	suite.Equal(int32(0), suite.sp.ToValue("unknown"))
	suite.Equal(int32(1), suite.sp.ToValue("online"))
	suite.Equal(int32(2), suite.sp.ToValue("offline"))
}

func (suite *stateParserTestSuite) TestToValueUnknown() {
	suite.Equal(int32(0), suite.sp.ToValue("xxxx"))
}

func TestStateParserTestSuite(t *testing.T) {
	suite.Run(t, new(stateParserTestSuite))
}
