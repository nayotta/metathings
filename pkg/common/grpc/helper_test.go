package grpc_helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessParseMethodDescription(t *testing.T) {
	desc, err := ParseMethodDescription("/a.b.c.S/M")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("a.b.c", desc.Package)
	assert.Equal("S", desc.Service)
	assert.Equal("M", desc.Method)
}

func TestFailedParseMethodDescription(t *testing.T) {
	desc, err := ParseMethodDescription("")

	assert := assert.New(t)
	assert.Nil(desc)
	assert.Equal(InvalidFullMethodName, err)
}
