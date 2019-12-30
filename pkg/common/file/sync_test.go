package file_helper

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FileSyncerTestSuite struct {
	fs       *FileSyncer
	src      string
	src_sha1 string
	dst      string
	suite.Suite
}

func (s *FileSyncerTestSuite) sha1_hash(path string) string {
	fp, err := os.Open(path)
	if err != nil {
		s.Failf("failed to sha1 data: %v", path)
	}
	defer fp.Close()

	hp := sha1.New()
	if _, err = io.Copy(hp, fp); err != nil {
		s.Failf("failed to sha1 data: %v", path)
	}

	return fmt.Sprintf("%x", hp.Sum(nil))
}

func (s *FileSyncerTestSuite) SetupTest() {
	var data_size int64 = 64*1024*1024 + 13
	data := make([]byte, data_size)
	rand.Read(data)
	src, err := ioutil.TempFile("", "fssrc")
	if err != nil {
		s.Fail("failed to create temp file")
	}

	s.src = src.Name()
	_, err = src.Write(data)
	if err != nil {
		s.Fail("failed to write test data")
	}
	src.Sync()
	defer src.Close()

	s.src_sha1 = s.sha1_hash(s.src)

	dst, err := ioutil.TempFile("", "fsdst")
	if err != nil {
		s.Fail("failed to create temp file")
	}
	s.dst = dst.Name()
	defer dst.Close()

	s.fs, err = NewFileSyncer(
		SetPath(s.dst),
		SetSize(data_size),
		SetSha1Hash(s.src_sha1),
	)
	if err != nil {
		s.Fail("failed to new file syncer")
	}
}

func (s *FileSyncerTestSuite) AfterTest(suiteName, testName string) {
	err := os.Remove(s.src)
	if err != nil {
		s.Fail("failed to remove src data")
	}

	err = os.Remove(s.dst)
	if err != nil {
		s.Fail("failed to remove dst data")
	}

	err = s.fs.Close()
	if err != nil {
		s.Fail("failed to close file syncer")
	}
}

func (s *FileSyncerTestSuite) TestSync() {
	fp, err := os.Open(s.src)
	s.Nil(err)
	defer fp.Close()
	data := make([]byte, DEFAULT_CHUNK_SIZE)
	var n int
_outer_loop:
	for {
		offsets, err := s.fs.Next(3)
		s.Nil(err)
		for _, offset := range offsets {
			fp.Seek(offset, 0)
			n, err = fp.Read(data)
			s.Nil(err)
			if err = s.fs.Sync(offset, data, n); err == DONE {
				break _outer_loop
			}
		}
	}
}

func TestFileSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(FileSyncerTestSuite))
}
