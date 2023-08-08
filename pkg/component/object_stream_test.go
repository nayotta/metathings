package metathings_component

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type ObjectStreamTestSuite struct {
	suite.Suite
}

func (ts *ObjectStreamTestSuite) TestReadOffset() {
	os := ts.newObjectStream()
	offset, err := os.Seek(1, io.SeekStart)
	ts.Require().Nil(err)
	ts.Require().Equal(int64(1), offset)

	offset = os.Offset()
	ts.Nil(err)
	ts.Equal(int64(1), offset)
}

func (ts *ObjectStreamTestSuite) TestSeek() {
	os := ts.newObjectStream()
	defer os.Close()
	n, err := os.Seek(100, io.SeekStart)
	ts.Nil(err)
	ts.Equal(int64(100), n)
	n = os.Offset()
	ts.Nil(err)
	ts.Equal(int64(100), n)
	ts.Equal(STATE_WRITABLE, os.state)
}

func (ts *ObjectStreamTestSuite) TestWrite() {
	buf := make([]byte, 1024)
	os := ts.newObjectStream()
	defer os.Close()
	_, err := os.Seek(0, io.SeekStart)
	ts.Require().Nil(err)
	n, err := os.Write(buf)
	ts.Nil(err)
	ts.Equal(1024, n)
	ts.Equal(STATE_READABLE, os.state)
	ts.Equal(int64(1024), os.Uploaded())
}

func (ts *ObjectStreamTestSuite) TestRead() {
	buf := make([]byte, 1024)
	os := ts.newObjectStream()
	defer os.Close()
	_, err := os.Seek(0, io.SeekStart)
	ts.Require().Nil(err)
	n, err := os.Write(buf)
	ts.Require().Nil(err)
	ts.Require().Equal(1024, n)

	buf1 := make([]byte, 1024)
	n, err = os.Read(buf1)
	ts.Equal(1024, n)
	ts.Nil(err)
	ts.Equal(buf, buf1)
}

func (ts *ObjectStreamTestSuite) TestAllInOne() {
	for _, st := range []struct {
		length      int64
		chunkLength int64
	}{
		{1024 * 1024, 512 * 1024},
		{1024 * 1024, 512},
		{1024*1024 + 42, 1024 * 1024},
		{1024*1024*1024 + 31, 16 * 1024 * 1024},
	} {
		src := make([]byte, st.length)
		n, err := rand.Read(src)
		ts.Require().Nil(err)
		ts.Require().Equal(st.length, int64(n))
		srcRd := bytes.NewReader(src)

		srcHash := sha1.New()
		srcHash.Write(src)
		srcSha1 := hex.EncodeToString(srcHash.Sum(nil))

		dstHash := sha1.New()

		os := ts.newObjectStream(WithLength(st.length), WithBufferLength(st.chunkLength))
		defer os.Close()

		for i := int64(0); i < st.length; i += st.chunkLength {
			sz := st.chunkLength
			if i+st.chunkLength > st.length {
				sz = st.length - i
			}

			srcBuf := make([]byte, sz)
			n, err := os.Seek(i, io.SeekStart)
			ts.Require().Nil(err)
			ts.Require().Equal(i, n)

			srcRd.Seek(i, io.SeekStart)
			srcRd.Read(srcBuf)

			n = os.Offset()
			ts.Equal(i, n)

			srcBufHash := sha1.New()
			srcBufHash.Write(srcBuf)
			srcBufSha1 := hex.EncodeToString(srcBufHash.Sum(nil))

			os.Write(srcBuf)
			dstBuf := make([]byte, sz)
			os.Read(dstBuf)
			dstBufHash := sha1.New()
			dstBufHash.Write(dstBuf)
			dstBufSha1 := hex.EncodeToString(dstBufHash.Sum(nil))

			ts.Require().Equal(srcBufSha1, dstBufSha1)

			dstHash.Write(dstBuf)
		}

		dstSha1 := hex.EncodeToString(dstHash.Sum(nil))

		ts.Equal(srcSha1, dstSha1)
		ts.Equal(st.length, os.Uploaded())
	}
}

func (ts *ObjectStreamTestSuite) TestUploadTimeout() {
	os := ts.newObjectStream(WithMaxAge(0))
	time.Sleep(1 * time.Millisecond)
	_, err := os.Seek(0, io.SeekStart)
	ts.Equal(ErrUploadTimeout, err)
}

func (ts *ObjectStreamTestSuite) newLogger() *logrus.Entry {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	return logrus.NewEntry(logger)
}

func (ts *ObjectStreamTestSuite) newObjectStream(opts ...NewObjectStreamOption) *objectStream {
	d := []NewObjectStreamOption{
		WithLogger(ts.newLogger()),
		WithName(""),
		WithSha1sum(""),
		WithMaxAge(1 * time.Hour),
		WithLength(1024 * 1024 * 1024),
		WithWaitTimeout(1 * time.Second),
	}
	os, err := NewObjectStream(append(d, opts...)...)
	ts.Require().Nil(err)
	return os.(*objectStream)
}

func TestObjectStreamTestSuite(t *testing.T) {
	suite.Run(t, new(ObjectStreamTestSuite))
}
