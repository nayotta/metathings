package metathings_plugin_evaluator

import (
	"context"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type LuaOperatorOption struct {
	Code string
}

type LuaOperator struct {
	dat_stor  dssdk.DataStorage
	smpl_stor dsdk.SimpleStorage
	caller    dsdk.Caller
	opt       *LuaOperatorOption
	env       *lua.LState
}

func (lo *LuaOperator) Run(gctx context.Context, ctx, dat esdk.Data) (esdk.Data, error) {
	core, err := newLuaMetathingsCore(
		"gocontext", gctx,
		"context", ctx,
		"data", dat,
		"data_storage", lo.dat_stor,
		"simple_storage", lo.smpl_stor,
		"caller", lo.caller,
	)
	if err != nil {
		return nil, err
	}
	luaInitOnceObject(lo.env, "metathings", core)

	err = lo.env.DoString(lo.opt.Code)
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
	var logger log.FieldLogger
	var ds dssdk.DataStorage
	var ss dsdk.SimpleStorage
	var caller dsdk.Caller

	opt := &LuaOperatorOption{}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"code":           opt_helper.ToString(&opt.Code),
		"logger":         opt_helper.ToLogger(&logger),
		"data_storage":   dssdk.ToDataStorage(&ds),
		"simple_storage": dsdk.ToSimpleStorage(&ss),
		"caller":         dsdk.ToCaller(&caller),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	op := &LuaOperator{
		dat_stor:  ds,
		smpl_stor: ss,
		caller:    caller,
		opt:       opt,
		env:       lua.NewState(),
	}

	return op, nil
}

func init() {
	registry_operator_factory("lua", NewLuaOperator)
	registry_operator_factory("default", NewLuaOperator)
}
