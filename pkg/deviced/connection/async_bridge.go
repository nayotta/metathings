package metathings_deviced_connection

type AsyncBridgeWrapper interface {
	Send() chan []byte
	Recv() chan []byte
	Err() error
	Close() error
}

type AsyncBridgeWrapperImpl struct {
	br      Bridge
	sender  chan []byte
	reciver chan []byte
	err     error
}

func (self *AsyncBridgeWrapperImpl) Send() chan []byte {
	go func() {
		buf := <-self.sender
		self.err = self.br.Send(buf)
	}()
	return self.sender
}

func (self *AsyncBridgeWrapperImpl) Recv() chan []byte {
	go func() {
		var buf []byte
		buf, self.err = self.br.Recv()
		if self.err != nil {
			self.reciver <- []byte{}
			return
		}
		self.reciver <- buf
	}()
	return self.reciver
}

func (self *AsyncBridgeWrapperImpl) Err() error {
	return self.err
}

func (self *AsyncBridgeWrapperImpl) Close() error {
	close(self.sender)
	close(self.reciver)
	return nil
}

func NewAsyncBridgeWrapper(br Bridge) AsyncBridgeWrapper {
	abr := &AsyncBridgeWrapperImpl{
		br:      br,
		sender:  make(chan []byte, 16),
		reciver: make(chan []byte, 16),
	}

	return abr
}
