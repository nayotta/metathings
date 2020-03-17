package metathings_plugin_evaluator

import (
	"bytes"
	"context"
	"encoding/gob"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cast"
	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
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
 *     storage:
 *       buckets:
 *       - bucket-id
 *       ...
 *       alias:
 *         bucket-alias: bucket-id
 *         ...
 */

type luaMetathingsCore struct {
	dat_stor dssdk.DataStorage

	context esdk.Data
	data    esdk.Data
}

func newLuaMetathingsCore(args ...interface{}) (*luaMetathingsCore, error) {
	var ctx, dat esdk.Data
	var ds dssdk.DataStorage

	if err := opt_helper.Setopt(map[string]func(string, interface {
	}) error{
		"context":      esdk.ToData(&ctx),
		"data":         esdk.ToData(&dat),
		"data_storage": dssdk.ToDataStorage(&ds),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCore{
		context:  ctx,
		data:     dat,
		dat_stor: ds,
	}, nil
}

func (c *luaMetathingsCore) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"data":    c.getData,
		"context": c.getContext,
		"storage": c.newStorage,
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
func (c *luaMetathingsCore) getData(L *lua.LState) int {
	core := c.check(L)
	key := L.CheckString(2)

	ival := core.data.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:context(key#string)
//   context body lookup github.com/nayotta/metathings/pkg/plugin/evaluator/evaluator_impl.go#L31
func (c *luaMetathingsCore) getContext(L *lua.LState) int {
	core := c.check(L)
	key := L.CheckString(2)

	ival := core.context.Get(key)
	val := parse_interface_to_lvalue(L, ival)
	L.Push(val)

	return 1
}

// LUA_FUNCTION: core:storage(msr#string, tags#table) storage
func (c *luaMetathingsCore) newStorage(L *lua.LState) int {
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

type luaMetathingsCoreStorage struct {
	dat_stor dssdk.DataStorage

	msr  string
	tags map[string]string
}

func newLuaMetathingsCoreStorage(args ...interface{}) (*luaMetathingsCoreStorage, error) {
	var ds dssdk.DataStorage
	var msr string
	var tags map[string]string

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"data_storage": dssdk.ToDataStorage(&ds),
		"measurement":  opt_helper.ToString(&msr),
		"tags":         opt_helper.ToStringMapString(&tags),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreStorage{
		dat_stor: ds,
		msr:      msr,
		tags:     tags,
	}, nil
}

func (s *luaMetathingsCoreStorage) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"with":  s.With,
		"write": s.Write,
	}
}

func (s *luaMetathingsCoreStorage) check(L *lua.LState) *luaMetathingsCoreStorage {
	ud := L.CheckUserData(1)
	v, ok := ud.Value.(*luaMetathingsCoreStorage)
	if !ok {
		L.ArgError(1, "core_storage expected")
		return nil
	}

	return v
}

// LUA_FUNCTION: storage:with(tags#table) storage
func (s *luaMetathingsCoreStorage) With(L *lua.LState) int {
	var pipe bytes.Buffer
	var tags map[string]string

	stor := s.check(L)

	enc := gob.NewEncoder(&pipe)
	enc.Encode(stor.tags)
	dec := gob.NewDecoder(&pipe)
	dec.Decode(&tags)

	tags_tb := L.CheckTable(2)
	exts := parse_ltable_to_string_map(tags_tb)

	for k, v := range exts {
		tags[k] = v.(string)
	}

	ns, err := newLuaMetathingsCoreStorage("data_storage", s.dat_stor, "measurement", stor.msr, "tags", tags)
	if err != nil {
		L.RaiseError("failed to new storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, ns)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: storage:write(data#table, [option#table])
//   option:
//     timestamp: data timestamp
func (s *luaMetathingsCoreStorage) Write(L *lua.LState) int {
	ctx := context.TODO()

	dat_tb := L.CheckTable(2)
	dat := parse_ltable_to_string_map(dat_tb)

	var opt_tb *lua.LTable
	if L.GetTop() > 2 {
		opt_tb = L.CheckTable(3)
	}
	opt := parse_ltable_to_string_map(opt_tb)

	var ts time.Time
	if tsi, ok := opt["timestamp"]; ok {
		ts = time.Unix(cast.ToInt64(tsi), 0)
	} else {
		ts = time.Now()
	}
	ctx = context.WithValue(ctx, "timestamp", ts)

	err := s.dat_stor.Write(ctx, s.msr, s.tags, dat)
	if err != nil {
		L.RaiseError("failed to write data to data storage")
		return 0
	}

	return 0
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
