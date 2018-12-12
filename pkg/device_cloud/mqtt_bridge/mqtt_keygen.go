package metathingsdevicecloudmqttbridge

import (
	"time"

	emitter "github.com/emitter-io/go"
)

// KeygenCenter KeygenCenter
type KeygenCenter interface {
	KeyGen(bridge MqttBridge) (string, error)
}

type keygenCenter struct {
	topic      string
	timeout    time.Duration
	kengenChan chan error
	key        string
}

func (that *keygenCenter) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
	if msg.Status != 200 {
		that.kengenChan <- ErrUnexpectedResponse
	}

	if msg.Channel != that.topic {
		that.kengenChan <- ErrMqttKeygenFailed
	}

	that.key = msg.Key
	that.kengenChan <- nil
}

// KeyGen KeyGen
func (that *keygenCenter) KeyGen(bridge MqttBridge) (string, error) {
	var err error

	opt := emitter.NewClientOptions()
	host, err := bridge.GetHost()
	if err != nil {
		return "", ErrInvalidArgument
	}
	rootkey, err := bridge.GetRootKey()

	if err != nil {
		return "", ErrInvalidArgument
	}
	opt.Servers = append(opt.Servers, host)
	opt.SetOnKeyGenHandler(that.keygenCallback)
	client := emitter.NewClient(opt)
	handle := client.Connect()
	if handle.Wait() && handle.Error() != nil {
		return "", ErrMqttConnectFailed
	}
	defer client.Disconnect(0)

	that.kengenChan = make(chan error)

	req := &emitter.KeyGenRequest{
		Key:     rootkey,
		Channel: that.topic,
		Type:    "rwslp",
		TTL:     0,
	}

	r := client.GenerateKey(req)
	if r.Wait() && r.Error() != nil {
		return "", ErrMqttKeygenFailed
	}

	select {
	case err := <-that.kengenChan:
		if err != nil {
			return "", err
		}
		return that.key, nil
	case <-time.After(that.timeout):
		return "", ErrMqttRequestTimeout
	}
}

// NewKeygen NewKeygen
func NewKeygen(timeout time.Duration, topic string) (KeygenCenter, error) {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	if topic == "" {
		return nil, ErrInvalidArgument
	}

	return &keygenCenter{
		timeout: timeout,
		topic:   topic,
	}, nil
}
