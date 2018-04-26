package cmd

type _serviceConfigOptions struct {
	Identityd _identitydServiceOptions
}

type _identitydServiceOptions struct {
	Address string
}

type _storageOptions struct {
	Driver string
	Uri    string
}
