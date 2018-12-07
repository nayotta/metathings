package metathingsmqttdconnection

import (
	"fmt"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	log "github.com/sirupsen/logrus"
)

// MqttBridgeOpt MqttBridgeOpt
type MqttBridgeOpt struct {
	server    *url.URL
	rootKey   string
	upKey     string
	downKey   string
	statusKey string
	client    emitter.Emitter
	logger    log.FieldLogger
	storage   Storage
}

// connect callback
func (that *MqttBridgeOpt) connectCallback(_ emitter.Emitter) {
	var err error

	fmt.Println("get connect")
	err = that.sub(that.upKey, "+/up/")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = that.sub(that.statusKey, "+/status/")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// connect lost callback
func (that *MqttBridgeOpt) connectLostCallback(_ emitter.Emitter, err error) {
	fmt.Println("lost connect:", err)
}

// presence callback
func (that *MqttBridgeOpt) presenceCallback(_ emitter.Emitter, msg emitter.PresenceEvent) {
	fmt.Println("get presence:", msg.Occupancy)
}

// message callback All message process here
func (that *MqttBridgeOpt) messageCallback(client emitter.Emitter, msg emitter.Message) {
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
func (that *MqttBridgeOpt) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
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
func (that *MqttBridgeOpt) Pub(msg []byte, path string) error {
	if msg == nil {
		return ErrMqttMsgBlank
	}

	r := that.client.Publish(that.downKey, path, msg)
	if r.Wait() && r.Error() != nil {
		return ErrMqttPubFailed
	}

	return nil
}

// sub sub +/down/
func (that *MqttBridgeOpt) sub(key, path string) error {
	r := that.client.Subscribe(key, path)
	if r.Wait() && r.Error() != nil {
		return ErrMqttSubFailed
	}

	return nil
}

// createSecretKey async create
func (that *MqttBridgeOpt) createSecretKey(path string) error {
	req := &emitter.KeyGenRequest{
		Key:     that.rootKey,
		Channel: path,
		Type:    "rwslp",
		TTL:     0,
	}

	r := that.client.GenerateKey(req)
	if r.Wait() && r.Error() != nil {
		return ErrMqttKeygenFailed
	}

	return nil
}

//createUpKey
func (that *MqttBridgeOpt) createUpKey() error {
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
func (that *MqttBridgeOpt) createDownKey() error {
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
func (that *MqttBridgeOpt) createStatusKey() error {
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

// CloseBridge CloseBridge
func (that *MqttBridgeOpt) CloseBridge() {
	if that.client.IsConnected() == true {
		that.client.Disconnect(0)
	}
}

// InitMqttBridge InitMqttBridge
func (that *MqttBridgeOpt) InitMqttBridge() error {
	var err error

	opt := emitter.NewClientOptions()
	opt.Servers = append(opt.Servers, that.server)

	opt.SetOnKeyGenHandler(that.keygenCallback)
	opt.SetOnMessageHandler(that.messageCallback)
	opt.SetOnConnectHandler(that.connectCallback)
	opt.SetOnPresenceHandler(that.presenceCallback)
	opt.SetOnConnectionLostHandler(that.connectLostCallback)

	that.client = emitter.NewClient(opt)
	handle := that.client.Connect()
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

	err = that.sub(that.upKey, "+/up/")
	if err != nil {
		return err
	}

	err = that.sub(that.statusKey, "+/status/")
	if err != nil {
		return err
	}

	return nil
}

// NewMqttBridge new mqtt client bridge
func NewMqttBridge(args []interface{}) (MqttBridge, error) {
	var ok bool
	var err error
	var key string
	var val interface{}
	var server *url.URL
	var rootKey string
	var logger log.FieldLogger

	if len(args)%2 != 0 {
		return nil, ErrInvalidArgument
	}

	for i := 0; i < len(args); i += 2 {
		key, ok = args[i].(string)
		if !ok {
			return nil, ErrInvalidArgument
		}
		val = args[i+1]

		switch key {
		case "url":
			urlStr, ok := val.(string)
			if !ok {
				return nil, ErrInvalidArgument
			}

			server, err = url.Parse(urlStr)
			if err != nil {
				return nil, ErrInvalidArgument
			}
		case "rootkey":
			rootKey, ok = val.(string)
			if !ok {
				return nil, ErrInvalidArgument
			}
		case "logger":
			if logger, ok = args[i+1].(log.FieldLogger); !ok {
				return nil, ErrInvalidArgument
			}
		}
	}
	return &MqttBridgeOpt{
		server:  server,
		rootKey: rootKey,
		logger:  logger,
	}, nil
}
