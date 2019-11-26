package cmd_contrib

import (
	"google.golang.org/grpc/credentials"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
)

type GetTransportCredentialOptioner interface {
	GetTransportCredential() TransportCredentialOptioner
}

type TransportCredentialOptioner interface {
	GetInsecureP() *bool
	GetInsecure() bool
	SetInsecure(bool)

	GetPlainTextP() *bool
	GetPlainText() bool
	SetPlainText(bool)

	GetCertFileP() *string
	GetCertFile() string
	SetCertFile(string)

	GetKeyFileP() *string
	GetKeyFile() string
	SetKeyFile(string)
}

type TransportCredentialOption struct {
	Insecure  bool
	PlainText bool   `mapstructure:"plain_text"`
	CertFile  string `mapstructure:"cert_file"`
	KeyFile   string `mapstructure:"key_file"`
}

func (self *TransportCredentialOption) GetInsecureP() *bool {
	return &self.Insecure
}

func (self *TransportCredentialOption) GetInsecure() bool {
	return self.Insecure
}

func (self *TransportCredentialOption) SetInsecure(insecure bool) {
	self.Insecure = insecure
}

func (self *TransportCredentialOption) GetPlainTextP() *bool {
	return &self.PlainText
}

func (self *TransportCredentialOption) GetPlainText() bool {
	return self.PlainText
}

func (self *TransportCredentialOption) SetPlainText(plaintext bool) {
	self.PlainText = plaintext
}

func (self *TransportCredentialOption) GetCertFileP() *string {
	return &self.CertFile
}

func (self *TransportCredentialOption) GetCertFile() string {
	return self.CertFile
}

func (self *TransportCredentialOption) SetCertFile(cert_file string) {
	self.CertFile = cert_file
}

func (self *TransportCredentialOption) GetKeyFileP() *string {
	return &self.KeyFile
}

func (self *TransportCredentialOption) GetKeyFile() string {
	return self.KeyFile
}

func (self *TransportCredentialOption) SetKeyFile(key_file string) {
	self.KeyFile = key_file
}

func NewClientTransportCredentials(opt TransportCredentialOptioner) (credentials.TransportCredentials, error) {
	return client_helper.NewClientTransportCredentials(
		opt.GetCertFile(), opt.GetKeyFile(),
		opt.GetPlainText(), opt.GetInsecure())
}

func NewServerTransportCredentials(opt TransportCredentialOptioner) (credentials.TransportCredentials, error) {
	return client_helper.NewServerTransportCredentials(opt.GetCertFile(), opt.GetKeyFile())
}
