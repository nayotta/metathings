package metathings_identityd2_contrib

import (
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func NewIssueTokenByCredentialRequest(domain, id, secret string) *pb.IssueTokenByCredentialRequest {
	ts := time.Now()
	nonce := rand.Int63()
	hmac := passwd_helper.MustParseHmac(secret, id, ts, nonce)

	pb_ts := pb_helper.FromTime(ts)

	req := &pb.IssueTokenByCredentialRequest{
		Credential: &pb.OpCredential{
			Domain: &pb.OpDomain{
				Id: &wrappers.StringValue{
					Value: domain,
				},
			},
			Id: &wrappers.StringValue{
				Value: id,
			},
		},
		Timestamp: &pb_ts,
		Nonce:     &wrappers.Int64Value{Value: nonce},
		Hmac:      &wrappers.StringValue{Value: hmac},
	}

	return req
}
