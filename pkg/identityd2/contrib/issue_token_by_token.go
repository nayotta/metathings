package metathings_identityd2_contrib

import (
	"github.com/golang/protobuf/ptypes/wrappers"

	pb "github.com/nayotta/metathings/proto/identityd2"
)

func NewIssueTokenByTokenRequest(domain, token string) *pb.IssueTokenByTokenRequest {
	req := &pb.IssueTokenByTokenRequest{
		Token: &pb.OpToken{
			Domain: &pb.OpDomain{
				Id: &wrappers.StringValue{Value: domain},
			},
			Text: &wrappers.StringValue{Value: token},
		},
	}

	return req
}
