package metathings_plugin_evaluator

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
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
 *   timestamp: <ts>  # type: int64
 *   tags:  # type: map[string]string
 *     ...
 */

type luaMetathingsCore struct {
	dat_stor  dssdk.DataStorage
	smpl_stor dsdk.SimpleStorage

	context esdk.Data
	data    esdk.Data
}

func newLuaMetathingsCore(args ...interface{}) (*luaMetathingsCore, error) {
	var ctx, dat esdk.Data
	var ds dssdk.DataStorage
	var ss dsdk.SimpleStorage

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"context":        esdk.ToData(&ctx),
		"data":           esdk.ToData(&dat),
		"data_storage":   dssdk.ToDataStorage(&ds),
		"simple_storage": dsdk.ToSimpleStorage(&ss),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCore{
		context:   ctx,
		data:      dat,
		dat_stor:  ds,
		smpl_stor: ss,
	}, nil
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

func (c *luaMetathingsCore) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"data":           c.luaGetData,
		"context":        c.luaGetContext,
		"storage":        c.luaNewStorage,
		"simple_storage": c.luaNewSimpleStorage,
		"device":         c.luaGetDevice,
	}
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

// LUA_FUNCTION: core:data(key#string)
func (c *luaMetathingsCore) luaGetData(L *lua.LState) int {
	core := c.check(L)
	key := L.CheckString(2)

	ival := core.GetData().Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:context(key#string)
//   context body lookup github.com/nayotta/metathings/pkg/plugin/evaluator/evaluator_impl.go#L31
func (c *luaMetathingsCore) luaGetContext(L *lua.LState) int {
	core := c.check(L)
	key := L.CheckString(2)

	ival := core.context.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:storage(msr#string, tags#table) storage
func (c *luaMetathingsCore) luaNewStorage(L *lua.LState) int {
	msr := L.CheckString(2)
	tags_tb := L.CheckTable(3)
	tags := cast.ToStringMapString(parse_ltable_to_string_map(tags_tb))
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
	if L.GetTop() > 1 {
		opt_tb := L.CheckTable(2)
		opt = parse_ltable_to_string_map(opt_tb)
	} else {
		opt = make(map[string]interface{})
	}
	s, err := newLuaMetathingsCoreSimpleStorage("simple_storage", c.smpl_stor, "immutable_option", opt)
	if err != nil {
		L.RaiseError("failed to new simple storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, s)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: core:device(name_or_alias#string) device
func (c *luaMetathingsCore) luaGetDevice(L *lua.LState) int {
	var dev_id string

	noa := L.CheckString(2)
	if noa == "self" {
		dev_id = cast.ToString(c.GetContext().Get("device.id"))
	} else {
		dev_id = cast.ToString(c.GetContext().Get("config.alias.device." + noa))
	}

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

func parse_pb_message_to_ltable(L *lua.LState, x proto.Message) (y *lua.LTable) {
	s, _ := new(jsonpb.Marshaler).MarshalToString(x)
	m := map[string]interface{}{}

	json.Unmarshal([]byte(s), &m)

	y = parse_string_map_to_ltable(L, m)

	return
}

func parse_pb_messages_to_ltable(L *lua.LState, xs []proto.Message) (ys *lua.LTable) {
	marshaler := new(jsonpb.Marshaler)

	var is []interface{}

	for _, x := range xs {
		s, _ := marshaler.MarshalToString(x)
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
