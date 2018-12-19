package metathingsdevicecloudmqttbridge

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	log "github.com/sirupsen/logrus"
)

// HeartBeatCenter HeartBeatCenter
type HeartBeatCenter struct {
	client            emitter.Emitter
	host              *url.URL
	timeout           time.Duration
	topicStatusSelect string
	downKey           string
	statusUpKey       string
	logger            log.FieldLogger
}

func (that *HeartBeatCenter) logE(componentID string, err error, text string) {
	that.logger.WithField("component_id", componentID).WithError(err).Errorf(text)
}

func (that *HeartBeatCenter) logD(componentID string, text string) {
	that.logger.WithField("component_id", componentID).Debugf(text)
}

func (that *HeartBeatCenter) heartBeatSelectProcess(componentID string) {
	var err error

	sessionIDStr, sessionID := newSessionID()

	selectTopic := componentID + "/statusup/" + sessionIDStr + "/"
	r := that.client.Subscribe(that.statusUpKey, selectTopic)
	if r.Wait() && r.Error() != nil {
		that.logE("", ErrMqttSubFailed, "mqtt sub failed")
		return
	}

	err = that.pubHeartBeatSelectResponse(componentID, sessionID)
	if err != nil {
		return
	}
}

func (that *HeartBeatCenter) heartBeatProcess(componentID string, sessionID int) {
	go that.pubHeartBeatResponse(componentID, sessionID)
}

func (that *HeartBeatCenter) heartBeatMsgCallback(client emitter.Emitter, msg emitter.Message) {
	//TODO(zh) call deviced heartbeat

	strs := strings.Split(msg.Topic(), "/")
	switch len(strs) {
	case 4:
		if strs[1] != "statusup" {
			that.logE("", ErrInvalidArgument, "unexcept topic get")
			return
		}

		if strs[3] != "" {
			that.logE("", ErrInvalidArgument, "unexcept topic get")
			return
		}

		componentID := strs[0]

		switch strs[2] {
		case "select":
			that.logD(strs[0], "get heartbeat select resquest")
			go that.heartBeatSelectProcess(componentID)
			return
		default:
			sessionID, err := strconv.Atoi(strs[2])
			if err != nil {
				that.logE("", ErrInvalidArgument, "unexcept topic sessionID get")
				return
			}
			that.logD(componentID, "get heartbeat request")
			go that.heartBeatProcess(componentID, sessionID)
			return
		}
	default:
		that.logE("", ErrInvalidArgument, "no seesion_id message get, should not appear")
		return
	}
}

func (that *HeartBeatCenter) heartBeatConnectCallback(_ emitter.Emitter) {
	var err error

	err = that.subStatusTopic()
	if err != nil {
		that.logE("", ErrMqttSubFailed, "heartbeat sub failed")
	}
}

func (that *HeartBeatCenter) heartBeatDisconnenctionCallback(client emitter.Emitter, err error) {
	that.logE("", ErrMqttDisconnectedError, "heartbeat connection disconnected")
}

func (that *HeartBeatCenter) connectMqtt() error {
	opt := emitter.NewClientOptions()
	opt.PingTimeout = 5 * time.Second
	opt.ConnectTimeout = 3 * time.Second
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnMessageHandler(that.heartBeatMsgCallback)
	opt.SetOnConnectHandler(that.heartBeatConnectCallback)
	opt.SetOnConnectionLostHandler(that.heartBeatDisconnenctionCallback)

	that.client = emitter.NewClient(opt)

	handle := that.client.Connect()
	if handle.Wait() && handle.Error() != nil {
		that.logE("", ErrMqttConnectFailed, "connect mqtt failed")
		return ErrMqttConnectFailed
	}

	return nil
}

func (that *HeartBeatCenter) subStatusTopic() error {
	r := that.client.Subscribe(that.statusUpKey, that.topicStatusSelect)
	if r.Wait() && r.Error() != nil {
		that.logE("", ErrMqttSubFailed, "mqtt sub failed")
		return ErrMqttSubFailed
	}

	return nil
}

func (that *HeartBeatCenter) pubHeartBeatSelectResponse(componentID string, sessionID int) error {
	componentTopic := componentID + "/down/"

	resMsg := &pb.MqttDeviceRequest{
		SessionId: (int32)(sessionID),
		Payload: &pb.MqttDeviceRequest_HeartbeatSelectRes{
			HeartbeatSelectRes: &pb.HeartbeatSelectRes{
				Status: 0,
			},
		},
	}

	res, err := proto.Marshal(resMsg)
	if err != nil {
		that.logE("", err, "mqtt marshal failed")
		return err
	}

	r := that.client.Publish(that.downKey, componentTopic, res)
	if r.Wait() && r.Error() != nil {
		that.logE(componentID, ErrMqttPubFailed, "heartbeat select pub failed")
		return ErrMqttPubFailed
	}

	that.logD(componentID, "send heartbeat select response")

	return nil
}

func (that *HeartBeatCenter) pubHeartBeatResponse(componentID string, sessionID int) error {
	//sessionIDStr := strconv.Itoa(sessionID)
	componentTopic := componentID + "/down/"

	// TODO(zh) call deviced heartbeat not response direct

	resMsg := &pb.MqttDeviceRequest{
		SessionId: (int32)(sessionID),
		Payload: &pb.MqttDeviceRequest_HeartbeatRes{
			HeartbeatRes: &pb.HeartbeatRes{
				Status: 0,
			},
		},
	}

	res, err := proto.Marshal(resMsg)
	if err != nil {
		that.logE("", err, "mqtt marshal failed")
		return err
	}

	r := that.client.Publish(that.downKey, componentTopic, res)
	if r.Wait() && r.Error() != nil {
		that.logE(componentID, ErrMqttPubFailed, "mqtt pub failed")
		return ErrMqttPubFailed
	}

	that.logD(componentID, "send heartbeat response")

	return nil
}

// HeartBeatLoop HeartBeatLoop
func (that *HeartBeatCenter) HeartBeatLoop() error {
	var err error

	err = that.connectMqtt()
	if err != nil {
		return err
	}

	return nil
}

// NewHeartBeatCenter NewHeartBeatCenter
func NewHeartBeatCenter(mqttBr *mqttBridge) (*HeartBeatCenter, error) {
	timeout := 10 * time.Second
	topicStatusSelect := "+/statusup/select/"

	return &HeartBeatCenter{
		host:              mqttBr.host,
		timeout:           timeout,
		topicStatusSelect: topicStatusSelect,
		statusUpKey:       mqttBr.statusUpKey,
		downKey:           mqttBr.downKey,
		logger:            mqttBr.logger,
	}, nil
}

func (that *mqttBridge) InitHeartBeatLoop() error {
	var err error

	hearBeatClient, err := NewHeartBeatCenter(that)
	if err != nil {
		that.logger.WithError(err).Errorf("InitHeartBeatLoop error")
		return err
	}

	err = hearBeatClient.HeartBeatLoop()
	if err != nil {
		that.logger.WithError(err).Errorf("HeartBeatLoop error")
		return err
	}

	return nil
}
