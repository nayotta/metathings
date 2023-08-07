package metathings_deviced_simple_storage

import (
	"io"
	"math/rand"
	"sync/atomic"

	logging "github.com/sirupsen/logrus"
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

func WrapedPipe(logger logging.FieldLogger) (*Reader, *Writer) {
	sess := rand.Int31()
	rd, wr := io.Pipe()
	return &Reader{
			PipeReader: rd,
			logger: logger.WithFields(logging.Fields{
				"#instance": "Reader",
				"session":   sess,
			}),
		}, &Writer{
			PipeWriter: wr,
			logger: logger.WithFields(logging.Fields{
				"#instance": "Writer",
				"session":   sess,
			}),
		}
}
