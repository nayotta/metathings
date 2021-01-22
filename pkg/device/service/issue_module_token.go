package metathings_device_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/proto/device"
	identityd2_pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsDeviceServiceImpl) IssueModuleToken(ctx context.Context, req *pb.IssueModuleTokenRequest) (*pb.IssueModuleTokenResponse, error) {
	logger := self.get_logger().WithField("method", "IssueModuleToken")

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to connect identityd2 service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	tkn, err := IssueModuleTokenWithClient(
		cli, context.TODO(),
		req.GetCredential().GetId().GetValue(),
		protobuf_helper.ToTime(*req.GetTimestamp()),
		req.GetNonce().GetValue(),
		req.GetHmac().GetValue(),
	)

	if err != nil {
		logger.WithError(err).Errorf("failed to issue token in identityd2 service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.IssueModuleTokenResponse{
		Token: tkn,
	}

	logger.Debugf("issue module token")

	return res, nil
}

func IssueModuleTokenWithClient(cli identityd2_pb.IdentitydServiceClient, ctx context.Context, credential_id string, timestamp time.Time, nonce int64, hmac string) (*identityd2_pb.Token, error) {
	ts_pb := protobuf_helper.FromTime(timestamp)
	itc_req := &identityd2_pb.IssueTokenByCredentialRequest{
		Credential: &identityd2_pb.OpCredential{
			Id: &wrappers.StringValue{
				Value: credential_id,
			},
		},
		Timestamp: &ts_pb,
		Nonce: &wrappers.Int64Value{
			Value: nonce,
		},
		Hmac: &wrappers.StringValue{
			Value: hmac,
		},
	}
	itc_res, err := cli.IssueTokenByCredential(context.Background(), itc_req)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return itt_res.GetToken(), nil
}
