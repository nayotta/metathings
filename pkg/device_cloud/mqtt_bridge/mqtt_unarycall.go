package metathingsdevicecloudmqttbridge

import (
	"context"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	deviced_storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

type unaryCallCenter struct {
	client        emitter.Emitter
	host          *url.URL
	sessionID     int
	unaryCallChan chan error
	resMsg        []byte
	timeout       time.Duration

	deviceID     string
	topicSession string
	topicDown    string

	downKey string
	upKey   string

	logger log.FieldLogger
}

func (that *unaryCallCenter) logE(err error, text string) {
	that.logger.WithField("device_id", that.deviceID).WithField("session_id", that.sessionID).WithError(err).Errorf(text)
}

func (that *unaryCallCenter) logD(text string) {
	that.logger.WithField("device_id", that.deviceID).WithField("session_id", that.sessionID).Debugf(text)
}

func (that *unaryCallCenter) unaryCallMsgCallback(client emitter.Emitter, msg emitter.Message) {
	if msg.Topic() != that.topicSession {
		that.logE(ErrUnexpectedResponse, "unexcept message get")
		that.unaryCallChan <- ErrUnexpectedResponse
	}

	that.logD("response message get")

	that.resMsg = msg.Payload()
	that.unaryCallChan <- nil
	that.client.Disconnect(0)
}

func (that *unaryCallCenter) unaryCallDisconnectCallback(client emitter.Emitter, err error) {
	that.logE(ErrMqttDisconnectedError, "mqtt broker unexcept disconnect")
	that.unaryCallChan <- ErrMqttDisconnectedError
}

func (that *unaryCallCenter) connectMqtt() error {
	opt := emitter.NewClientOptions()
	opt.PingTimeout = 5 * time.Second
	opt.ConnectTimeout = 3 * time.Second
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnMessageHandler(that.unaryCallMsgCallback)
	opt.SetOnConnectionLostHandler(that.unaryCallDisconnectCallback)

	that.client = emitter.NewClient(opt)

	handle := that.client.Connect()
	if handle.Wait() && handle.Error() != nil {
		that.logE(ErrMqttConnectFailed, "connect mqtt failed")
		return ErrMqttConnectFailed
	}

	return nil
}

func (that *unaryCallCenter) subUnaryCallTopic() error {
	r := that.client.Subscribe(that.upKey, that.topicSession)
	if r.Wait() && r.Error() != nil {
		that.logE(ErrMqttSubFailed, "mqtt sub failed")
		return ErrMqttSubFailed
	}

	return nil
}

func (that *unaryCallCenter) PubMsgAndWaitRes(msg []byte) ([]byte, error) {
	var err error
	var reqPayload pb.MqttDeviceRequest

	// connect mqtt broker
	err = that.connectMqtt()
	if err != nil {
		return nil, err
	}

	// sub topic
	err = that.subUnaryCallTopic()
	if err != nil {
		return nil, err
	}

	// msg add seesionid
	err = proto.Unmarshal(msg, &reqPayload)
	if err != nil {
		that.logE(err, "proto unmarshal failed")
		return nil, err
	}

	reqPayload.SessionId = (int32)(that.sessionID)
	sendMsg, err := proto.Marshal(&reqPayload)
	if err != nil {
		return nil, err
	}

	// pub msg
	r := that.client.Publish(that.downKey, that.topicDown, sendMsg)
	if r.Wait() && r.Error() != nil {
		that.logE(ErrMqttPubFailed, "mqtt pub failed")
		return nil, ErrMqttPubFailed
	}

	that.logD("unarycall pub")

	select {
	case err = <-that.unaryCallChan:
		if err != nil {
			return nil, err
		}
		return that.resMsg, nil
	case <-time.After(that.timeout):
		that.logE(ErrMqttRequestTimeout, "wait unarycall response timeout")
		return nil, ErrMqttRequestTimeout
	}
}

func newUnaryCallCenter(mqttBr *mqttBridge, deviceID string) (*unaryCallCenter, error) {
	unaryCallChan := make(chan error)
	timeout := 10 * time.Second

	sessionIDStr, sessionID := newSessionID()
	topicSession := deviceID + "/up/" + sessionIDStr + "/"
	topicDown := deviceID + "/down/"

	return &unaryCallCenter{
		host:          mqttBr.host,
		sessionID:     sessionID,
		unaryCallChan: unaryCallChan,
		timeout:       timeout,
		deviceID:      deviceID,
		topicSession:  topicSession,
		topicDown:     topicDown,
		downKey:       mqttBr.downKey,
		upKey:         mqttBr.upKey,
		logger:        mqttBr.logger,
	}, nil
}

func (that *mqttBridge) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var err error

	devID := req.GetDeviceId().GetValue()
	msg := req.GetPayload().GetValue()

	unaryCallClient, err := newUnaryCallCenter(that, devID)
	if err != nil {
		return nil, err
	}

	resMsg, err := unaryCallClient.PubMsgAndWaitRes(msg)
	if err != nil {
		return nil, err
	}

	return &pb.UnaryCallResponse{
		Payload: &any.Any{
			Value: resMsg,
		},
	}, nil
}

func (that *mqttBridge) UnaryCallForDeviced(dev *deviced_storage.Device, req *deviced_pb.OpUnaryCallValue) (*deviced_pb.UnaryCallValue, error) {
	var err error

	msg := req.GetValue().GetValue()
	unaryCallClient, err := newUnaryCallCenter(that, *dev.Id)
	if err != nil {
		return nil, err
	}

	resMsg, err := unaryCallClient.PubMsgAndWaitRes(msg)
	if err != nil {
		return nil, err
	}

	return &deviced_pb.UnaryCallValue{
		Name:      req.GetName().GetValue(),
		Component: req.GetComponent().GetValue(),
		Method:    req.GetMethod().GetValue(),
		Value: &any.Any{
			Value: resMsg,
		},
	}, nil
}
