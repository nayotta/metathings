package metathings_deviced_connection

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type dummyBridge struct {
	pipe chan []byte
}

func (self *dummyBridge) Id() string {
	return "dummy"
}

func (self *dummyBridge) Send(buf []byte) error {
	self.pipe <- buf
	return nil
}

func (self *dummyBridge) Recv() ([]byte, error) {
	buf := <-self.pipe
	return buf, nil
}

func newDummyBridge() Bridge {
	return &dummyBridge{
		pipe: make(chan []byte, 16),
	}
}

type asyncBridgeWrapperTestSuite struct {
	suite.Suite
	br  Bridge
	abr AsyncBridgeWrapper
}

func (self *asyncBridgeWrapperTestSuite) SetupTest() {
	self.br = newDummyBridge()
	self.abr = NewAsyncBridgeWrapper(self.br)
}

func (self *asyncBridgeWrapperTestSuite) TestSend() {
	dat := []byte("hello")
	self.abr.Send() <- dat
	buf, err := self.br.Recv()
	self.Nil(err)
	self.Equal(dat, buf)
}

func (self *asyncBridgeWrapperTestSuite) TestRecv() {
	dat := []byte("hello")
	self.br.Send(dat)
	buf := <-self.abr.Recv()
	self.Equal(dat, buf)
}

func TestAsyncBridgeWrapperTestSuite(t *testing.T) {
	suite.Run(t, new(asyncBridgeWrapperTestSuite))
}
