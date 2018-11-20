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
	clientID  string
	rootKey   string
	secretKey string
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

// message callback
func (that *MqttBridgeOpt) messageCallback(client emitter.Emitter, msg emitter.Message) {
	fmt.Println("get message:", msg.Payload())
}

// generate key callback
func (that *MqttBridgeOpt) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
	if msg.Status == 200 {
		fmt.Println("keygen res:", msg.Key)
		that.secretKey = msg.Key
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

	r := that.client.Publish(that.secretKey, that.pubPath, msg)
	fmt.Println("pub msg:", that.pubPath)
	if r.Wait() && r.Error() != nil {
		err = errors.New("mqtt failed pub")
		that.logger.WithError(err).Errorf("mqtt failed pub")
		return err
	}

	return nil
}

// Sub Sub
func (that *MqttBridgeOpt) Sub() error {
	var err error

	r := that.client.Subscribe(that.secretKey, that.subPath)
	fmt.Println("sub msg:", that.subPath)
	if r.Wait() && r.Error() != nil {
		err = errors.New("mqtt failed sub")
		that.logger.WithError(err).Errorf("mqtt failed sub")
		return err
	}

	return nil
}

func (that *MqttBridgeOpt) createSecretKey() error {
	var err error

	req := &emitter.KeyGenRequest{
		Key:     that.rootKey,
		Channel: that.subPath,
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

	err = that.createSecretKey()
	if err != nil {
		that.logger.WithError(err).Errorf("createKey failed")
		return err
	}

	// wait 5s for key response
	for i := 0; i < 100; i++ {
		if that.secretKey != "" {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	err = that.Sub()
	if err != nil {
		that.logger.WithError(err).Errorf("Sub failed")
		return err
	}

	return nil
}

// NewMqttBridge new mqtt client bridge
func NewMqttBridge(
	server *url.URL,
	clientID string,
	rootKey string,
	pubPath string,
	subPath string,
	logger log.FieldLogger,
	storage Storage,
) (MqttBridge, error) {
	if pubPath == "" {
		pubPath = fmt.Sprintf("%v/up/", clientID)
	}
	if subPath == "" {
		subPath = fmt.Sprintf("%v/down/", clientID)
	}

	return &MqttBridgeOpt{
		server:   server,
		clientID: clientID,
		rootKey:  rootKey,
		pubPath:  pubPath,
		subPath:  subPath,
		logger:   logger,
		storage:  storage,
	}, nil
}
