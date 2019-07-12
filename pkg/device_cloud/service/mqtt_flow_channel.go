package metathings_device_cloud_service

import (
	"crypto/tls"
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	component_pb "github.com/nayotta/metathings/pkg/proto/component"
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
	Module struct {
		Id string
	}
	Channel struct {
		Session string
	}
}

type MQTTPushFrameToFlowChannel struct {
	opt    *MQTTPushFrameToFlowChannelOption
	cli    mqtt.Client
	logger log.FieldLogger

	init_client_once sync.Once
	err              error
	frm_ch           chan *pb.OpFrame
}

func (fc *MQTTPushFrameToFlowChannel) get_logger() log.FieldLogger {
	return fc.logger
}

func (fc *MQTTPushFrameToFlowChannel) error() error {
	return fc.err
}

func (fc *MQTTPushFrameToFlowChannel) mqtt_topic(dir string) string {
	return fmt.Sprintf("mt/modules/%v/flow_channel/sessions/%v/%v", fc.opt.Module.Id, fc.opt.Channel.Session, dir)
}

func (fc *MQTTPushFrameToFlowChannel) init_client() {
	fc.init_client_once.Do(func() {
		opts := mqtt.NewClientOptions().
			AddBroker(fc.opt.MQTT.Address).
			SetUsername(fc.opt.MQTT.Username).
			SetPassword(fc.opt.MQTT.Password).
			SetClientID(fc.opt.MQTT.ClientId).
			SetCleanSession(true).
			SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}).
			SetOnConnectHandler(func(c mqtt.Client) {
				if tkn := c.Subscribe(fc.mqtt_topic("upstream"), fc.opt.MQTT.QoS, fc.handle_message); tkn.Wait() && tkn.Error() != nil {
					fc.err = tkn.Error()
				}
			})

		fc.cli = mqtt.NewClient(opts)
		if tkn := fc.cli.Connect(); tkn.Wait() && tkn.Error() != nil {
			fc.err = tkn.Error()
		}
	})
}

func (fc *MQTTPushFrameToFlowChannel) Channel() <-chan *pb.OpFrame {
	fc.frm_ch = make(chan *pb.OpFrame)
	return fc.frm_ch
}

func (fc *MQTTPushFrameToFlowChannel) handle_message(c mqtt.Client, m mqtt.Message) {
	var req component_pb.PushFrameToFlowRequest
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
	case *component_pb.PushFrameToFlowRequest_Ping_:
		res := &component_pb.PushFrameToFlowResponse{
			Id: id,
			Response: &component_pb.PushFrameToFlowResponse_Pong_{
				Pong: &component_pb.PushFrameToFlowResponse_Pong{},
			},
		}
		buf, err := proto.Marshal(res)
		if err != nil {
			fc.get_logger().WithError(err).Warningf("failed to pong response")
			return
		}
		tkn := fc.cli.Publish(fc.mqtt_topic("downstream"), fc.opt.MQTT.QoS, false, buf)
		if tkn.Wait() && tkn.Error() != nil {
			fc.get_logger().WithError(err).Warningf("failed to send pong response")
			return
		}
	case *component_pb.PushFrameToFlowRequest_Frame:
		fc.frm_ch <- req.GetFrame()
	}
}

func (fc *MQTTPushFrameToFlowChannel) Close() error {
	if tkn := fc.cli.Unsubscribe(fc.mqtt_topic("upstream")); tkn.Wait() && tkn.Error() != nil {
		return tkn.Error()
	}

	close(fc.frm_ch)

	return nil
}

type MQTTPushFrameToFlowChannelFactory struct{}

func (f *MQTTPushFrameToFlowChannelFactory) New(args ...interface{}) (PushFrameToFlowChannel, error) {
	var opt MQTTPushFrameToFlowChannelOption
	opt.MQTT.ClientId = id_helper.NewId()
	opt.MQTT.QoS = byte(0)

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"mqtt_address":    opt_helper.ToString(&opt.MQTT.Address),
		"mqtt_username":   opt_helper.ToString(&opt.MQTT.Username),
		"mqtt_password":   opt_helper.ToString(&opt.MQTT.Password),
		"mqtt_clientid":   opt_helper.ToString(&opt.MQTT.ClientId),
		"mqtt_qos":        opt_helper.ToByte(&opt.MQTT.QoS),
		"module_id":       opt_helper.ToString(&opt.Module.Id),
		"channel_session": opt_helper.ToString(&opt.Channel.Session),
	})(args...); err != nil {
		return nil, err
	}

	panic("unimplemented")
}

func init() {
	register_push_frame_to_flow_channel_factory("mqtt", new(MQTTPushFrameToFlowChannelFactory))
}
