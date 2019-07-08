package metathings_component

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	MQTT_UPSTREAM   = "upstream"
	MQTT_DOWNSTREAM = "downstream"
)

type MqttModuleProxyOption struct {
	Module struct {
		Id string
	}
	Session struct {
		Id int64
	}
	Config struct {
		UnaryCallTimeout           time.Duration
		StreamCallConfigAckTimeout time.Duration
		MQTTConnectTimeout         time.Duration
		MQTTDisconnectTimeout      time.Duration
	}
	MQTT struct {
		Address  string
		Username string
		Password string
		ClientId string
		QoS      byte
	}
}

type MqttModuleProxy struct {
	opt        *MqttModuleProxyOption
	c          mqtt.Client
	sess_chans map[int64]chan *pb.UpStreamFrame
	mtx        sync.Mutex
	logger     log.FieldLogger
}

func (p *MqttModuleProxy) get_logger() log.FieldLogger {
	return p.logger
}

func (p *MqttModuleProxy) get_client() (mqtt.Client, error) {
	var err error

	if p.c == nil {
		p.c, err = p.new_client()
		if err != nil {
			return nil, err
		}
	}

	return p.c, nil
}

func (p *MqttModuleProxy) new_client() (mqtt.Client, error) {
	var err error
	var ok bool

	topic := p.mqtt_topic(p.opt.Module.Id, "+", MQTT_UPSTREAM)
	errs := make(chan error, 1)
	opts := mqtt.NewClientOptions().
		SetUsername(p.opt.MQTT.Username).
		SetPassword(p.opt.MQTT.Password).
		AddBroker(p.opt.MQTT.Address).
		SetClientID(p.opt.MQTT.ClientId).
		SetCleanSession(true).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}).
		SetOnConnectHandler(func(c mqtt.Client) {
			if tkn := c.Subscribe(topic, p.opt.MQTT.QoS, p.handle_message); tkn.Wait() && tkn.Error() != nil {
				errs <- tkn.Error()
			} else {
				close(errs)
			}
		})

	cli := mqtt.NewClient(opts)
	if tkn := cli.Connect(); tkn.Wait() && tkn.Error() != nil {
		return nil, tkn.Error()
	}

	select {
	case err, ok = <-errs:
		if ok {
			return nil, err
		}
		return cli, nil
	case <-time.After(p.opt.Config.MQTTConnectTimeout):
		return nil, ErrStartTimeout
	}
}

func (p *MqttModuleProxy) extra_session_from_topic(topic string) int64 {
	var sess int64

	if n, err := fmt.Sscanf(topic, "mt/modules/"+p.opt.Module.Id+"/sessions/%d/upstream", &sess); err != nil || n != 1 {
		return -1
	}

	return sess
}

func (p *MqttModuleProxy) handle_message(c mqtt.Client, m mqtt.Message) {
	if sess := p.extra_session_from_topic(m.Topic()); sess != -1 {
		var res pb.UpStreamFrame
		if err := proto.Unmarshal(m.Payload(), &res); err != nil {
			p.get_logger().WithError(err).Warningf("failed to unmarshal UpStreamFrame")
			return

		}

		p.mtx.Lock()
		defer p.mtx.Unlock()

		if ch, ok := p.sess_chans[sess]; ok {
			ch <- &res
		} else {
			p.get_logger().Warningf("receive message with unscribed session, drop it")
		}
	}
}

func (p *MqttModuleProxy) subscribe_session(sess int64) (chan *pb.UpStreamFrame, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if _, ok := p.sess_chans[sess]; ok {
		return nil, ErrSubscribedSession
	}

	ch := make(chan *pb.UpStreamFrame)
	p.sess_chans[sess] = ch

	return ch, nil
}

func (p *MqttModuleProxy) unsubscribe_session(sess int64) error {
	var ch chan *pb.UpStreamFrame
	var ok bool

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if ch, ok = p.sess_chans[sess]; !ok {
		return ErrUnsubscribedSession
	}

	close(ch)
	delete(p.sess_chans, sess)

	return nil
}

