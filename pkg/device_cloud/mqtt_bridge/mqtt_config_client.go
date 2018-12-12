package metathingsdevicecloudmqttbridge

import (
	"fmt"
	"time"

	emitter "github.com/emitter-io/go"
)

// sub sub +/down/
func (that *mqttBridge) Sub(key, path string) error {
	r := that.configClient.Subscribe(key, path)
	if r.Wait() && r.Error() != nil {
		return ErrMqttSubFailed
	}

	return nil
}

// connect callback
func (that *mqttBridge) connectCallback(_ emitter.Emitter) {
	var err error

	fmt.Println("get connect")

	err = that.Sub(that.statusKey, "+/status/")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// connect lost callback
func (that *mqttBridge) connectLostCallback(_ emitter.Emitter, err error) {
	fmt.Println("lost connect:", err)
}

// presence callback
func (that *mqttBridge) presenceCallback(_ emitter.Emitter, msg emitter.PresenceEvent) {
}

// message callback All message process here
func (that *mqttBridge) messageCallback(client emitter.Emitter, msg emitter.Message) {
	topicStr := msg.Topic()
	deviceID := getTopicDeviceID(topicStr)
	msgType := getTopicType(topicStr)

	go func() {
		switch msgType {
		case "up":
			fmt.Printf("up from device:%s message:%s\n", deviceID, msg.Payload())
		case "status":
			fmt.Printf("status from device:%s message:%s\n", deviceID, msg.Payload())
		default:
			fmt.Printf("unknown message type:%s\n", msgType)
		}
	}()
}

// generate key callback
func (that *mqttBridge) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
	if msg.Status == 200 {
		topicStr := msg.Channel
		switch topicStr {
		case "+/up/#/":
			that.upKey = msg.Key
			fmt.Println("upKey res:", msg.Key)
		case "+/status/#/":
			that.statusKey = msg.Key
			fmt.Println("statusKey res:", msg.Key)
		case "+/down/#/":
			that.downKey = msg.Key
			fmt.Println("downKey res:", msg.Key)
		default:
			// TODO(zh) device channel key
		}
	}
}

// Pub Publish
func (that *mqttBridge) Pub(msg []byte, path string) error {
	if msg == nil {
		return ErrMqttMsgBlank
	}

	r := that.configClient.Publish(that.downKey, path, msg)
	if r.Wait() && r.Error() != nil {
		return ErrMqttPubFailed
	}

	return nil
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

//createStatusKey
func (that *mqttBridge) createStatusKey() error {
	err := that.createSecretKey("+/status/#/")
	if err != nil {
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.statusKey != "" {
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
	opt.SetOnMessageHandler(that.messageCallback)
	opt.SetOnConnectHandler(that.connectCallback)
	opt.SetOnPresenceHandler(that.presenceCallback)
	opt.SetOnConnectionLostHandler(that.connectLostCallback)

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

	err = that.createStatusKey()
	if err != nil {
		return err
	}

	err = that.Sub(that.statusKey, "+/status/")
	if err != nil {
		return err
	}

	return nil
}

func (that *mqttBridge) HeartBeatSelect() {

}

func (that *mqttBridge) KeyGen() {

}
