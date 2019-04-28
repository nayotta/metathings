package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
	identityd2_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsDeviceServiceImpl) IssueModuleToken(ctx context.Context, req *pb.IssueModuleTokenRequest) (*pb.IssueModuleTokenResponse, error) {
	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect identityd2 service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	itc_req := &identityd2_pb.IssueTokenByCredentialRequest{
		Credential: req.GetCredential(),
		Timestamp:  req.GetTimestamp(),
		Nonce:      req.GetNonce(),
		Hmac:       req.GetHmac(),
	}
	itc_res, err := cli.IssueTokenByCredential(context.Background(), itc_req)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to issue token by credential from identityd2 service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	tkn0 := itc_res.GetToken()
	itt_req := &identityd2_pb.IssueTokenByTokenRequest{
		Token: &identityd2_pb.OpToken{
			Domain: &identityd2_pb.OpDomain{
				Id: &wrappers.StringValue{
					Value: tkn0.GetDomain().GetId(),
				},
			},
			Text: &wrappers.StringValue{
				Value: tkn0.GetText(),
			},
		},
	}
	itt_res, err := cli.IssueTokenByToken(context.Background(), itt_req)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to issue token by token from identityd2 service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.IssueModuleTokenResponse{
		Token: itt_res.GetToken(),
	}

	self.logger.Debugf("issue module token")

	return res, nil
}
