package stream_manager

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

type LuaEngine struct {
	logger log.FieldLogger
	ls     *lua.LState
}

func (self *LuaEngine) Close() {
	self.ls.Close()
}

func (self *LuaEngine) SetContext(ctx context.Context) {
	self.ls.SetContext(ctx)
}

func (self *LuaEngine) Filter(filter string, metadata StreamData, data StreamData) (bool, error) {
	luanch_str := fmt.Sprintf(`
function luanch(metadata, data)
  return { ok = (%v) }
end
`, filter)

	err := self.luanch(luanch_str, metadata, data)
	if err != nil {
		return false, err
	}

	lv_result := self.ls.GetGlobal("result")
	if lv_result.Type() != lua.LTTable {
		return false, ErrUnexpectedResultType
	}

	result := lv_result.(*lua.LTable).RawGetString("ok") == lua.LTrue

	self.logger.WithField("result", result).Debugf("engine.Filter")
	return result, nil
}

func (self LuaEngine) Luanch(luanch_str string, metadata StreamData, data StreamData) (StreamData, error) {
	err := self.luanch(luanch_str, metadata, data)
	if err != nil {
		return nil, err
	}

	lv_result := self.ls.GetGlobal("result")
	if lv_result.Type() != lua.LTTable {
		return nil, ErrUnexpectedResultType
	}

	result := self.ltable_to_streamdata(lv_result)

	self.logger.WithField("result", result.Data()).Debugf("engine.Luanch")
	return result, nil
}

func (self LuaEngine) luanch(luanch_str string, metadata StreamData, data StreamData) error {
	self.load_metadata(metadata)
	self.load_data(data)

	lua_str := fmt.Sprintf(`
%v

result = luanch(metadata, data)
`, luanch_str)

	self.logger.WithFields(log.Fields{
		"script":   lua_str,
		"metadata": metadata.Data(),
		"data":     data.Data(),
	}).Debugf("luanch lua engine")

	err := self.ls.DoString(lua_str)
	if err != nil {
		return err
	}

	return nil

}

func (self *LuaEngine) streamdata_to_ltable(x StreamData) *lua.LTable {
	tbl := self.ls.NewTable()
	for _, k := range x.Keys() {
		tbl.RawSetString(k, lua.LString(x.AsString(k)))
	}
	return tbl
}

func (self *LuaEngine) load_metadata(metadata StreamData) {
	self.ls.SetGlobal("metadata", self.streamdata_to_ltable(metadata))
}

func (self *LuaEngine) load_data(data StreamData) {
	self.ls.SetGlobal("data", self.streamdata_to_ltable(data))
}

func (self *LuaEngine) ltable_to_streamdata(val lua.LValue) StreamData {
	sd := NewStreamData(nil)

	if val.Type() != lua.LTTable {
		return sd
	}

	tbl := val.(*lua.LTable)
	tbl.ForEach(func(k, v lua.LValue) {
		sd.Set(k.String(), v.String())
	})

	return sd
}

func NewLuaEngine(logger log.FieldLogger) *LuaEngine {
	return &LuaEngine{
		logger: logger.WithField("#module", "lua_engine"),
		ls:     lua.NewState(),
	}
}
