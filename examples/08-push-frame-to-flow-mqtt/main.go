package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	device_pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	mqtt_addr    string
	dc_ep_opt    cmd_contrib.ServiceEndpointOption
	base_opt     cmd_contrib.BaseOption
	cred_id      string
	cred_srt     string
	mqtt_qos     int
	flow_name    string
	interval     int
	request_file string
)

func get_full_url(ep cmd_contrib.ServiceEndpointOption, p string) string {
	u, err := url.Parse("http://" + path.Join(ep.Address, "/v1/device_cloud", p))
	if err != nil {
		panic(err)
	}

	if ep.PlainText {
		u.Scheme = "http"
	} else {
		u.Scheme = "https"
	}

	return u.String()
}

func get_http_client(ep cmd_contrib.ServiceEndpointOption) *http.Client {
	if ep.PlainText {
		return http.DefaultClient
	} else if ep.Insecure {
		return &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	} else {
		return &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: nil,
			},
		}
	}

}

func main() {
	pflag.StringVar(&dc_ep_opt.Address, "addr", "", "Device Cloud Service Address")
	pflag.BoolVar(&dc_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&dc_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&dc_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&dc_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&mqtt_addr, "mqtt-addr", "", "MQTT Server Address")
	pflag.IntVar(&mqtt_qos, "mqtt-qos", 0, "MQTT QoS")
	pflag.StringVar(&cred_id, "credential-id", "", "Credential ID")
	pflag.StringVar(&cred_srt, "credential-secret", "", "Credential Secret")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")
	pflag.IntVar(&interval, "interval", 1, "Send data interval")
	pflag.StringVar(&request_file, "request-file", "", "Request File(json)")
	pflag.Parse()

	httpcli := get_http_client(dc_ep_opt)
	dec := jsonpb.Unmarshaler{}

	base_opt.Token = issue_module_token(httpcli)

	req, err := http.NewRequest("POST", get_full_url(dc_ep_opt, "/actions/show_module"), strings.NewReader("{}"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+base_opt.Token)

	res, err := httpcli.Do(req)
	if err != nil {
		panic(err)
	}

	bbuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("show module response: ", string(bbuf))
	var sm_res device_pb.ShowModuleResponse
	err = dec.Unmarshal(bytes.NewReader(bbuf), &sm_res)
	if err != nil {
		panic(err)
	}

	device_id := sm_res.GetModule().GetDeviceId()
	hostname, _ := os.Hostname()
	mqtt_username := cred_id
	mqtt_password := mosquitto_service.ParseMosquittoPluginPassword(cred_id, cred_srt)
	mqtt_clientid := hostname + strconv.Itoa(time.Now().Second())

	fmt.Printf(`token=%v
device_cloud_addr=%v
mqtt_addr=%v
mqtt_qos=%v
cred_id=%v
cred_srt=%v
mqtt_usr=%v
mqtt_pwd=%v
mqtt_cli=%v
device=%v
flow=%v
interval=%v
`, base_opt.Token, dc_ep_opt.Address, mqtt_addr, mqtt_qos, cred_id, cred_srt, mqtt_username, mqtt_password, mqtt_clientid, device_id, flow_name, interval)

	pftf_cfg_req := &device_pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Request: &device_pb.PushFrameToFlowRequest_Config_{
			Config: &device_pb.PushFrameToFlowRequest_Config{
				Flow: &deviced_pb.OpFlow{
					Name: &wrappers.StringValue{Value: flow_name},
				},
				ConfigAck: &wrappers.BoolValue{Value: true},
				PushAck:   &wrappers.BoolValue{Value: true},
			},
		},
	}
	enc := jsonpb.Marshaler{}
	buf, err := enc.MarshalToString(pftf_cfg_req)
	if err != nil {
		panic(err)
	}
	fmt.Println("config request: ", buf)

	req, err = http.NewRequest("POST", get_full_url(dc_ep_opt, "/actions/push_frame_to_flow"), strings.NewReader(buf))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+base_opt.Token)

	var pftf_cfg_res device_pb.PushFrameToFlowResponse

	res, err = httpcli.Do(req)
	if err != nil {
		panic(err)
	}

	bbuf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("config response: ", string(bbuf))

	err = dec.Unmarshal(bytes.NewReader(bbuf), &pftf_cfg_res)
	if err != nil {
		panic(err)
	}

	cfg := pftf_cfg_res.GetConfig()
	if cfg == nil {
		panic("bad response")
	}

	sess := cfg.GetSession()
	pub_tpc := fmt.Sprintf("mt/devices/%v/flow_channel/sessions/%v/%v", device_id, sess, "upstream")
	sub_tpc := fmt.Sprintf("mt/devices/%v/flow_channel/sessions/%v/%v", device_id, sess, "downstream")
	fmt.Println("publish topic: ", pub_tpc)
	fmt.Println("subscribe topic: ", sub_tpc)

	mqtt_opts := mqtt.NewClientOptions().
		AddBroker(mqtt_addr).
		SetUsername(mqtt_username).
		SetPassword(mqtt_password).
		SetClientID(mqtt_clientid).
		SetCleanSession(true).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}).
		SetOnConnectHandler(func(c mqtt.Client) {
			if tkn := c.Subscribe(sub_tpc, byte(mqtt_qos), new_message_handler()); tkn.Wait() && tkn.Error() != nil {
				panic(tkn.Error())
			}
			fmt.Println("sbuscribe topic: ", sub_tpc)
		})
	mqtt_cli := mqtt.NewClient(mqtt_opts)
	if tkn := mqtt_cli.Connect(); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
	fmt.Println("mqtt client connected")

	go ping_loop(mqtt_cli, pub_tpc)
	go send_data_loop(mqtt_cli, pub_tpc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
}

