package application_credential_manager

import (
	"context"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	identityd_pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type ApplicationCredentialManager interface {
	GetToken() string
}

type applicationCredentialManager struct {
	identityd_addr                string
	application_credential_id     string
	application_credential_secret string
	application_credential_token  string
}

func (mgr *applicationCredentialManager) GetToken() string {
	return "mt " + mgr.application_credential_token
}

func NewApplicationCredentialManager(identityd_addr, application_credential_id, application_credential_secret string) (ApplicationCredentialManager, error) {
	log.WithFields(log.Fields{
		"identiyd_address":              identityd_addr,
		"application_credential_id":     application_credential_id,
		"application_credential_secret": application_credential_secret,
	}).Debugf("login via application credential")

	var header metadata.MD
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()

	req := &identityd_pb.IssueTokenRequest{}
	req.Method = identityd_pb.AUTH_METHOD_APPLICATION_CREDENTIAL
	req.Payload = &identityd_pb.IssueTokenRequest_ApplicationCredential{
		&identityd_pb.ApplicationCredentialPayload{
			Id:     &gpb.StringValue{application_credential_id},
			Secret: &gpb.StringValue{application_credential_secret},
		},
	}

	conn, err := grpc.Dial(identityd_addr, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := identityd_pb.NewIdentityServiceClient(conn)

	_, err = cli.IssueToken(ctx, req, grpc.Header(&header))
	if err != nil {
		return nil, err
	}

	application_credential_token := header["authorization"][0]
	application_credential_token = application_credential_token[3:len(application_credential_token)]

	mgr := &applicationCredentialManager{
		identityd_addr,
		application_credential_id,
		application_credential_secret,
		application_credential_token,
	}

	return mgr, nil
}
