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
	nonexpire         bool
	credential_domain string
	credential_id     string
	credential_secret string
	credential_token  string
}

func (self *tokener) GetToken() string {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	return "mt " + self.credential_token
}

func (self *tokener) issueToken() error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	itbc_req := identityd2_contrib.NewIssueTokenByCredentialRequest(const_helper.DEFAULT_DOMAIN, self.credential_id, self.credential_secret)
	itbc_res, err := cli.IssueTokenByCredential(context.Background(), itbc_req)
	if err != nil {
		return err
	}
	txt := itbc_res.Token.Text

	if self.nonexpire {
		itbt_req := identityd2_contrib.NewIssueTokenByTokenRequest(self.credential_domain, txt)
		itbt_res, err := cli.IssueTokenByToken(context.Background(), itbt_req)
		if err != nil {
			return err
		}
		txt = itbt_res.Token.Text
	}

	self.credential_token = txt

	return nil
}

func NewTokener(cli_fty *client_helper.ClientFactory, credential_domain, credential_id, credential_secret string) (Tokener, error) {
	tknr := &tokener{
		mtx:               new(sync.Mutex),
		cli_fty:           cli_fty,
		nonexpire:         false,
		credential_domain: credential_domain,
		credential_id:     credential_id,
		credential_secret: credential_secret,
	}

	if err := tknr.issueToken(); err != nil {
		return nil, err
	}

	return tknr, nil
}

func NewNoExpireTokener(cli_fty *client_helper.ClientFactory, credential_domain, credential_id, credential_secret string) (Tokener, error) {
	tknr := &tokener{
		mtx:               new(sync.Mutex),
		cli_fty:           cli_fty,
		nonexpire:         true,
		credential_domain: credential_domain,
		credential_id:     credential_id,
		credential_secret: credential_secret,
	}

	if err := tknr.issueToken(); err != nil {
		return nil, err
	}

	return tknr, nil
}
