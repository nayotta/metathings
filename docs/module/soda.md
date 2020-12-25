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

实现一个南向设备, 能够接收 `Metathings` 平台的北向请求, 并且作出回应.

#### 4.1.2.  流程预览

在`Metathings` 平台创建相关的设备, 用于接收北向命令和转发消息到南向设备.

使用 `Python` 构建一个简单的 `HTTP Server`, 实现一个 `Hello` 服务.

启动该服务程序, 然后再启动`Device` 和`Soda Module` 程序.

然后通过测试程序测试请求是否正常响应.

#### 4.1.3. 详细流程

1. 创建一个工作目录

```bash
$ mkdir 01hello
$ export WORKDIR=`pwd`/01hello
$ cd $WORKDIR
```

2. 下载 Metathings 客户端

进入 `https://github.com/nayotta/metathings/releases/latest`

下载 `metathings_<version>_linux_amd64.tar.gz` 文件.

```bash
# 修改对应的 version
$ wget https://github.com/nayotta/metathings/releases/download/v<version>/metathings_<version>_linux_amd64.tar.gz
$ tar zxvf metathings_<version>_linux_amd64.tar.gz
```


3. 实现 `HTTP Server`

```python
import flask
from flask import request

app = flask.Flask("hello")


@app.route("/hello", methods=["POST"])
def hello():
    return flask.make_response({"message": "Hello, {0}".format(request.json["name"])})


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000, debug=True)
```

保存为 `hello.py`

4.  运行 `HTTP Server`

```bash
$ python3 hello.py &
```

5. 配置 `Device`

```yaml
debug:  # 与MT_STAGE环境变量一致
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

6. 运行 `Device`

```bash
$ MT_STAGE=debug ./metathings device run -c device.yaml &
```

7. 配置 `Module`

```yaml
debug:  # 与MTC_STAGE环境变量一致
  name: hello  # Module 的名字, 在 Metathings平台创建时定义
  service:
    scheme: mtp+soda  # Soda Module的协议, 在 Metathings平台创建时定义
    host: 0.0.0.0  # Soda Module的监听地址, 在 Metathings平台创建时定义
    port: 13401  # Soda Module的监听端口, 在 Metathings平台创建时定义
  log:
    level: debug
  heartbeat:
    interval: 15  # Module 与 Device 的心跳间隔时间
  credential:
    id: <credential-id>  # Module 的 Credential ID, 在 Metathings 平台生成
    secret: <credential-secret>  # Module 的 Credential Secret, 在 Metathings 平台生成
  service_endpoint:
    device:
      address: 127.0.0.1:5002  # Device 的通讯地址, 参考上面 Device 的配置
      plain_text: true  # Device 与 Module 是内部通讯, 所以采用 PlainText模式即可
    default:
      address:  api.metathings.nayotta.com:443  # Metathings的地址, 默认即可
  backend:
    name: http  # Soda Module 的 Backend 驱动, 暂时只支持 http
    host: 0.0.0.0  # http 服务的监听地址
    port: 8001  # http 服务的监听端口
    target:
      url: http://127.0.0.1:8000  # 指向 Python 的 HTTP Server
    downstreams:
      hello:  # 暴露 Python HTTP Server 的服务, 名字叫 hello
        path: /hello  # 指向 Python HTTP Server 的路径
```

保存为 `module.yaml`

8. 运行 `Module`

```bash
$ MTC_STAGE=debug ./metathings module run -c module.yaml &
```

9. 测试

调用 `metathings`的测试命令, 可以发送消息到 `Python HTTP Server`, 可获得 `Python HTTP Server` 的返回值.

```bash
$ ./metathings device unary-call --soda --device <device-id> --module <module-name> --method hello --data '{"name": "World"}'

{"message":"Hello, World"}
```

### 4.2. 点亮二极管

TBD

### 4.3. 采集传感器数据

TBD

### 4.4. 设备上线与下线

TBD

### 4.5. 认证机制

TBD

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

### 5.2. Soda Module Plugin API

#### 5.2.1. Healthz

如果`Heartbeat Strategy`选择了`reverse`之后, `Soda Module Plugin` 端需要实现检查是否健康的函数,

供`Soda Module` 调用检查是否存活.

uri: `/healthz`

method: `GET`

response body:

`ok`

response code: `200`
