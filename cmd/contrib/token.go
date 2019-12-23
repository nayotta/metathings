package cmd_contrib

type TokenOptioner interface {
	GetTokenP() *string
	GetToken() string
	SetToken(string)
}

type TokenOption struct {
	Token string
}

func (self *TokenOption) GetTokenP() *string {
	return &self.Token
}

func (self *TokenOption) GetToken() string {
	return self.Token
}

func (self *TokenOption) SetToken(tkn string) {
	self.Token = tkn
}
