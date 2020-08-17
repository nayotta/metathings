package metathings_plugin_evaluator

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/objx"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type LuaOperatorTestSuite struct {
	suite.Suite

	op        *LuaOperator
	dat_stor  *dssdk.MockDataStorage
	smpl_stor *dsdk.MockSimpleStorage
	flow      *dsdk.MockFlow
	caller    *dsdk.MockCaller
	gctx      context.Context
	ctx       esdk.Data
	dat       esdk.Data
}

func (s *LuaOperatorTestSuite) SetupTest() {
	s.dat_stor = new(dssdk.MockDataStorage)
	s.smpl_stor = new(dsdk.MockSimpleStorage)
	s.flow = new(dsdk.MockFlow)
	s.caller = new(dsdk.MockCaller)
	s.gctx = context.TODO()
}

func (s *LuaOperatorTestSuite) BeforeTest(suiteName, testName string) {
	if fn, ok := map[string]func(){
		"TestRun":                           s.setupTestRun,
		"TestRunWithDataStorage":            s.setupTestRunWithDataStorage,
		"TestRunWithDeviceDataStorage":      s.setupTestRunWithDeviceDataStorage,
		"TestRunWithAliasDeviceDataStorage": s.setupTestRunWithAliasDeviceDataStorage,
		"TestRunWithSimpleStorage":          s.setupTestRunWithSimpleStorage,
		"TestRunWithDeviceSimpleStorage":    s.setupTestRunWithDeviceSimpleStorage,
		"TestRunWithFlow":                   s.setupTestRunWithFlow,
		"TestRunWithDeviceFlow":             s.setupTestRunWithDeviceFlow,
		"TestRunWithDeviceCaller":           s.setupTestRunWithDeviceCaller,
	}[testName]; ok {
		fn()
	}
}

func (s *LuaOperatorTestSuite) setupOperator(code string) {
	op, err := NewLuaOperator(
		"code", code,
		"data_storage", s.dat_stor,
		"simple_storage", s.smpl_stor,
		"flow", s.flow,
		"caller", s.caller,
	)
	s.Require().Nil(err)
	s.op = op.(*LuaOperator)
}

func (s *LuaOperatorTestSuite) runMainTest() {
	if s.ctx == nil {
		s.ctx, _ = esdk.DataFromMap(nil)
	}
	if s.dat == nil {
		s.dat, _ = esdk.DataFromMap(nil)
	}

	_, err := s.op.Run(s.gctx, s.ctx, s.dat)
	s.Require().Nil(err)
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

	s.setupOperator(code)
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

	dat, err := s.op.Run(s.gctx, ctx, dat)
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

	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a": "b",
			"c": "d",
		}, map[string]interface{}{
			"e": "f",
			"g": float64(42),
		}).Return(nil)

	s.setupOperator(code)
}

func (s *LuaOperatorTestSuite) TestRunWithDataStorage() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithDeviceDataStorage() {
	code := `
local dev = metathings:device("self")
local s = dev:storage("msr", {["a"] = "b"})
s:write({["c"] = "d"})

return {}
`

	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a":            "b",
			"$device_id":   "test",
			"$source_id":   "xxx",
			"$source_type": "yyy",
		}, map[string]interface{}{
			"c": "d",
		}).Return(nil)

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "test",
		},
		"source": map[string]interface{}{
			"id":   "xxx",
			"type": "yyy",
		},
	})
	s.dat, _ = esdk.DataFromMap(nil)
}

func (s *LuaOperatorTestSuite) TestRunWithDeviceDataStorage() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithAliasDeviceDataStorage() {
	code := `
local dev = metathings:device("light")
local s = dev:storage("msr", {["a"] = "b"})
s:write({["c"] = "d"})

return {}
`

	s.dat_stor.On("Write", mock.Anything, "msr",
		map[string]string{
			"a":            "b",
			"$device_id":   "light-id",
			"$source_id":   "xxx",
			"$source_type": "yyy",
		}, map[string]interface{}{
			"c": "d",
		}).Return(nil)

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
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
	s.dat, _ = esdk.DataFromMap(nil)

}

func (s *LuaOperatorTestSuite) TestRunWithAliasDeviceDataStorage() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithSimpleStorage() {
	code := `
local s = metathings:simple_storage()
s:put({
  ["device"] = "device",
  ["prefix"] = "/prefix",
  ["name"] = "name"
}, "hello, world")

s:remove({
  ["device"] = "device",
  ["prefix"] = "/prefix",
  ["name"] = "name"
})

s:rename({
  ["device"] = "device",
  ["prefix"] = "/prefix-src",
  ["name"] = "src"
}, {
  ["device"] = "device",
  ["prefix"] = "/prefix-dst",
  ["name"] = "dst"
})

s:get({
  ["device"] = "device",
  ["prefix"] = "/prefix",
  ["name"] = "name"
})

s:get_content({
  ["device"] = "device",
  ["prefix"] = "/prefix",
  ["name"] = "name"
})

s:list({
  ["device"] = "device",
  ["prefix"] = "/prefix",
  ["name"] = "name",
  ["recursive"] = true,
  ["depth"] = 42
})

return {}
`

	obj := map[string]interface{}{
		"device": map[string]interface{}{
			"id": "device",
		},
		"prefix": "/prefix",
		"name":   "name",
	}

	get_ret := &deviced_pb.Object{}
	get_content_ret := "hello, world"
	list_ret := []*deviced_pb.Object{}

	s.smpl_stor.On("Put", mock.Anything, obj, "hello, world").Return(nil)
	s.smpl_stor.On("Remove", mock.Anything, obj).Return(nil)
	s.smpl_stor.On("Rename", mock.Anything, map[string]interface{}{
		"device": map[string]interface{}{
			"id": "device",
		},
		"prefix": "/prefix-src",
		"name":   "src",
	}, map[string]interface{}{
		"device": map[string]interface{}{
			"id": "device",
		},
		"prefix": "/prefix-dst",
		"name":   "dst",
	}).Return(nil)
	s.smpl_stor.On("Get", mock.Anything, obj).Return(get_ret, nil)
	s.smpl_stor.On("GetContent", mock.Anything, obj).Return(get_content_ret, nil)
	s.smpl_stor.On("List", mock.Anything, map[string]interface{}{
		"device": map[string]interface{}{
			"id": "device",
		},
		"prefix":    "/prefix",
		"name":      "name",
		"recursive": true,
		"depth":     42,
	}).Return(list_ret, nil)

	s.setupOperator(code)
}

