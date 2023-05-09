package metathings_component

import (
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PeerXu/option-go"
	logging "github.com/sirupsen/logrus"

	option_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	STATE_SEEKABLE = "seekable"
	STATE_READABLE = "readable"
	STATE_WRITABLE = "writable"
)

func NewDefaultObjectStreamOption() option.Option {
	return option.NewOption(map[string]any{
		OPTION_BUFFER_LENGTH: 4 * 1024 * 1024, // 4MiB
		OPTION_WAIT_TIMEOUT:  31 * time.Second,
	})
}

type ObjectStream interface {
	Name() string
	Sha1sum() string
	MaxAge() time.Duration
	Remained() time.Duration
	Length() int64
	Uploaded() int64
	io.Writer
	io.Reader
	io.Seeker
	io.Closer
}

func NewObjectStream(opts ...NewObjectStreamOption) (ObjectStream, error) {
	o := option.ApplyWithDefault(NewDefaultObjectStreamOption(), opts...)

	logger, err := option_helper.GetLogger(o)
	if err != nil {
		return nil, err
	}

	name, err := GetName(o)
	if err != nil {
		return nil, err
	}

	sha1sum, err := GetSha1sum(o)
	if err != nil {
		return nil, err
	}

	maxAge, err := GetMaxAge(o)
	if err != nil {
		return nil, err
	}

	length, err := GetLength(o)
	if err != nil {
		return nil, err
	}

	bufferLength, err := GetBufferLength(o)
	if err != nil {
		return nil, err
	}

	waitTimeout, err := GetWaitTimeout(o)
	if err != nil {
		return nil, err
	}

	os := &objectStream{
		name:    name,
		sha1sum: sha1sum,
		maxAge:  maxAge,
		length:  length,

		offset:    0,
		uploaded:  0,
		createdAt: time.Now(),

		bufferLength: bufferLength,
		buffer:       make([]byte, bufferLength),

		logger:      logger,
		opMtx:       &sync.Mutex{},
		state:       STATE_SEEKABLE,
		seekCh:      make(chan struct{}, 1),
		readCh:      make(chan struct{}, 1),
		writeCh:     make(chan struct{}, 1),
		closed:      make(chan struct{}),
		waitTimeout: waitTimeout,
	}

	os.seekCh <- struct{}{}

	return os, nil
}

type objectStream struct {
	name      string
	sha1sum   string
	maxAge    time.Duration
	length    int64
	offset    int64
	uploaded  int64
	createdAt time.Time

	bufferLength int
	bufferSize   int
	buffer       []byte

	logger      *logging.Entry
	opMtx       sync.Locker
	state       string
	seekCh      chan struct{}
	readCh      chan struct{}
	writeCh     chan struct{}
	closed      chan struct{}
	closeOnce   sync.Once
	waitTimeout time.Duration
}

func (os *objectStream) Name() string          { return os.name }
func (os *objectStream) Sha1sum() string       { return os.sha1sum }
func (os *objectStream) MaxAge() time.Duration { return os.maxAge }
func (os *objectStream) Length() int64         { return os.length }
func (os *objectStream) Uploaded() int64       { return os.uploaded }

func (os *objectStream) Remained() time.Duration {
	remained := os.maxAge - time.Now().Sub(os.createdAt)
	if remained < 0 {
		remained = 0
	}
	return remained
}

func (os *objectStream) Write(b []byte) (int, error) {
	var err error

	logger := os.Logger().WithFields(logging.Fields{
		"#method": "Write",
	})

	if err = os.wait(os.writable); err != nil {
		return 0, err
	}
	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	os.bufferSize = copy(os.buffer, b)
	atomic.AddInt64(&os.uploaded, int64(os.bufferSize))

	if err = os.notify(STATE_READABLE, os.readable); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("write")

	return os.bufferSize, nil
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

	n = copy(b, os.buffer[:os.bufferSize])
	if n != os.bufferSize {
		err = ErrInvalidBuffer
		logger.WithError(err).Debugf("invalid argument")
		return 0, err
	}

	if err = os.notify(STATE_SEEKABLE, os.seekable); err != nil {
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

	if err = os.notify(STATE_WRITABLE, os.writable); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("seeked")

	return offset, nil
}

func (os *objectStream) Close() error { return os.close() }

func (os *objectStream) Logger() *logging.Entry {
	return os.logger.WithFields(logging.Fields{
		"#instance": "objectStream",
		"name":      os.Name(),
		"length":    os.Length(),
		"uploaded":  os.Uploaded(),
		"sha1sum":   os.Sha1sum(),
		"maxAge":    os.MaxAge(),
		"remained":  os.Remained(),
		"offset":    os.offset,
		"createdAt": os.createdAt,
		"state":     os.state,
	})
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
	if os.Remained() <= 0 {
		return ErrUploadTimeout
	}

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
