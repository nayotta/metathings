package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/echo"
)

type streamingEchoStream struct {
	pb.EchoService_StreamingEchoClient
}

func (s streamingEchoStream) Send(req *any.Any) error {
	req1 := new(pb.EchoRequest)
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return err
	}

	err = s.EchoService_StreamingEchoClient.Send(req1)
	if err != nil {
		return err
	}

	return nil
}

func (s streamingEchoStream) Recv() (*any.Any, error) {
	res, err := s.EchoService_StreamingEchoClient.Recv()
	if err != nil {
		return nil, err
	}

	res1, err := ptypes.MarshalAny(res)
	if err != nil {
		return nil, err
	}

	return res1, nil
}

func stream_streaming_echo(cli pb.EchoServiceClient, ctx context.Context) (mt_plugin.Stream, error) {
	stream, err := cli.StreamingEcho(ctx)
	if err != nil {
		return nil, err
	}

	return streamingEchoStream{stream}, nil
}

func init() {
	stream_call_methods["StreamingEcho"] = stream_streaming_echo
}
