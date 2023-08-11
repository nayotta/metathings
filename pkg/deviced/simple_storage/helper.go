package metathings_deviced_simple_storage

import (
	"io"
	"math/rand"
	"os"
	"sync/atomic"

	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type Reader struct {
	*io.PipeReader
	logger logging.FieldLogger
	cnt    atomic.Int64
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.PipeReader.Read(p)
	r.logger.WithFields(logging.Fields{
		"#method": "Read",
		"size":    n,
		"err":     err,
		"count":   r.cnt.Add(1),
	}).Tracef("read")
	return n, err
}

type Writer struct {
	*io.PipeWriter
	logger logging.FieldLogger
	cnt    atomic.Int64
}

func (w *Writer) Write(p []byte) (int, error) {
	n, err := w.PipeWriter.Write(p)
	w.logger.WithFields(logging.Fields{
		"#method": "Write",
		"size":    n,
		"err":     err,
		"count":   w.cnt.Add(1),
	}).Tracef("write")
	return n, err
}

func WrapedPipe(logger logging.FieldLogger) (io.ReadCloser, io.WriteCloser) {
	rd, wr := io.Pipe()
	return wrapReader(rd, logger), wrapWriter(wr, logger)
}

var wrapReader func(*io.PipeReader, logging.FieldLogger) io.ReadCloser
var wrapWriter func(*io.PipeWriter, logging.FieldLogger) io.WriteCloser

func init() {
	s := os.Getenv("MT_DEBUG_ENABLE_OBJECT_STREAM_TRACING")
	enableObjectStreamTracing := cast.ToBool(s)
	if enableObjectStreamTracing {
		sess := rand.Int31()
		wrapReader = func(rd *io.PipeReader, logger logging.FieldLogger) io.ReadCloser {
			return &Reader{
				PipeReader: rd,
				logger: logger.WithFields(logging.Fields{
					"#instance": "Reader",
					"session":   sess,
				}),
			}
		}
		wrapWriter = func(wr *io.PipeWriter, logger logging.FieldLogger) io.WriteCloser {
			return &Writer{
				PipeWriter: wr,
				logger: logger.WithFields(logging.Fields{
					"#instance": "Writer",
					"session":   sess,
				}),
			}
		}
	} else {
		wrapReader = func(rd *io.PipeReader, logger logging.FieldLogger) io.ReadCloser { return rd }
		wrapWriter = func(wr *io.PipeWriter, logger logging.FieldLogger) io.WriteCloser { return wr }
	}
}