func ping_loop(c mqtt.Client, pub_tpc string) {
	for {
		time.Sleep(13 * time.Second)
		ping_once(c, pub_tpc)
	}
}

func ping_once(c mqtt.Client, pub_tpc string) {
	req := &device_pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Request: &device_pb.PushFrameToFlowRequest_Ping_{
			Ping: &device_pb.PushFrameToFlowRequest_Ping{},
		},
	}

	bbuf, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	tkn := c.Publish(pub_tpc, byte(mqtt_qos), false, string(bbuf))
	if tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}

	fmt.Println("ping data")
}

func send_data_loop(c mqtt.Client, pub_tpc string) {
	for {
		send_data_once(c, pub_tpc)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func send_data_once(c mqtt.Client, pub_tpc string) {
	dat_js, err := ioutil.ReadFile(request_file)
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	err = json.Unmarshal(dat_js, &dat)
	if err != nil {
		panic(err)
	}

	var dat_st stpb.Struct
	err = jsonpb.UnmarshalString(string(dat_js), &dat_st)
	if err != nil {
		panic(err)
	}

	req := &device_pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Request: &device_pb.PushFrameToFlowRequest_Frame{
			Frame: &deviced_pb.OpFrame{
				Ts:   ptypes.TimestampNow(),
				Data: &dat_st,
			},
		},
	}

	bbuf, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	if tkn := c.Publish(pub_tpc, byte(mqtt_qos), false, string(bbuf)); tkn.Wait() && tkn.Error() != nil {
		panic(tkn.Error())
	}
	fmt.Println("send data: ", dat)
	fmt.Println("send data topic: ", pub_tpc)
	fmt.Println("send data req: ", req.GetId().GetValue())
}

func random_data() map[string]interface{} {
	return map[string]interface{}{
		"temperature": 35.0 + rand.Int63n(20)/10.,
	}
}

func new_message_handler() func(mqtt.Client, mqtt.Message) {
	return func(c mqtt.Client, m mqtt.Message) {
		var res device_pb.PushFrameToFlowResponse
		err := proto.Unmarshal(m.Payload(), &res)
		if err != nil {
			panic(err)
		}

		if res.GetAck() != nil {
			fmt.Println("recv ack res: ", res.GetId())
		}

		if res.GetPong() != nil {
			fmt.Println("recv pong res")
		}
	}
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

	imt_url := get_full_url(dc_ep_opt, "/actions/issue_module_token")

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
