package metathings_component

import (
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"

	logging "github.com/sirupsen/logrus"
)

type ObjectStream interface {
	Name() string
	Sha1sum() string
	MaxAge() time.Duration
	Remained() time.Duration
	Length() int64
	Uploaeded() int64
	io.Writer
	io.Reader
	io.Seeker
	io.Closer
}

func NewObjectStream() (ObjectStream, error) {
	panic("unimplemented")
}

type objectStream struct {
	offset int64

	uploaded     int64
	buffer       []byte
	bufferSize   int
	bufferLength int

	opMtx       sync.Locker
	state       string
	seekCh      chan struct{}
	readCh      chan struct{}
	writeCh     chan struct{}
	closed      chan struct{}
	closeOnce   sync.Once
	waitTimeout time.Duration
}

func (os *objectStream) Name() string            { panic("unimplemented") }
func (os *objectStream) Sha1sum() string         { panic("unimplemented") }
func (os *objectStream) MaxAge() time.Duration   { panic("unimplemented") }
func (os *objectStream) Remained() time.Duration { panic("unimplemented") }
func (os *objectStream) Length() int64           { panic("unimplemented") }
func (os *objectStream) Uploaded() int64         { panic("unimplemented") }

func (os *objectStream) Write(b []byte) (n int, err error) {
	logger := os.Logger().WithFields(logging.Fields{
		"#method": "Write",
	})

	if err = os.wait(os.writable); err != nil {
		return 0, err
	}
	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	n = copy(os.buffer, b)
	atomic.AddInt64(&os.uploaded, int64(n))

	if err = os.notify("readable", os.readable); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("write")

	return n, nil
}

func (os *objectStream) Read(b []byte) (n int, err error) {
	logger := os.Logger().WithFields(logging.Fields{
		"#method": "Read",
	})

	if err = os.wait(os.readable); err != nil {
		return 0, err
	}
	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	n = copy(b, os.buffer[:os.bufferLength])
	if n != os.bufferLength {
		err = ErrInvalidBuffer
		logger.WithError(err).Debugf("invalid argument")
		return 0, err
	}

	if err = os.notify("seekable", os.seekable); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("read")

	return n, nil
}

func (os *objectStream) Seek(offset int64, whence int) (n int64, err error) {
	logger := os.Logger().WithFields(logging.Fields{
		"#method": "Seek",
		"offset":  offset,
		"whence":  whence,
	})

	if offset == 0 && whence == io.SeekCurrent {
		return os.offset, nil
	}

	if whence != io.SeekStart {
		err = ErrOnlySupportSeekStartWhenOffsetNotEqual0
		logger.WithError(err).Debugf("invalid argument")
		return 0, err
	}

	if err = os.wait(os.seekable); err != nil {
		logger.WithError(err).Debugf("failed to wait seekable")
		return 0, err
	}
	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	atomic.StoreInt64(&os.offset, offset)

	if err = os.notify("writable", os.writable); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("seeked")

	return offset, nil
}

func (os *objectStream) Close() error { return os.close() }

func (os *objectStream) Logger() *logging.Entry {
	panic("unimplemented")
}

func (os *objectStream) close() error {
	os.closeOnce.Do(func() {
		logger := os.Logger().WithFields(logging.Fields{
			"#method": "close",
		})

		close(os.seekCh)
		close(os.readCh)
		close(os.writeCh)

		close(os.closed)

		logger.Tracef("closed")
	})
	return nil
}

func (os *objectStream) wait(opable func() chan struct{}) error {
	select {
	case <-os.closed:
		return ErrClosed
	case <-time.After(os.waitTimeout):
		return ErrWaitTimeout
	case <-opable():
		return nil
	}
}

func (os *objectStream) notify(state string, opable func() chan struct{}) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
			if errors.Is(err, ErrSendOnClosedChannel) {
				err = ErrClosed
			}
		}
	}()

	if os.isClosed() {
		return ErrClosed
	}

	opable() <- struct{}{}
	os.state = state

	return nil
}

func (os *objectStream) isClosed() bool {
	select {
	case <-os.closed:
		return true
	default:
		return false
	}
}

func (os *objectStream) seekable() chan struct{} { return os.seekCh }
func (os *objectStream) readable() chan struct{} { return os.readCh }
func (os *objectStream) writable() chan struct{} { return os.writeCh }
