package metathings_deviced_connection

type kafkaBridge struct{}

func (self *kafkaBridge) Send([]byte) error {
	panic("unimplemented")
}

func (self *kafkaBridge) Recv() ([]byte, error) {
	panic("unimplemented")
}

type kafkaBridgeFactory struct{}

func (self *kafkaBridgeFactory) BuildBridge(device_id string, session int32) (Bridge, error) {
	panic("unimplemented")
}

func (self *kafkaBridgeFactory) GetBridge(br_id string) (Bridge, error) {
	panic("unimplemented")
}

func new_kafka_bridge_factory(args ...interface{}) (BridgeFactory, error) {
	return &kafkaBridgeFactory{}, nil
}

func init() {
	register_bridge_factory_factory("kafka", new_kafka_bridge_factory)
}
