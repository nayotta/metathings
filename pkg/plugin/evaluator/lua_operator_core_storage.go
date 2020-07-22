package metathings_plugin_evaluator

import (
	"context"
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	"github.com/spf13/cast"
	lua "github.com/yuin/gopher-lua"
)

type luaMetathingsCoreStorage struct {
	dat_stor dssdk.DataStorage

	msr            string
	immutable_tags map[string]string
	tags           map[string]string
}

func newLuaMetathingsCoreStorage(args ...interface{}) (*luaMetathingsCoreStorage, error) {
	var ds dssdk.DataStorage
	var msr string
	var immutable_tags map[string]string
	var tags map[string]string

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"data_storage":   dssdk.ToDataStorage(&ds),
		"measurement":    opt_helper.ToString(&msr),
		"immutable_tags": opt_helper.ToStringMapString(&immutable_tags),
		"tags":           opt_helper.ToStringMapString(&tags),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreStorage{
		dat_stor:       ds,
		msr:            msr,
		immutable_tags: immutable_tags,
		tags:           tags,
	}, nil
}

func (s *luaMetathingsCoreStorage) check(L *lua.LState) *luaMetathingsCoreStorage {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreStorage)
	if !ok {
		L.ArgError(1, "storage expected")
		return nil
	}

	return v
}

func (s *luaMetathingsCoreStorage) get_context() context.Context {
	return context.TODO()
}

func (s *luaMetathingsCoreStorage) get_tags() map[string]string {
	tags := map[string]string{}

	for k, v := range s.tags {
		tags[k] = v
	}

	for k, v := range s.immutable_tags {
		tags[k] = v
	}

	return tags
}

func (s *luaMetathingsCoreStorage) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"with":  s.luaWith,
		"write": s.luaWrite,
	}
}

// LUA_FUNCTION: storage:with(tags#table) storage
func (s *luaMetathingsCoreStorage) luaWith(L *lua.LState) int {
	s.check(L)

	tags := map[string]string{}
	stor := s.check(L)
	exts_tb := L.CheckTable(2)
	exts := parse_ltable_to_string_map(exts_tb)

	for k, v := range stor.tags {
		tags[k] = v
	}

	for k, v := range exts {
		tags[k] = v.(string)
	}

	ns, err := newLuaMetathingsCoreStorage("data_storage", stor.dat_stor, "measurement", stor.msr, "immutable_tags", stor.immutable_tags, "tags", tags)
	if err != nil {
		L.RaiseError("failed to new storage")
		return 0
	}

	_, ud := luaBindingObjectMethods(L, ns)
	L.Push(ud)

	return 1
}

// LUA_FUNCTION: storage:write(data#table, option#table<optional>)
//   option:
//     timestamp: data timestamp
func (s *luaMetathingsCoreStorage) luaWrite(L *lua.LState) int {
	s.check(L)

	ctx := s.get_context()

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

	err := s.dat_stor.Write(ctx, s.msr, s.get_tags(), dat)
	if err != nil {
		L.RaiseError("failed to write data to data storage")
		return 0
	}

	return 0
}
