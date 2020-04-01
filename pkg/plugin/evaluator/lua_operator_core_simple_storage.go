package metathings_plugin_evaluator

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	lua "github.com/yuin/gopher-lua"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
)

type luaMetathingsCoreSimpleStorageOption struct {
	ImmutableOption map[string]interface{}
}

type luaMetathingsCoreSimpleStorage struct {
	opt        *luaMetathingsCoreSimpleStorageOption
	immutables objx.Map
	core       *luaMetathingsCore
}

func (ss *luaMetathingsCoreSimpleStorage) check(L *lua.LState) *luaMetathingsCoreSimpleStorage {
	ud := L.CheckUserData(1)

	v, ok := ud.Value.(*luaMetathingsCoreSimpleStorage)
	if !ok {
		L.ArgError(1, "simple_storage expected")
		return nil
	}

	return v
}

func (ss *luaMetathingsCoreSimpleStorage) MetatableIndex() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"put":         ss.luaPut,
		"get":         ss.luaGet,
		"remove":      ss.luaRemove,
		"rename":      ss.luaRename,
		"get_content": ss.luaGetContent,
		"list":        ss.luaList,
	}
}

func (ss *luaMetathingsCoreSimpleStorage) GetImmutableOption() objx.Map {
	if ss.immutables == nil {
		ss.immutables = objx.New(ss.opt.ImmutableOption)
	}
	return ss.immutables
}

func (ss *luaMetathingsCoreSimpleStorage) get_context() context.Context {
	return context_helper.WithToken(context.TODO(), cast.ToString(ss.core.GetContext().Get("token")))
}

// LUA_FUNCTION: simple_storage:put(option#table, content#string)
//   option:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
func (ss *luaMetathingsCoreSimpleStorage) luaPut(L *lua.LState) int {
	ss.check(L)

	opt_tb := L.CheckTable(2)
	cnt := L.CheckString(3)
	obj := parse_ltable_to_pb_object(ss, opt_tb)

	if err := ss.core.GetSimpleStorage().Put(ss.get_context(), obj, strings.NewReader(cnt)); err != nil {
		L.RaiseError("failed to put object to simple storage")
		return 0
	}

	return 0
}

// LUA_FUNCTION: simple_storage:remove(option#table)
//   option:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
func (ss *luaMetathingsCoreSimpleStorage) luaRemove(L *lua.LState) int {
	ss.check(L)

	opt_tb := L.CheckTable(2)
	obj := parse_ltable_to_pb_object(ss, opt_tb)

	if err := ss.core.GetSimpleStorage().Remove(ss.get_context(), obj); err != nil {
		L.RaiseError("failed to remove object from simple storage")
		return 0
	}

	return 0
}

// LUA_FUNCTION: simple_storage:rename(src#table, dst#table)
//   src:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...   # type: string
//   dst:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
func (ss *luaMetathingsCoreSimpleStorage) luaRename(L *lua.LState) int {
	ss.check(L)

	src_tb := L.CheckTable(2)
	dst_tb := L.CheckTable(3)
	src := parse_ltable_to_pb_object(ss, src_tb)
	dst := parse_ltable_to_pb_object(ss, dst_tb)

	if err := ss.core.GetSimpleStorage().Rename(ss.get_context(), src, dst); err != nil {
		L.RaiseError("failed to rename object in simple storage")
		return 0
	}

	return 0
}

// LUA_FUNCTION: simple_storage:get(option#table) table
//   option:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
//   return:
//     table:
//       device: ...  # type: string
//       prefix: ...  # type: string
//       name: ...  # type: string
//       length: ...  # type: number
//       etag: ...  # type: string
//       last_modified: ...  # type: number
func (ss *luaMetathingsCoreSimpleStorage) luaGet(L *lua.LState) int {
	ss.check(L)

	opt_tb := L.CheckTable(2)
	obj := parse_ltable_to_pb_object(ss, opt_tb)

	obj, err := ss.core.GetSimpleStorage().Get(ss.get_context(), obj)
	if err != nil {
		L.RaiseError("failed to get object from simple storage")
		return 0
	}

	ret_tb := parse_pb_message_to_ltable(L, obj)
	L.Push(ret_tb)

	return 1
}

// LUA_FUNCTION: simple_storage:get_content(option#table) string
//   option:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
func (ss *luaMetathingsCoreSimpleStorage) luaGetContent(L *lua.LState) int {
	ss.check(L)

	opt_tb := L.CheckTable(2)
	obj := parse_ltable_to_pb_object(ss, opt_tb)

	buf, err := ss.core.GetSimpleStorage().GetContent(ss.get_context(), obj)
	if err != nil {
		L.RaiseError("failed to get object content from simple storage")
		return 0
	}

	L.Push(lua.LString(string(buf)))
	return 1
}

// LUA_FUNCTION: simple_storage:list(option#table) table
//   option:
//     device: ...  # type: string
//     prefix: ...  # type: string
//     name: ...  # type: string
//     recursive: ...  # type: bool
//     depth: ...  # type: number
//   return:
//     table[ object ]
//       object:
//         device: ...  # type: string
//         prefix: ...  # type: string
//         name: ...   # type: string
//         length: ...  # type: number
//         etag: ...  # type: string
//         last_modified: ...  # type: number
func (ss *luaMetathingsCoreSimpleStorage) luaList(L *lua.LState) int {
	ss.check(L)

	opt_tb := L.CheckTable(2)
	obj := parse_ltable_to_pb_object(ss, opt_tb)
	opt := objx.New(parse_ltable_to_string_map(opt_tb))
	opts := []dsdk.SimpleStorageListOption{
		dsdk.SimpleStorageListOption_SetRecursive(opt.Get("recursive").Bool()),
		dsdk.SimpleStorageListOption_SetDepth(int(opt.Get("depth").Float64())),
	}

	objs, err := ss.core.GetSimpleStorage().List(ss.get_context(), obj, opts...)
	if err != nil {
		L.RaiseError("failed to list objects from simple storage")
		return 0
	}

	msgs := []proto.Message{}
	for _, obj := range objs {
		msgs = append(msgs, obj)
	}

	ret_tb := parse_pb_messages_to_ltable(L, msgs)
	L.Push(ret_tb)

	return 1
}

func newLuaMetathingsCoreSimpleStorageOption() *luaMetathingsCoreSimpleStorageOption {
	return &luaMetathingsCoreSimpleStorageOption{
		ImmutableOption: make(map[string]interface{}),
	}
}

func newLuaMetathingsCoreSimpleStorage(args ...interface{}) (*luaMetathingsCoreSimpleStorage, error) {
	var core *luaMetathingsCore
	opt := newLuaMetathingsCoreSimpleStorageOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"immutable_option": opt_helper.ToStringMap(&opt.ImmutableOption),
		"core":             toLuaMetathingsCore(&core),
	})(args...); err != nil {
		return nil, err
	}

	return &luaMetathingsCoreSimpleStorage{
		opt:  opt,
		core: core,
	}, nil
}