func (p *MqttModuleProxy) publish_message(sess int64, msg *pb.DownStreamFrame) error {
	cli, err := p.get_client()
	if err != nil {
		return err
	}

	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	topic := p.mqtt_topic(p.opt.Module.Id, sess, MQTT_DOWNSTREAM)
	if token := cli.Publish(topic, p.opt.MQTT.QoS, false, buf); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (p *MqttModuleProxy) startup_session() int32 {
	return session_helper.GetStartupSession(p.opt.Session.Id)
}

func (p *MqttModuleProxy) generate_temp_session() int64 {
	return session_helper.NewSession(p.startup_session(), session_helper.GenerateTempSession())
}

func (p *MqttModuleProxy) generate_minor_session() int64 {
	return session_helper.NewSession(p.startup_session(), session_helper.GenerateMinorSession())
}

func (p *MqttModuleProxy) mqtt_topic(mdl_id string, sess interface{}, dir string) string {
	var s string
	switch sess.(type) {
	case int64:
		s = fmt.Sprintf("%v", sess.(int64))
	case string:
		s = sess.(string)
	}

	return fmt.Sprintf("mt/modules/%v/sessions/%v/%v", mdl_id, s, dir)
}

func (p *MqttModuleProxy) UnaryCall(ctx context.Context, method string, value *any.Any) (*any.Any, error) {
	temp_sess := p.generate_temp_session()

	req := &pb.DownStreamFrame{
		Kind: pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.DownStreamFrame_UnaryCall{
			UnaryCall: &pb.OpUnaryCallValue{
				Session: &wrappers.Int64Value{Value: temp_sess},
				Method:  &wrappers.StringValue{Value: method},
				Value:   value,
			},
		},
	}

	err := p.publish_message(p.opt.Session.Id, req)
	if err != nil {
		return nil, err
	}

	ch, err := p.subscribe_session(temp_sess)
	if err != nil {
		return nil, err
	}
	defer p.unsubscribe_session(temp_sess)

	var res *pb.UpStreamFrame
	select {
	case res = <-ch:
	case <-time.After(p.opt.Config.UnaryCallTimeout):
		return nil, ErrUnaryCallTimeout
	}

	uc := res.GetUnaryCall()
	if uc == nil {
		return nil, ErrUnexceptedResponse
	}

	return uc.GetValue(), nil
}

func (p *MqttModuleProxy) init_downstream(sess int64, ds_recv_ch chan *pb.UpStreamFrame, method string) error {
	ack := rand.Int63()
	cfg_req := &pb.DownStreamFrame{
		Kind: pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.DownStreamFrame_StreamCall{
			StreamCall: &pb.OpStreamCallValue{
				Union: &pb.OpStreamCallValue_Config{
					Config: &pb.OpStreamCallConfig{
						Session: &wrappers.Int64Value{Value: sess},
						Method:  &wrappers.StringValue{Value: method},
						Ack:     &wrappers.Int64Value{Value: ack},
					},
				},
			},
		},
	}

	err := p.publish_message(p.opt.Session.Id, cfg_req)
	if err != nil {
		return err
	}

	var cfg_res *pb.UpStreamFrame
	select {
	case cfg_res = <-ds_recv_ch:
	case <-time.After(p.opt.Config.StreamCallConfigAckTimeout):
		return ErrStreamCallConfigAckTimeout
	}

	var res_ack int64
	if stm_res := cfg_res.GetStreamCall(); stm_res != nil {
		if ack_res := stm_res.GetAck(); ack_res != nil {
			res_ack = ack_res.GetValue()
		}
	}

	if res_ack != ack {
		return ErrStreamCallConfig
	}

	return nil
}

func (p *MqttModuleProxy) StreamCall(ctx context.Context, method string, upstm ModuleProxyStream) error {
	minor_sess := p.generate_minor_session()
	ch, err := p.subscribe_session(minor_sess)
	if err != nil {
		return err
	}
	defer p.unsubscribe_session(minor_sess)

	if err = p.init_downstream(minor_sess, ch, method); err != nil {
		return err
	}
	p.get_logger().Debugf("downstream initialized")

	north2south_wait := make(chan struct{}, 1)
	south2north_wait := make(chan struct{}, 1)
	go p.stm_north2south(upstm, minor_sess, north2south_wait)
	go p.stm_south2north(upstm, minor_sess, ch, south2north_wait)

	p.get_logger().Debugf("stream call started")
	select {
	case <-north2south_wait:
	case <-south2north_wait:
	}

	p.get_logger().Debugf("stream call done")

	return nil
}

func (p *MqttModuleProxy) stm_north2south(north ModuleProxyStream, sess int64, wait chan struct{}) {
	// var val *any.Any
	// var err error

	// logger := p.logger.WithFields(log.Fields{
	// 	"dir": "NS",
	// })

	// defer close(wait)
	// for epoch := uint64(0); ; epoch++ {
	// 	logger := logger.WithFields(log.Fields{
	// 		"epoch": epoch,
	// 	})

	// 	if val, err = north.Recv(); err != nil {
	// 		logger.WithError(err).Debugf("failed to recv msg from north side stream")
	// 		return
	// 	}
	// 	logger.Debugf("recv msg from north side stream")

	// 	req := &pb.DownStreamFrame{}
	// }

	panic("unimplemented")
}

func (p *MqttModuleProxy) stm_south2north(north ModuleProxyStream, sess int64, downch chan *pb.UpStreamFrame, wait chan struct{}) {
	panic("unimplemented")
}

func (p *MqttModuleProxy) Close() error {
	if p.c != nil {
		p.c.Disconnect(uint(p.opt.Config.MQTTDisconnectTimeout / time.Millisecond))
	}

	return nil
}

type MqttModuleProxyFactory struct{}

func (f *MqttModuleProxyFactory) NewModuleProxy(args ...interface{}) (ModuleProxy, error) {
	opt := &MqttModuleProxyOption{}
	opt.Config.UnaryCallTimeout = 5 * time.Second
	opt.Config.StreamCallConfigAckTimeout = 5 * time.Second
	opt.Config.MQTTConnectTimeout = 3 * time.Second
	opt.Config.MQTTDisconnectTimeout = 3 * time.Second
	opt.MQTT.ClientId = id_helper.NewId()
	opt.MQTT.QoS = byte(0)

	p := &MqttModuleProxy{
		opt:        opt,
		sess_chans: make(map[int64]chan *pb.UpStreamFrame),
	}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":        opt_helper.ToLogger(&p.logger),
		"module_id":     opt_helper.ToString(&p.opt.Module.Id),
		"session_id":    opt_helper.ToInt64(&p.opt.Session.Id),
		"mqtt_address":  opt_helper.ToString(&p.opt.MQTT.Address),
		"mqtt_username": opt_helper.ToString(&p.opt.MQTT.Username),
		"mqtt_password": opt_helper.ToString(&p.opt.MQTT.Password),
		"mqtt_clientid": opt_helper.ToString(&p.opt.MQTT.ClientId),
		"mqtt_qos":      opt_helper.ToByte(&p.opt.MQTT.QoS),
	})(args...); err != nil {
		return nil, err
	}

	return p, nil
}

func init() {
	register_module_proxy_factory("mqtt", new(MqttModuleProxyFactory))
}
