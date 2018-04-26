package cmd_helper

type LogOptions struct {
	Level string
}

type ApplicationCredentialOptions struct {
	Id     string
	Secret string
}

type TokenOptions struct {
	Token string
}

type RootOptions struct {
	Config                string
	Stage                 string
	Verbose               bool
	Log                   LogOptions
	ApplicationCredential ApplicationCredentialOptions `mapstructure:"application_credential"`
}
