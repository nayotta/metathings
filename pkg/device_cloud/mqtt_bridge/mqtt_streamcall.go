package metathingsdevicecloudmqttbridge

import (
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
	log "github.com/sirupsen/logrus"
)

type streamCallCenter struct {
	client             emitter.Emitter
	handle             emitter.Token
	host               *url.URL
	sessionID          int32
	streamCallSendChan chan error
	streamCallRecvChan chan error
	streamMsgChan      chan []byte
	streamConfigChan   chan interface{}
	timeout            time.Duration

	componentID    string
	topicSession   string
	topicHeartBeat string
	topicNotify    string
	topicDown      string

	downKey   string
	upKey     string
	statusKey string

	heartbeatChan     chan interface{}
	heartbeat         time.Time
	heartbeatInterval time.Duration
	heartbeatTimeout  time.Duration

	logger log.FieldLogger
}

func (that *streamCallCenter) streamCallMsgCallback(client emitter.Emitter, msg emitter.Message) {
	switch msg.Topic() {
	case that.topicSession:
		that.streamMsgChan <- msg.Payload()
		that.logger.WithField("component_id", that.componentID).Debugf("mqtt session response recv")
		break
	case that.topicNotify:
		that.streamMsgChan <- msg.Payload()
		that.logger.WithField("component_id", that.componentID).Debugf("mqtt notify response recv")
		break
	case that.topicHeartBeat:
		that.heartbeat = time.Now()
		that.logger.WithField("component_id", that.componentID).Debugf("heatbeat recv")
		that.streamConfigChan <- nil
		break
	default:
		that.logger.WithField("component_id", that.componentID).WithError(ErrUnexpectedResponse).Errorf("unknown streamcall message:%v", msg.Topic())
		that.streamCallSendChan <- ErrUnexpectedResponse
	}
}

func (that *streamCallCenter) streamCallDisconnectCallback(client emitter.Emitter, err error) {
	that.logger.WithField("component_id", that.componentID).WithError(ErrMqttDisconnectedError).Errorf("mqtt disconnected")
	that.streamCallSendChan <- ErrMqttDisconnectedError
}

func (that *streamCallCenter) streamCallHeartBeat() error {
	var err error

	req := &pb.MqttDeviceRequest{
		SessionId: that.sessionID,
		Payload: &pb.MqttDeviceRequest_Heartbeat{
			Heartbeat: &pb.Heartbeat{
				Status: 0,
			},
		},
	}

	msg, err := proto.Marshal(req)
	if err != nil {
		return ErrInvalidArgument
	}

	r := that.client.Publish(that.downKey, that.topicDown, msg)
	if r.Wait() && r.Error() != nil {
		that.logger.WithField("component_id", that.componentID).WithError(ErrMqttPubFailed).Errorf("heatbeat pub failed")
		return ErrMqttPubFailed
	}

	that.logger.WithField("component_id", that.componentID).Debugf("heatbeat pub")

	return nil
}

// StreamCallHeartBeatLoop heartbeat loop for stream(if any improve??)
func (that *streamCallCenter) StreamCallHeartBeatLoop() {
	go func() {
		for {
			err := that.streamCallHeartBeat()
			if err != nil {
				that.logger.WithField("component_id", that.componentID).WithError(err).Errorf("heatbeat failed")
				return
			}
			select {
			case <-time.After(that.heartbeatInterval):
				continue
			case <-that.heartbeatChan:
				return
			}
		}
	}()

	go func() {
		for {
			if time.Now().Sub(that.heartbeat) >= that.heartbeatTimeout {
				that.logger.WithField("component_id", that.componentID).WithError(ErrMqttStreamCallHearBeatTimeoutError).Errorf("heatbeat timeout")
				that.streamCallSendChan <- ErrMqttStreamCallHearBeatTimeoutError
				return
			}
			select {
			case <-time.After(1 * time.Second):
				continue
			case <-that.heartbeatChan:
				return
			}
		}
	}()
}

