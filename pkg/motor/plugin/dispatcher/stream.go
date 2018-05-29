package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type streamStream struct {
	pb.MotorService_StreamClient
	mt_plugin.Closer
}

func (s streamStream) Send(req *any.Any) error {
	req1 := new(pb.StreamRequests)
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return err
	}

	return s.MotorService_StreamClient.Send(req1)
}

func (s streamStream) Recv() (*any.Any, error) {
	res, err := s.MotorService_StreamClient.Recv()
	if err != nil {
		return nil, err
	}

	return ptypes.MarshalAny(res)
}

func stream_stream(cli pb.MotorServiceClient, ctx context.Context, cbs ...func()) (mt_plugin.Stream, error) {
	stream, err := cli.Stream(ctx)
	if err != nil {
		return nil, err
	}

	return streamStream{
		MotorService_StreamClient: stream,
		Closer: mt_plugin.Closer{cbs},
	}, nil
}

func init() {
	stream_call_methods["Stream"] = stream_stream
}
