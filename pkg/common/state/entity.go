package state_helper

import (
	state "github.com/nayotta/metathings/proto/common/state"
)

type EntityStateParser struct {
	parser StateParser
}

func (p EntityStateParser) ToString(s state.EntityState) string {
	return p.parser.ToString(int32(s))
}

func (p EntityStateParser) ToValue(x string) state.EntityState {
	return state.EntityState(p.parser.ToValue(x))
}

func NewEntityStateParser() EntityStateParser {
	return EntityStateParser{
		parser: NewStateParser("entity_state", state.EntityState_name, state.EntityState_value),
	}
}