// SubStreamTopic sub topic to recieve device msg
func (that *streamCallCenter) SubStreamTopic() error {
	// sub session channel
	r := that.client.Subscribe(that.upKey, that.topicSession)
	if r.Wait() && r.Error() != nil {
		that.logger.WithField("component_id", that.componentID).WithError(ErrMqttSubFailed).Errorf("sub failed")
		return ErrMqttSubFailed
	}

	// sub notify channel
	r = that.client.Subscribe(that.upKey, that.topicNotify)
	if r.Wait() && r.Error() != nil {
		that.logger.WithField("component_id", that.componentID).WithError(ErrMqttSubFailed).Errorf("sub failed")
		return ErrMqttSubFailed
	}

	// sub heartbeat channel
	r = that.client.Subscribe(that.statusKey, that.topicHeartBeat)
	if r.Wait() && r.Error() != nil {
		that.logger.WithField("component_id", that.componentID).WithError(ErrMqttSubFailed).Errorf("sub failed")
		return ErrMqttSubFailed
	}

	return nil
}

// ConnectMqtt connect mqtt broker
func (that *streamCallCenter) ConnectMqtt() error {
	opt := emitter.NewClientOptions()
	opt.PingTimeout = 5 * time.Second
	opt.ConnectTimeout = 3 * time.Second
	opt.Servers = append(opt.Servers, that.host)
	opt.SetOnMessageHandler(that.streamCallMsgCallback)
	opt.SetOnConnectionLostHandler(that.streamCallDisconnectCallback)

	that.client = emitter.NewClient(opt)

	that.handle = that.client.Connect()
	if that.handle.Wait() && that.handle.Error() != nil {
		that.logger.WithField("component_id", that.componentID).WithError(ErrMqttConnectFailed).Errorf("connect mqtt failed")
		return ErrMqttConnectFailed
	}

	return nil
}

// PubMsg pubmsg to device
func (that *streamCallCenter) PubMsg(msg []byte) error {
	var err error
	var reqPayload pb.MqttDeviceRequest

	err = proto.Unmarshal(msg, &reqPayload)
	if err != nil {
		that.logger.WithField("component_id", that.componentID).WithError(err).Errorf("proto unmarshal failed")
		return err
	}
	reqPayload.SessionId = that.sessionID
	msg, err = proto.Marshal(&reqPayload)
	if err != nil {
		that.logger.WithField("component_id", that.componentID).WithError(err).Errorf("proto marshal failed")
		return err
	}

	r := that.client.Publish(that.downKey, that.topicDown, msg)
	if r.Wait() && r.Error() != nil {
		err = ErrMqttPubFailed
		that.logger.WithField("component_id", that.componentID).WithError(err).Errorf("mqtt pub failed")
		return err
	}

	return nil
}

// NewStreamCallCenter streamcall center init
func newStreamCallCenter(mqttBr *mqttBridge, componentID string) *streamCallCenter {
	timeout := 10 * time.Second
	sessionIDStr, sessionID := newSessionID()

	host := mqttBr.host

	streamcallRecvChan := make(chan error)
	streamcallSendChan := make(chan error)
	streamMsgChan := make(chan []byte)
	streamConfigChan := make(chan interface{})

	topicSession := componentID + "/up/" + sessionIDStr + "/"
	topicNotify := componentID + "/up/notify/"
	topicDown := componentID + "/down/"
	topicHeartBeat := componentID + "/status/" + sessionIDStr + "/"

	downKey := mqttBr.downKey
	upKey := mqttBr.upKey
	statusKey := mqttBr.statusKey

	heartbeat := time.Now()
	heartbeatChan := make(chan interface{})
	heartbeatInterval := 10 * time.Second
	heartbeatTimeout := 30 * time.Second

	loggerBr := mqttBr.logger

	return &streamCallCenter{
		host:               host,
		timeout:            timeout,
		sessionID:          (int32)(sessionID),
		streamCallRecvChan: streamcallRecvChan,
		streamCallSendChan: streamcallSendChan,
		streamMsgChan:      streamMsgChan,
		streamConfigChan:   streamConfigChan,
		componentID:        componentID,
		topicSession:       topicSession,
		topicNotify:        topicNotify,
		topicDown:          topicDown,
		topicHeartBeat:     topicHeartBeat,
		downKey:            downKey,
		upKey:              upKey,
		statusKey:          statusKey,
		heartbeat:          heartbeat,
		heartbeatChan:      heartbeatChan,
		heartbeatInterval:  heartbeatInterval,
		heartbeatTimeout:   heartbeatTimeout,
		logger:             loggerBr,
	}
}

