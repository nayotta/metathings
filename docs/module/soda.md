# Soda Module 开发指南

## 1. 简介

`Module` 是 `Metathings` 平台下的一个概念, 可以理解为南向设备的业务实现程序.

由于 `Metathings` 平台是采用 `Golang` 编写的, 所以 `SDK` 暂时也只提供了 `Golang` 版本的.

所以导致部分对 `Golang` 语言不熟悉的开发者无法高效地完成开发工作.

于是乎, `Soda Module` 应运而生.

`Soda Module` 的目的是为了让开发者可以更加容易开发接入`Metathings` 平台的南向设备而设计的一套接口.

暂时只支持`HTTP`协议. (将来会支持更多的协议, 例如`MQTT` 等)

## 2. 文章结构

本文分为三个部分:

1. 相关概念

简单介绍重要概念, 方便后续理解.

2. 快速入门

通过几个简单的案例, 由浅入深地讲解相关的机制运作.

通过本章的讲解, 开发者可以根据自身的业务需求, 开发南向设备的程序.

3. `API`详解

深入讲解`Soda Module` 提供的 `API`, 可以完全释放 `Metathings` 平台提供的能力.

开发者可以通过 `Metathings` 平台提供的功能实现更加复杂的业务需求.

##  3. 相关概念

### 3.1. Metathings

TBD

### 3.2. Device

TBD

### 3.3. Module

TBD

### 3.4. Flow

TBD

### 3.5. SimpleStorage

TBD

### 3.6. Soda Module

TBD

## 4. 快速入门

### 4.1. Hello World

#### 4.1.1. 目的

实现一个南向设备, 能够接收 `Metathings` 平台的北向`Hello`请求, 并且作出回应.

#### 4.1.2.  流程预览

在`Metathings` 平台创建相关的设备, 用于接收北向命令和转发消息到南向设备.

使用 `Python` 构建一个简单的, 基于`HTTP Server` 的`Soda Module Plugin` 程序, 实现一个 `Hello` 服务, 服务包含 `hello` 方法.

启动该服务程序, 然后再启动`Device` 和`Soda Module` 程序.

然后通过测试程序测试请求是否正常响应.

#### 4.1.3. 详细流程

#### 4.1.3.1. 创建一个工作目录

```bash
$ mkdir 01hello
$ export WORKDIR=`pwd`/01hello
$ cd $WORKDIR
```

#### 4.1.3.2. 下载 Metathings 客户端

进入 `https://github.com/nayotta/metathings/releases/latest`

下载 `metathings_<version>_linux_amd64.tar.gz` 文件.

```bash
# 修改对应的 version
$ wget https://github.com/nayotta/metathings/releases/download/v<version>/metathings_<version>_linux_amd64.tar.gz
$ tar zxvf metathings_<version>_linux_amd64.tar.gz
```


#### 4.1.3.3. 实现 `Hello`服务

这里 采用 `Python` 的 `Flask` Web框架实现 `Hello` 服务.

暴露的接口为 `/hello`,  `HTTP`请求方法为 `POST`(`Soda Module Plugin`规定),

接收的参数为`json` 格式.

这里暴露的接口后续会与`Module` 的配置有关联.

```python
import flask
from flask import request

app = flask.Flask("hello")


@app.route("/hello", methods=["POST"])
def hello():
    name = request.json["name"]
    hello_str = "Hello, {0}".format(name)
    return flask.make_response({"message": hello_str})


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000, debug=True)
```

保存为 `hello.py`

#### 4.1.3.4.  运行 `Soda Module Plugin`

```bash
$ python3 hello.py &
```

#### 4.1.3.5. 配置 `Device`

```yaml
prod:  # 与MT_STAGE环境变量一致, 默认为 prod
  listen: 0.0.0.0:5002  # Device的与Module的通讯地址, 默认即可
  verbose: true  # 显示详细信息
  log:
    level: debug  # 日志等级, 可设置为 debug, info, warning, error
  credential:
    id: <credential-id>  # Device的Credential ID, 在 Metathings 平台生成
    secret: <credential-secret>  # Device的Credential Secret, 在 Metathings 平台生成
    domain: default  # Device的Credential Domain, 暂时只能填default
  transport_credential:
    plain_text: true  # Device与Module的通信传输协议设置, 默认即可
  service_endpoint:
    default:
      address: api.metathings.nayotta.com:443  # Metathings的地址, 默认即可
```

 保存为 `device.yaml`

