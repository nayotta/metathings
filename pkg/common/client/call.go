package client_helper

import (
	gpb "github.com/golang/protobuf/ptypes/wrappers"
)

func NewString(s string) *gpb.StringValue {
	return &gpb.StringValue{Value: s}
}

func NewUInt64(x uint64) *gpb.UInt64Value {
	return &gpb.UInt64Value{Value: x}
}

func NewFloat32(x float32) *gpb.FloatValue {
	return &gpb.FloatValue{Value: x}
}
