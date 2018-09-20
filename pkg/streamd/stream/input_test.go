package stream_manager

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type inputDataCodecTestSuite struct {
	suite.Suite

	input_data *InputData
	codec      *InputDataCodec
}

func (self *inputDataCodecTestSuite) SetupTest() {
	self.input_data = NewInputData(map[string]interface{}{
		"foo": "bar",
	}, map[string]interface{}{
		"baz": "qux",
	})
	self.codec = new(InputDataCodec)
}

func (self *inputDataCodecTestSuite) TestEncodeDecode() {
	buf, err := self.codec.Encode(self.input_data)
	self.Nil(err)

	val, err := self.codec.Decode(buf)
	self.Nil(err)

	dat, ok := val.(*InputData)
	self.True(ok)

	self.Equal("bar", dat.AsString("foo"))
	self.Equal("qux", dat.Metadata().AsString("baz"))
}

func TestInputDataCodecTestSuite(t *testing.T) {
	suite.Run(t, new(inputDataCodecTestSuite))
}
