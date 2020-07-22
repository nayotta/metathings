package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/pflag"

	echo_pb "github.com/nayotta/metathings-component-echo/proto"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	pb "github.com/nayotta/metathings/pkg/proto/component"
)

var (
	domain    string
	cred_id   string
	cred_srt  string
	token     string
	prefix    string
	mqtt_addr string
	mdl_sess  int64
)

func main() {
	pflag.StringVar(&domain, "domain", "default", "Credential Domain")
	pflag.StringVar(&cred_id, "credential-id", "", "Credential ID")
	pflag.StringVar(&cred_srt, "credential-secret", "", "Credential Secret")
	pflag.StringVar(&prefix, "prefix", "http://metathings.ai:21733", "Device Cloud Service Address Prefix")
	pflag.StringVar(&mqtt_addr, "mqtt-addr", "mqtt.metathings.ai:1883", "Module MQTT Service Address")
	pflag.Int64Var(&mdl_sess, "module-session", 0, "Module Session")

	pflag.Parse()

	if mdl_sess == 0 {
		mdl_sess = rand.Int63()
	}

	cli := http.DefaultClient
	if token == "" {
		token = issue_module_token(cli)
	}

	go func() {
		for {
			heartbeat(cli)
			time.Sleep(30 * time.Second)
		}
	}()

	go start(cli)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch
}

func issue_module_token(cli *http.Client) string {
	ts := time.Now()
	nonce := rand.Int63()
	hmac := passwd_helper.MustParseHmac(cred_srt, cred_id, ts, nonce)

	imt_req := map[string]interface{}{
		"credential": map[string]interface{}{
			"id": cred_id,
		},
		"timestamp": ts.Format(time.RFC3339Nano),
		"nonce":     nonce,
		"hmac":      hmac,
	}

	buf, err := json.Marshal(imt_req)
	if err != nil {
		panic(err)
	}

	imt_url := fmt.Sprintf("%v/actions/issue_module_token", prefix)
	req, err := http.NewRequest("POST", imt_url, bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := cli.Do(req)
	if err != nil {
		panic(err)
	}

	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	tmp_tkn_res := struct{ Token struct{ Text string } }{}
	err = json.Unmarshal(buf, &tmp_tkn_res)
	if err != nil {
		panic(err)
	}

	return tmp_tkn_res.Token.Text
}

func heartbeat(cli *http.Client) {
	hb_url := fmt.Sprintf("%v/actions/heartbeat", prefix)
	req, err := http.NewRequest("POST", hb_url, strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("MT-Module-Session", fmt.Sprintf("%v", mdl_sess))

	_, err = cli.Do(req)
	if err != nil {
		panic(err)
	}
}

func start(http_cli *http.Client) {
	mdl_id := show_module_id(http_cli)
	hostname, _ := os.Hostname()
	mqtt_username := cred_id
	mqtt_password := mosquitto_service.ParseMosquittoPluginPassword(cred_id, cred_srt)
	topic := fmt.Sprintf("mt/modules/%v/proxy/sessions/+/downstream", mdl_id)
	fmt.Println(topic)
	opts := mqtt.NewClientOptions().
		AddBroker(mqtt_addr).
		SetUsername(mqtt_username).
		SetPassword(mqtt_password).
		SetClientID(hostname + strconv.Itoa(time.Now().Second())).
		SetCleanSession(true).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}).
		SetOnConnectHandler(func(c mqtt.Client) {
			if tkn := c.Subscribe(topic, byte(0), new_handler(c, mdl_id)); tkn.Wait() && tkn.Error() != nil {
				panic(tkn.Error())
			}
		})

	cli := mqtt.NewClient(opts)
	if tkn := cli.Connect(); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
}

func new_handler(cli mqtt.Client, mdl_id string) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		sess := extra_session_from_topic(m.Topic(), mdl_id, "downstream")
		var frm pb.DownStreamFrame
		if err := proto.Unmarshal(m.Payload(), &frm); err != nil {
			panic(err)
		}

		go handle_frame(c, mdl_id, sess, &frm)
	}
}

func extra_session_from_topic(topic string, mdl_id string, dir string) (ret int64) {
	if n, err := fmt.Sscanf(topic, "mt/modules/"+mdl_id+"/proxy/sessions/%d/"+dir, &ret); err != nil || n != 1 {
		return -1
	}
	return
}

func handle_frame(cli mqtt.Client, mdl_id string, sess int64, dstm_frm *pb.DownStreamFrame) {
	if dstm_frm.Kind != pb.StreamFrameKind_STREAM_FRAME_KIND_USER {
		panic("unexpected stream frame type")
	}

	switch dstm_frm.Union.(type) {
	case *pb.DownStreamFrame_UnaryCall:
		handle_unary_call_frame(cli, mdl_id, sess, dstm_frm)
	case *pb.DownStreamFrame_StreamCall:
		handle_stream_call_frame(cli, mdl_id, sess, dstm_frm)
	default:
		panic("unexpected stream frame type")
	}
}

