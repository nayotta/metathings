package context_helper

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc/metadata"

	identityd2_pb "github.com/nayotta/metathings/proto/identityd2"
)

func WithToken(ctx context.Context, token string) context.Context {
	return NewOutgoingContext(ctx, WithTokenOp(token))
}

func WithTokenOp(token string) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("Authorization", TokenToAuthorization(token))
		return md
	}
}

func WithSessionOp(sess int64) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("MT-Session", strconv.FormatInt(sess, 10))
		return md
	}
}

func WithDeviceOp(dev string) func(metadata.MD) metadata.MD {
	return func(md metadata.MD) metadata.MD {
		md.Append("MT-Device", dev)
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

func TokenToAuthorization(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		return token

	}
	return "Bearer " + strings.Trim(token, " ")
}

func AuthorizationToToken(authorization string) string {
	return strings.Trim(strings.Trim(authorization, "Bearer"), " ")
}
