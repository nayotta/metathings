package metathings_component

import (
	"errors"
	"io"
	"math/rand"
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
		OPTION_MAX_AGE:       301 * time.Second,
		OPTION_BUFFER_LENGTH: int64(4 * 1024 * 1024), // 4MiB
		OPTION_WAIT_TIMEOUT:  31 * time.Second,
	})
}

type ObjectStream interface {
	Name() string
	Sha1sum() string
	MaxAge() time.Duration
	Remained() time.Duration
	Length() int64
	BufferLength() int64
	Uploaded() int64
	Wait(string) error
	Offset() int64
	io.Writer
	io.Reader
	io.Seeker
	io.Closer
	CloseWithError(error) error
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
		closed:      make(chan struct{}),
		waitTimeout: waitTimeout,
	}

	os.registerState(STATE_SEEKABLE)
	os.registerState(STATE_READABLE)
	os.registerState(STATE_WRITABLE)

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

	bufferLength int64
	bufferSize   int
	buffer       []byte

	waits sync.Map

	logger      *logging.Entry
	opMtx       sync.Locker
	state       string
	closed      chan struct{}
	closeOnce   sync.Once
	err         error
	waitTimeout time.Duration
}

func (os *objectStream) Name() string          { return os.name }
func (os *objectStream) Sha1sum() string       { return os.sha1sum }
func (os *objectStream) MaxAge() time.Duration { return os.maxAge }
func (os *objectStream) Length() int64         { return os.length }
func (os *objectStream) BufferLength() int64   { return os.bufferLength }
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

	if err = os.wait(STATE_WRITABLE); err != nil {
		return 0, err
	}

	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	os.bufferSize = copy(os.buffer, b)
	atomic.AddInt64(&os.uploaded, int64(os.bufferSize))

	if err = os.notify(STATE_READABLE); err != nil {
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

	if err = os.wait(STATE_READABLE); err != nil {
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

	if err = os.notify(STATE_SEEKABLE); err != nil {
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

	if whence != io.SeekStart {
		err = ErrUnsupportedWhenceFn(whence)
		logger.WithError(err).Debugf("failed to seek")
		return 0, err
	}

	if err = os.wait(STATE_SEEKABLE); err != nil {
		logger.WithError(err).Debugf("failed to wait seekable")
		return 0, err
	}
	os.opMtx.Lock()
	defer os.opMtx.Unlock()

	atomic.StoreInt64(&os.offset, offset)

	if err = os.notify(STATE_WRITABLE); err != nil {
		logger.WithError(err).Debugf("failed to notify")
		return 0, err
	}

	logger.Tracef("seeked")

	return offset, nil
}

func (os *objectStream) Close() error { return os.close(nil) }

func (os *objectStream) CloseWithError(err error) error {
	return os.close(err)
}

func (os *objectStream) Wait(state string) error {
	return os.wait(state)
}

func (os *objectStream) Offset() int64 {
	return os.offset
}

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

func (os *objectStream) close(err error) error {
	os.closeOnce.Do(func() {
		logger := os.Logger().WithFields(logging.Fields{
			"#method": "close",
		})

		os.waits.Range(func(k, v any) bool {
			v.(*sync.Map).Range(func(k1, v1 any) bool {
				defer func() { recover() }()
				close(v1.(chan struct{}))
				return true
			})
			return true
		})

		close(os.closed)
		os.err = err

		logger.Tracef("closed")
	})
	return nil
}

func (os *objectStream) wait(state string) error {
	logger := os.Logger().WithFields(logging.Fields{
		"#method": "wait",
		"expect":  state,
		"current": os.state,
	})

	if os.isClosed() {
		logger.WithError(os.err).Debugf("closed")

		if os.err != nil {
			return os.err
		}
		return ErrClosed
	}

	if os.Remained() <= 0 {
		err := ErrUploadTimeout
		logger.WithError(err).Debugf("upload timeout")
		return err
	}

	if os.state == state {
		return nil
	}

	opable, cancel, err := os.lowWait(state)
	if err != nil {
		logger.WithError(err).Debugf("failed to low wait")
		return err
	}
	defer cancel()

	select {
	case <-os.closed:
		logger.Tracef("closed")

		if os.err != nil {
			return os.err
		}
		return ErrClosed
	case <-time.After(os.waitTimeout):
		logger.Tracef("wait timeout")
		return ErrWaitTimeout
	case _, ok := <-opable:
		if !ok {
			logger.Tracef("closed when waiting")
			if os.err != nil {
				return os.err
			}
			return ErrClosed
		}
		logger.Tracef("opable")
		return nil
	}
}

func (os *objectStream) notify(state string) error {
	logger := os.Logger().WithFields(logging.Fields{
		"#method": "notify",
		"state":   state,
	})

	if os.isClosed() {
		logger.WithError(os.err).Debugf("closed")

		if os.err != nil {
			return os.err
		}
		return ErrClosed
	}

	v, ok := os.waits.Load(state)
	if !ok {
		return ErrUnregisteredState
	}
	v.(*sync.Map).Range(func(k1, v1 any) bool {
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				if !errors.Is(err, ErrSendOnClosedChannel) {
					panic(err)
				}
			}
		}()
		v1.(chan struct{}) <- struct{}{}
		return true
	})

	os.state = state
	logger.Tracef("notify")

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

func (os *objectStream) registerState(s string) {
	os.waits.Store(s, &sync.Map{})
}

func (os *objectStream) lowWait(state string) (chan struct{}, func(), error) {
	v, ok := os.waits.Load(state)
	if !ok {
		return nil, nil, ErrUnregisteredState
	}

	m := v.(*sync.Map)
	t := rand.Int63()
	ch := make(chan struct{})
	c := func() {
		defer func() { recover() }()
		m.Delete(t)
		close(ch)
	}

	m.Store(t, ch)

	if os.state == state {
		go func() { ch <- struct{}{} }()
	}

	return ch, c, nil
}