func handle_unary_call_frame(cli mqtt.Client, mdl_id string, sess int64, dstm_frm *pb.DownStreamFrame) {
	unary := dstm_frm.GetUnaryCall()
	if unary == nil {
		panic("unary data is null")
	}

	method := unary.GetMethod().GetValue()
	if method != "Echo" {
		panic("unsupported method: " + method)
	}

	var req echo_pb.EchoRequest
	err := ptypes.UnmarshalAny(unary.GetValue(), &req)
	if err != nil {
		panic(err)
	}

	txt := req.GetText().GetValue()
	fmt.Println("recv msg: ", txt)

	res := echo_pb.EchoResponse{Text: txt}
	res_any, err := ptypes.MarshalAny(&res)
	if err != nil {
		panic(err)
	}

	temp_sess := unary.GetSession().GetValue()
	ustm_frm := pb.UpStreamFrame{
		Kind: pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.UpStreamFrame_UnaryCall{
			UnaryCall: &pb.UnaryCallValue{
				Session: temp_sess,
				Method:  method,
				Value:   res_any,
			},
		},
	}

	buf, err := proto.Marshal(&ustm_frm)
	if err != nil {
		panic(err)
	}

	topic := fmt.Sprintf("mt/modules/%v/proxy/sessions/%v/upstream", mdl_id, temp_sess)
	if tkn := cli.Publish(topic, byte(0), false, buf); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
}

func handle_stream_call_frame(cli mqtt.Client, mdl_id string, sess int64, dstm_frm *pb.DownStreamFrame) {
	switch sess {
	case mdl_sess:
		handle_stream_call_config_frame(cli, mdl_id, sess, dstm_frm)
	default:
		handle_stream_call_data_frame(cli, mdl_id, sess, dstm_frm)
	}
}

func handle_stream_call_config_frame(cli mqtt.Client, mdl_id string, sess int64, dstm_frm *pb.DownStreamFrame) {
	stream := dstm_frm.GetStreamCall()
	if stream == nil {
		panic("stream config is null")
	}

	cfg := stream.GetConfig()

	method := cfg.GetMethod().GetValue()
	if method != "StreamingEcho" {
		panic("unsupported method: " + method)
	}

	ack := cfg.GetAck().GetValue()

	ustm_frm := pb.UpStreamFrame{
		Kind: pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.UpStreamFrame_StreamCall{
			StreamCall: &pb.StreamCallValue{
				Union: &pb.StreamCallValue_Ack{
					Ack: &pb.StreamCallAck{Value: ack},
				},
			},
		},
	}

	buf, err := proto.Marshal(&ustm_frm)
	if err != nil {
		panic(err)
	}

	topic := fmt.Sprintf("mt/modules/%v/proxy/sessions/%v/upstream", mdl_id, cfg.GetSession().GetValue())
	if tkn := cli.Publish(topic, byte(0), false, buf); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
}

func handle_stream_call_data_frame(cli mqtt.Client, mdl_id string, sess int64, dstm_frm *pb.DownStreamFrame) {
	stream := dstm_frm.GetStreamCall()
	if stream == nil {
		panic("stream data is null")
	}

	var req echo_pb.StreamingEchoRequest
	err := ptypes.UnmarshalAny(stream.GetValue(), &req)
	if err != nil {
		panic(err)
	}

	txt := req.GetText().GetValue()
	fmt.Println("recv msg: ", txt)

	res := echo_pb.StreamingEchoResponse{Text: txt}
	res_any, err := ptypes.MarshalAny(&res)
	if err != nil {
		panic(err)
	}

	ustm_frm := pb.UpStreamFrame{
		Kind: pb.StreamFrameKind_STREAM_FRAME_KIND_USER,
		Union: &pb.UpStreamFrame_StreamCall{
			StreamCall: &pb.StreamCallValue{
				Union: &pb.StreamCallValue_Value{
					Value: res_any,
				},
			},
		},
	}

	buf, err := proto.Marshal(&ustm_frm)
	if err != nil {
		panic(err)
	}

	topic := fmt.Sprintf("mt/modules/%v/proxy/sessions/%v/upstream", mdl_id, sess)
	if tkn := cli.Publish(topic, byte(0), false, buf); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
}

func show_module_id(cli *http.Client) string {
	sm_url := fmt.Sprintf("%v/actions/show_module", prefix)
	req, err := http.NewRequest("POST", sm_url, strings.NewReader(`{}`))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	res, err := cli.Do(req)
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var tmp_mdl_res struct{ Module struct{ Id string } }
	err = json.Unmarshal(buf, &tmp_mdl_res)
	if err != nil {
		panic(err)
	}

	return tmp_mdl_res.Module.Id
}