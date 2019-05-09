package context_helper

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc/metadata"

	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func WithToken(ctx context.Context, token string) context.Context {
	return NewOutgoingContext(ctx, WithTokenOp(token))
}

func WithTokenOp(token string) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		if !strings.HasPrefix(token, "mt") {
			token = "mt " + strings.Trim(token, " ")
		}

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
		md.Append("session", strconv.FormatInt(sess, 10))
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

func ExtractToken(ctx context.Context) *identityd2_pb.Token {
	return ctx.Value("token").(*identityd2_pb.Token)
}
