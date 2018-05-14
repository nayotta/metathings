package context_helper

import (
	"context"

	"google.golang.org/grpc/metadata"

	identityd_pb "github.com/nayotta/metathings/pkg/proto/identity"
)

func WithToken(ctx context.Context, token string) context.Context {
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}

func Credential(ctx context.Context) (cred *identityd_pb.Token) {
	var ok bool

	if cred, ok = ctx.Value("credential").(*identityd_pb.Token); !ok {
		return nil
	}

	return cred
}
