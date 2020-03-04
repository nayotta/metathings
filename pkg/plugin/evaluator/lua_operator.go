package metathings_plugin_evaluator

import (
	"strconv"
	"sync"

	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type LuaOperatorOption struct {
	Code string
}

type LuaOperator struct {
	opt *LuaOperatorOption
	env *lua.LState
}

func (lo *LuaOperator) Run(cfg, dat esdk.Data) (esdk.Data, error) {
	mt_tb := lo.env.NewTable()

	cfg_tb := lo.parse_string_map_to_ltable(cfg.Iter())
	mt_tb.RawSetString("config", cfg_tb)

	dat_tb := lo.parse_string_map_to_ltable(dat.Iter())
	mt_tb.RawSetString("data", dat_tb)

	lo.env.SetGlobal("metathings", mt_tb)

	err := lo.env.DoString(lo.opt.Code)
	if err != nil {
		return nil, err
	}

	ret_lv := lo.env.Get(-1)
	if ret_lv.Type() != lua.LTTable {
		return nil, ErrUnexpectedOperatorResult
	}

	dat_m := lo.parse_ltable_to_string_map(ret_lv.(*lua.LTable))

	ret, err := esdk.DataFromMap(dat_m)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (lo *LuaOperator) Close() error {
	lo.env.Close()

	return nil
}

func (lo *LuaOperator) parse_string_map_to_ltable(xs map[string]interface{}) *lua.LTable {
	tb := lo.env.NewTable()

	for xk, xv := range xs {
		tb.RawSetString(xk, lo.parse_interface_to_lvalue(xv))
	}

	return tb
}

func (lo *LuaOperator) parse_interface_to_lvalue(x interface{}) lua.LValue {
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
		return lo.parse_string_map_to_ltable(x.(map[string]interface{}))
	case []interface{}:
		return lo.parse_interface_slice_to_ltable(x.([]interface{}))
	default:
		panic("unimplemented")
	}
}

func (lo *LuaOperator) parse_interface_slice_to_ltable(xs []interface{}) *lua.LTable {
	tb := lo.env.NewTable()

	for _, x := range xs {
		tb.Append(lo.parse_interface_to_lvalue(x))
	}

	return tb
}

func (lo *LuaOperator) parse_ltable_to_string_map(x *lua.LTable) (y map[string]interface{}) {
	y, _ = lo.parse_ltable_to_interface(x).(map[string]interface{})
	return
}

func (lo *LuaOperator) parse_ltable_to_interface(x *lua.LTable) interface{} {
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
			ys = append(ys, lo.parse_lvalue_to_interface(val))
		} else {
			copy_array_to_map_once.Do(func() {
				for i, y := range ys {
					mys[strconv.Itoa(i)] = y
				}
			})
			mys[lua.LVAsString(key)] = lo.parse_lvalue_to_interface(val)
		}
	})

	if is_array {
		return ys
	} else {
		return mys
	}
}

func (lo *LuaOperator) parse_lvalue_to_interface(x lua.LValue) interface{} {
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
		return lo.parse_ltable_to_interface(x.(*lua.LTable))
	default:
		panic("unimplemented")
	}
}

func NewLuaOperator(args ...interface{}) (Operator, error) {
	opt := &LuaOperatorOption{}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"code": opt_helper.ToString(&opt.Code),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	op := &LuaOperator{
		opt: opt,
		env: lua.NewState(),
	}

	return op, nil
}

func init() {
	registry_operator_factory("lua", NewLuaOperator)
	registry_operator_factory("default", NewLuaOperator)
}
