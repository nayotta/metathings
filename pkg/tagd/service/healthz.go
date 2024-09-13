package metathings_tagd_service

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (ts *MetathingsTagdService) Healthz(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: "OK"}, nil
}