#### 4.1.3.6. 运行 `Device`

```bash
$ ./metathings device run -c device.yaml &
```

#### 4.1.3.7. 配置 `Module`

```yaml
prod:  # 与MTC_STAGE环境变量一致, 默认为 prod
  name: hello  # Module 的名字, 在 Metathings平台创建时定义
  service:
    scheme: mtp+soda  # Soda Module的协议, 在 Metathings平台创建时定义
    host: 0.0.0.0  # Soda Module的监听地址, 在 Metathings平台创建时定义
    port: 13401  # Soda Module的监听端口, 在 Metathings平台创建时定义
  log:
    level: debug
  heartbeat:
    interval: 15  # Module 与 Device 的心跳间隔时间
    strategy: auto  #  心跳策略, 暂时采用 auto.
  credential:
    id: <credential-id>  # Module 的 Credential ID, 在 Metathings 平台生成
    secret: <credential-secret>  # Module 的 Credential Secret, 在 Metathings 平台生成
  service_endpoint:
    device:
      address: 127.0.0.1:5002  # Device 的通讯地址, 参考上面 Device 的配置
      plain_text: true  # Device 与 Module 是内部通讯, 所以采用 PlainText模式即可
    default:
      address:  api.metathings.nayotta.com:443  # Metathings的地址.
  backend:
    name: http  # Soda Module 的 Backend 驱动, 暂时只支持 http
    host: 0.0.0.0  # http 服务的监听地址
    port: 8001  # http 服务的监听端口
    auth:
      name: dummy  # 认证机制驱动, 暂时不启用, 即相信 Soda Module Plugin是安全可信的, 所以采用 dummy 驱动.
    target:
      url: http://127.0.0.1:8000  # 指向 Python 的 HTTP Server
    downstreams:
      my_hello:  # 暴露 Python HTTP Server 的服务, 名字叫 my_hello
        path: /hello  # 指向 Python HTTP Server 的路径
```

保存为 `module.yaml`

#### 4.1.3.8. 运行 `Module`

```bash
$ ./metathings module run -c module.yaml &
```

#### 4.1.3.9. 测试

调用 `metathings`的测试命令, 可以发送消息到 `Soda Module Plugin`, 可获得对应的返回值.

*注意:* `--method`的参数与上面 `Module`的配置项 `backend.downstreams`的key值匹配.

```bash
$ eval $(./metathings token issue --domain-id default --username <username> --password <password> --env)
$ ./metathings device unary-call --soda --device <device-id> --module <module-name> --method my_hello --data '{"name": "World"}'

{"message":"Hello, World"}
```

### 4.2. 点亮二极管

#### 4.2.1. 目的

实现一个南向设备, 能够接收 `Metathings`平台的北向控制请求, 控制与树莓派连接的LED灯.

#### 4.2.2. 流程预览

在`Metathings` 平台创建相关的设备, 用于接收北向命令和转发消息到南向设备.

使用 `Python` 构建一个简单的, 基于`HTTTP Server`的 `Soda Module Plugin` 程序, 实现一个 `LED` 服务, 服务包含 `on` 和 `off` 两个方法.

启动该服务程序, 然后再启动 `Device` 和 `Soda Module` 程序.

然后通过测试程序测试请求是否正常工作.

#### 4.2.3. 详细流程

#### 4.2.3.1.  创建工作目录

```bash
$ mkdir 02led
$ export WORKDIR=`pwd`/02led
$ cd $WORKDIR
```

#### 4.2.3.2. 下载 Metathings 客户端

参考 `4.1.3.2.` 章节.

#### 4.2.3.3. 实现 `LED` 服务

这里采用 `Python` 的`Flask` Web 框架实现 `LED` 服务.

暴露的接口有两个:

1. `/on`
2. `/off`

两个接口的 `HTTP` 请求方法为`POST`.

接收参数为`json` 格式.

这里暴露的接口后续会与`Module` 的配置有关联.

```python
#! /usr/bin/env python3

import flask
import gpiozero

app = flask.Flask("led")
_LED = None


def get_led():
    global _LED

    if _LED is None:
        _LED = gpiozero.LED(4)  # 采用 GPIO4 为 LED正极连接针脚, LED负极连接树莓派的GPIO GND针脚.

    return _LED


@app.route("/on")
def on():
    led = get_led()

    led.on()

    return flask.make_response({"status": "on"})


@app.route("/off")
def off():
    led = get_led()

    led.off()

    return flask.make_response({"status": "off"})


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000, debug=True)
```

