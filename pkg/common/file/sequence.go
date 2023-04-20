package file_helper

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
)

type SequenceFileSyncerOption struct {
	size       int64
	sha1_hash  string
	chunk_size int64
}

type SequenceFileSyncer struct {
	opt           *SequenceFileSyncerOption
	writeCloser   io.WriteCloser
	sha1Hash      hash.Hash
	currentOffset int64
}

func NewSequenceFileSyncer(writeCloser io.WriteCloser, size int64, sha1_hash string, chunk_size int64) *SequenceFileSyncer {
	return &SequenceFileSyncer{
		opt: &SequenceFileSyncerOption{
			size:       size,
			sha1_hash:  sha1_hash,
			chunk_size: chunk_size,
		},
		writeCloser:   writeCloser,
		sha1Hash:      sha1.New(),
		currentOffset: 0,
	}
}

// HACK: ignore count
func (sfs *SequenceFileSyncer) Next(_ int) (offsets []int64, err error) {
	if sfs.is_done() {
		return nil, DONE
	}

	offsets = []int64{sfs.currentOffset}
	sfs.currentOffset += sfs.opt.chunk_size

	return offsets, nil
}

// TODO: verify chunk hash
func (sfs *SequenceFileSyncer) Sync(offset int64, data []byte, size int) (err error) {
	_, err = sfs.writeCloser.Write(data[:size])
	if err != nil {
		return
	}

	_, err = sfs.sha1Hash.Write(data[:size])
	if err != nil {
		return
	}

	if sfs.is_done() {
		defer sfs.Close()

		if fmt.Sprintf("%x", sfs.sha1Hash.Sum(nil)) != sfs.opt.sha1_hash {
			return ErrHashNotMatch
		}

		return DONE
	}

	return
}

func (sfs *SequenceFileSyncer) Close() error {
	return sfs.writeCloser.Close()
}

func (sfs *SequenceFileSyncer) is_done() bool {
	return sfs.currentOffset >= sfs.opt.size
}
