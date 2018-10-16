package cmd_contrib

import "net"

type ListenOptioner interface {
	GetListenP() *string
	GetListen() string
	SetListen(string)
}

type ListenOption struct {
	Listen string
}

func (self *ListenOption) GetListenP() *string {
	return &self.Listen
}

func (self *ListenOption) GetListen() string {
	return self.Listen
}

func (self *ListenOption) SetListen(lis string) {
	self.Listen = lis
}

func NewListener(opt ListenOptioner) (net.Listener, error) {
	return net.Listen("tcp", *opt.GetListenP())
}
