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
	mt_plugin.Closer
}

func (s streamingEchoStream) Send(req *any.Any) error {
	req1 := new(pb.EchoRequest)
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return err
	}

	return s.EchoService_StreamingEchoClient.Send(req1)
}

func (s streamingEchoStream) Recv() (*any.Any, error) {
	res, err := s.EchoService_StreamingEchoClient.Recv()
	if err != nil {
		return nil, err
	}

	return ptypes.MarshalAny(res)
}

func stream_streaming_echo(cli pb.EchoServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error) {
	stream, err := cli.StreamingEcho(ctx)
	if err != nil {
		return nil, err
	}

	return streamingEchoStream{
		EchoService_StreamingEchoClient: stream,
		Closer: mt_plugin.Closer{cbs},
	}, nil
}

func init() {
	stream_call_methods["StreamingEcho"] = stream_streaming_echo
}
