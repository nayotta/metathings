package metathingsdevicecloudmqttbridge

import (
	"context"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	log "github.com/sirupsen/logrus"
)

type keygenCenter struct {
	client     emitter.Emitter
	host       *url.URL
	keygenChan chan error
	retKey     string
	timeout    time.Duration

	deviceID string
	topic    string
	rootkey  string

	logger log.FieldLogger
}

func (that *keygenCenter) logE(err error, text string) {
	that.logger.WithField("device_id", that.deviceID).WithError(err).Errorf(text)
}

func (that *keygenCenter) logD(text string) {
	that.logger.WithField("device_id", that.deviceID).Debugf(text)
}

func (that *keygenCenter) keygenCallback(_ emitter.Emitter, msg emitter.KeyGenResponse) {
	if msg.Status != 200 {
		that.logE(ErrUnexpectedResponse, "keygen ret code not 200")
		that.keygenChan <- ErrUnexpectedResponse
	}

	if msg.Channel != that.topic {
		that.logE(ErrMqttKeygenFailed, "keygen topic not match")
		that.keygenChan <- ErrMqttKeygenFailed
	}

	that.retKey = msg.Key
	that.client.Disconnect(0)
	that.keygenChan <- nil
}

func (that *keygenCenter) keygenDisconnectCallback(client emitter.Emitter, err error) {
	that.logE(ErrMqttDisconnectedError, "mqtt broker unexcept disconnect")
	that.keygenChan <- ErrMqttDisconnectedError
}

func (that *keygenCenter) connectMqtt() error {
	opt := emitter.NewClientOptions()
	opt.PingTimeout = 5 * time.Second
	opt.ConnectTimeout = 3 * time.Second
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnKeyGenHandler(that.keygenCallback)
	opt.SetOnConnectionLostHandler(that.keygenDisconnectCallback)

	that.client = emitter.NewClient(opt)

	handle := that.client.Connect()
	if handle.Wait() && handle.Error() != nil {
		that.logE(ErrMqttConnectFailed, "connect mqtt failed")
		return ErrMqttConnectFailed
	}

	return nil
}

// KeyGen KeyGen
func (that *keygenCenter) KeyGen() (string, error) {
	var err error

	err = that.connectMqtt()
	if err != nil {
		return "", err
	}

	req := &emitter.KeyGenRequest{
		Key:     that.rootkey,
		Channel: that.topic,
		Type:    "rwslp",
		TTL:     0,
	}

	that.logD("keygen request")
	r := that.client.GenerateKey(req)
	if r.Wait() && r.Error() != nil {
		that.logE(ErrMqttKeygenFailed, "mqtt keygen error")
		return "", ErrMqttKeygenFailed
	}

	select {
	case err := <-that.keygenChan:
		if err != nil {
			return "", err
		}
		that.logD("keygen success")
		return that.retKey, nil
	case <-time.After(that.timeout):
		that.logE(ErrMqttRequestTimeout, "keygen timeout")
		return "", ErrMqttRequestTimeout
	}
}

func newKeygenCenter(mqttBr *mqttBridge, deviceID string) (*keygenCenter, error) {
	timeout := 10 * time.Second

	keygenChan := make(chan error)

	if deviceID == "" {
		mqttBr.logger.WithField("device_id", deviceID).WithError(ErrInvalidArgument).Errorf("device_id null")
		return nil, ErrInvalidArgument
	}
	topic := deviceID + "/#/"

	return &keygenCenter{
		host:       mqttBr.host,
		keygenChan: keygenChan,
		timeout:    timeout,
		deviceID:   deviceID,
		topic:      topic,
		rootkey:    mqttBr.rootKey,
		logger:     mqttBr.logger,
	}, nil
}

func (that *mqttBridge) KeyGen(ctx context.Context, req *pb.GenKeyRequest) (*pb.GenKeyResponse, error) {
	var err error

	cpID := req.GetDeviceId().GetValue()

	keygenClient, err := newKeygenCenter(that, cpID)
	if err != nil {
		return nil, err
	}

	retKey, err := keygenClient.KeyGen()
	if err != nil {
		return nil, err
	}

	return &pb.GenKeyResponse{
		Key: retKey,
	}, nil
}
