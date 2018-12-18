package metathingsdevicecloudmqttbridge

import (
	"context"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	log "github.com/sirupsen/logrus"
)

type unaryCallCenter struct {
	client        emitter.Emitter
	host          *url.URL
	sessionID     int
	unaryCallChan chan error
	resMsg        []byte
	timeout       time.Duration

	componentID  string
	topicSession string
	topicDown    string

	downKey string
	upKey   string

	logger log.FieldLogger
}

func (that *unaryCallCenter) logE(err error, text string) {
	that.logger.WithField("component_id", that.componentID).WithField("session_id", that.sessionID).WithError(err).Errorf(text)
}

func (that *unaryCallCenter) logD(text string) {
	that.logger.WithField("component_id", that.componentID).WithField("session_id", that.sessionID).Debugf(text)
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

func (that *unaryCallCenter) PubMsg(msg []byte) error {
	var err error
	var reqPayload pb.MqttDeviceRequest

	// connect mqtt broker
	err = that.connectMqtt()
	if err != nil {
		return err
	}

	// sub topic
	err = that.subUnaryCallTopic()
	if err != nil {
		return err
	}

	// msg add seesionid
	err = proto.Unmarshal(msg, &reqPayload)
	if err != nil {
		that.logE(err, "proto unmarshal failed")
		return err
	}

	reqPayload.SessionId = (int32)(that.sessionID)
	sendMsg, err := proto.Marshal(&reqPayload)
	if err != nil {
		return err
	}

	// pub msg
	r := that.client.Publish(that.downKey, that.topicDown, sendMsg)
	if r.Wait() && r.Error() != nil {
		that.logE(ErrMqttPubFailed, "mqtt pub failed")
		return ErrMqttPubFailed
	}

	that.logD("unarycall pub")

	return nil
}

func newUnaryCallCenter(mqttBr *mqttBridge, componentID string) (*unaryCallCenter, error) {
	unaryCallChan := make(chan error)
	timeout := 10 * time.Second

	sessionIDStr, sessionID := newSessionID()
	topicSession := componentID + "/up/" + sessionIDStr + "/"
	topicDown := componentID + "/down/"

	return &unaryCallCenter{
		host:          mqttBr.host,
		sessionID:     sessionID,
		unaryCallChan: unaryCallChan,
		timeout:       timeout,
		componentID:   componentID,
		topicSession:  topicSession,
		topicDown:     topicDown,
		downKey:       mqttBr.downKey,
		upKey:         mqttBr.upKey,
		logger:        mqttBr.logger,
	}, nil
}

func (that *mqttBridge) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var err error

	cpID := req.GetComponentId().GetValue()
	msg := req.GetPayload().GetValue()

	unaryCallClient, err := newUnaryCallCenter(that, cpID)
	if err != nil {
		return nil, err
	}

	err = unaryCallClient.PubMsg(msg)
	if err != nil {
		return nil, err
	}

	select {
	case err = <-unaryCallClient.unaryCallChan:
		if err != nil {
			return nil, err
		}

		return &pb.UnaryCallResponse{
			Payload: &any.Any{
				Value: unaryCallClient.resMsg,
			},
		}, nil
	case <-time.After(unaryCallClient.timeout):
		that.logger.WithField("component_id", cpID).WithField("session_id", unaryCallClient.sessionID).WithError(ErrMqttRequestTimeout).Errorf("wait unarycall response timeout")
		return nil, ErrMqttRequestTimeout
	}
}
