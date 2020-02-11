package metathings_plugin_evaluator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LuaOperatorTestSuite struct {
	suite.Suite

	op *LuaOperator
}

func (s *LuaOperatorTestSuite) SetupTest() {
	code := `
local a, b, c, d, e, ret

a = metathings.data.a
b = metathings.data.b
c = metathings.data["c"]["d"]
e = metathings.data["e"][1]
ret = a + b + c + e

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

	op, _ := NewLuaOperator("code", code)
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) TearDownTest() {
	s.op.Close()
}

func (s *LuaOperatorTestSuite) TestRun() {
	dat, _ := DataFromMap(map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": map[string]interface{}{
			"d": 3,
		},
		"e": []interface{}{4},
	})
	cfg, _ := ConfigFromMap(nil)

	dat, err := s.op.Run(dat, cfg)
	s.Nil(err)
	result_i := dat.Get("result")
	s.NotNil(result_i)
	s.Equal(float64(10), result_i)

	m_i := dat.Get("map")
	s.NotNil(m_i)
	m, ok := m_i.(map[string]interface{})
	s.True(ok)
	test_i, ok := m["test"]
	s.True(ok)
	s.Equal("hello", test_i)

	arr_i := dat.Get("array")
	s.NotNil(arr_i)
	arr, ok := arr_i.([]interface{})
	s.True(ok)
	s.Len(arr, 2)
	s.Equal(float64(1), arr[0])
	s.Equal("world", arr[1])
}

func TestLuaOperatorTestSuite(t *testing.T) {
	suite.Run(t, new(LuaOperatorTestSuite))
}
