package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

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
	insecure     bool
	plaintext    bool
	certfile     string
	keyfile      string
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&module, "module", "", "Module Name")
	pflag.StringVar(&component, "component", "echo", "Component Name")
	pflag.StringVar(&method, "method", "Echo", "Method Name")
	pflag.StringVar(&request, "request", "", "JSON Request File")
	pflag.BoolVar(&insecure, "insecure", false, "Insecure")
	pflag.BoolVar(&plaintext, "plaintext", false, "Plaintext")
	pflag.StringVar(&certfile, "certfile", "", "CertFile")
	pflag.StringVar(&keyfile, "keyfile", "", "KeyFile")

	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	fmt.Printf("addr=%v\ndevice=%v\nmodule=%v\ncomponent=%v\nmethod=%v\nrequest=%v\n", deviced_addr, device_id, module, component, method, request)

	var any_req *any.Any

	switch component {
	case "echo":
		switch method {
		case "Echo":
			echo_req := &echo_pb.EchoRequest{
				Text: &wrappers.StringValue{Value: request},
			}
			any_req, _ = ptypes.MarshalAny(echo_req)
		}
	case "camera":
		switch method {
		case "Start":
			any_req, _ = ptypes.MarshalAny(&empty.Empty{})
		case "Stop":
			any_req, _ = ptypes.MarshalAny(&empty.Empty{})
		}
	case "switch":
		switch method {
		case "Start":
			any_req, _ = ptypes.MarshalAny(&empty.Empty{})
		case "Stop":
			any_req, _ = ptypes.MarshalAny(&empty.Empty{})
		}
	}

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

	ctx := context_helper.WithToken(context.Background(), "Bearer "+token)

	var opts []grpc.DialOption
	if plaintext {
		opts = append(opts, grpc.WithInsecure())
	} else if insecure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})))
	} else if certfile != "" && keyfile != "" {
		cred, err := credentials.NewServerTLSFromFile(certfile, keyfile)
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}

	conn, err := grpc.Dial(deviced_addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)
	res, err := cli.UnaryCall(ctx, req)
	if err != nil {
		panic(err)
	}

	switch component {
	case "echo":
		switch method {
		case "Echo":
			var echo_res echo_pb.EchoResponse
			err = ptypes.UnmarshalAny(res.Value.Value, &echo_res)
			if err != nil {
				panic(err)
			}

			fmt.Println(echo_res.Text)
		}
	case "camera":
		switch method {
		case "Start":
			fmt.Println("camera started")
		case "Stop":
			fmt.Println("camera stoped")
		}
	}
}
