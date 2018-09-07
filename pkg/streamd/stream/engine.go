package stream_manager

import (
	"context"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type LuaEngine struct {
	ls *lua.LState
}

func (self *LuaEngine) Close() {
	self.ls.Close()
}

func (self *LuaEngine) SetContext(ctx context.Context) {
	self.ls.SetContext(ctx)
}

func (self *LuaEngine) Filter(filter string, metadata StreamData, data StreamData) (bool, error) {
	self.load_metadata(metadata)
	self.load_data(data)

	lua_str := fmt.Sprintf(`ret = (function() return %v end)()`, filter)
	err := self.ls.DoString(lua_str)
	if err != nil {
		return false, err
	}

	return self.ls.GetGlobal("ret") == lua.LTrue, nil
}

func (self *LuaEngine) load_metadata(metadata StreamData) {
	mt_tbl := self.ls.NewTable()
	for _, k := range metadata.Keys() {
		mt_tbl.RawSetString(k, lua.LString(metadata.AsString(k)))
	}
	self.ls.SetGlobal("metadata", mt_tbl)
}

func (self *LuaEngine) load_data(data StreamData) {
	for _, k := range data.Keys() {
		self.ls.SetGlobal(k, lua.LString(data.AsString(k)))
	}
}

func NewLuaEngine() *LuaEngine {
	return &LuaEngine{
		ls: lua.NewState(),
	}
}
