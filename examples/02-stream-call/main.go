package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

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
	times        uint
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&module, "module", "", "Module Name")
	pflag.StringVar(&component, "component", "echo", "Component Name")
	pflag.StringVar(&method, "method", "StreamingEcho", "Method Name")
	pflag.StringVar(&request, "request", "", "JSON data")
	pflag.UintVar(&times, "times", 3, "repeat to send request")

	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	fmt.Printf("addr=%v\ndevice=%v\nmodule=%v\ncomponent=%v\nmethod=%v\nrequest=%v\n", deviced_addr, device_id, module, component, method, request)

	if component != "echo" && method != "StreamingEcho" {
		panic("unsupported call")
	}

	ctx := context_helper.WithToken(context.Background(), "mt "+token)
	conn, err := grpc.Dial(deviced_addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)
	stm, err := cli.StreamCall(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(1)

	cfg := &pb.StreamCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device_id},
		},
		Value: &pb.OpStreamCallValue{
			Union: &pb.OpStreamCallValue_Config{
				Config: &pb.OpStreamCallConfig{
					Name:      &wrappers.StringValue{Value: module},
					Component: &wrappers.StringValue{Value: component},
					Method:    &wrappers.StringValue{Value: method},
				},
			},
		},
	}

	err = stm.Send(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(2)

	res, err := stm.Recv()
	if err != nil {
		panic(err)
	}

	if res.GetValue().GetConfigAck() != nil {
		switch res.GetValue().Union.(type) {
		case *pb.StreamCallValue_ConfigAck:
		default:
			panic("unexpected response")
		}
	}

	echo_req := &echo_pb.EchoRequest{
		Text: &wrappers.StringValue{Value: request},
	}
	any_req, err := ptypes.MarshalAny(echo_req)
	if err != nil {
		panic(err)
	}
	fmt.Println(3)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := uint(0); i < times; i++ {
			res, err := stm.Recv()
			if err != nil {
				panic(err)
			}
			any_res := res.GetValue().GetValue()

			fmt.Printf("type=%v\nvalue=%v\n", any_res.TypeUrl, any_res.Value)

			var echo_res echo_pb.EchoResponse
			err = ptypes.UnmarshalAny(any_res, &echo_res)
			if err != nil {
				panic(err)
			}

			fmt.Println(echo_res.Text)
		}
	}()

	req := &pb.StreamCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device_id},
		},
		Value: &pb.OpStreamCallValue{
			Union: &pb.OpStreamCallValue_Value{
				Value: any_req,
			},
		},
	}

	for i := uint(0); i < times; i++ {
		time.Sleep(time.Duration(rand.Int63n(200))*time.Millisecond + 350*time.Millisecond)
		if err = stm.Send(req); err != nil {
			panic(err)
		}
		fmt.Printf("4.%v\n", i)
	}

	wg.Wait()
}
