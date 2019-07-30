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
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	token             string
	device_cloud_addr string
	mqtt_addr         string
	cred_id           string
	cred_srt          string
	mqtt_qos          int
	flow_name         string
	times             int
)

func main() {
	pflag.StringVar(&device_cloud_addr, "device-cloud-addr", "", "Device Cloud Service Address")
	pflag.StringVar(&mqtt_addr, "mqtt-addr", "", "MQTT Server Address")
	pflag.IntVar(&mqtt_qos, "mqtt-qos", 0, "MQTT QoS")
	pflag.StringVar(&cred_id, "credential-id", "", "Credential ID")
	pflag.StringVar(&cred_srt, "credential-secret", "", "Credential Secret")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")
	pflag.IntVar(&times, "times", 3, "Repeat to send data times")
	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	dec := jsonpb.Unmarshaler{}

	req, err := http.NewRequest("POST", device_cloud_addr+"/actions/show_module", strings.NewReader("{}"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
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
times=%v
`, token, device_cloud_addr, mqtt_addr, mqtt_qos, cred_id, cred_srt, mqtt_username, mqtt_password, mqtt_clientid, device_id, flow_name, times)

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

	req, err = http.NewRequest("POST", device_cloud_addr+"/actions/push_frame_to_flow", strings.NewReader(buf))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	var pftf_cfg_res device_pb.PushFrameToFlowResponse

	res, err = http.DefaultClient.Do(req)
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
		time.Sleep(time.Duration(500+rand.Int63n(1000)) * time.Millisecond)
		send_data_once(c, pub_tpc)
	}
}

func send_data_once(c mqtt.Client, pub_tpc string) {
	dat := random_data()
	dat_js, err := json.Marshal(dat)
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
