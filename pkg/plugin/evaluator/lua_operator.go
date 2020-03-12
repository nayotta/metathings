package metathings_plugin_evaluator

import (
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
	registerLuaMetathingsCore(lo.env)

	mt_ud := newLuaMetathingsCore(lo.env, cfg, dat)
	lo.env.SetGlobal("metathings", mt_ud)

	err := lo.env.DoString(lo.opt.Code)
	if err != nil {
		return nil, err
	}

	ret_lv := lo.env.Get(-1)
	if ret_lv.Type() != lua.LTTable {
		return nil, ErrUnexpectedOperatorResult
	}

	dat_m := parse_ltable_to_string_map(ret_lv.(*lua.LTable))

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
