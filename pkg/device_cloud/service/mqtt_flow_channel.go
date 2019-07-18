package metathings_device_cloud_service

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type MQTTPushFrameToFlowChannelOption struct {
	MQTT struct {
		Address  string
		Username string
		Password string
		ClientId string
		QoS      byte
	}
	Device struct {
		Id string
	}
	Channel struct {
		Session string
		PushAck bool
	}
	Config struct {
		AliveInterval time.Duration
		PingTimeout   time.Duration
	}
}

type MQTTPushFrameToFlowChannel struct {
	opt    *MQTTPushFrameToFlowChannelOption
	cli    mqtt.Client
	logger log.FieldLogger

	err         error
	frm_ch      chan *pb.OpFrame
	frm_ch_once sync.Once
	ping_at     time.Time
}

func (fc *MQTTPushFrameToFlowChannel) get_logger() log.FieldLogger {
	return fc.logger
}

func (fc *MQTTPushFrameToFlowChannel) error() error {
	return fc.err
}

func (fc *MQTTPushFrameToFlowChannel) mqtt_topic(dir string) string {
	return fmt.Sprintf("mt/devices/%v/flow_channel/sessions/%v/%v", fc.opt.Device.Id, fc.opt.Channel.Session, dir)
}

func (fc *MQTTPushFrameToFlowChannel) init_client() error {
	opts := mqtt.NewClientOptions().
		AddBroker(fc.opt.MQTT.Address).
		SetUsername(fc.opt.MQTT.Username).
		SetPassword(fc.opt.MQTT.Password).
		SetClientID(fc.opt.MQTT.ClientId).
		SetCleanSession(true).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}).
		SetOnConnectHandler(func(c mqtt.Client) {
			sub_tpc := fc.mqtt_topic("upstream")

			if tkn := c.Subscribe(sub_tpc, fc.opt.MQTT.QoS, fc.handle_message); tkn.Wait() && tkn.Error() != nil {
				fc.err = tkn.Error()
			}
			fc.get_logger().WithField("topic", sub_tpc).Debugf("mqtt client subscribe topic")
		})

	fc.cli = mqtt.NewClient(opts)
	if tkn := fc.cli.Connect(); tkn.Wait() && tkn.Error() != nil {
		return tkn.Error()
	}
	fc.get_logger().WithFields(log.Fields{
		"broker":   fc.opt.MQTT.Address,
		"username": fc.opt.MQTT.Username,
		"password": fc.opt.MQTT.Password,
		"cilentid": fc.opt.MQTT.ClientId,
	}).Debugf("mqtt client connected")

	go fc.alive_loop()

	return nil
}

func (fc *MQTTPushFrameToFlowChannel) alive_loop() {
	for {
		time.Sleep(fc.opt.Config.AliveInterval)
		if time.Since(fc.ping_at) > fc.opt.Config.PingTimeout {
			fc.get_logger().Warningf("mqtt push frame to flow channel is dead")
			fc.Close()
			return
		}
	}
}

func (fc *MQTTPushFrameToFlowChannel) Channel() <-chan *pb.OpFrame {
	fc.frm_ch_once.Do(func() {
		fc.frm_ch = make(chan *pb.OpFrame)
	})
	return fc.frm_ch
}

func (fc *MQTTPushFrameToFlowChannel) handle_message(c mqtt.Client, m mqtt.Message) {
	var req device_pb.PushFrameToFlowRequest
	err := proto.Unmarshal(m.Payload(), &req)
	if err != nil {
		fc.get_logger().Warningf("failed to unmarshal push frame to flow request")
		return
	}

	if req.Id == nil || req.GetId().GetValue() == "" {
		fc.get_logger().Warningf("bad push frame to flow request id")
		return
	}
	id := req.GetId().GetValue()

	switch req.Request.(type) {
	case *device_pb.PushFrameToFlowRequest_Ping_:
		res := &device_pb.PushFrameToFlowResponse{
			Id: id,
			Response: &device_pb.PushFrameToFlowResponse_Pong_{
				Pong: &device_pb.PushFrameToFlowResponse_Pong{},
			},
		}

		err := fc.send_response_message(res)
		if err != nil {
			fc.get_logger().WithError(err).Warningf("failed to send pong response")
			return
		}

		fc.ping_at = time.Now()
	case *device_pb.PushFrameToFlowRequest_Frame:
		if fc.frm_ch != nil {
			fc.frm_ch <- req.GetFrame()
			if fc.opt.Channel.PushAck {
				res := &device_pb.PushFrameToFlowResponse{
					Id: req.GetId().GetValue(),
					Response: &device_pb.PushFrameToFlowResponse_Ack_{
						Ack: &device_pb.PushFrameToFlowResponse_Ack{},
					},
				}

				err := fc.send_response_message(res)
				if err != nil {
					fc.get_logger().WithError(err).Warningf("failed to send ack response")
					return
				}
			}
		}
	}
}

func (fc *MQTTPushFrameToFlowChannel) send_response_message(res *device_pb.PushFrameToFlowResponse) error {
	buf, err := proto.Marshal(res)
	if err != nil {
		return err
	}

	tkn := fc.cli.Publish(fc.mqtt_topic("downstream"), fc.opt.MQTT.QoS, false, buf)
	if tkn.Wait() && tkn.Error() != nil {
		return tkn.Error()
	}

	return nil
}

func (fc *MQTTPushFrameToFlowChannel) Close() error {
	if tkn := fc.cli.Unsubscribe(fc.mqtt_topic("upstream")); tkn.Wait() && tkn.Error() != nil {
		return tkn.Error()
	}

	if fc.frm_ch != nil {
		close(fc.frm_ch)
		fc.frm_ch = nil
	}

	return nil
}

type MQTTPushFrameToFlowChannelFactory struct{}

func (f *MQTTPushFrameToFlowChannelFactory) New(args ...interface{}) (PushFrameToFlowChannel, error) {
	var logger log.FieldLogger
	var opt MQTTPushFrameToFlowChannelOption
	opt.MQTT.ClientId = id_helper.NewId()
	opt.MQTT.QoS = byte(0)
	opt.Config.AliveInterval = 11 * time.Second
	opt.Config.PingTimeout = 131 * time.Second

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"mqtt_address":    opt_helper.ToString(&opt.MQTT.Address),
		"mqtt_username":   opt_helper.ToString(&opt.MQTT.Username),
		"mqtt_password":   opt_helper.ToString(&opt.MQTT.Password),
		"mqtt_clientid":   opt_helper.ToString(&opt.MQTT.ClientId),
		"mqtt_qos":        opt_helper.ToByte(&opt.MQTT.QoS),
		"device_id":       opt_helper.ToString(&opt.Device.Id),
		"channel_session": opt_helper.ToString(&opt.Channel.Session),
		"push_ack":        opt_helper.ToBool(&opt.Channel.PushAck),
		"logger":          opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	c := &MQTTPushFrameToFlowChannel{
		opt:     &opt,
		logger:  logger,
		ping_at: time.Now(),
	}
	if err := c.init_client(); err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	register_push_frame_to_flow_channel_factory("mqtt", new(MQTTPushFrameToFlowChannelFactory))
}
