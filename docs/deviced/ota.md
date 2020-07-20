# Metathings Device & Module OTA 指南

## 1. 简介

OTA([Over-the-air programming](https://zh.wikipedia.org/wiki/%E7%A9%BA%E4%B8%AD%E7%BC%96%E7%A8%8B))功能允许 Device 与 Module 的程序支持远程更新.

OTA 带来的优势是可以不用物理接触物理设备, 即可对设备进行软件的升级(修复 bug 或升级功能)操作.

## 2. 术语介绍

### 2.1. Metathings Deviced Service

`Deviced Service` 是 `Device` 接入的服务, 负责转发指令请求和状态维护等功能的服务.

### 2.2. Device & Module

`Module` 是指令的执行终端, 负责大多数的行为执行操作.

通常假设 `Module` 无法直接连接到 `Deviced Service`上, 所以添加了 `Device` 来负责通信中转功能.

通常情况下, 客户端调用指令以 `Device`为最终调用实体, 至于 `Device` 与 `Module`之间的通信协议可以通过不同的实现(http, grpc, bluetooth等)来完成.

### 2.3. FirmwareHub

`FirmwareHub` 是一个抽象的概念, 主要是为了管理 `Device` 与 `FirmwareDescriptor` 之前的关系而存在的概念.

`FirmwareHub` 包含了 `Device` 与 `FirmwareDescriptor`, 在同一个`FirmwareHub`下的`Device`, 可以设置`FirmwareDescriptor` 为该 `FirmwareHub`下的 `FirmwareDescriptor`.

### 2.4. firmwareDescriptor

`FirmwareDescriptor`主要负责描述具体的版本信息, 例如 `name`, `uri`, `sha256`等.

当然版本信息内容是开放的, 任何想使用OTA功能的实现者都允许自定义版本信息内容.

## 3. OTA协议

### 3.1. 简介

`OTA`协议分为两部分.

一部分是 `Deviced`与 `Device`的 `OTA`协议, 负责更新 `Device`.

一部分是 `Device` 与 `Module`的 `OTA`协议, 负责更新 `Module`.

`FirmwareDescriptor` 原则上应当包含描述 `Device` 与 `Module`的版本信息, 但是如果特殊情况(例如 Simple 类型的设备, 与 `Device`要求更新的设备)都是允许 Device 的版本信息为空, 意思是, `Device`的 `OTA`行为不发生.

### 3.2. Deviced 与 Device 的协议

`Deviced` 发送特殊的 `UnaryCall Request`到 `Device`上完成 `Device` 的OTA实现.

特殊的 `UnaryCall Request`结构如下:

```
{
  "component": "system",
  "name": "system",
  "method": "sync_firmware"
}
```

详情请参考 [connectionCenter.SyncFirmware](https://github.com/nayotta/metathings/blob/master/pkg/deviced/connection/connection.go) 与 [MetathingsDeviceServiceImpl.handle_system_unary_request_sync_firmware](https://github.com/nayotta/metathings/blob/master/pkg/device/service/handle.go)实现.

*注意* `Device`有义务在启动时, 把当前的版本写入 SimpleStorage(`/sys/firmware/device/version/current`) 内.

*注意2* `Simple` 类型的 `Device` 采用的是 `Device Cloud`代理 `Device`, 此类型设备的 `Device`不支持 OTA(`Module`是支持的).

### 3.3. Device 与 Module 的协议

`Device` 发送特殊的 `UnaryCall Value` 到 `Module`上完成 `Module` 的OTA实现.

 特殊的 `UnaryCall Value`结构如下:

```
{
  "method": "SyncFirmware"
}
```

 详情请参考 [MetathingsDeviceServiceImpl.do_sync_modules_firmware](https://github.com/nayotta/metathings/blob/master/pkg/device/service/handle.go)(`Advanced Device`), [DeviceConnection.handle_system_unary_request_sync_firmware](https://github.com/nayotta/metathings/blob/master/pkg/device_cloud/service/handle.go)(`Simple Device`) 和 [ComponentExtFirmwareService.SyncFirmware](https://github.com/nayotta/metathings/blob/master/pkg/component_ext/firmware/service/service.go) 等.

*注意* `Advanced Module`的 `FirmwareService` 是以插件的形式编写的, 所以只需要引入并且初始化就会获取OTA的功能. [Example](https://github.com/nayotta/metathings/blob/master/pkg/component_ext/firmware/service/example_test.go)

## 4. OTA流程

### 4.1. `Device` 流程概览

### 4.2. `Device` 详细流程

### 4.3. `Module` 流程概述

### 4.4. `Module` 详细流程

## 5. 总结
