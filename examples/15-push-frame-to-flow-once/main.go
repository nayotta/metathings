package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	stpb "github.com/golang/protobuf/ptypes/struct"
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
	device     string
	flow       string
	file       string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Deviced Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&device, "device", "", "Device ID")
	pflag.StringVar(&flow, "flow", "", "Flow Name")
	pflag.StringVar(&file, "file", "", "Request Files(json)")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	fmt.Printf(`
addr=%v
device=%v
flow=%v
file=%v
`, srv_ep_opt.Address, device, flow, file)

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

	cli := pb.NewDevicedServiceClient(conn)

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var payload stpb.Struct
	err = jsonpb.UnmarshalString(string(buf), &payload)
	if err != nil {
		panic(err)
	}

	frm := &pb.OpFrame{
		Ts:   ptypes.TimestampNow(),
		Data: &payload,
	}

	req := &pb.PushFrameToFlowOnceRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device},
			Flows: []*pb.OpFlow{
				{Name: &wrappers.StringValue{Value: flow}},
			},
		},
		Frame: frm,
	}

	if _, err = cli.PushFrameToFlowOnce(ctx, req); err != nil {
		panic(err)
	}

	fmt.Println("send frame")
}
