package cmd_contrib

import "google.golang.org/grpc/credentials"

type TransportCredentialOptioner interface {
	GetCertFileP() *string
	GetCertFile() string
	SetCertFile(string)

	GetKeyFileP() *string
	GetKeyFile() string
	SetKeyFile(string)
}

type TransportCredentialOption struct {
	CertFile string
	KeyFile  string
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
	}
	return nil, nil
}
