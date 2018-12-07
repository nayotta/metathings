package context_helper

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func WithToken(ctx context.Context, token string) context.Context {
	return NewOutgoingContext(ctx, WithTokenOp(token))
}

func WithTokenOp(token string) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("authorization", token)
		return md
	}
}

func WithSessionIdOp(sess_id string) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("session-id", sess_id)
		return md
	}
}

func WithSessionOp(sess int64) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("session", fmt.Sprintf("%v", sess))
		return md
	}
}

func NewOutgoingContext(ctx context.Context, fns ...func(metadata.MD) metadata.MD) context.Context {
	md := metadata.MD{}
	for _, fn := range fns {
		md = fn(md)
	}
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

func ExtractToken(ctx context.Context) *identityd2_pb.Token {
	return ctx.Value("token").(*identityd2_pb.Token)
}
