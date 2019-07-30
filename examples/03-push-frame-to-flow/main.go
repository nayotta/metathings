package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	token        string
	deviced_addr string
	device_id    string
	flow_name    string
	input        string
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")
	// TODO(Peer): make it work.
	pflag.StringVar(&input, "input", "-", "Input File")

	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	fmt.Printf(`addr=%v
device=%v
flow=%v
`, deviced_addr, device_id, flow_name)

	ctx := context_helper.WithToken(context.Background(), "Bearer "+token)
	conn, err := grpc.Dial(deviced_addr, grpc.WithInsecure())
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

	// {"getSwitchDataRes":{"switchAddr":1,"votage":2,"leakCurrent":3,"power":0,"temp":0,"current":0}}
	src_js := `{"getSwitchDataRes":{"switchAddr":1,"votage":2,"leakCurrent":3,"power":0,"temp":0,"current":0}}`

	var src_st stpb.Struct
	err = jsonpb.UnmarshalString(src_js, &src_st)
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
