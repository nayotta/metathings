package main

import (
	"context"
	"os"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	component "github.com/nayotta/metathings/pkg/component"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

var (
	srv_ep_opt  cmd_contrib.ServiceEndpointOption
	base_opt    cmd_contrib.BaseOption
	device      string
	source      string
	destination string
)

func main() {
	pflag.StringVar(&srv_ep_opt.Address, "addr", "", "Deviced Service Address")
	pflag.BoolVar(&srv_ep_opt.Insecure, "insecure", false, "Insecure Connection")
	pflag.BoolVar(&srv_ep_opt.PlainText, "plaintext", false, "Plain Text Connection")
	pflag.StringVar(&srv_ep_opt.CertFile, "certfile", "", "Cert File for connect to Deviced")
	pflag.StringVar(&srv_ep_opt.KeyFile, "keyfile", "", "Key File for connect to Deviced")
	pflag.StringVar(&device, "device", "", "Device ID")
	pflag.StringVar(&source, "source", "", "Read file from source path")
	pflag.StringVar(&destination, "destination", "", "Write file to destination path")

	pflag.Parse()

	base_opt.Token = os.Getenv("MT_TOKEN")

	fp, err := os.Open(source)
	if err != nil {
		panic(err)
	}

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

	cli := pb.NewDevicedServiceClient(conn)
	stm, err := cli.PutObjectStreaming(ctx)
	if err != nil {
		panic(err)
	}

	pos_opt, err := component.NewPutObjectStreamingOptionFromPath(source)
	if err != nil {
		panic(err)
	}

	cfg_req := &pb.PutObjectStreamingRequest{
		Id: &wrappers.StringValue{Value: id_helper.NewId()},
		Request: &pb.PutObjectStreamingRequest_Metadata_{
			Metadata: &pb.PutObjectStreamingRequest_Metadata{
				Object: &pb.OpObject{
					Name:   &wrappers.StringValue{Value: destination},
					Length: &wrappers.Int64Value{Value: pos_opt.Length},
				},
				Sha1: &wrappers.StringValue{Value: pos_opt.Sha1},
			},
		},
	}

	err = stm.Send(cfg_req)
	if err != nil {
		panic(err)
	}

	errs := make(chan error)
	defer close(errs)

	go func() {
		for {
			res, err := stm.Recv()
			if err != nil {
				errs <- err
				return
			}

			chunks := res.GetChunks()
			if chunks == nil {
				continue
			}

			chk_req := &pb.PutObjectStreamingRequest{
				Id: &wrappers.StringValue{Value: res.GetId()},
			}
			req_chks := []*pb.OpObjectChunk{}
			for _, chk := range chunks.GetChunks() {
				offset := chk.GetOffset()
				length := chk.GetLength()
				buf := make([]byte, length)

				_, err = fp.Seek(offset, 0)
				if err != nil {
					errs <- err
					return
				}

				n, err := fp.Read(buf)
				if err != nil {
					errs <- err
					return
				}
				req_chks = append(req_chks, &pb.OpObjectChunk{
					Offset: &wrappers.Int64Value{Value: offset},
					Data:   &wrappers.BytesValue{Value: buf},
					Length: &wrappers.Int64Value{Value: int64(n)},
				})
			}
			chk_req.Request = &pb.PutObjectStreamingRequest_Chunks{
				Chunks: &pb.OpObjectChunks{
					Chunks: req_chks,
				},
			}

			err = stm.Send(chk_req)
			if err != nil {
				errs <- err
				return
			}
		}
	}()

	err = <-errs
	if err != nil {
		panic(err)
	}
}
