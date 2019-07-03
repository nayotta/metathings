package metathings_component

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/goiiot/libmqtt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	pb "github.com/nayotta/metathings/pkg/proto/component"
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
	}
}

type MqttModuleProxy struct {
	opt        *MqttModuleProxyOption
	c          libmqtt.Client
	sess_chans map[int64]chan []byte
	mtx        sync.Mutex
	logger     log.FieldLogger
}

func (p *MqttModuleProxy) start() error {
	topic := fmt.Sprintf("mt/modules/%v/sessions/+/upstream", p.opt.Module.Id)

	p.c.Handle(topic, p.handle_message)
	p.c.Subscribe([]*libmqtt.Topic{
		{Name: topic},
	}...)

	return nil
}

func (p *MqttModuleProxy) extra_session_from_topic(topic string) int64 {
	var sess int64

	if n, err := fmt.Sscanf(topic, "mt/modules/"+p.opt.Module.Id+"/sessions/%d/upstream", &sess); err != nil || n != 1 {
		return -1
	}

	return sess
}

func (p *MqttModuleProxy) handle_message(topic string, qos libmqtt.QosLevel, msg []byte) {
	sess := p.extra_session_from_topic(topic)
	if sess != -1 {
		p.dispatch_message(sess, msg)
	}
}

func (p *MqttModuleProxy) dispatch_message(sess int64, msg []byte) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	ch, ok := p.sess_chans[sess]
	if ok {
		ch <- msg
	}
}

func (p *MqttModuleProxy) subscribe_session(sess int64) (chan []byte, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if _, ok := p.sess_chans[sess]; ok {
		return nil, ErrSubscribedSession
	}

	ch := make(chan []byte)
	p.sess_chans[sess] = ch

	return ch, nil
}

func (p *MqttModuleProxy) unsubscribe_session(sess int64) error {
	var ch chan []byte
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

func (p *MqttModuleProxy) publish_message(sess int64, msg []byte) error {
	p.c.Publish([]*libmqtt.PublishPacket{
		{TopicName: p.mqtt_topic(p.opt.Module.Id, sess), Payload: msg},
	}...)
	return nil
}

func (p *MqttModuleProxy) startup_session() int32 {
	return session_helper.GetStartupSession(p.opt.Session.Id)
}

func (p *MqttModuleProxy) generate_temp_session() int64 {
	return session_helper.NewSession(p.startup_session(), session_helper.GenerateTempSession())
}

func (p *MqttModuleProxy) mqtt_topic(mdl_id string, sess int64) string {
	return fmt.Sprintf("mtv1/modules/%v/sessions/%v", mdl_id, sess)
}

func (p *MqttModuleProxy) major_topic() string {
	return p.mqtt_topic(p.opt.Module.Id, p.opt.Session.Id)
}

func (p *MqttModuleProxy) UnaryCall(ctx context.Context, method string, value *any.Any) (*any.Any, error) {
	temp_sess := p.generate_temp_session()

	req := &pb.DownStreamFrame{
		SessionId: &wrappers.Int64Value{Value: temp_sess},
		Kind:      pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.DownStreamFrame_UnaryCall{
			UnaryCall: &pb.OpUnaryCallValue{
				Method: &wrappers.StringValue{Value: method},
				Value:  value,
			},
		},
	}

	buf, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = p.publish_message(p.opt.Session.Id, buf)
	if err != nil {
		return nil, err
	}

	ch, err := p.subscribe_session(temp_sess)
	if err != nil {
		return nil, err
	}
	defer p.unsubscribe_session(temp_sess)

	select {
	case buf = <-ch:
	case <-time.After(p.opt.Config.UnaryCallTimeout):
		return nil, ErrUnaryCallTimeout
	}

	res := pb.UpStreamFrame{}
	err = proto.Unmarshal(buf, &res)
	if err != nil {
		return nil, err
	}

	uc := res.GetUnaryCall()
	if uc == nil {
		return nil, ErrUnexceptedResponse
	}

	return uc.GetValue(), nil
}

func (p *MqttModuleProxy) StreamCall(ctx context.Context, method string, upstm ModuleProxyStream) error {
	panic("unimplemented")
}

type MqttModuleProxyFactory struct{}

func (f *MqttModuleProxyFactory) NewModuleProxy(args ...interface{}) (ModuleProxy, error) {
	opt := &MqttModuleProxyOption{}
	opt.Config.UnaryCallTimeout = 5 * time.Second
	opt.Config.StreamCallConfigAckTimeout = 5 * time.Second

	p := &MqttModuleProxy{
		opt:        opt,
		sess_chans: make(map[int64]chan []byte),
	}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":     opt_helper.ToLogger(&p.logger),
		"module_id":  opt_helper.ToString(&p.opt.Module.Id),
		"session_id": opt_helper.ToInt64(&p.opt.Session.Id),
		"mqtt_client": func(key string, val interface{}) error {
			var ok bool
			if p.c, ok = val.(libmqtt.Client); !ok {
				return opt_helper.InvalidArgument(key)
			}
			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return p, nil
}

func init() {
	register_module_proxy_factory("mqtt", new(MqttModuleProxyFactory))
}
