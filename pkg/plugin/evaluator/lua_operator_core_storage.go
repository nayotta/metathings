package metathings_plugin_evaluator

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	"github.com/spf13/cast"
	lua "github.com/yuin/gopher-lua"
)

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
		"with":  s.luaWith,
		"write": s.luaWrite,
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
func (s *luaMetathingsCoreStorage) luaWith(L *lua.LState) int {
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
func (s *luaMetathingsCoreStorage) luaWrite(L *lua.LState) int {
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
