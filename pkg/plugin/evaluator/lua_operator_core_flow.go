package metathings_plugin_evaluator

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	lua "github.com/yuin/gopher-lua"
)

type luaMetathingsCoreFlowOption struct {
	Device string
	Flow   string
}

type luaMetathingsCoreFlow struct {
	opt  *luaMetathingsCoreFlowOption
	core *luaMetathingsCore
}

func (f *luaMetathingsCoreFlow) check(L *lua.LState) *luaMetathingsCoreFlow {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreFlow)
	if !ok {
		L.ArgError(1, "flow expected")
		return nil
	}

	return v
}

// LUA_FUNCTION: flow:push_frame(data#table)
func (f *luaMetathingsCoreFlow) luaPushFrame(L *lua.LState) int {
	f.check(L)

	dat_tb := L.CheckTable(2)
	dat := parse_ltable_to_string_map(dat_tb)

	err := f.core.GetFlow().PushFrame(f.core.GetGoContextWithToken(), f.opt.Device, f.opt.Flow, dat)
	if err != nil {
		L.RaiseError("failed to push frame to flow")
		return 0
	}

	return 0
}

func (f *luaMetathingsCoreFlow) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"push_frame": f.luaPushFrame,
	}
}

func newLuaMetathingsCoreFlow(args ...interface{}) (*luaMetathingsCoreFlow, error) {
	var opt luaMetathingsCoreFlowOption
	var core *luaMetathingsCore

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"device": opt_helper.ToString(&opt.Device),
		"flow":   opt_helper.ToString(&opt.Flow),
		"core":   toLuaMetathingsCore(&core),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreFlow{
		opt:  &opt,
		core: core,
	}, nil
}
