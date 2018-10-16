package state_helper

import (
	"testing"

	state "github.com/nayotta/metathings/pkg/proto/common/state"
	. "github.com/smartystreets/goconvey/convey"
)

var testCoreStateParser = NewCoreStateParser()

func TestCoreStateParserToString(t *testing.T) {
	Convey("TestCoreStateParserToString", t, func() {
		So(testCoreStateParser.ToString(state.CoreState(state.CoreState_value["CORE_STATE_UNKNOWN"])), ShouldEqual, "unknown")
		So(testCoreStateParser.ToString(state.CoreState(state.CoreState_value["CORE_STATE_ONLINE"])), ShouldEqual, "online")
		So(testCoreStateParser.ToString(state.CoreState(state.CoreState_value["CORE_STATE_OFFLINE"])), ShouldEqual, "offline")
	})

	Convey("TestCoreStateParserToString Unknown", t, func() {
		So(testCoreStateParser.ToString(state.CoreState(-1)), ShouldEqual, "unknown")
	})
}

func TestCoreStateParserToValue(t *testing.T) {
	Convey("TestCoreStateParserToValue", t, func() {
		So(testCoreStateParser.ToValue("unknown"), ShouldEqual, 0)
		So(testCoreStateParser.ToValue("online"), ShouldEqual, 1)
		So(testCoreStateParser.ToValue("offline"), ShouldEqual, 2)
	})

	Convey("TestCoreStateParserToValue Unknown", t, func() {
		So(testCoreStateParser.ToValue("xxxxx"), ShouldEqual, 0)
	})
}