保存为 `led.py`

#### 4.2.3.4. 运行 `Soda Module Plugin`

```bash
$ python3 led.py &
```

#### 4.2.3.5. 配置 `Device`

参考 `4.1.3.5.`章节.

#### 4.2.3.6. 运行 `Device`

```bash
$ ./metathings device run -c device.yaml &
```

#### 4.2.3.7. 配置 `Module`

这里通过配置文件, 暴露了两个方法, 分别是 `on` 和 `off`, 指向了 `Soda Module Pulgin` 的`/on` 和`/off` 接口.

```yaml
prod:
  ...
  backend:
    ...
    downstreams:
      on:
        path: /on
      off:
        path: /off
    ...
  ...
```

未细讲部分参考 `4.1.3.7.`章节.

 保存为 `module.yaml`

#### 4.2.3.8. 运行 `Module`

```bash
$ ./metathings module run -c module.yaml &
```

#### 4.2.3.9. 测试

在所有的配置正确的情况下(包括软件和硬件的), 通过 `metathings` 工具发送 `on` 请求时, LED灯就会被点亮, 发送`off` 请求时, LED灯就会被熄灭.

```bash
$ eval $(./metathings token issue --domain-id default --username <username> --password <password> --env)
$ ./metathings device unary-call --soda --device <device-id> --module <module-name> --method on  # or off
```

### 4.3. 采集传感器数据

#### 4.3.1. 目的

从传感器读取数据, 并且上传数据到设备所属的`Flow`.

#### 4.3.2. 流程预览

在`Metathings` 平台创建相关的设备, 用于接收北向命令和转发消息到南向设备.

启动`Device` 和`Soda Module` 程序.

程序从传感器读取数据后, 通过 `Soda Module API` 接口发送数据到`Flow`.

#### 4.3.3. 详细流程

#### 4.3.3.1. 创建工作目录

```bash
$ mkdir 03upload_data
$ export WORKDIR=`pwd`/03upload_data
$ cd $WORKDIR
```

#### 4.3.3.2. 下载 Metathings 客户端

参考`4.1.3.2.` 章节

#### 4.3.3.3. 实现采集程序

这里先构建一个模拟的温度传感器接口, 然后读取到数据之后, 采用`Soda Module API` 接口发送数据.

发送数据采用的接口是 `5.1.2.` 章节的接口 `/v1/actions/push_frame_to_flow_once`,

该接口的详细用法请参考该章节.

```python
#! /usr/bin/env python3

from urllib.parse import urljoin
import random

import requests

SODA_MODULE_API_HOST = "localhost"  # 指定 Soda Module API 的地址
SODA_MODULE_API_PORT = 8001  # 指定 Soda Module API 的端口
SODA_MODULE_API_ADDR = "http://{0}:{1}".format(SODA_MODULE_API_HOST, SODA_MODULE_API_PORT)
PUSH_FRAME_TO_FLOW_ONCE = "/v1/actions/push_frame_to_flow_once"


def get_temperature():
    return 30 + 3 * random.random()


def upload_data():
    url = urljoin(SODA_MODULE_ADDR, PUSH_FRAME_TO_FLOW_ONCE)
    data = {"temperature": get_temperature()}
    # 具体的请求数据内容参考接口文档
    req = {
        "flow": {
            "name": "temperature"  # 指定数据往名字为`temperature` 的`Flow` 推送.
        },
        "frame": {
            "data": data
        }
    }
    res = requests.post(url, json=req)
    assert res.status_code == 204

if __name__ == "__main__":
    upload_data()
```

 保存为 `upload.py`

#### 4.3.3.4. 配置 `Device`

参考 `4.1.3.5.`章节

#### 4.3.3.5. 运行 `Device`

参考 `4.1.3.6.`章节

#### 4.3.3.6. 配置 `Module`

```yaml
prod:
  ...
  backend:
    ...
    name: http
    host: 0.0.0.0  # Soda Module API 的监听地址
    port: 8001  # Soda Module API 的监听端口
    ...
  ...
```

未细讲部分参考 `4.1.3.7.`章节

#### 4.3.3.7. 运行 `Module`

参考 `4.1.3.8.`章节

#### 4.3.3.8. 测试

调用 `metathings`的`pull-flow`命令, 监听`Flow` 的数据, 然后再通过上面的程序发送数据到`Flow`, 可以观察到数据成功发送.

