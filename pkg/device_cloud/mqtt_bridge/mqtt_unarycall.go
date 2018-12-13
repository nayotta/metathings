package metathingsdevicecloudmqttbridge

import (
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
)

type unaryCallCenter struct {
	timeout       time.Duration
	unaryCallChan chan error
	resMsg        []byte
	topic         string
}

func (that *unaryCallCenter) unaryCallMsgCallback(client emitter.Emitter, msg emitter.Message) {
	if msg.Topic() != that.topic {
		that.unaryCallChan <- ErrUnexpectedResponse
	}

	that.resMsg = msg.Payload()
	that.unaryCallChan <- nil
}

func newUnaryCallCenter(timeout time.Duration, unaryCallChan chan error, topic string) *unaryCallCenter {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	return &unaryCallCenter{
		timeout:       timeout,
		unaryCallChan: unaryCallChan,
		topic:         topic,
	}
}

func (that *mqttBridge) UnaryCall(req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var err error
	var unarycallChan chan error
	var reqPayload pb.MqttDeviceRequest
	unarycallChan = make(chan error)

	cpID := req.GetComponentId().GetValue()
	sessionIDStr, sessionID := newSessionID()

	topicUp := cpID + "/up/" + sessionIDStr + "/"
	topicDown := cpID + "/down/"
	unaryCallClient := newUnaryCallCenter(0, unarycallChan, topicUp)

	opt := emitter.NewClientOptions()
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnMessageHandler(unaryCallClient.unaryCallMsgCallback)

	client := emitter.NewClient(opt)
	handle := client.Connect()
	if handle.Wait() && handle.Error() != nil {
		return nil, ErrMqttConnectFailed
	}
	defer client.Disconnect(0)

	r := client.Subscribe(that.upKey, topicUp)
	if r.Wait() && r.Error() != nil {
		return nil, ErrMqttSubFailed
	}

	msg := req.GetPayload().GetValue()

	err = proto.Unmarshal(msg, &reqPayload)
	if err != nil {
		return nil, err
	}
	reqPayload.SessionId = (int64)(sessionID)
	msg, err = proto.Marshal(&reqPayload)
	if err != nil {
		return nil, err
	}

	r = client.Publish(that.downKey, topicDown, msg)
	if r.Wait() && r.Error() != nil {
		return nil, ErrMqttPubFailed
	}

	select {
	case err = <-unarycallChan:
		if err != nil {
			return nil, err
		}
		return &pb.UnaryCallResponse{
			Payload: &any.Any{
				Value: unaryCallClient.resMsg,
			},
		}, nil
	case <-time.After(unaryCallClient.timeout):
		return nil, ErrMqttRequestTimeout
	}
}
