package cmd

type _serviceConfigOptions struct {
	Identityd _identitydServiceOptions
}

type _identitydServiceOptions struct {
	Address string
}

type _applicationCredentialOptions struct {
	Id     string
	Secret string
}

type _logOptions struct {
	Level string
}