```bash
$ eval $(./metathings token issue --domain-id default --username <username> --password <password> --env)
$ ./metathings device pull-flow --device <device-id> --flow <flow-name>
```

打开另外一个终端, 运行发送数据的程序.

```bash
$ python3 upload.py
```

可以观察到上一个终端打印出对应的数据.

### 4.4. 设备在线状态管理

#### 4.4.1. 目的

管理设备的在线状态, 介绍各种心跳策略`Heartbeat Strategy`的区别.

#### 4.4.2. `Auto Strategy`

默认的心跳策略, 由 `Soda Module`程序自行管理.

当`Soda Module` 运行并且能够与`Device` 进行`Heartbeat` 的话, 就可以上线.

这种模式的好处是 `Soda Module Plugin` 的开发者不需要关心相关的内容, 可以尽快地进行`Soda Module Plugin` 的开发.

但是这样带来的问题就是, `Soda Module Plugin`开发者无法主动控制设备的在线状态, 可能会导致一些应用无法根据状态进行正确的逻辑判断.

配置方式:

```yaml
prod:
  ...
  heartbeat:
    strategy: auto  # 默认为 auto, 可选.
  ...
```

#### 4.4.3. `Manual Strategy`

心跳的行为交由`Soda Module Plugin`开发者控制, 只需调用 `/v1/actions/heartbeat`.

接口详细说明参考 `5.1.9.`章节.

这种模式的好处是可以由`Soda Module Plugin`开发者管理心跳的频率, 并且可以最大的自由度控制设备的在线情况.

例如定义传感器无法获取数据后, `Soda Module` 就离线的话, 就可以采用这种方式.

这种模式带来的问题是, 开发者需要自行管理该接口的调用, 对开发者造成一定的不便.

配置方式:

```yaml
prod:
  ...
  heartbeat:
    strategy: manual
  ...
```

#### 4.4.4. `Reverse Strategy`

心跳的检测逻辑交由`Soda Module Plugin`开发者实现, 但是心跳的发起由`Soda Module`实现.

心跳的检测逻辑实现接口可以阅读 `5.2.1.`章节了解更多.

与`manual strategy` 相比, `reverse strategy`对开发者的心智负担更加轻, 开发者只需要关心实现检测逻辑即可.

而什么时候, 怎么通知其他组件的工作, 交由了已经实现的`Soda Module` 完成.

配置方法:

```yaml
prod:
  ...
  heartbeat:
    strategy: reverse
  ...
```

### 4.5. 认证机制

#### 4.5.1. 目的

`Metathings` 平台与 `Device`, `Device` 与 `Module` 之间的通信是采用认证机制, 所以之间的通讯是经过认证的.

而以上的场景都是默认`Soda Module` 与`Soda Module Plugin` 处于受信环境下的通信.

可能存在第三方攻击的情况, 所以加入了`Soda Module` 对`Soda Module Plugin` 的认证机制.

暂时的认证机制是`Soda Module Plugin` 与`Soda Module` 之间采用 `secret` 机制(即双方在运行前保存相同的 `secret`, 供通讯时使用).

将来会提供更加丰富的认证机制.

#### 4.5.2. `Module` 配置

```yaml
prod:
  ...
  backend:
    auth:
      name: secret
      secret: <secret>  # 这里配置预设的 `secret`, 等下通信时需要带上.
  ...
```

#### 4.5.3. `Soda Module Plugin` 开发更改

需要在调用 `Soda Module API` 时, 需要在`HTTP Request Header`上设置 `Authorization` 字段.

格式如下:

HTTP Header: `Authorization: Bearer <secret>`

## 5. `API`详解

### 5.1. Soda Module HTTP API

#### 5.1.1. Show

获取 Module 的相关信息.

uri: `/v1/actions/show`

method: `POST`

response body:

```json
{
    "module": {
        "alias": <alias>,  # Alias: Module的别名, 显示作用.
        "component": <component>,  # Component: 描述 Module 所属的 Component 值, 该值暂时没有意义.
        "deviceId": <device-id>,  # DeviceID: Module 所属的 Device 的 ID.
        "endpoint": <endpoint>,  # Endpoint: Device 与 Module 通信的连接地址, SodaModule 通常是以 mtp+soda开始.
        "id": <id>,  # ID: Module 的 ID.
        "name": <name>,  # Name: Module 的名字, 通常情况下接口调用采用 Name 来作为索引.
        "state": "MODULE_STATE_OFFLINE"  # 不可靠值, 该值暂时没有意义.
    }
}
```

response code: `200`

