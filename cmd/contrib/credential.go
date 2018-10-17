package cmd_contrib

import (
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type CredentialOptioner interface {
	GetCredentialIdP() *string
	GetCredentialId() string
	SetCredentialId(string)

	GetCredentialSecretP() *string
	GetCredentialSecret() string
	SetCredentialSecret(string)
}

type CredentialOption struct {
	Credential struct {
		Id     string
		Secret string
	}
}

func (self *CredentialOption) GetCredentialIdP() *string {
	return &self.Credential.Id
}

func (self *CredentialOption) GetCredentialId() string {
	return self.Credential.Id
}

func (self *CredentialOption) SetCredentialId(id string) {
	self.Credential.Id = id
}

func (self *CredentialOption) GetCredentialSecretP() *string {
	return &self.Credential.Secret
}

func (self *CredentialOption) GetCredentialSecret() string {
	return self.Credential.Secret
}

func (self *CredentialOption) SetCredentialSecret(srt string) {
	self.Credential.Secret = srt
}

func NewCredentialManager(cli_fty *client_helper.ClientFactory, opt CredentialOptioner) (app_cred_mgr.ApplicationCredentialManager, error) {
	return app_cred_mgr.NewApplicationCredentialManager(cli_fty, opt.GetCredentialId(), opt.GetCredentialSecret())
}
