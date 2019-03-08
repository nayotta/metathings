package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) AuthorizeToken(ctx context.Context, req *pb.AuthorizeTokenRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
