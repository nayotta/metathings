package metathings_deviced_connection

type AsyncBridgeWrapper interface {
	Send() chan []byte
	Recv() chan []byte
}

type AsyncBridgeWrapperImpl struct {
	br Bridge
}

func (self *AsyncBridgeWrapperImpl) Send() chan []byte {
	ch := make(chan []byte)
	go func() {
		buf := <-ch
		self.br.Send(buf)
	}()
	return ch
}

func (self *AsyncBridgeWrapperImpl) Recv() chan []byte {
	ch := make(chan []byte)
	go func() {
		buf, _ := self.br.Recv()
		ch <- buf
	}()
	return ch
}

func NewAsyncBridgeWrapper(br Bridge) AsyncBridgeWrapper {
	return &AsyncBridgeWrapperImpl{br: br}
}
