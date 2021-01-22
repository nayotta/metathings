package metathings_component

func init() {
	register_module_proxy_factory("soda", new(GrpcModuleProxyFactory))
}
