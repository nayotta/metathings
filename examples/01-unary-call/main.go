package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	stpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	token        string
	deviced_addr string
	device_id    string
	module       string
	method       string
	request      string
	protobufset  string
	insecure     bool
	plaintext    bool
	certfile     string
	keyfile      string
	soda         bool
)

func main() {
	pflag.StringVar(&deviced_addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&device_id, "device", "", "Device ID")
	pflag.StringVar(&module, "module", "", "Module Name")
	pflag.StringVar(&method, "method", "", "Method Name")
	pflag.StringVar(&request, "request", "", "JSON Request File")
	pflag.StringVar(&protobufset, "protobufset", "", "ProtobufSet File")
	pflag.BoolVar(&insecure, "insecure", false, "Insecure")
	pflag.BoolVar(&plaintext, "plaintext", false, "Plaintext")
	pflag.StringVar(&certfile, "certfile", "", "CertFile")
	pflag.StringVar(&keyfile, "keyfile", "", "KeyFile")
	pflag.BoolVar(&soda, "soda", false, "Enable soda mode")

	pflag.Parse()

	token = os.Getenv("MT_TOKEN")

	fmt.Printf("token=%v\naddr=%v\ndevice=%v\nmodule=%v\nprotobufset=%v\nmethod=%v\nrequest=%v\n", token, deviced_addr, device_id, module, protobufset, method, request)

	var err error
	var req_buf []byte
	var any_req *any.Any
	var md *desc.MethodDescriptor

	if request != "-" {
		req_buf, err = ioutil.ReadFile(request)
		if err != nil {
			panic(err)
		}
	} else {
		req_buf, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	if soda {
		var st stpb.Struct

		if err := grpc_helper.JSONPBUnmarshaler.Unmarshal(bytes.NewReader(req_buf), &st); err != nil {
			panic(err)
		}

		if any_req, err = ptypes.MarshalAny(&st); err != nil {
			panic(err)
		}
	} else {
		var fds dpb.FileDescriptorSet

		buf, err := ioutil.ReadFile(protobufset)
		if err != nil {
			panic(err)
		}

		if err = proto.Unmarshal(buf, &fds); err != nil {
			panic(err)
		}

		fd, err := desc.CreateFileDescriptorFromSet(&fds)
		if err != nil {
			panic(err)
		}

		srvs := fd.GetServices()
		if len(srvs) == 0 {
			panic("unexpected protobufset")
		}

		md = srvs[0].FindMethodByName(method)
		msg_req := dynamic.NewMessage(md.GetInputType())

		if err = msg_req.UnmarshalJSON(req_buf); err != nil {
			panic(err)
		}

		if any_req, err = ptypes.MarshalAny(msg_req); err != nil {
			panic(err)
		}
	}

	req := &pb.UnaryCallRequest{
		Device: &pb.OpDevice{
			Id: &wrappers.StringValue{Value: device_id},
		},
		Value: &pb.OpUnaryCallValue{
			Name:   &wrappers.StringValue{Value: module},
			Method: &wrappers.StringValue{Value: method},
			Value:  any_req,
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

	var msg_res proto.Message

	if soda {
		msg_res = new(stpb.Struct)
	} else {
		msg_res = dynamic.NewMessage(md.GetOutputType())
	}

	err = ptypes.UnmarshalAny(res.Value.Value, msg_res)
	if err != nil {
		panic(err)
	}

	out, err := grpc_helper.JSONPBMarshaler.MarshalToString(msg_res)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
