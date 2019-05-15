package metathingsdevicecloudmqttbridge

import (
	"context"
	"net/url"
	"strings"
	"time"

	emitter "github.com/emitter-io/go"
	json_pb "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
)

// PushFrameToFlowCenter PushFrameToFlowCenter
type PushFrameToFlowCenter struct {
	client  emitter.Emitter
	host    *url.URL
	timeout time.Duration
	downKey string
	upKey   string
	upTopic string
	cliFty  *client_helper.ClientFactory
	logger  log.FieldLogger
}

func (that *PushFrameToFlowCenter) logE(flowID string, err error, text string) {
	that.logger.WithField("flowID", flowID).WithError(err).Errorf(text)
}

func (that *PushFrameToFlowCenter) logD(flowID string, text string) {
	that.logger.WithField("flowID", flowID).Debugf(text)
}

func (that *PushFrameToFlowCenter) pushFrameToFlowProcess(deviceID string, flowID string, frame []byte) {
	go that.pubPushFrameToFlowResponse(deviceID, flowID, frame)
}

func (that *PushFrameToFlowCenter) pushFrameToFlowMsgCallback(client emitter.Emitter, msg emitter.Message) {
	strs := strings.Split(msg.Topic(), "/")
	switch len(strs) {
	case 5:
		if strs[0] != "flow" {
			that.logE("", ErrInvalidArgument, "unexcept topic 0 get")
			return
		}

		if strs[3] != "up" {
			that.logE("", ErrInvalidArgument, "unexcept topic 2 get")
			return
		}

		flowID := strs[2]
		deviceID := strs[1]
		go that.pushFrameToFlowProcess(deviceID, flowID, msg.Payload())
	default:
		that.logE("", ErrInvalidArgument, "unexcept topic size get")
		return
	}
}

func (that *PushFrameToFlowCenter) pushFrameToFlowConnectCallback(_ emitter.Emitter) {
	var err error

	err = that.subUpTopic()
	if err != nil {
		that.logE("", ErrMqttSubFailed, "pushFrameToFlow sub failed")
	}
}

func (that *PushFrameToFlowCenter) pushFrameToFlowDisconnenctionCallback(client emitter.Emitter, err error) {
	that.logE("", ErrMqttDisconnectedError, "pushFrameToFlow connection disconnected")
}

func (that *PushFrameToFlowCenter) connectMqtt() error {
	opt := emitter.NewClientOptions()
	opt.PingTimeout = 5 * time.Second
	opt.ConnectTimeout = 3 * time.Second
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnMessageHandler(that.pushFrameToFlowMsgCallback)
	opt.SetOnConnectHandler(that.pushFrameToFlowConnectCallback)
	opt.SetOnConnectionLostHandler(that.pushFrameToFlowDisconnenctionCallback)

	that.client = emitter.NewClient(opt)

	handle := that.client.Connect()
	if handle.Wait() && handle.Error() != nil {
		that.logE("", ErrMqttConnectFailed, "connect mqtt failed")
		return ErrMqttConnectFailed
	}

	return nil
}

func (that *PushFrameToFlowCenter) subUpTopic() error {
	r := that.client.Subscribe(that.upKey, that.upTopic)
	if r.Wait() && r.Error() != nil {
		that.logE("", ErrMqttSubFailed, "mqtt sub failed")
		return ErrMqttSubFailed
	}

	return nil
}

func (that *PushFrameToFlowCenter) pubPushFrameToFlowResponse(deviceID string, flowID string, frame []byte) error {
	var frameStruct _struct.Struct

	deviceTopic := "flow/" + deviceID + "/" + flowID + "/down/"

	err := json_pb.UnmarshalString(string(frame), &frameStruct)
	if err != nil {
		return err
	}

	now := pb_helper.Now()
	hbreq := &deviced_pb.MqttPushFrameToFlowRequest{
		DeviceId: &wrappers.StringValue{Value: deviceID},
		FlowId:   &wrappers.StringValue{Value: flowID},
		Frame: &deviced_pb.OpFrame{
			Ts:   &now,
			Data: &frameStruct,
		},
	}

	cli, cfn, err := that.cliFty.NewDevicedServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	resMsg, err := cli.MqttPushFrameToFlow(context.Background(), hbreq)
	if err != nil {
		return err
	}

	res, err := proto.Marshal(resMsg)
	if err != nil {
		that.logE("", err, "mqtt marshal failed")
		return err
	}

	r := that.client.Publish(that.downKey, deviceTopic, res)
	if r.Wait() && r.Error() != nil {
		that.logE(flowID, ErrMqttPubFailed, "mqtt pub failed")
		return ErrMqttPubFailed
	}

	that.logD(flowID, "send PushFrameToFlow response")

	return nil
}

// PushFrameToFlowLoop PushFrameToFlowLoop
func (that *PushFrameToFlowCenter) PushFrameToFlowLoop() error {
	var err error

	err = that.connectMqtt()
	if err != nil {
		return err
	}

	return nil
}

// NewPushFrameToFlowCenter NewPushFrameToFlowCenter
func NewPushFrameToFlowCenter(mqttBr *mqttBridge) (*PushFrameToFlowCenter, error) {
	timeout := 10 * time.Second
	upTopic := "flow/+/+/up/"

	return &PushFrameToFlowCenter{
		host:    mqttBr.host,
		timeout: timeout,
		downKey: mqttBr.downKey,
		upKey:   mqttBr.flowUpKey,
		upTopic: upTopic,
		cliFty:  mqttBr.cliFty,
		logger:  mqttBr.logger,
	}, nil
}

func (that *mqttBridge) InitPushFrameToFlowLoop() error {
	var err error

	that.pushFrame, err = NewPushFrameToFlowCenter(that)
	if err != nil {
		that.logger.WithError(err).Errorf("InitPushFrameToFlowLoop error")
		return err
	}

	err = that.pushFrame.PushFrameToFlowLoop()
	if err != nil {
		that.logger.WithError(err).Errorf("PushFrameToFlowLoop error")
		return err
	}

	return nil
}
