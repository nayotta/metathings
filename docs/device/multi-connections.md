# 设备连接冗余设计

支持多路连接, 防止某个deviced crash导致设备无法快速上线的问题.

暂时未暴露任何参数接口, 需要重新编译device端程序支持.

修改`NewMetathingsDeviceServiceOption`函数的`ExpectedConnections`参数, 即支持多链路连接.
