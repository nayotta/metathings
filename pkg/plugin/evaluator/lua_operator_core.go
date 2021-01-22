package metathings_plugin_evaluator

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	lua "github.com/yuin/gopher-lua"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
	pb "github.com/nayotta/metathings/proto/deviced"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type luaBaseObject interface {
	MetatableIndex() map[string]lua.LGFunction
}

func luaBindingObjectMethods(L *lua.LState, obj luaBaseObject) (*lua.LTable, *lua.LUserData) {
	mt := L.NewTable()
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), obj.MetatableIndex()))
	ud := L.NewUserData()
	ud.Value = obj
	L.SetMetatable(ud, mt)
	return mt, ud
}

func luaInitOnceObject(L *lua.LState, name string, obj luaBaseObject) {
	_, ud := luaBindingObjectMethods(L, obj)
	L.SetGlobal(name, ud)
}

/*
 * context:
 *   config:
 *     alias:
 *       <type>:
 *         <name>: <id>
 *     # examples:
 *     # device:
 *     #   light: light-id
 *     ...
 *   source:
 *     id: <id>  # type: string
 *     type: <type>  # type: string
 *   token: <token>  # type: string
 *   timestamp: <ts>  # type: int64
 *   tags:  # type: map[string]string
 *     ...
 */

type luaMetathingsCore struct {
	gocontext context.Context

	dat_stor  dssdk.DataStorage
	smpl_stor dsdk.SimpleStorage
	flow      dsdk.Flow
	caller    dsdk.Caller

	context esdk.Data
	data    esdk.Data
}

func newLuaMetathingsCore(args ...interface{}) (*luaMetathingsCore, error) {
	var gctx context.Context
	var ctx, dat esdk.Data
	var ds dssdk.DataStorage
	var ss dsdk.SimpleStorage
	var fw dsdk.Flow
	var caller dsdk.Caller

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"gocontext":      opt_helper.ToContext(&gctx),
		"context":        esdk.ToData(&ctx),
		"data":           esdk.ToData(&dat),
		"data_storage":   dssdk.ToDataStorage(&ds),
		"simple_storage": dsdk.ToSimpleStorage(&ss),
		"flow":           dsdk.ToFlow(&fw),
		"caller":         dsdk.ToCaller(&caller),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCore{
		gocontext: gctx,
		context:   ctx,
		data:      dat,
		dat_stor:  ds,
		smpl_stor: ss,
		flow:      fw,
		caller:    caller,
	}, nil
}

func (c *luaMetathingsCore) check(L *lua.LState) *luaMetathingsCore {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCore)
	if !ok {
		L.ArgError(1, "core expected")
		return nil
	}

	return v
}

func (c *luaMetathingsCore) GetGoContext() context.Context {
	return c.gocontext
}

func (c *luaMetathingsCore) GetGoContextWithToken() context.Context {
	return context_helper.WithToken(c.GetGoContext(), cast.ToString(c.GetContext().Get("token")))
}

func (c *luaMetathingsCore) GetContext() esdk.Data {
	return c.context
}

func (c *luaMetathingsCore) GetData() esdk.Data {
	return c.data
}

func (c *luaMetathingsCore) GetDataStorage() dssdk.DataStorage {
	return c.dat_stor
}

func (c *luaMetathingsCore) GetSimpleStorage() dsdk.SimpleStorage {
	return c.smpl_stor
}

func (c *luaMetathingsCore) GetFlow() dsdk.Flow {
	return c.flow
}

func (c *luaMetathingsCore) GetCaller() dsdk.Caller {
	return c.caller
}

func (c *luaMetathingsCore) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"data":           c.luaGetData,
		"context":        c.luaGetContext,
		"storage":        c.luaNewStorage,
		"simple_storage": c.luaNewSimpleStorage,
		"flow":           c.luaNewFlow,
		"callback":       c.luaNewCallback,
		"device":         c.luaGetDevice,
	}
}

