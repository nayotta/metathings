package file_helper

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SequenceFileSyncerTestSuite struct {
	suite.Suite
}

func (s *SequenceFileSyncerTestSuite) runCase1(bufferSize int64, chunkSize int64) {
	data := make([]byte, bufferSize)
	n, err := rand.Read(data)
	s.Require().Nil(err)
	s.Require().Equal(bufferSize, int64(n))

	h := sha1.New()
	h.Write(data)
	sha1Str := fmt.Sprintf("%x", h.Sum(nil))

	rd, wr := net.Pipe()
	sfs := NewSequenceFileSyncer(wr, bufferSize, sha1Str, chunkSize)

	var wg sync.WaitGroup
	var bb bytes.Buffer

	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := io.Copy(&bb, rd)
		s.Require().Nil(err)
		s.Require().Equal(bufferSize, total)
	}()

	for {
		ofs, err := sfs.Next(0)
		if err != nil {
			s.Require().Equal(DONE, err)
			break
		}

		for _, of := range ofs {
			chkSz := chunkSize
			if of+chkSz > bufferSize {
				chkSz = bufferSize - of
			}
			err = sfs.Sync(of, data[of:of+chkSz], int(chkSz))
			if err != nil {
				s.Require().Equal(DONE, err)
				break
			}
		}
	}
	sfs.Close()

	wg.Wait()

	data1 := bb.Bytes()
	h1 := sha1.New()
	h1.Write(data1)
	sha1Str2 := fmt.Sprintf("%x", h1.Sum(nil))
	s.Equal(sha1Str, sha1Str2)
	s.T().Logf("bufSz=%v,chkSz=%v,sha1(data)=%v,sha1(dataFromReader)=%v", bufferSize, chunkSize, sha1Str, sha1Str2)
}

func (s *SequenceFileSyncerTestSuite) TestCase1() {
	for _, st := range []struct {
		bufferSize int64
		chunkSize  int64
	}{
		{1024 * 1024 * 1024, 4 * 1024 * 1024},
		{16*1024*1024 + 3, 16},
	} {
		s.runCase1(st.bufferSize, st.chunkSize)
	}

}

func TestSequenceFileSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(SequenceFileSyncerTestSuite))
}
