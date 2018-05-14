package switcher_state_helper

import (
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/switcher"
)

type SwitcherStateParser struct {
	parser state_helper.StateParser
}

func (p SwitcherStateParser) ToString(s pb.SwitcherState) string {
	return p.parser.ToString(int32(s))
}

func (p SwitcherStateParser) ToValue(x string) pb.SwitcherState {
	return pb.SwitcherState(p.parser.ToValue(x))
}

func NewSwitcherStateParser() SwitcherStateParser {
	return SwitcherStateParser{
		parser: state_helper.NewStateParser("switcher_state", pb.SwitcherState_name, pb.SwitcherState_value),
	}
}
