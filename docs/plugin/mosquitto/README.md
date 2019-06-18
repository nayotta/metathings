# Mosquitto Plugin 使用指南

## 概要

`Mosquitto` 插件, 在 `identityd2` 创建 `credential` 和删除 `credential`时, 会通知 `Mosquitto` 服务增加或删除对应的用户.

## 使用步骤

### 1. 配置 Mosquitto Plugin Service

可以参考 `thrid_party/mosquitto/mosquitto-plugin.template.yaml`

**注意**: `webhhok.secret` 需要和 `identityd2` 配置的 `webhook.mosquitto-plugin.secret` 字段相同.

### 2. 配置 Identityd2 Webhook 功能

启用 `identityd2 webhook` 功能

```
STAGE:
  ...
  webhook_service:
    default:
      content_type: application/json
    mosquitto-plugin:
      url: http://<mosquitto-plugin-service>/webhook
      secret: <webhook-secret>
  ...
```

### 3. 启动 Mosquitto Plugin Service

可以采取多种启动方式.

可以参考 `thrid_party/mosquitto/docker-compose.yaml`

### 4. 启动 Identityd2 服务

可以采取多种启动方式.


### 5. 创建设备

创建 `Device` , 然后再创建相应的 `Credential`.

创建 `Credential` 后, 获取得到 `CredentialId` 和 `CredentialSecret`.

### 6. 设备连接到 Mosquitto 服务

调用 `examples/05-get-mqtt-client-password/main.go` 程序获得连接 `Mosquitto` 的账户名和密码.

其他语言的实现可以参考上面的程序移植实现.