func (s *LuaOperatorTestSuite) TestRunWithSimpleStorage() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithDeviceSimpleStorage() {
	code := `
local dev = metathings:device("self")
local s = dev:simple_storage()
s:put({
  ["prefix"] = "/prefix",
  ["name"] = "name"
}, "hello, world")

return {}
`

	s.smpl_stor.On("Put", mock.Anything, map[string]interface{}{
		"device": map[string]interface{}{
			"id": "light",
		},
		"prefix": "/prefix",
		"name":   "name",
	}, "hello, world").Return(nil)

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "light",
		},
	})
	s.dat, _ = esdk.DataFromMap(nil)
}

func (s *LuaOperatorTestSuite) TestRunWithDeviceSimpleStorage() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithFlow() {
	code := `
local flow = metathings:flow("self", "greeting")
flow:push_frame({
  ["text"] = "hello, world!"
})
return {}
`

	s.flow.On("PushFrame", mock.Anything, "hello", "greeting", map[string]interface{}{
		"text": "hello, world!",
	})

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "hello",
		},
	})
	s.dat, _ = esdk.DataFromMap(nil)

}

func (s *LuaOperatorTestSuite) TestRunWithFlow() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) setupTestRunWithDeviceFlow() {
	code := `
local dev = metathings:device("self")
local flow = dev:flow("greeting")
flow:push_frame({
  ["text"] = "hello, world!"
})
return {}
`

	s.flow.On("PushFrame", mock.Anything, "hello", "greeting", map[string]interface{}{
		"text": "hello, world!",
	})

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "hello",
		},
	})
}

func (s *LuaOperatorTestSuite) TestRunWithDeviceFlow() {
	s.runMainTest()
}

func (s *LuaOperatorTestSuite) TestRunWithCallback() {
	tag_prefix := "custom_tag-prefix-"
	code := `
local cb = metathings:callback()
cb:emit({
  ["text"] = "hello, world!",
})
return {}
`
	s.setupOperator(code)

	rr := http.NewServeMux()
	rr.HandleFunc("/webhook", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal("light", r.Header.Get("Custom-Tag-Prefix-Device"))
		s.Equal("sensor", r.Header.Get("Custom-Tag-Prefix-Source"))
		s.Equal("flow", r.Header.Get("Custom-Tag-Prefix-Source-Type"))
		s.Equal("bearer youshallnotpass", r.Header.Get("Token"))

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.Require().Nil(err)
		}
		defer r.Body.Close()

		body := map[string]interface{}{}
		err = json.Unmarshal(buf, &body)
		s.Require().Nil(err)

		bodyx := objx.New(body)
		s.Equal("hello, world!", bodyx.Get("text").String())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))

	ts := httptest.NewServer(rr)
	defer ts.Close()

	s.ctx, _ = esdk.DataFromBytes([]byte(fmt.Sprintf(`
{
  "device": {
    "id": "light"
  },
  "source": {
    "id": "sensor",
    "type": "flow"
  },
  "config": {
    "callback": {
      "name": "webhook",
      "allow_plain_text": true,
      "tag_prefix": "%s",
      "url": "%s",
      "custom_headers": {
        "token": "bearer youshallnotpass"
      }
    }
  }
}
`, tag_prefix, ts.URL+"/webhook")))
	s.dat, _ = esdk.DataFromMap(nil)
	_, err := s.op.Run(s.gctx, s.ctx, s.dat)
	s.Require().Nil(err)
}

func (s *LuaOperatorTestSuite) setupTestRunWithDeviceCaller() {
	code := `
local dev = metathings:device("self")
dev:unary_call("switch", "turn", {
  ["state"] = "on",
})
return {}
`

	s.caller.On("UnaryCall", mock.Anything, "light", "switch", "turn", map[string]interface{}{
		"state": "on",
	})

	s.setupOperator(code)

	s.ctx, _ = esdk.DataFromMap(map[string]interface{}{
		"device": map[string]interface{}{
			"id": "light",
		},
	})
	s.dat, _ = esdk.DataFromMap(nil)
}

func (s *LuaOperatorTestSuite) TestRunWithDeviceCaller() {
	s.runMainTest()
}

func TestLuaOperatorTestSuite(t *testing.T) {
	suite.Run(t, new(LuaOperatorTestSuite))
}
