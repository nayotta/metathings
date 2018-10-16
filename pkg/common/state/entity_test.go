package state_helper

import (
	"testing"

	state "github.com/nayotta/metathings/pkg/proto/common/state"
	. "github.com/smartystreets/goconvey/convey"
)

var testEntityStateParser = NewEntityStateParser()

func TestEntityStateParserToString(t *testing.T) {
	Convey("TestEntityStateParserToString", t, func() {
		So(testEntityStateParser.ToString(state.EntityState(state.EntityState_value["ENTITY_STATE_UNKNOWN"])), ShouldEqual, "unknown")
		So(testEntityStateParser.ToString(state.EntityState(state.EntityState_value["ENTITY_STATE_ONLINE"])), ShouldEqual, "online")
		So(testEntityStateParser.ToString(state.EntityState(state.EntityState_value["ENTITY_STATE_OFFLINE"])), ShouldEqual, "offline")
	})

	Convey("TestEntityStateParserToString Unknown", t, func() {
		So(testEntityStateParser.ToString(state.EntityState(-1)), ShouldEqual, "unknown")
	})
}

func TestEntityStateParserToValue(t *testing.T) {
	Convey("TestEntityStateParserToValue", t, func() {
		So(testEntityStateParser.ToValue("unknown"), ShouldEqual, 0)
		So(testEntityStateParser.ToValue("online"), ShouldEqual, 1)
		So(testEntityStateParser.ToValue("offline"), ShouldEqual, 2)
	})

	Convey("TestEntityStateParserToValue Unknown", t, func() {
		So(testEntityStateParser.ToValue("xxxxx"), ShouldEqual, 0)
	})
}
