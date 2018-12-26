package cmd_contrib

// MqttBridgeOptioner MqttBridgeOptioner
type MqttBridgeOptioner interface {
	GetBrokerP() *string
	GetBroker() string
	SetBroker(string)

	GetRootkeyP() *string
	GetRootkey() string
	SetRootKey(string)
}

// MqttBridgeOption MqttBridgeOption
type MqttBridgeOption struct {
	MqttBridge struct {
		Broker  string
		Rootkey string
	}
}

// GetBrokerP GetBrokerP
func (that *MqttBridgeOption) GetBrokerP() *string {
	return &that.MqttBridge.Broker
}

// GetBroker GetBroker
func (that *MqttBridgeOption) GetBroker() string {
	return that.MqttBridge.Broker
}

// SetBroker SetBroker
func (that *MqttBridgeOption) SetBroker(broker string) {
	that.MqttBridge.Broker = broker
}

// GetRootkeyP GetRootkeyP
func (that *MqttBridgeOption) GetRootkeyP() *string {
	return &that.MqttBridge.Rootkey
}

// GetRootkey GetRootkey
func (that *MqttBridgeOption) GetRootkey() string {
	return that.MqttBridge.Rootkey
}

// SetRootKey SetRootKey
func (that *MqttBridgeOption) SetRootKey(rootkey string) {
	that.MqttBridge.Rootkey = rootkey
}
