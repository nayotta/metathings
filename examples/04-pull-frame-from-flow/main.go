package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	token        string
	deviced_addr string
	device_id    string
	flow_name    string
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")

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

	req_id := id_helper.NewId()
	req := &pb.PullFrameFromFlowRequest{
		Id: &wrappers.StringValue{Value: req_id},
		Request: &pb.PullFrameFromFlowRequest_Config_{
			Config: &pb.PullFrameFromFlowRequest_Config{
				Device: &pb.OpDevice{
					Id: &wrappers.StringValue{Value: device_id},
					Flows: []*pb.OpFlow{
						{Name: &wrappers.StringValue{Value: flow_name}},
					},
				},
				ConfigAck: &wrappers.BoolValue{Value: true},
			},
		},
	}

	cli := pb.NewDevicedServiceClient(conn)
	stm, err := cli.PullFrameFromFlow(ctx, req)
	if err != nil {
		panic(err)
	}

	res, err := stm.Recv()
	if err != nil {
		panic(err)
	}

	if res.GetAck() == nil {
		panic("unexpected response")
	}

	for {
		res, err = stm.Recv()
		if err != nil {
			panic(err)
		}

		pack := res.GetPack()
		if pack == nil {
			panic("unexpected response")
		}

		frms := pack.GetFrames()
		if len(frms) == 0 {
			continue
		}

		fmt.Println(frms[0].GetData().String())
	}
}
