package binary_synchronizer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type BinarySynchronizer interface {
	Sync(ctx context.Context, src_filepath, dst_filename, uri, sha256sum string) error
}

type BinarySynchronizerOption struct {
	TempPath       string
	SyncTimeout    time.Duration
	ForceRelink    bool
	IgnoreChecksum bool
}

func NewBinarySynchronizerOption() *BinarySynchronizerOption {
	return &BinarySynchronizerOption{
		TempPath:       os.TempDir(),
		SyncTimeout:    300 * time.Second,
		ForceRelink:    false,
		IgnoreChecksum: false,
	}
}

type binarySynchronizer struct {
	opt    *BinarySynchronizerOption
	logger logrus.FieldLogger
}

func (bs *binarySynchronizer) Sync(ctx context.Context, src_filepath, dst_filename, uri, sha256sum string) error {
	var err error
	var cancel context.CancelFunc
	var real_source_filepath string

	logger := bs.logger.WithFields(logrus.Fields{
		"source":      src_filepath,
		"destination": dst_filename,
		"uri":         uri,
		"sha256sum":   sha256sum,
	})

	real_source_filepath, err = os.Readlink(src_filepath)
	if err != nil {
		if !bs.opt.ForceRelink {
			logger.WithError(err).Debugf("failed to readlink")
			return err
		}

		real_source_filepath = src_filepath
		logger.Debugf("source file not symlink")
	}

	cli := grab.NewClient()
	if bs.opt.SyncTimeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, bs.opt.SyncTimeout)
		defer cancel()
	}
	req, err := grab.NewRequest(bs.opt.TempPath, uri)
	if err != nil {
		logger.WithError(err).Debugf("failed to new download request")
		return err
	}

	if !bs.opt.IgnoreChecksum {
		sum, err := hex.DecodeString(sha256sum)
		if err != nil {
			logger.WithError(err).Debugf("failed to decode sha256sum")
			return err
		}

		req.SetChecksum(sha256.New(), sum, true)
	}

	req = req.WithContext(ctx)
	res := cli.Do(req)
	if err = res.Err(); err != nil {
		defer os.Remove(res.Filename)
		logger.WithError(err).Debugf("failed to download file")
		return err
	}

	if err = os.Chmod(res.Filename, 0755); err != nil {
		logger.WithError(err).Debugf("failed to chmod downloaded file")
		return err
	}

	real_source_path := filepath.Dir(real_source_filepath)
	dst_filepath := filepath.Join(real_source_path, dst_filename)

	if err = os.Rename(res.Filename, dst_filepath); err != nil {
		defer os.Remove(res.Filename)
		logger.WithError(err).Debugf("failed to rename downloaded file to destination")
		return err
	}

	if err = os.Remove(src_filepath); err != nil {
		defer os.Remove(dst_filepath)
		logger.WithError(err).Debugf("failed to remove source file")
		return err
	}

	if err = os.Symlink(dst_filepath, src_filepath); err != nil {
		logger.WithError(err).Debugf("failed to symlink")
		return err
	}

	if src_filepath != real_source_filepath {
		if err = os.Remove(real_source_filepath); err != nil {
			logger.WithError(err).Debugf("failed to remove remove real source filepath")
		}
	}

	logger.Debugf("binary synchronized")

	return nil
}

func NewBinarySynchronizer(args ...interface{}) (BinarySynchronizer, error) {
	var logger logrus.FieldLogger
	opt := NewBinarySynchronizerOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":          opt_helper.ToLogger(&logger),
		"temp_path":       opt_helper.ToString(&opt.TempPath),
		"sync_timeout":    opt_helper.ToDuration(&opt.SyncTimeout),
		"force_relink":    opt_helper.ToBool(&opt.ForceRelink),
		"ignore_checksum": opt_helper.ToBool(&opt.IgnoreChecksum),
	})(args...); err != nil {
		return nil, err
	}

	return &binarySynchronizer{
		opt:    opt,
		logger: logger,
	}, nil
}
