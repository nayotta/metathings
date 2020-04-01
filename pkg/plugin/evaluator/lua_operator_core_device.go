package metathings_plugin_evaluator

import (
	"github.com/spf13/cast"
	lua "github.com/yuin/gopher-lua"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type luaMetathingsCoreDeviceOption struct {
	Id string
}

type luaMetathingsCoreDevice struct {
	opt  *luaMetathingsCoreDeviceOption
	core *luaMetathingsCore
}

func (d *luaMetathingsCoreDevice) check(L *lua.LState) *luaMetathingsCoreDevice {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreDevice)
	if !ok {
		L.ArgError(1, "device expected")
		return nil
	}

	return v
}

func (d *luaMetathingsCoreDevice) Id() string {
	return d.opt.Id
}

// LUA_FUNCTION: device:id() string
func (d *luaMetathingsCoreDevice) luaGetId(L *lua.LState) int {
	d.check(L)

	L.Push(lua.LString(d.opt.Id))

	return 1
}

// LUA_FUNCTION: device:storage(msr#string, tags#table) storage
func (d *luaMetathingsCoreDevice) luaNewStorage(L *lua.LState) int {
	d.check(L)

	msr := L.CheckString(2)
	tags_tb := L.CheckTable(3)
	tags := cast.ToStringMapString(parse_ltable_to_string_map(tags_tb))

	immutable_tags := map[string]string{
		"$source_id":   cast.ToString(d.core.GetContext().Get("source.id")),
		"$source_type": cast.ToString(d.core.GetContext().Get("source.type")),
		"$device_id":   d.opt.Id,
	}

	stor, err := newLuaMetathingsCoreStorage("data_storage", d.core.GetDataStorage(), "measurement", msr, "tags", tags, "immutable_tags", immutable_tags)
	if err != nil {
		L.RaiseError("failed to new device storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, stor)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: device:simple_storage(option#table) simple_storage
func (d *luaMetathingsCoreDevice) luaNewSimpleStorage(L *lua.LState) int {
	var opt map[string]interface{}

	d.check(L)

	if L.GetTop() > 1 {
		opt_tb := L.CheckTable(2)
		opt = parse_ltable_to_string_map(opt_tb)
	} else {
		opt = make(map[string]interface{})
	}
	opt["device"] = d.opt.Id

	s, err := newLuaMetathingsCoreSimpleStorage("immutable_option", opt, "core", d.core)
	if err != nil {
		L.RaiseError("failed to new device simple storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, s)
	L.Push(ud)

	return 1
}

func (d *luaMetathingsCoreDevice) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"id":             d.luaGetId,
		"storage":        d.luaNewStorage,
		"simple_storage": d.luaNewSimpleStorage,
	}
}

func newLuaMetathingsCoreDevice(args ...interface{}) (*luaMetathingsCoreDevice, error) {
	var opt luaMetathingsCoreDeviceOption
	var core *luaMetathingsCore

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"id":   opt_helper.ToString(&opt.Id),
		"core": toLuaMetathingsCore(&core),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreDevice{
		opt:  &opt,
		core: core,
	}, nil
}
