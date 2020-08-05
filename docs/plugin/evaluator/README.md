# Evaluator Plugin 使用指南

## 1. 概要

`Evaluator` 插件, 是提供一个运行时环境, 对数据通道上传的数据进行在线处理的平台.

## 2. 使用步骤

TBD.

## 3. 接口

### 3.1. metathings

#### 3.1.1. metathings:data

#### 3.1.1.1. 描述

获取触发本次`eval`事件的数据.

`metathings:data(key#string)`

#### 3.1.1.2. 范例

```lua
...
local val = metathings:data("key")
...
```

#### 3.1.2. metathings:context

#### 3.1.2.1. 描述

获取 `evaluator`的 `context`内容.

[context结构](https://github.com/nayotta/metathings/blob/master/pkg/plugin/evaluator/lua_operator_core.go#L41)描述.

`metathings:context(key#string)`

#### 3.1.2.2.  范例

```lua
...
local val = metathings:context("key")
...
```

#### 3.1.3. metathings:device

#### 3.1.3.1. 描述

获取 `evaluator`的 `device`对象.

`metathings:device(alias#string)`

当 `alias` 指定为`"self"`时, 如果此次的请求是由 `device` 发起的话, 就可以获取当前请求的 `device`.

#### 3.1.3.2. 范例

```lua
...
local dev = metathings:device("self")
...
```

#### 3.1.4. metathings:storage

#### 3.1.4.1. 描述

获取 `metathings`的 `storage`对象.

`metathings:storage(msr#string, tags#table<optional>) storage`

`storage` 对象的具体使用方法可以参考 `3.3`章节.

#### 3.1.4.2. 范例

```lua
...
local stor = metathings:storage("temperature", {["device"] = "<device-id>"})
...
```

#### 3.1.5. metathings:simple_storage

#### 3.1.5.1. 描述

获取 `metathings`的 `simple_storage`对象.

`metathings:simple_storage(option#table) simple_storage`

`simple_storage` 对象的具体使用方法可以参考 `3.4`章节.

#### 3.1.5.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
...
```

#### 3.1.6. metathings:flow

#### 3.1.6.1. 描述

获取 `metathings`的 `flow`对象.

`metathings:flow(alias#string, name#string) flow`

`flow` 对象的具体使用方法可以参考 `3.5`章节.

#### 3.1.6.2. 范例

```lua
...
local flw = metathings:flow("self", "greeting")
...
```

### 3.2. device

#### 3.2.1. device:id

#### 3.2.1.1. 描述

获取 `device`的`id`.

`device:id()`

#### 3.2.1.2. 范例

```lua
...
local dev = metathings:device("self")
local dev_id = dev:id()
...
```

### 3.2.2. device:storage

#### 3.2.2.1. 描述

获取 `device`的`storage`对象.

`device:storage(measurement#string, tags#table<optional>) storage`

`storage` 对象的具体使用方法可以参考 `3.3`章节.

#### 3.2.2.2. 范例

```lua
...
local dev = metathings:device("self")
local stor = dev:storage("sensors", {
  ["device"] = metathings:context("device.id"),
})
...
```

### 3.2.3. device:simple_storage

#### 3.2.3.1. 描述

获取 `device`的`simple_storage`对象.

`device:simple_storage(option#table) simple_storage`

`simple_storage` 对象的具体使用方法可以参考 `3.4`章节.

#### 3.2.3.2. 范例

```lua
...
local dev = metathings:device("self")
local smpl_stor = dev:simple_storage()
...
```

### 3.2.4. device:flow

#### 3.2.4.1. 描述

获取 `device`的`flow`对象.

`device:flow(name#string) flow`

`flow`  对象的具体使用方法可以参考 `3.4`章节.

#### 3.2.4.2. 范例

```lua
...
local dev = metathings:device("self")
local flw = dev:flow("greeting")
...
```

### 3.2.5. device:unary_call

#### 3.2.5.1. 描述

调用 `device`的`unary_call` 方法.

`device:unary_call(module#string, method#string, argument#table) table`

#### 3.2.5.2. 范例

```lua
...
local dev = metathings:device("self")
dev:unary_call("echo", "Echo", {
  ["text"] = "hello, world!",
})
...
```

#### 3.2.5.3. 备注

`unary_call`接口需要预先设置相关配置.

 配置流程:
 
1. 通过 `UploadDescriptor` 接口上传 [`protoset` 文件](https://developers.google.com/protocol-buffers/docs/techniques), 并且获得该 `protoset`的 `SHA1`值.

2. 通过 `ListConfigsByDevice` 获取到 `device` 的 `Configs`, 如果不存在 `_sys_descriptor`, 则通过 `CreateConfig` 创建 `_sys_descriptor`配置, 并且通过 `AddConfigsToDevice`接口把`config`与`device`关联(允许多个设备关联一个配置文件), 格式如下:

```
{
  "modules": {
    "<module1-name>": {
      "sha1": "<module1-SHA1>",
    },
    "<module2-name>": {
      "sha1": "<module2-SHA1>",
    },
    ...
  },
}
```

详细的配置过程可以[参考](https://github.com/nayotta/metathings/examples/14-setup-device-configs/main.go).

### 3.3. storage

#### 3.3.1. storage:with

#### 3.3.1.1. 描述

获取包含了 `tags`的 `storage`对象.

`storage:with(tags#table) storage`

#### 3.3.1.2. 范例

```lua
...
local stor = metathings:storage(...)
local new_stor = stor:with({["location"] = "guangzhou"})
...
```

#### 3.3.2. storage:write

#### 3.3.2.1. 描述

写入数据到`storage`.

`storage:write(data#table, option#table<optional>)`

#### 3.3.2.2. 范例

```lua
...
local stor = metathings:storage(...)
stor:write({
  ["temperature"] = 37.2
})
...
```

### 3.4. simple_storage

#### 3.4.1. simple_storage:put

#### 3.4.1.1. 描述

写入数据到 `simple_storage`.

`simple_storage:put(option#table, content#string)`

option:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`

#### 3.4.1.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
smpl_stor:put({
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.txt"
}, "hello, world!")
...
```

#### 3.4.2. simple_storage:remove

#### 3.4.2.1. 描述

从`simple_storage`删除数据文件.

`simple_storage:remove(option#table)`

option:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`

#### 3.4.2.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
smpl_stor:remove({
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.txt"
})
...
```

#### 3.4.3. simple_storage:rename

#### 3.4.3.1. 描述

重命名`simple_storage`的数据文件.

`simple_storage:rename(src#table, dst#table)`

src:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`

dst:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`

#### 3.4.3.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
smpl_stor:rename({
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.txt"
}, {
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.log"
})
...
```

#### 3.4.4. simple_storage:get

#### 3.4.4.1. 描述

获取`simple_storage`的数据文件元数据.

`simple_storage:get(option#table) table`

option:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`

return:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`
* `name`, 上传数据的文件名, type: `string`
* `length`, 文件长度, type: `number`
* `etag`, 文件hash值, type: `string`
* `last_modified`, 文件最后修改时间, type: `number`

#### 3.4.4.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
local md = smpl_stor:get({
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.txt"
})
...
```

#### 3.4.5. simple_storage:get_content

#### 3.4.5.1. 描述

获取`simple_storage`的数据文件内容.

`simple_storage:get_content(option#table) string`

option:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`


#### 3.4.5.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
local content = smpl_stor:get_content({
  ["device"] = "<device-id>",
  ["prefix"] = "/",
  ["name"] = "hello.txt"
})
...
```

#### 3.4.6. simple_storage:list

#### 3.4.6.1. 描述

列出`simple_storage`包含的数据文件的元数据.

`simple_storage:list(option#table) table`

option:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`, optional: `true`
* `name`, 上传数据的文件名, type: `string`
* `recursive`, 是否递归查询下级目录, type: `bool`, default: `false`
* `depth`, 查询下级目录深度, type: `number`, default: 16

return:

* `table[object]`, 元数据的表格数据

object:

* `device`, 设备的ID, type: `string`
* `prefix`, 上传数据的路径, type: `string`
* `name`, 上传数据的文件名, type: `string`
* `length`, 文件长度, type: `number`
* `etag`, 文件hash值, type: `string`
* `last_modified`, 文件最后修改时间, type: `number`

#### 3.4.6.2. 范例

```lua
...
local smpl_stor = metathings:simple_storage()
local objs = smpl_stor:list({
  ["device"] = "<device-id>",
  ["name"] = "/",
  ["recursive"] = true,
  ["depth"] = 1
})
...
```

### 3.4. flow

#### 3.4.1. flow:push_frame

#### 3.4.1.1. 描述

推送数据到流.

`flow:push_frame(data#table)`

#### 3.4.1.2. 范例

```lua
...
local flw = metathings:flow("self", "greeting")
flw:push_frame({
  ["text"] = "hello, world!",
})
...
```

## 4. Timer(定时器)

### 4.1. 描述

定时器允许定时执行指定的`Evaluator`.

### 4.2. 使用步骤

1. 创建需要被执行的`Evaluator`.
2. 创建对应的定时器`Timer`.
3. 将`Timer`添加到`Evaluator`.

### 4.3.  详细步骤

#### 4.3.1. 创建Evaluator

使用`ai.metathings.service.evaluatord.EvaluatordService/CreateEvaluator`接口创建`Evaluator`.

范例:

```bash
// set token to MT_TOKEN
// set address to MT_ADDR
$ echo ${MT_TOKEN}
"<token>"
$ cat create_evaluator.json
{
  "evaluator": {
    "alias":  "echo",
    "timezone": "Asia/Shanghai",
    "enabled": false
  }
}
$ cat create_timer.json | grpcurl -protoset pkg/proto/evaluatord/service.protoset -H "authorization: Bearer ${MT_TOKEN}" -d @ "${MT_ADDR}" ai.metathings.service.evaluatord.EvaluatordService/CreateTimer
{
  "timer": {
    "id": "c6e1e19eb5274040a7d5b2bfed7ee613",
    "schedule": "@every 10s",
    "timezone": "Asia/Shanghai",
    "configs": [],
    "alias": "c6e1e19eb5274040a7d5b2bfed7ee613"
  }
}
$ cat create_timer_config.json
{
  "config": {
    "alias": "default",
    "body": {
      "version": "v1",
      "device": {
        "id": "c2a652e6de95483d968b30eeb82f384a"
      },
      "data": {
        "text": "hello, world!"
      }
    }
  }
}
$ cat create_timer_config.json | grpcurl -protoset pkg/proto/deviced/service.protoset -H "authorization: Bearer ${MT_TOKEN}" -d @ "${MT_ADDR}" ai.metathings.service.deviced.DevicedService/CreateConfig
{
  "config": {
    "id": "44e5317df2664bd7912a688e8c217393",
    "alias": "default",
    "body": {
      "data": {
        "text": "hello, world!"
      },
      "device": {
        "id": "c2a652e6de95483d968b30eeb82f384a"
      },
      "version": "v1"
    }
  }
}
$ cat add_configs_to_timer.json
{
  "timer": {
    "id": "c6e1e19eb5274040a7d5b2bfed7ee613"
  },
  "configs": [
    {
      "id": "44e5317df2664bd7912a688e8c217393"
    }
  ]
}
$ cat add_configs_to_timer.json | grpcurl -protoset pkg/proto/evaluatord/service.protoset -H "authorization: Bearer ${MT_TOKEN}" -d @ "${MT_ADDR}" ai.metathings.service.evaluatord.EvaluatordService/AddConfigsToTimer
{}
$ cat add_timers_to_evaluator.json
{
  "evaluator": {
    "id": "d84735aa83e1488294f6483578f92055"
  }, "sources": [
    {
      "id": "c6e1e19eb5274040a7d5b2bfed7ee613",
      "type": "timer"
    }
  ]
}
```

#### 4.3.3. 添加Timer到Evaluator

1. 使用`ai.metathings.service.evaluatord.EvaluatordService/AddSourcesToEvaluator`接口添加`Timer`到`Evlauator`.
2. 使用`ai.metathings.service.evaluatord.EvaluatordService/PatchTimer`接口启用`Timer`.

```bash
$ cat add_timers_to_evaluator.json | grpcurl -protoset pkg/proto/evaluatord/service.protoset -H "authorization: Bearer ${MT_TOKEN}" -d @ "${MT_ADDR}" ai.metathings.service.evaluatord.EvaluatordService/AddSourcesToEvaluator
{}
$ cat enable_timer.json
{
  "timer": {
    "id": "c6e1e19eb5274040a7d5b2bfed7ee613",
    "enabled": true
  }
}
$ cat enable_timer.json | grpcurl -protoset pkg/proto/evaluatord/service.protoset -H "authorization: Bearer ${MT_TOKEN}" -d @ "${MT_ADDR}" ai.metathings.service.evaluatord.EvaluatordService/PatchTimer
{
  "timer": {
    "id": "c6e1e19eb5274040a7d5b2bfed7ee613",
    "schedule": "@every 10s",
    "timezone": "Asia/Shanghai",
    "enabled": true,
    "configs": [
      {
        "id": "44e5317df2664bd7912a688e8c217393"
      }
    ],
    "alias": "c6e1e19eb5274040a7d5b2bfed7ee613"
  }
}
```
