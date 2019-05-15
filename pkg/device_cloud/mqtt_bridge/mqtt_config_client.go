package metathingsdevicecloudmqttbridge

import (
	"time"

	emitter "github.com/emitter-io/go"
)

// generate key callback
func (that *mqttBridge) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
	if msg.Status == 200 {
		topicStr := msg.Channel
		switch topicStr {
		case "flow/+/+/up/":
			that.flowUpKey = msg.Key
		case "+/up/#/":
			that.upKey = msg.Key
		case "+/statusup/#/":
			that.statusUpKey = msg.Key
		case "+/down/#/":
			that.downKey = msg.Key
		}
	}
}

// createSecretKey async create
func (that *mqttBridge) createSecretKey(path string) error {
	req := &emitter.KeyGenRequest{
		Key:     that.rootKey,
		Channel: path,
		Type:    "rwslp",
		TTL:     0,
	}

	r := that.configClient.GenerateKey(req)
	if r.Wait() && r.Error() != nil {
		return ErrMqttKeygenFailed
	}

	return nil
}

//createUpKey
func (that *mqttBridge) createUpKey() error {
	err := that.createSecretKey("+/up/#/")
	if err != nil {
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.upKey != "" {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return ErrMqttUpKeygenFailed
}

//createFlowUpKey
func (that *mqttBridge) createFlowUpKey() error {
	err := that.createSecretKey("flow/+/+/up/")
	if err != nil {
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.flowUpKey != "" {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return ErrMqttUpKeygenFailed
}

//createDownKey
func (that *mqttBridge) createDownKey() error {
	err := that.createSecretKey("+/down/#/")
	if err != nil {
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.downKey != "" {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return ErrMqttDownKeygenFailed
}

//createStatusUpKey
func (that *mqttBridge) createStatusUpKey() error {
	err := that.createSecretKey("+/statusup/#/")
	if err != nil {
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.statusUpKey != "" {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return ErrMqttStatusKeygenFailed
}

// InitMqttBridge InitMqttBridge
func (that *mqttBridge) InitMqttBridge() error {
	var err error

	opt := emitter.NewClientOptions()
	opt.Servers = append(opt.Servers, that.host)

	opt.SetOnKeyGenHandler(that.keygenCallback)

	that.configClient = emitter.NewClient(opt)
	handle := that.configClient.Connect()
	if handle.Wait() && handle.Error() != nil {
		return ErrMqttConnectFailed
	}

	err = that.createUpKey()
	if err != nil {
		return err
	}

	err = that.createDownKey()
	if err != nil {
		return err
	}

	err = that.createStatusUpKey()
	if err != nil {
		return err
	}

	err = that.createFlowUpKey()
	if err != nil {
		return err
	}

	return nil
}
