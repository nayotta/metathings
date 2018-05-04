package application_credential_manager

import (
	"context"
	"sync"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	identityd_pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type ApplicationCredentialManager interface {
	GetToken() string
}

type applicationCredentialManager struct {
	mtx_token                     *sync.Mutex
	client_factory                *client_helper.ClientFactory
	application_credential_id     string
	application_credential_secret string
	application_credential_token  string

	refresh_interval int64
}

func (mgr *applicationCredentialManager) GetToken() string {
	mgr.mtx_token.Lock()
	defer mgr.mtx_token.Unlock()

	return "mt " + mgr.application_credential_token
}

func (mgr *applicationCredentialManager) refreshToken() error {
	mgr.mtx_token.Lock()
	defer mgr.mtx_token.Unlock()

	var header metadata.MD
	ctx := context.Background()
	req := &identityd_pb.IssueTokenRequest{
		Method: identityd_pb.AUTH_METHOD_APPLICATION_CREDENTIAL,
		Payload: &identityd_pb.IssueTokenRequest_ApplicationCredential{
			&identityd_pb.ApplicationCredentialPayload{
				Id:     &gpb.StringValue{mgr.application_credential_id},
				Secret: &gpb.StringValue{mgr.application_credential_secret},
			},
		},
	}

	cli, fn, err := mgr.client_factory.NewIdentityServiceClient()
	if err != nil {
		return err
	}
	defer fn()

	_, err = cli.IssueToken(ctx, req, grpc.Header(&header))
	if err != nil {
		return err
	}

	application_credential_token := header["authorization"][0]
	application_credential_token = application_credential_token[3:len(application_credential_token)]
	mgr.application_credential_token = application_credential_token

	return nil
}

func NewApplicationCredentialManager(cli_fty *client_helper.ClientFactory, application_credential_id, application_credential_secret string) (ApplicationCredentialManager, error) {
	log.WithFields(log.Fields{
		"application_credential_id":     application_credential_id,
		"application_credential_secret": application_credential_secret,
	}).Debugf("login via application credential")

	mgr := &applicationCredentialManager{
		refresh_interval:              360,
		mtx_token:                     new(sync.Mutex),
		client_factory:                cli_fty,
		application_credential_id:     application_credential_id,
		application_credential_secret: application_credential_secret,
	}

	err := mgr.refreshToken()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			<-time.After(time.Duration(mgr.refresh_interval) * time.Second)
			mgr.refreshToken()
		}
	}()

	return mgr, nil
}
