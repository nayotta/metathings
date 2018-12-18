package metathingsdevicecloudmqttbridge

import (
	"context"
	"net/url"

	emitter "github.com/emitter-io/go"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	log "github.com/sirupsen/logrus"
)

// MqttBridge MqttBridge
type MqttBridge interface {
	GetRootKey() (string, error)
	GetHost() (*url.URL, error)

	InitMqttBridge() error
	KeyGen(context.Context, *pb.GenKeyRequest) (*pb.GenKeyResponse, error)
	HeartBeatSelect()
	UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error)
	StreamCall(stm pb.DeviceCloudService_StreamCallServer) error
}

type mqttBridge struct {
	host         *url.URL
	rootKey      string
	upKey        string
	downKey      string
	statusKey    string
	configClient emitter.Emitter
	logger       log.FieldLogger
}

func (that *mqttBridge) GetRootKey() (string, error) {
	if that.rootKey == "" {
		return "", ErrInvalidArgument
	}
	return that.rootKey, nil
}

func (that *mqttBridge) GetHost() (*url.URL, error) {
	if that.host == nil {
		return nil, ErrInvalidArgument
	}
	return that.host, nil
}

// NewMqttBridge NewMqttBridge
func NewMqttBridge(args []interface{}) (MqttBridge, error) {
	var ok bool
	var err error
	var key string
	var val interface{}
	var host *url.URL
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

			host, err = url.Parse(urlStr)
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
	return &mqttBridge{
		host:    host,
		rootKey: rootKey,
		logger:  logger,
	}, nil
}
