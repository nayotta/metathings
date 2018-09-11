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
	luanch0    string
	luanch1    string
	luanch2    string
	luanch3    string
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
	self.filter1 = `tonumber(data.value)==1`
	self.filter2 = `tonumber(data.value)<1`

	self.luanch0 = `
function luanch(metadata, data)
return { test = "test" }
end`
	self.luanch1 = `
function luanch(metadata, data)
return { test = metadata.test }
end`
	self.luanch2 = `
function luanch(metadata, data)
return { value = data.value }
end`
	self.luanch3 = `
function luanch(metadata, data)
return { value = data.value+1 }
end`
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

func (self *testLuaEngineSuite) TestLuanch0() {
	ret, err := self.engine.Luanch(self.luanch0, self.metadata0, self.data0)
	self.Nil(err)
	self.Equal("test", ret.AsString("test"))
}

func (self *testLuaEngineSuite) TestLuanch1() {
	ret, err := self.engine.Luanch(self.luanch1, self.metadata0, self.data0)
	self.Nil(err)
	self.Equal(self.metadata0.AsString("test"), ret.AsString("test"))
}

func (self *testLuaEngineSuite) TestLuanch2() {
	ret, err := self.engine.Luanch(self.luanch2, self.metadata0, self.data0)
	self.Nil(err)
	self.Equal(self.data0.AsInt("value"), ret.AsInt("value"))
}

func (self *testLuaEngineSuite) TestLuanch3() {
	ret, err := self.engine.Luanch(self.luanch3, self.metadata0, self.data0)
	self.Nil(err)
	self.Equal(self.data0.AsInt("value")+1, ret.AsInt("value"))
}

func TestLuaEngine(t *testing.T) {
	suite.Run(t, new(testLuaEngineSuite))
}
