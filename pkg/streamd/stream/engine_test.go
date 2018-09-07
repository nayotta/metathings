package stream_manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testLuaEngineSuite struct {
	suite.Suite
	engine     *LuaEngine
	filter_bad string
	metadata0  StreamData
	data0      StreamData
	filter0    string
	filter1    string
	filter2    string
}

func (self *testLuaEngineSuite) SetupTest() {
	self.engine = NewLuaEngine()

	self.filter_bad = "x x x"
	self.metadata0 = NewStreamData(map[string]interface{}{
		"test": "test",
	})
	self.data0 = NewStreamData(map[string]interface{}{
		"value": "1",
	})

	self.filter0 = `metadata.test=="test"`
	self.filter1 = `tonumber(value)==1`
	self.filter2 = `tonumber(value)<1`
}

func (self *testLuaEngineSuite) TestBadFilter() {
	_, err := self.engine.Filter(self.filter_bad, self.metadata0, self.data0)
	self.Error(err)
}

func (self *testLuaEngineSuite) TestPassFilter0() {
	pass, err := self.engine.Filter(self.filter0, self.metadata0, self.data0)
	self.Nil(err)
	self.True(pass)
}

func (self *testLuaEngineSuite) TestPassFilter1() {
	pass, err := self.engine.Filter(self.filter1, self.metadata0, self.data0)
	self.Nil(err)
	self.True(pass)
}

func (self *testLuaEngineSuite) TestNotPassFilter2() {
	pass, err := self.engine.Filter(self.filter2, self.metadata0, self.data0)
	self.Nil(err)
	self.False(pass)
}

func TestLuaEngine(t *testing.T) {
	suite.Run(t, new(testLuaEngineSuite))
}
