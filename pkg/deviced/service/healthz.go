package metathings_deviced_service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (self *MetathingsDevicedService) Healthz(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: "OK"}, nil
}
