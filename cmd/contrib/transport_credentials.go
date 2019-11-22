package cmd_contrib

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
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
	PlainText bool
	CertFile  string
	KeyFile   string
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

func NewTransportCredentials(opt TransportCredentialOptioner) (credentials.TransportCredentials, error) {
	cert_file := opt.GetCertFile()
	key_file := opt.GetKeyFile()
	if cert_file != "" && key_file != "" {
		return credentials.NewServerTLSFromFile(cert_file, key_file)
	} else if opt.GetInsecure() {
		return credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		}), nil
	} else if opt.GetPlainText() {
		return nil, nil
	}
	return credentials.NewTLS(nil), nil
}
