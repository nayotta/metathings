package main

import (
	"context"
	"fmt"
	"os"

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
	srv_ep_opt  cmd_contrib.ServiceEndpointOption
	base_opt    cmd_contrib.BaseOption
	flow_set_id string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Deviced Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&srv_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&srv_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&flow_set_id, "flow-set-id", "", "Flow Set Id")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	fmt.Printf(`addr=%v
flow_set=%v
`, srv_ep_opt.Address, flow_set_id)

	ctx := context_helper.WithToken(context.TODO(), base_opt.Token)
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
	req := &pb.PullFrameFromFlowSetRequest{
		Id: &wrappers.StringValue{Value: req_id},
		Request: &pb.PullFrameFromFlowSetRequest_Config_{
			Config: &pb.PullFrameFromFlowSetRequest_Config{
				FlowSet: &pb.OpFlowSet{
					Id: &wrappers.StringValue{Value: flow_set_id},
				},
				ConfigAck: &wrappers.BoolValue{Value: true},
			},
		},
	}

	cli := pb.NewDevicedServiceClient(conn)
	stm, err := cli.PullFrameFromFlowSet(ctx, req)
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
			fmt.Println(res.GetId())
			continue
		}

		dev := pack.GetDevice()
		frms := pack.GetFrames()

		for _, frm := range frms {
			fmt.Printf("device=%v flow=%v frame=%v\n", dev.GetId(), dev.GetFlows()[0].GetName(), frm)
		}
	}
}
