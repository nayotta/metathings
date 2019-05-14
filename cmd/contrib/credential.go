package cmd_contrib

type CredentialOptioner interface {
	GetCredentialDomainP() *string
	GetCredentialDomain() string
	SetCredentialDomain(string)

	GetCredentialIdP() *string
	GetCredentialId() string
	SetCredentialId(string)

	GetCredentialSecretP() *string
	GetCredentialSecret() string
	SetCredentialSecret(string)
}

type CredentialOption struct {
	Credential struct {
		Domain string
		Id     string
		Secret string
	}
}

func (self *CredentialOption) GetCredentialDomainP() *string {
	return &self.Credential.Domain
}

func (self *CredentialOption) GetCredentialDomain() string {
	return self.Credential.Domain
}

func (self *CredentialOption) SetCredentialDomain(dom string) {
	self.Credential.Domain = dom
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
