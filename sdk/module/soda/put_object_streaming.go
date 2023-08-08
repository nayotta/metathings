package metathings_module_soda_sdk

import (
	"io"

	"github.com/sirupsen/logrus"
)

func (cli *sodaClient) PutObjectStreaming(name string, src io.ReadSeeker, length int64, opts PutObjectStreamingOption) error {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method":  "PutObjectStreaming",
		"fileName": name,
		"length":   length,
		"sha1sum":  opts.Sha1sum,
	})

	osName, err := cli.LLPutObjectStreaming(name, length, opts.Sha1sum)
	if err != nil {
		logger.WithError(err).Debugf("failed to put object streaming")
		return err
	}
	logger = logger.WithField("name", osName)

	remained, offset, length, err := cli.LLObjectStreamNextChunk(osName)
	if err != nil {
		logger.WithError(err).Debugf("failed to get object stream next chunk")
		return err
	}
	src.Seek(offset, io.SeekStart)
	for {
		if remained < 0 {
			err = ErrPutObjectTimeout
			logger.WithError(err).Debugf("put object streaming timeout")
			return err
		}

		buf := make([]byte, length)
		n, err := src.Read(buf)
		if err != nil {
			logger.WithError(err).Debugf("failed to read from source")
			return err
		}

		err = cli.LLObjectStreamWriteChunk(osName, offset, int64(n), buf[:n])
		if err != nil {
			logger.WithError(err).Debugf("failed to write chunk to object stream")
			return err
		}

		remained, offset, length, err = cli.LLObjectStreamNextChunk(osName)
		if err != nil {
			logger.WithError(err).Debugf("failed to get object stream next chunk")
			return err
		}

		if offset != offset {
			offset, err = src.Seek(offset, io.SeekStart)
			if err != nil {
				logger.WithError(err).Debugf("failed to seek source")
				return err
			}
		}
	}
}
