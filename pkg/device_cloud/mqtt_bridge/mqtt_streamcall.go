package metathingsdevicecloudmqttbridge

import (
	"fmt"
	"net/url"
	"time"

	emitter "github.com/emitter-io/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/nayotta/metathings/pkg/proto/device_cloud"
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
}

func (that *streamCallCenter) streamCallMsgCallback(client emitter.Emitter, msg emitter.Message) {
	switch msg.Topic() {
	case that.topicSession:
		that.streamMsgChan <- msg.Payload()
		break
	case that.topicNotify:
		that.streamMsgChan <- msg.Payload()
		break
	case that.topicHeartBeat:
		that.heartbeat = time.Now()
		fmt.Println("heartbeat res at:", time.Now())
		that.streamConfigChan <- nil
		break
	default:
		that.streamCallRecvChan <- ErrUnexpectedResponse
		that.streamCallSendChan <- ErrUnexpectedResponse
	}
}

func (that *streamCallCenter) streamCallDisconnectCallback(client emitter.Emitter, err error) {
	that.streamCallRecvChan <- ErrMqttDisconnectedError
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
		return ErrMqttPubFailed
	}

	return nil
}

// StreamCallHeartBeatLoop heartbeat loop for stream(if any improve??)
func (that *streamCallCenter) StreamCallHeartBeatLoop() {
	go func() {

		defer fmt.Println("heartbeat loop send deinit")

		for {
			err := that.streamCallHeartBeat()
			fmt.Println("heartbeat req at:", time.Now())
			if err != nil {
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
		defer fmt.Println("heartbeat loop check deinit")
		for {
			if time.Now().Sub(that.heartbeat) >= that.heartbeatTimeout {
				that.streamCallRecvChan <- ErrMqttStreamCallHearBeatTimeoutError
				that.streamCallSendChan <- ErrMqttStreamCallHearBeatTimeoutError
				fmt.Println("heartbeat check error", time.Now())
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
		return ErrMqttSubFailed
	}

	// sub notify channel
	r = that.client.Subscribe(that.upKey, that.topicNotify)
	if r.Wait() && r.Error() != nil {
		return ErrMqttSubFailed
	}

	// sub heartbeat channel
	r = that.client.Subscribe(that.statusKey, that.topicHeartBeat)
	if r.Wait() && r.Error() != nil {
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
		return err
	}
	reqPayload.SessionId = that.sessionID
	msg, err = proto.Marshal(&reqPayload)
	if err != nil {
		return err
	}

	r := that.client.Publish(that.downKey, that.topicDown, msg)
	if r.Wait() && r.Error() != nil {
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
	topicNotify := componentID + "/up/"
	topicDown := componentID + "/down/"
	topicHeartBeat := componentID + "/status/" + sessionIDStr + "/"

	downKey := mqttBr.downKey
	upKey := mqttBr.upKey
	statusKey := mqttBr.statusKey

	heartbeat := time.Now()
	heartbeatChan := make(chan interface{})
	heartbeatInterval := 10 * time.Second
	heartbeatTimeout := 30 * time.Second

	return &streamCallCenter{
		host:               host,
		timeout:            timeout,
		sessionID:          (int32)(sessionID),
		streamCallRecvChan: streamcallRecvChan,
		streamCallSendChan: streamcallSendChan,
		streamMsgChan:      streamMsgChan,
		streamConfigChan:   streamConfigChan,
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
	}
}

func (that *mqttBridge) StreamCall(stm pb.DeviceCloudService_StreamCallServer) error {
	var err error

	fmt.Println("device_cloud streamcall")
	defer fmt.Println("streamcall deinit")

	req, err := stm.Recv()
	if err != nil {
		return ErrMqttStreamCallConfigError
	}

	cpID := req.GetComponentId().GetValue()

	// init
	fmt.Printf("init ID:%v\n", cpID)
	streamCallClient := newStreamCallCenter(that, cpID)

	// connect mqtt
	fmt.Println("connect mqtt")
	err = streamCallClient.ConnectMqtt()
	if err != nil {
		return err
	}
	defer streamCallClient.client.Disconnect(0)

	// Sub
	fmt.Println("sub mqtt")
	err = streamCallClient.SubStreamTopic()
	if err != nil {
		return err
	}

	// heartbeat loop
	fmt.Println("heartbeat loop")
	streamCallClient.StreamCallHeartBeatLoop()
	defer func() {
		streamCallClient.heartbeatChan <- nil
		streamCallClient.streamCallRecvChan <- nil
		streamCallClient.streamCallSendChan <- nil
	}()

	// wait config ok signal
	fmt.Println("wait config ok signal")
	select {
	case <-streamCallClient.streamConfigChan:
		break
	case <-time.After(streamCallClient.timeout):
		fmt.Println("wait config ok signal timeout")
		streamCallClient.heartbeatChan <- nil
		return ErrMqttRequestTimeout
	}

	errs := make(chan error)

	fmt.Println("stream loop")
	// send loop
	go func() {
		var err error
		var req *pb.StreamCallRequest

		defer func() {
			fmt.Println("sent loop deinit")
			streamCallClient.heartbeatChan <- nil
			streamCallClient.streamCallRecvChan <- nil
			streamCallClient.streamCallSendChan <- nil
			errs <- err
		}()

		recvStm := NewAsyncDeviceCloudGrpc(stm)
		recvCh := recvStm.Recv()

		for {
			select {
			case req = <-recvCh:
				if req == nil {
					fmt.Println("send loop close err")
					return
				}
				fmt.Println("stream request msg", req)
				break
			case err := <-streamCallClient.streamCallRecvChan:
				fmt.Printf("send loop close, err:%v\n", err)
				return
			}

			msg := req.GetPayload().GetValue()

			err = streamCallClient.PubMsg(msg)
			if err != nil {
				fmt.Println("send loop close send error")
				return
			}
		}
	}()

	//recieve loop
	go func() {
		var err error

		defer func() {
			fmt.Println("recieve loop deinit")
			streamCallClient.heartbeatChan <- nil
			errs <- err
		}()

		for {
			select {
			case <-streamCallClient.streamConfigChan:
				continue
			case msg := <-streamCallClient.streamMsgChan:
				fmt.Println("stream response msg")
				err = stm.Send(&pb.StreamCallResponse{
					Payload: &any.Any{
						Value: msg,
					},
				})
				if err != nil {
					fmt.Printf("pub error:%v\n", err)
					return
				}
			case err = <-streamCallClient.streamCallSendChan:
				fmt.Println("recieve loop close")
				return
			}
		}
	}()

	err = <-errs
	if err != nil {
		fmt.Printf("stream error:%v\n", err)
		return err
	}

	fmt.Println("stream loop done")

	return nil
}
