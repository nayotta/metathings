package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	srv_ep_opt cmd_contrib.ServiceEndpointOption
	base_opt   cmd_contrib.BaseOption
	device_id  string
	flow_name  string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Deviced Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&srv_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&srv_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&flow_name, "flow", "", "Flow Name")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	fmt.Printf(`addr=%v
device=%v
flow=%v
`, srv_ep_opt.Address, device_id, flow_name)

	ctx := context_helper.WithToken(context.TODO(), "Bearer "+base_opt.Token)

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

		fmt.Println(time.Now())

		pack := res.GetPack()
		if pack == nil {
			fmt.Println(res.GetId())
			continue
		}

		frms := pack.GetFrames()
		if len(frms) == 0 {
			continue
		}

		fmt.Println(frms[0].GetData().String())
	}
}
