package file_helper

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path"
	"sync"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
)

const (
	DEFAULT_CHUNK_SIZE int64 = 512 * 1024
)

type SetFileSyncerOption func(*FileSyncerOption)

type FileSyncerOption struct {
	path       string
	size       int64
	cache_path string
	sha1_hash  string
	chunk_size int64
}

func SetPath(path string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.path = path
	}
}

func SetSize(size int64) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.size = size
	}
}

func SetCachePath(path string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.cache_path = path
	}
}

func SetSha1Hash(hash string) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.sha1_hash = hash
	}
}

func SetChunkSize(size int64) SetFileSyncerOption {
	return func(o *FileSyncerOption) {
		o.chunk_size = size
	}
}

func NewFileSyncerOption() *FileSyncerOption {
	o := &FileSyncerOption{}
	o.chunk_size = DEFAULT_CHUNK_SIZE
	return o
}

type FileSyncer struct {
	opt *FileSyncerOption
	db  sync.Map
	fp  *os.File

	stat struct {
		Chunks int64
	}
}

func (fs *FileSyncer) init_db() error {
	var i int64
	for i = 0; i*fs.opt.chunk_size < fs.opt.size; i++ {
		fs.db.Store(i*fs.opt.chunk_size, "")
	}

	fs.stat.Chunks = i

	return nil
}

func (fs *FileSyncer) Close() (err error) {
	if _, err = os.Stat(fs.opt.cache_path); err == nil {
		return os.Remove(fs.opt.cache_path)
	}

	return nil
}

func (fs *FileSyncer) is_done() bool {
	done := true
	fs.db.Range(func(key, val interface{}) bool {
		done = false
		return false
	})
	return done
}

func (fs *FileSyncer) debug() {
	logger := log_helper.NewDebugLogger()
	var wtchks int64
	fs.db.Range(func(_, _ interface{}) bool {
		wtchks++
		return true
	})

	logger.WithField("sum", fmt.Sprintf("%v/%v", fs.stat.Chunks-wtchks, fs.stat.Chunks)).Debugf("chunk state")
}

func (fs *FileSyncer) Next(batch int) (offsets []int64, err error) {
	if batch < 0 {
		return nil, ErrInvalidArgument
	}

	if fs.is_done() {
		return nil, DONE
	}

	fs.db.Range(func(key, val interface{}) bool {
		if batch <= 0 {
			return false
		}
		batch -= 1
		offsets = append(offsets, key.(int64))
		return true
	})

	return offsets, nil
}

func (fs *FileSyncer) Sync(offset int64, data []byte, size int) (err error) {
	if fs.fp == nil {
		if fs.fp, err = os.OpenFile(fs.opt.cache_path, os.O_WRONLY, 0); err != nil {
			return err
		}
	}

	if _, err = fs.fp.Seek(offset, 0); err != nil {
		return err
	}

	if _, err = fs.fp.Write(data[:size]); err != nil {
		return err
	}

	fs.db.Delete(offset)

	if fs.is_done() {
		defer fs.Close()
		if err = fs.post_sync(); err != nil {
			return err
		}

		return DONE
	}

	return nil
}

func (fs *FileSyncer) post_sync() (err error) {
	if err = fs.validate(fs.opt.cache_path); err != nil {
		return err
	}

	if err = os.Rename(fs.opt.cache_path, fs.opt.path); err != nil {
		return err
	}

	return nil
}

func (fs *FileSyncer) validate(path string) error {
	fp, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrHashNotMatch
		}

		return err
	}
	defer fp.Close()

	hp := sha1.New()
	if _, err = io.Copy(hp, fp); err != nil {
		return err
	}

	if fs.opt.sha1_hash != fmt.Sprintf("%x", hp.Sum(nil)) {
		return ErrHashNotMatch
	}

	return nil
}

func (fs *FileSyncer) Validate() error {
	return fs.validate(fs.opt.path)
}

func (fs *FileSyncer) create_empty_cache_file() error {
	f, err := os.Create(fs.opt.cache_path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = f.Truncate(fs.opt.size)
	if err != nil {
		os.Remove(fs.opt.cache_path)
		return err
	}

	return nil
}

func (fs *FileSyncer) Initialize() error {
	if fs.opt.cache_path == "" {
		fs.opt.cache_path = path.Join(path.Dir(fs.opt.path), "."+path.Base(fs.opt.path))
	}

	if err := fs.init_db(); err != nil {
		return err
	}

	if _, err := os.Stat(fs.opt.cache_path); err != nil {
		if os.IsNotExist(err) {
			if err = fs.create_empty_cache_file(); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func NewFileSyncer(opts ...SetFileSyncerOption) (*FileSyncer, error) {
	o := NewFileSyncerOption()
	for _, opt := range opts {
		opt(o)
	}

	fs := &FileSyncer{opt: o}
	if err := fs.Validate(); err != nil {
		if err != ErrHashNotMatch {
			return nil, err
		}

		if err = fs.Initialize(); err != nil {
			return nil, err
		}
	}

	return fs, nil
}