func (that *mqttBridge) StreamCall(stm pb.DeviceCloudService_StreamCallServer) error {
	var err error

	req, err := stm.Recv()
	if err != nil {
		that.logger.WithError(err).Errorf("stream recv failed")
		return ErrMqttStreamCallConfigError
	}

	cpID := req.GetComponentId().GetValue()

	// init
	streamCallClient := newStreamCallCenter(that, cpID)

	// connect mqtt
	err = streamCallClient.ConnectMqtt()
	if err != nil {
		that.logger.WithField("component_id", cpID).WithError(err).Errorf("mqtt connect failed")
		return err
	}
	defer func() {
		that.logger.WithField("component_id", cpID).Debugf("mqtt client disconnect")
		streamCallClient.client.Disconnect(0)
	}()

	// Sub
	err = streamCallClient.SubStreamTopic()
	if err != nil {
		that.logger.WithField("component_id", cpID).WithError(err).Errorf("mqtt sub failed")
		return err
	}

	// heartbeat loop
	streamCallClient.StreamCallHeartBeatLoop()

	// wait config ok signal
	select {
	case <-streamCallClient.streamConfigChan:
		break
	case <-time.After(streamCallClient.timeout):
		that.logger.WithField("component_id", cpID).WithError(err).Errorf("stream config timeout")
		return ErrMqttRequestTimeout
	}

	errs := make(chan error)

	// send loop
	go func() {
		var err error
		var req *pb.StreamCallRequest

		defer func() {
			streamCallClient.heartbeatChan <- nil
			streamCallClient.streamCallRecvChan <- nil
			errs <- err
		}()

		recvStm := NewAsyncDeviceCloudGrpc(stm)
		recvCh := recvStm.Recv()

		for {
			select {
			case req = <-recvCh:
				if req == nil {
					err = ErrMqttStreamCallRecvError
					that.logger.WithField("component_id", cpID).WithError(err).Errorf("streamcall recv error")
					return
				}
				that.logger.WithField("component_id", cpID).Debugf("message to be send")
				break
			case err = <-streamCallClient.streamCallSendChan:
				return
			}

			msg := req.GetPayload().GetValue()

			err = streamCallClient.PubMsg(msg)
			if err != nil {
				that.logger.WithField("component_id", cpID).WithError(err).Errorf("mqtt pub failed")
				return
			}
		}
	}()

	//recieve loop
	go func() {
		var err error

		defer func() {
			streamCallClient.heartbeatChan <- nil
			streamCallClient.streamCallSendChan <- nil
			errs <- err
		}()

		for {
			select {
			case <-streamCallClient.streamConfigChan:
				continue
			case msg := <-streamCallClient.streamMsgChan:
				err = stm.Send(&pb.StreamCallResponse{
					Payload: &any.Any{
						Value: msg,
					},
				})
				if err != nil {
					that.logger.WithField("component_id", cpID).WithError(err).Errorf("streamcall send failed")
					return
				}
			case err = <-streamCallClient.streamCallRecvChan:
				return
			}
		}
	}()

	err = <-errs
	if err != nil {
		return err
	}

	return nil
}
