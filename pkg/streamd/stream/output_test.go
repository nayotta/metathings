package stream_manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type outputDataCodecTestSuite struct {
	suite.Suite

	output_data *OutputData
	codec       *OutputDataCodec
}

func (self *outputDataCodecTestSuite) SetupTest() {
	self.output_data = NewOutputData(map[string]interface{}{
		"foo": "bar",
	}, map[string]interface{}{
		"baz": "qux",
	})
	self.codec = new(OutputDataCodec)
}

func (self *outputDataCodecTestSuite) TestEncodeDecode() {
	buf, err := self.codec.Encode(self.output_data)
	self.Nil(err)

	val, err := self.codec.Decode(buf)
	self.Nil(err)

	dat, ok := val.(*OutputData)
	self.True(ok)

	self.Equal("bar", dat.AsString("foo"))
	self.Equal("qux", dat.Metadata().AsString("baz"))
}

func TestOutputDataCodecTestSuite(t *testing.T) {
	suite.Run(t, new(outputDataCodecTestSuite))
}
