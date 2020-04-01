package metathings_plugin_evaluator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type DataStorageMock struct {
	mock.Mock
}

func (m *DataStorageMock) Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error {
	m.Called(ctx, measurement, tags, data)
	return nil
}

type LuaOperatorTestSuite struct {
	suite.Suite

	op       *LuaOperator
	dat_stor *DataStorageMock
}

func (s *LuaOperatorTestSuite) BeforeTest(suiteName, testName string) {
	map[string]func(){
		"TestRun":                           s.setupTestRun,
		"TestRunWithDataStorage":            s.setupTestRunWithDataStorage,
		"TestRunWithDeviceDataStorage":      s.setupTestRunWithDeviceDataStorage,
		"TestRunWithAliasDeviceDataStorage": s.setupTestRunWithAliasDeviceDataStorage,
	}[testName]()
}

func (s *LuaOperatorTestSuite) setupTestRun() {
	code := `
local a, b, c, d, e, ret

ca = metathings:context("a")
a = metathings:data("a")
b = metathings:data("b")
c = metathings:data("c.d")
e = metathings:data("e.[0]")
ret = ca + a + b + c + e

return {
  ["result"] = ret,
  ["map"] = {
    ["test"] = "hello",
  },
  ["array"] = {
    [1] = 1,
    [2] = "world"
  }
}
`

	op, _ := NewLuaOperator("code", code, "data_storage", new(DataStorageMock))
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) TearDownTest() {
	s.op.Close()
}

func (s *LuaOperatorTestSuite) TestRun() {
	ctx, _ := esdk.DataFromMap(map[string]interface{}{
		"a": 1,
	})
	dat, _ := esdk.DataFromMap(map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": map[string]interface{}{
			"d": 3,
		},
		"e": []interface{}{4},
	})

	dat, err := s.op.Run(ctx, dat)
	s.Require().Nil(err)
	result_i := dat.Get("result")
	s.NotNil(result_i)
	s.Equal(float64(11), result_i)

	m_i := dat.Get("map")
	s.Require().NotNil(m_i)
	m, ok := m_i.(map[string]interface{})
	s.True(ok)
	test_i, ok := m["test"]
	s.True(ok)
	s.Equal("hello", test_i)

	arr_i := dat.Get("array")
	s.Require().NotNil(arr_i)
	arr, ok := arr_i.([]interface{})
	s.True(ok)
	s.Len(arr, 2)
	s.Equal(float64(1), arr[0])
	s.Equal("world", arr[1])
}

func (s *LuaOperatorTestSuite) setupTestRunWithDataStorage() {
	code := `
local s = metathings:storage("msr", {["a"] = "b"})
s = s:with({ ["c"] = "d" })
s:write({ ["e"] = "f", ["g"] = 42 })

return {}
`

	s.dat_stor = new(DataStorageMock)
	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a": "b",
			"c": "d",
		}, map[string]interface{}{
			"e": "f",
			"g": float64(42),
		}).Return(nil)
	op, err := NewLuaOperator("code", code, "data_storage", s.dat_stor)
	s.Require().Nil(err)
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) TestRunWithDataStorage() {
	ctx, _ := esdk.DataFromMap(nil)
	dat, _ := esdk.DataFromMap(nil)

	_, err := s.op.Run(ctx, dat)
	s.Require().Nil(err)
}

func (s *LuaOperatorTestSuite) setupTestRunWithDeviceDataStorage() {
	code := `
local dev = metathings:device("self")
local s = dev:storage("msr", {["a"] = "b"})
s:write({["c"] = "d"})

return {}
`

	s.dat_stor = new(DataStorageMock)
	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a":            "b",
			"$device_id":   "test",
			"$source_id":   "xxx",
			"$source_type": "yyy",
		}, map[string]interface{}{
			"c": "d",
		}).Return(nil)
	op, err := NewLuaOperator("code", code, "data_storage", s.dat_stor)
	s.Require().Nil(err)
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) TestRunWithDeviceDataStorage() {
	ctx, _ := esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "test",
		},
		"source": map[string]interface{}{
			"id":   "xxx",
			"type": "yyy",
		},
	})
	dat, _ := esdk.DataFromMap(nil)

	_, err := s.op.Run(ctx, dat)
	s.Require().Nil(err)
}

func (s *LuaOperatorTestSuite) setupTestRunWithAliasDeviceDataStorage() {
	code := `
local dev = metathings:device("light")
local s = dev:storage("msr", {["a"] = "b"})
s:write({["c"] = "d"})

return {}
`

	s.dat_stor = new(DataStorageMock)
	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a":            "b",
			"$device_id":   "light-id",
			"$source_id":   "xxx",
			"$source_type": "yyy",
		}, map[string]interface{}{
			"c": "d",
		}).Return(nil)
	op, err := NewLuaOperator("code", code, "data_storage", s.dat_stor)
	s.Require().Nil(err)
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) TestRunWithAliasDeviceDataStorage() {
	ctx, _ := esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "light-id",
		},
		"source": map[string]interface{}{
			"id":   "xxx",
			"type": "yyy",
		},
		"config": map[string]interface{}{
			"alias": map[string]interface{}{
				"device": map[string]interface{}{
					"light": "light-id",
				},
			},
		},
	})
	dat, _ := esdk.DataFromMap(nil)

	_, err := s.op.Run(ctx, dat)
	s.Require().Nil(err)
}

func TestLuaOperatorTestSuite(t *testing.T) {
	suite.Run(t, new(LuaOperatorTestSuite))
}
