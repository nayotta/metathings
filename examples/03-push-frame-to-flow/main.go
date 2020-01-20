package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	srv_ep_opt   cmd_contrib.ServiceEndpointOption
	base_opt     cmd_contrib.BaseOption
	device_id    string
	flow_name    string
	request_file string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Device Cloud Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&srv_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&srv_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")
	pflag.StringVar(&request_file, "request-file", "", "Request Files(json)")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	fmt.Printf(`addr=%v
device=%v
flow=%v
`, srv_ep_opt.Address, device_id, flow_name)

	ctx := context_helper.WithToken(context.Background(), base_opt.Token)
	ts, err := cmd_contrib.NewClientTransportCredentials(&srv_ep_opt)
	if err != nil {
		panic(err)
	}

	opts := grpc_helper.NewDialOptionWithTransportCredentials(ts)

	conn, err := grpc.Dial(srv_ep_opt.Address, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)
	stm, err := cli.PushFrameToFlow(ctx)
	if err != nil {
		panic(err)
	}

	cfg_id := id_helper.NewId()
	cfg := &pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: cfg_id},
		Request: &pb.PushFrameToFlowRequest_Config_{
			Config: &pb.PushFrameToFlowRequest_Config{
				Device: &pb.OpDevice{
					Id: &wrappers.StringValue{Value: device_id},
					Flows: []*pb.OpFlow{
						{Name: &wrappers.StringValue{Value: flow_name}},
					},
				},
				ConfigAck: &wrappers.BoolValue{Value: true},
				PushAck:   &wrappers.BoolValue{Value: true},
			},
		},
	}

	wg_cfg := sync.WaitGroup{}
	go func() {
		res, err := stm.Recv()
		if err != nil {
			panic(err)
		}

		if res.GetAck() == nil {
			panic("unexpected response")
		}

		if cfg_id != res.GetId() {
			panic("unexpected request id")
		}
		fmt.Println("ack config")
		wg_cfg.Done()
	}()

	wg_cfg.Add(1)
	err = stm.Send(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("send config")
	wg_cfg.Wait()

	src_js, err := ioutil.ReadFile(request_file)
	if err != nil {
		panic(err)
	}

	var src map[string]interface{}
	err = json.Unmarshal(src_js, &src)
	if err != nil {
		panic(err)
	}

	var src_st stpb.Struct
	err = jsonpb.UnmarshalString(string(src_js), &src_st)
	if err != nil {
		panic(err)
	}

	frm_id := id_helper.NewId()
	now := pb_helper.Now()
	req := &pb.PushFrameToFlowRequest{
		Id: &wrappers.StringValue{Value: frm_id},
		Request: &pb.PushFrameToFlowRequest_Frame{
			Frame: &pb.OpFrame{
				Ts:   &now,
				Data: &src_st,
			},
		},
	}

	wg_frm := sync.WaitGroup{}
	go func() {
		res, err := stm.Recv()
		if err != nil {
			panic(err)
		}

		if res.GetAck() == nil {
			panic("unexpected response")
		}
		fmt.Println("ack frame")
		wg_frm.Done()
	}()

	wg_frm.Add(1)
	err = stm.Send(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("send frame")
	wg_frm.Wait()

	fmt.Println("done")
}
