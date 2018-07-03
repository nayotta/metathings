package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
)

type streamStream struct {
	pb.ServoService_StreamClient
	mt_plugin.Closer
}

func (s streamStream) Send(req *any.Any) error {
	req1 := new(pb.StreamRequests)
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return err
	}

	return s.ServoService_StreamClient.Send(req1)
}

func (s streamStream) Recv() (*any.Any, error) {
	res, err := s.ServoService_StreamClient.Recv()
	if err != nil {
		return nil, err
	}

	return ptypes.MarshalAny(res)
}

func stream_stream(cli pb.ServoServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error) {
	stream, err := cli.Stream(ctx)
	if err != nil {
		return nil, err
	}

	return streamStream{
		ServoService_StreamClient: stream,
		Closer: mt_plugin.Closer{cbs},
	}, nil
}

func init() {
	stream_call_methods["Stream"] = stream_stream
}
