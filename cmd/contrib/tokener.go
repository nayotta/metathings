package cmd_contrib

import (
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
)

func NewTokener(opt CredentialOptioner, cli_fty *client_helper.ClientFactory) (token_helper.Tokener, error) {
	var tknr token_helper.Tokener
	var err error

	if tknr, err = token_helper.NewTokener(cli_fty, opt.GetCredentialId(), opt.GetCredentialSecret()); err != nil {
		return nil, err
	}

	return tknr, nil
}
