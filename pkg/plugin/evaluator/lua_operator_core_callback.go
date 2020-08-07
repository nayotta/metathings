package metathings_plugin_evaluator

import (
	"github.com/spf13/cast"
	lua "github.com/yuin/gopher-lua"

	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	cbsdk "github.com/nayotta/metathings/sdk/callback"
)

type luaMetathingsCoreCallbackOption struct {
	Tags map[string]string
}

type luaMetathingsCoreCallback struct {
	opt  *luaMetathingsCoreCallbackOption
	core *luaMetathingsCore
}

func (cb *luaMetathingsCoreCallback) check(L *lua.LState) *luaMetathingsCoreCallback {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreCallback)
	if !ok {
		L.ArgError(1, "callback expected")
		return nil
	}

	return v
}

// LUA_FUNCTION: callback:emit(data#table)
func (cb *luaMetathingsCoreCallback) luaEmit(L *lua.LState) int {
	cb.check(L)

	cfg := cast.ToStringMap(cb.core.GetContext().Get("config.callback"))
	name, args, err := cfg_helper.ParseConfigOption("name", cfg)
	if err != nil {
		if err != cfg_helper.ErrExpectedKeyNotFound {
			L.RaiseError("bad callback configs")
			return 0
		}
		name = "dummy"
	}

	c, err := cbsdk.NewCallback(name, args...)
	if err != nil {
		L.RaiseError("failed to new callback client")
		return 0
	}

	dat_tb := L.CheckTable(2)
	dat := parse_ltable_to_string_map(dat_tb)

	if err := c.Emit(dat, cb.opt.Tags); err != nil {
		L.RaiseError("failed to emit callback data")
		return 0
	}

	return 0
}

func (cb *luaMetathingsCoreCallback) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"emit": cb.luaEmit,
	}
}

func newLuaMetathingsCoreCallback(args ...interface{}) (*luaMetathingsCoreCallback, error) {
	var opt luaMetathingsCoreCallbackOption
	var core *luaMetathingsCore

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"tags": opt_helper.ToStringMapString(&opt.Tags),
		"core": toLuaMetathingsCore(&core),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreCallback{
		opt:  &opt,
		core: core,
	}, nil
}
