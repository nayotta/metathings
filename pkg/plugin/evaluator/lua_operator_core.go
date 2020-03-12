package metathings_plugin_evaluator

import (
	"strconv"
	"sync"

	lua "github.com/yuin/gopher-lua"

	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

const luaMetathingsCoreTypeName = "core"

type luaMetathingsCore struct {
	config esdk.Data
	data   esdk.Data
}

func registerLuaMetathingsCore(L *lua.LState) {
	mt := L.NewTypeMetatable(luaMetathingsCoreTypeName)
	L.SetGlobal("core", mt)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), luaMetathingsCoreMethods))
}

func newLuaMetathingsCore(L *lua.LState, cfg, dat esdk.Data) *lua.LUserData {
	core := &luaMetathingsCore{
		config: cfg,
		data:   dat,
	}
	ud := L.NewUserData()
	ud.Value = core
	L.SetMetatable(ud, L.GetTypeMetatable(luaMetathingsCoreTypeName))
	return ud
}

func checkLuaMetathingsCore(L *lua.LState) *luaMetathingsCore {
	ud := L.CheckUserData(1)
	v, ok := ud.Value.(*luaMetathingsCore)
	if !ok {
		L.ArgError(1, "core expected")
	}

	return v
}

var luaMetathingsCoreMethods = map[string]lua.LGFunction{
	"data":   luaMetathingsCoreGetData,
	"config": luaMetathingsCoreGetConfig,
}

func luaMetathingsCoreGetData(L *lua.LState) int {
	core := checkLuaMetathingsCore(L)
	key := L.CheckString(2)

	ival := core.data.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

func luaMetathingsCoreGetConfig(L *lua.LState) int {
	core := checkLuaMetathingsCore(L)
	key := L.CheckString(2)

	ival := core.config.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

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
