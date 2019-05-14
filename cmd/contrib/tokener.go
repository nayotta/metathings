package cmd_contrib

import (
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
)

func NewTokener(opt CredentialOptioner, cli_fty *client_helper.ClientFactory) (tknr token_helper.Tokener, err error) {
	tknr, err = token_helper.NewTokener(cli_fty, opt.GetCredentialDomain(), opt.GetCredentialId(), opt.GetCredentialSecret())
	return
}

func NewNoExpireTokener(opt CredentialOptioner, cli_fty *client_helper.ClientFactory) (tknr token_helper.Tokener, err error) {
	tknr, err = token_helper.NewNoExpireTokener(cli_fty, opt.GetCredentialDomain(), opt.GetCredentialId(), opt.GetCredentialSecret())
	return
}
