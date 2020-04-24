# Evaluator Plugin  使用指南

## 1. 概要

`Evaluator` 插件, 是提供一个运行时环境, 对数据通道上传的数据进行一个在线处理平台.

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

[context结构](https://github.com/nayotta/metathings/pkg/plugin/evaluator/lua_operator_core.go#L40)描述.

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

`metathings:device(name_or_alias#string)`

`"self"`是一个别名, 当此次的请求是由 `device` 发起的话, 就可以获取当前请求的 `device`.

#### 3.1.3.2. 范例

```lua
...
local dev = metathings:device("self")
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

#### 3.2.2.2. 范例

```lua
...
local dev = metathings:device("self")
local stor = dev:storage("temperature", {
  ["device"] = metathings:context("device.id"),
})
...
```

### 3.2.3. device:simple_storage

#### 3.2.3.1. 描述

获取 `device`的`simple_storage`对象.

`device:simple_storage(option#table) simple_storage`

#### 3.2.3.2. 范例

```lua
...
local dev = metathings:device("self")
local smpl_stor = dev:simple_storage()
...
```

### 3.2.4. device:unary_call

#### 3.2.4.1. 描述

调用 `device`的`unary_call` 方法.

`device:unary_call(module#string, method#string, argument#table) table`

#### 3.2.4.2 范例

```lua
...
local dev = metathings:device("self")
dev:unary_call("echo", "Echo", {
  ["text"] = "hello, world!",
})
...
```

#### 3.2.4.3 备注

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
