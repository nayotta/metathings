package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ListActions(ctx context.Context, req *pb.ListActionsRequest) (*pb.ListActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
