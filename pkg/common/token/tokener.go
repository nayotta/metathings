package token_helper

import (
	"context"
	"sync"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
	identityd2_contrib "github.com/nayotta/metathings/pkg/identityd2/contrib"
)

type Tokener interface {
	GetToken() string
}

type tokener struct {
	mtx               *sync.Mutex
	cli_fty           *client_helper.ClientFactory
	credential_id     string
	credential_secret string
	credential_token  string
}

func (self *tokener) GetToken() string {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	return "mt " + self.credential_token
}

func (self *tokener) refreshToken() error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	req := identityd2_contrib.NewIssueTokenByCredentialRequest(const_helper.DEFAULT_DOMAIN, self.credential_id, self.credential_secret)

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	res, err := cli.IssueTokenByCredential(context.Background(), req)
	if err != nil {
		return err
	}

	self.credential_token = res.Token.Text

	return nil
}

func NewTokener(cli_fty *client_helper.ClientFactory, credential_id, credential_secret string) (Tokener, error) {
	tknr := &tokener{
		mtx:               new(sync.Mutex),
		cli_fty:           cli_fty,
		credential_id:     credential_id,
		credential_secret: credential_secret,
	}

	if err := tknr.refreshToken(); err != nil {
		return nil, err
	}

	return tknr, nil
}
