package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	echo_pb "github.com/nayotta/metathings-component-echo/proto"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	token        string
	deviced_addr string
	device_id    string
	module       string
	component    string
	method       string
	request      string
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&module, "module", "", "Module Name")
	pflag.StringVar(&component, "component", "echo", "Component Name")
	pflag.StringVar(&method, "method", "Echo", "Method Name")
	pflag.StringVar(&request, "request", "", "JSON Request File")

	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	fmt.Printf("addr=%v\ndevice=%v\nmodule=%v\ncomponent=%v\nmethod=%v\nrequest=%v\n", deviced_addr, device_id, module, component, method, request)

	if component != "echo" && method != "Echo" {
		panic("unsupported call")
	}

	echo_req := &echo_pb.EchoRequest{
		Text: &wrappers.StringValue{Value: request},
	}

	any_req, _ := ptypes.MarshalAny(echo_req)

	req := &pb.UnaryCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device_id},
		},
		Value: &pb.OpUnaryCallValue{
			Name:      &wrappers.StringValue{Value: module},
			Component: &wrappers.StringValue{Value: component},
			Method:    &wrappers.StringValue{Value: method},
			Value:     any_req,
		},
	}

	ctx := context_helper.WithToken(context.Background(), "mt "+token)

	conn, err := grpc.Dial(deviced_addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)
	res, err := cli.UnaryCall(ctx, req)
	if err != nil {
		panic(err)
	}

	var echo_res echo_pb.EchoResponse
	err = ptypes.UnmarshalAny(res.Value.Value, &echo_res)
	if err != nil {
		panic(err)
	}

	fmt.Println(echo_res.Text)
}
