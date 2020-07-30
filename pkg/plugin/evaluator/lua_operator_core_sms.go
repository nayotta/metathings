package metathings_plugin_evaluator

import (
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	smssdk "github.com/nayotta/metathings/sdk/sms"
)

type luaMetathingsCoreSmsOption struct {
}

func newLuaMetathingsCoreSmsOption() *luaMetathingsCoreSmsOption {
	return &luaMetathingsCoreSmsOption{}
}

type luaMetathingsCoreSms struct {
	opt        *luaMetathingsCoreSmsOption
	core       *luaMetathingsCore
	sms_sender smssdk.SmsSender
}

func (p *luaMetathingsCoreSms) check(L *lua.LState) *luaMetathingsCoreSms {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreSms)
	if !ok {
		L.ArgError(1, "sms expected")
		return nil
	}

	return v
}

// LUA_FUNCTION: sms:send(id_or_alias#string, arguments#table<optional>)
func (p *luaMetathingsCoreSms) luaSend(L *lua.LState) int {
	var sms_id string
	var sms_cfg map[string]interface{}
	var args map[string]string
	var err error

	p.check(L)

	// id or alias
	ioa := L.CheckString(2)

	cctx := p.core.GetContext()

	if sms_id = cast.ToString(cctx.Get("config.alias.sms_sender." + ioa)); sms_id == "" {
		sms_id = ioa
	}

	sms_cfg = cast.ToStringMap(cctx.Get("config.sms_sender." + sms_id))
	if sms_cfg == nil {
		L.ArgError(1, "ip or alias expected")
		return 0
	}

	cfgx := objx.New(sms_cfg)
	nums := cast.ToStringSlice(cfgx.Get("numbers").Data())
	if len(nums) == 0 {
		L.RaiseError("numbers is empty")
		return 0
	}

	args = cast.ToStringMapString(cfgx.Get("arguments").Data())

	if L.GetTop() > 2 {
		args_tb := L.CheckTable(3)
		for k, v := range cast.ToStringMapString(parse_ltable_to_string_map(args_tb)) {
			args[k] = v
		}
	}

	if err = p.sms_sender.SendSms(p.core.GetGoContext(), sms_id, nums, args); err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	return 0
}

func (p *luaMetathingsCoreSms) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"send": p.luaSend,
	}
}

func newLuaMetathingsCoreSms(args ...interface{}) (*luaMetathingsCoreSms, error) {
	var core *luaMetathingsCore
	var ss smssdk.SmsSender
	opt := newLuaMetathingsCoreSmsOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"core":       toLuaMetathingsCore(&core),
		"sms_sender": smssdk.ToSmsSender(&ss),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreSms{
		opt:        opt,
		core:       core,
		sms_sender: ss,
	}, nil
}
