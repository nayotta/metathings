package streamd_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
)

type StreamStateParser struct {
	parser state_helper.StateParser
}

func (p StreamStateParser) ToString(s pb.StreamState) string {
	return p.parser.ToString(int32(s))
}

func (p StreamStateParser) ToValue(x string) pb.StreamState {
	return pb.StreamState(p.parser.ToValue(x))
}

func NewStreamStateParser() StreamStateParser {
	return StreamStateParser{
		parser: state_helper.NewStateParser("stream_state", pb.StreamState_name, pb.StreamState_value),
	}
}

var (
	STREAM_STATE_PARSER = NewStreamStateParser()
)