// LUA_FUNCTION: core:data(key#string<optional>)
func (c *luaMetathingsCore) luaGetData(L *lua.LState) int {
	c.check(L)

	var ival interface{}
	if L.GetTop() > 1 {
		key := L.CheckString(2)
		ival = c.GetData().Get(key)
	} else {
		ival = c.GetData().Iter()
	}
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:context(key#string)
//   context body: github.com/nayotta/metathings/pkg/plugin/evaluator/lua_operator_core.go#L40
func (c *luaMetathingsCore) luaGetContext(L *lua.LState) int {
	c.check(L)

	key := L.CheckString(2)

	ival := c.context.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:storage(msr#string, tags#table<optional>) storage
func (c *luaMetathingsCore) luaNewStorage(L *lua.LState) int {
	var tags map[string]string

	c.check(L)

	msr := L.CheckString(2)
	if L.GetTop() > 2 {
		tags_tb := L.CheckTable(3)
		tags = cast.ToStringMapString(parse_ltable_to_string_map(tags_tb))
	} else {
		tags = make(map[string]string)
	}
	s, err := newLuaMetathingsCoreStorage("data_storage", c.dat_stor, "measurement", msr, "tags", tags)
	if err != nil {
		L.RaiseError("failed to new storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, s)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: core:simple_storage(option#table) simple_storage
func (c *luaMetathingsCore) luaNewSimpleStorage(L *lua.LState) int {
	var opt map[string]interface{}

	c.check(L)

	if L.GetTop() > 1 {
		opt_tb := L.CheckTable(2)
		opt = parse_ltable_to_string_map(opt_tb)
	} else {
		opt = make(map[string]interface{})
	}
	s, err := newLuaMetathingsCoreSimpleStorage("immutable_option", opt, "core", c)
	if err != nil {
		L.RaiseError("failed to new simple storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, s)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: core:flow(device_alias#string, flow_name#string) flow
func (c *luaMetathingsCore) luaNewFlow(L *lua.LState) int {
	c.check(L)

	dev_id := c.get_device_id_by_alias(L.CheckString(2))
	if dev_id == "" {
		L.ArgError(2, "unsupported device name or alias")
		return 0
	}

	flw_name := L.CheckString(3)

	f, err := newLuaMetathingsCoreFlow(
		"device", dev_id,
		"flow", flw_name,
		"core", c,
	)
	if err != nil {
		L.RaiseError("failed to new flow")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, f)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: core:callback() callback
func (c *luaMetathingsCore) luaNewCallback(L *lua.LState) int {
	c.check(L)

	ctx := c.GetContext()
	tags := map[string]string{
		"source":      cast.ToString(ctx.Get("source.id")),
		"source_type": cast.ToString(ctx.Get("source.type")),
	}

	if dev_id := c.get_device_id_by_alias("self"); dev_id != "" {
		tags["device"] = dev_id
	}

	cb, err := newLuaMetathingsCoreCallback("tags", tags, "core", c)
	if err != nil {
		L.RaiseError("failed to new callback")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, cb)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: core:device(alias#string) device
func (c *luaMetathingsCore) luaGetDevice(L *lua.LState) int {
	c.check(L)

	dev_id := c.get_device_id_by_alias(L.CheckString(2))
	if dev_id == "" {
		L.ArgError(2, "unsupported device name or alias")
		return 0
	}

	dev, err := newLuaMetathingsCoreDevice("id", dev_id, "core", c)
	if err != nil {
		L.RaiseError("failed to get device")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, dev)
	L.Push(ud)

	return 1
}

func (c *luaMetathingsCore) get_device_id_by_alias(x string) string {
	var y string

	if x == "self" {
		y = cast.ToString(c.GetContext().Get("device.id"))
	} else {
		y = cast.ToString(c.GetContext().Get("config.alias.device." + x))
	}

	return y
}

func parse_string_map_to_ltable(L *lua.LState, xs map[string]interface{}) *lua.LTable {
	tb := L.NewTable()

	for xk, xv := range xs {
		tb.RawSetString(xk, parse_interface_to_lvalue(L, xv))
	}

	return tb
}

func parse_interface_to_lvalue(L *lua.LState, x interface{}) lua.LValue {
	switch x.(type) {
	case nil:
		return lua.LNil
	case bool:
		return lua.LBool(x.(bool))
	case int:
		return lua.LNumber(x.(int))
	case int64:
		return lua.LNumber(x.(int64))
	case float64:
		return lua.LNumber(x.(float64))
	case string:
		return lua.LString(x.(string))
	case map[string]interface{}:
		return parse_string_map_to_ltable(L, x.(map[string]interface{}))
	case []interface{}:
		return parse_interface_slice_to_ltable(L, x.([]interface{}))
	default:
		panic("unimplemented")
	}
}

func parse_interface_slice_to_ltable(L *lua.LState, xs []interface{}) *lua.LTable {
	tb := L.NewTable()

	for _, x := range xs {
		tb.Append(parse_interface_to_lvalue(L, x))
	}

	return tb
}

func parse_ltable_to_string_map(x *lua.LTable) (y map[string]interface{}) {
	y, _ = parse_ltable_to_interface(x).(map[string]interface{})
	return
}

func parse_ltable_to_interface(x *lua.LTable) interface{} {
	if x == nil {
		return nil
	}

	var ys []interface{}
	mys := make(map[string]interface{})
	var copy_array_to_map_once sync.Once
	is_array := true
	next_index := 1

	x.ForEach(func(key, val lua.LValue) {
		if key.Type() != lua.LTNumber {
			is_array = false
		} else {
			if int(lua.LVAsNumber(key)) != next_index {
				is_array = false
			} else {
				next_index++
			}
		}

		if is_array {
			ys = append(ys, parse_lvalue_to_interface(val))
		} else {
			copy_array_to_map_once.Do(func() {
				for i, y := range ys {
					mys[strconv.Itoa(i)] = y
				}
			})
			mys[lua.LVAsString(key)] = parse_lvalue_to_interface(val)
		}
	})

	if is_array {
		return ys
	} else {
		return mys
	}
}

func parse_lvalue_to_interface(x lua.LValue) interface{} {
	switch x.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return lua.LVAsBool(x)
	case lua.LTNumber:
		return float64(lua.LVAsNumber(x))
	case lua.LTString:
		return lua.LVAsString(x)
	case lua.LTTable:
		return parse_ltable_to_interface(x.(*lua.LTable))
	default:
		panic("unimplemented")
	}
}

func parse_ltable_to_pb_object(getter opt_helper.GetImmutableOptioner, x *lua.LTable) (y *deviced_pb.Object) {
	opt := objx.New(parse_ltable_to_string_map(x))
	y = &pb.Object{}

	if dev_id := opt_helper.GetValueWithImmutableOption(getter, opt, "device").Str(); dev_id != "" {
		y.Device = &pb.Device{Id: dev_id}
	}

	if pre := opt_helper.GetValueWithImmutableOption(getter, opt, "prefix").Str(); pre != "" {
		y.Prefix = pre
	}

	if name := opt_helper.GetValueWithImmutableOption(getter, opt, "name").Str(); name != "" {
		y.Name = name
	}

	return y
}

func parse_ltable_to_string_slice(x *lua.LTable) []string {
	panic("unimplemnetd")
}

func parse_pb_message_to_ltable(L *lua.LState, x proto.Message) (y *lua.LTable) {
	s, _ := grpc_helper.JSONPBMarshaler.MarshalToString(x)
	m := map[string]interface{}{}

	json.Unmarshal([]byte(s), &m)

	y = parse_string_map_to_ltable(L, m)

	return
}

func parse_pb_messages_to_ltable(L *lua.LState, xs []proto.Message) (ys *lua.LTable) {
	var is []interface{}

	for _, x := range xs {
		s, _ := grpc_helper.JSONPBMarshaler.MarshalToString(x)
		m := map[string]interface{}{}
		json.Unmarshal([]byte(s), &m)
		is = append(is, m)
	}

	ys = parse_interface_slice_to_ltable(L, is)

	return
}

func toLuaMetathingsCore(c **luaMetathingsCore) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		*c, ok = val.(*luaMetathingsCore)
		if !ok {
			return opt_helper.InvalidArgument(key)
		}
		return nil
	}
}