#### 5.1.2. PushFrameToFlowOnce

推送数据到指定的`Flow`.

uri: `/v1/actions/push_frame_to_flow_once`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
  "id": <id>,  # RequestID: 请求的 ID, 可选.
  "flow": {
    "name": <name>  # Name: 推送到对应的 Flow 的名字.
  },
  "frame": {
    "data": <data>,  # Data: 推送的数据, 以 json 格式填入.
    "ts": <ts>  # Timestamp: 数据生成时间, 可选. 格式为 RFC3339.
  }
}
```

response code: `204`

#### 5.1.3. PutObject

传输文件到`SimpleStorage`.

uri: `/v1/actions/put_object`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "object": {
        "name": <name>,  # Name: 文件的名字, 包含路径.
        "content": <content>  # Content: 文件的内容.
    }
}
```

response code: `200`

#### 5.1.4. GetObject

获取`SimpleStorage`上的文件的元信息.

uri: `/v1/actions/get_object`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "object": {
        "name": <name>  # Name: 文件的名字, 包含路径.
    }
}
```

response body:

```json
{
    "object": {
        "device": {
            "id": <id>  # DeviceID: 所属的 Device 的 ID.
        },
        "etag": <etag>,  # Etag: 文件的哈希值, 用来判断文件是否发生更改.
        "lastModified": "2020-12-24T08:54:26.329487636Z",  # LastModified: 最后修改时间.
        "length": <length>,  # Length: 文件长度, 注意: 该值为 string.
        "name": <name>,  # Name: 文件名, 不包含路径.
        "prefix": <prefix>  # Prefix: 文件所在路径.
    }
}
```

response code: `200`

#### 5.1.5. GetObjectContent

获取`SimpleStorage`上的文件的内容.

uri: `/v1/actions/get_object_content`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "object": {
        "name": <name>  # Name: 文件的名字, 包含路径.
    }
}
```

response body:

```json
{
    "content": <content>  # Content: 文件的内容, 以 base64编码, 读取实际内容需要通过 base64解码.
}
```

response code: `200`

#### 5.1.6. RemoveObject

删除`SimpleStorage`上的文件.

暂时不支持递归删除, 需要一个一个删除.

uri: `/v1/actions/remove_object`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "object": {
        "name": <name>  # Name: 文件的名字, 包含路径.
    }
}
```

response code: `204`

#### 5.1.7. RenameObject

重命名`SimpleStorage`上的文件.

uri: `/v1/actions/rename_object`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "source": {
        "name": <name>  # Name: 将要被重命名的文件, 包含路径.
    },
    "destination": {
        "name": <name>  # Name: 重命名文件的目的地址, 包含路径.
    }
}
```

response code: `204`

#### 5.1.8. ListObjects

列出所属设备的`SimpleStorage`下的文件.

uri: `/v1/actions/list_objects`

method: `POST`

request headers:

`Content-Type: application/json`

request body:

```json
{
    "object": {
        "name": <name>  # 需要列出的文件或文件夹的路径
    },
    "recursive": <recursive>,  # 是否递归列出, 可选, 默认为 false.
    "dpeth": <depth>,  # 递归查询的深度, 可选,  默认为 16.
}
```

response body:

```json
{
    "objects": [
        {
            "device": {
                "id": <device-id>  # DeviceID: 所属 device 的 id.
            },
            "etag": <etag>,  # Etag: 文件的哈希值, 用来判断文件是否发生更改.
            "lastModified": "2020-12-24T08:54:26.329487636Z",  # LastModified: 最后修改时间.
            "length": <length>,  # Length: 文件长度, 注意: 该值为 string.
            "name": <name>,  # Name: 文件名, 不包含路径.
            "prefix": <prefix>  # Prefix: 文件所在路径.
        },
        # ...
    ]
}
```

response code: `200`

#### 5.1.9. Heartbeat

`Soda Module`的`heartbeat strategy`配置为`manual` 时, 会暴露的接口.

本接口提供主动的心跳接口, 供`Soda Module Plugin`主动心跳使用.

uri: `/v1/actions/heartbeat`

method: `POST`

response code: `204`

### 5.2. Soda Module Plugin API

#### 5.2.1. Healthz

如果`Heartbeat Strategy`选择了`reverse`之后, `Soda Module Plugin` 端需要实现检查是否健康的函数,

供`Soda Module` 调用检查是否存活.

uri: `/healthz`

method: `GET`

response body:

`ok`

response code: `200`
