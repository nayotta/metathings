package metathingsmqttdconnection

import (
	"errors"
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
	statusKey string
	client    emitter.Emitter
	pubPath   string
	subPath   string
	logger    log.FieldLogger
	storage   Storage
}

// connect callback
func (that *MqttBridgeOpt) connectCallback(_ emitter.Emitter) {
	fmt.Println("get connect")
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
		fmt.Println("keygen res:", msg.Key)
		topicStr := msg.Channel
		switch topicStr {
		case "+/up/":
			that.upKey = msg.Key
		case "+/status/":
			that.statusKey = msg.Key
		default:
			// TODO(zh) device channel key
		}
	}
}

// Pub Publish
func (that *MqttBridgeOpt) Pub(msg string) error {
	var err error

	if msg == "" {
		err = errors.New("mqtt failed pub blank msg")
		that.logger.WithError(err).Errorf("mqtt failed pub blank msg")
		return err
	}

	r := that.client.Publish(that.upKey, that.pubPath, msg)
	fmt.Println("pub msg:", that.pubPath)
	if r.Wait() && r.Error() != nil {
		err = errors.New("mqtt failed pub")
		that.logger.WithError(err).Errorf("mqtt failed pub")
		return err
	}

	return nil
}

// sub sub +/down/
func (that *MqttBridgeOpt) sub(key, path string) error {
	var err error

	r := that.client.Subscribe(key, path)
	fmt.Println("sub msg:", path)
	if r.Wait() && r.Error() != nil {
		err = errors.New("mqtt failed sub")
		that.logger.WithError(err).Errorf("mqtt failed sub")
		return err
	}

	return nil
}

// GetUpKey get key for channel
func (that *MqttBridgeOpt) GetUpKey() (string, error) {
	var err error

	if that.upKey == "" {
		err = errors.New("no sercret key found")
		that.logger.WithError(err).Errorf("no sercret key found")
		return "", err
	}

	return that.upKey, nil
}

// createSecretKey async create
func (that *MqttBridgeOpt) createSecretKey(path string) error {
	var err error

	req := &emitter.KeyGenRequest{
		Key:     that.rootKey,
		Channel: path,
		Type:    "rwslp",
		TTL:     0,
	}

	r := that.client.GenerateKey(req)
	if r.Wait() && r.Error() != nil {
		err = errors.New("keygen failed")
		that.logger.WithError(err).Errorf("keygen failed")
		return err
	}

	return nil
}

//createUpKey
func (that *MqttBridgeOpt) createUpKey() error {
	var err error

	//create key
	err = that.createSecretKey("+/up/")
	if err != nil {
		that.logger.WithError(err).Errorf("createUpKey failed")
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.upKey != "" {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

//createStatusKey
func (that *MqttBridgeOpt) createStatusKey() error {
	var err error

	//create key
	err = that.createSecretKey("+/status/")
	if err != nil {
		that.logger.WithError(err).Errorf("createStatusKey failed")
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.statusKey != "" {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	return nil
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
		err = errors.New("mqtt failed connect client")
		that.logger.WithError(err).Errorf("mqtt failed connect client")
		return err
	}

	err = that.createUpKey()
	if err != nil {
		that.logger.WithError(err).Errorf("create up key failed")
		return err
	}

	err = that.createStatusKey()
	if err != nil {
		that.logger.WithError(err).Errorf("create status key failed")
		return err
	}

	err = that.sub(that.upKey, "+/up/")
	if err != nil {
		that.logger.WithError(err).Errorf("Sub +/up/ failed")
		return err
	}

	err = that.sub(that.statusKey, "+/status/")
	if err != nil {
		that.logger.WithError(err).Errorf("Sub +/status/ failed")
		return err
	}

	return nil
}

// NewMqttBridge new mqtt client bridge
func NewMqttBridge(
	server *url.URL,
	rootKey string,
	pubPath string,
	subPath string,
	logger log.FieldLogger,
	storage Storage,
) (MqttBridge, error) {
	if pubPath == "" {
		pubPath = "+/up/"
	}
	if subPath == "" {
		subPath = "+/down/"
	}

	return &MqttBridgeOpt{
		server:  server,
		rootKey: rootKey,
		pubPath: pubPath,
		subPath: subPath,
		logger:  logger,
		storage: storage,
	}, nil
}
