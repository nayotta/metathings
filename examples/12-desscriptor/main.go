package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	token     string
	addr      string
	command   string
	protoset  string
	sha1      string
	insecure  bool
	plaintext bool
)

func main() {
	pflag.StringVar(&addr, "addr", "", "Deviced Service Address")
	pflag.StringVar(&token, "token", "", "Token")
	pflag.StringVarP(&command, "command", "c", "", "Command[upload, download]")
	pflag.StringVar(&protoset, "protoset", "", "Protoset path")
	pflag.StringVar(&sha1, "sha1", "", "SHA1 hash code")
	pflag.BoolVar(&insecure, "insecure", false, "Insecure")
	pflag.BoolVar(&plaintext, "plaintext", false, "Plaintext")

	pflag.Parse()

	if token == "" {
		token = os.Getenv("MT_TOKEN")
	}

	ctx := context_helper.WithToken(context.TODO(), "Bearer "+token)
	var opts []grpc.DialOption
	if plaintext {
		opts = append(opts, grpc.WithInsecure())
	} else if insecure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := pb.NewDevicedServiceClient(conn)

	switch command[0] {
	case 'u':
		buf, err := ioutil.ReadFile(protoset)
		if err != nil {
			panic(err)
		}

		req := &pb.UploadDescriptorRequest{
			Descriptor_: &pb.OpDescriptor{
				Body: &wrappers.BytesValue{
					Value: buf,
				},
			},
		}

		res, err := cli.UploadDescriptor(ctx, req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("sha1: %v\n", res.GetDescriptor_().GetSha1())
	case 'd':
		req := &pb.GetDescriptorRequest{
			Descriptor_: &pb.OpDescriptor{
				Sha1: &wrappers.StringValue{
					Value: sha1,
				},
			},
		}

		res, err := cli.GetDescriptor(ctx, req)
		if err != nil {
			panic(err)
		}

		if err = ioutil.WriteFile(protoset, res.GetDescriptor_().GetBody(), 0644); err != nil {
			panic(err)
		}
	}
}
