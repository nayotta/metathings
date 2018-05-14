package state_helper

import (
	state "github.com/nayotta/metathings/pkg/proto/common/state"
)

type CoreStateParser struct {
	parser StateParser
}

func (p CoreStateParser) ToString(s state.CoreState) string {
	return p.parser.ToString(int32(s))
}

func (p CoreStateParser) ToValue(x string) state.CoreState {
	return state.CoreState(p.parser.ToValue(x))
}

func NewCoreStateParser() CoreStateParser {
	return CoreStateParser{
		parser: NewStateParser("core_state", state.CoreState_name, state.CoreState_value),
	}
}
