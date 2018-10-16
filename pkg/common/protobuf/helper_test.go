package protobuf_helper

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProtobufHelper(t *testing.T) {
	var protobufTime timestamp.Timestamp
	unixTime := time.Unix(100, 200)
	protobufTime.Seconds = 100
	protobufTime.Nanos = 200
	fmt.Println(unixTime)
	fmt.Println(protobufTime)

	Convey("TestToTime", t, func() {
		timeRet := ToTime(protobufTime)
		So(timeRet, ShouldEqual, unixTime)
	})

	Convey("FromTime", t, func() {
		protoRet := FromTime(unixTime)
		So(protoRet.Seconds, ShouldEqual, protobufTime.Seconds)
		So(protoRet.Nanos, ShouldEqual, protobufTime.Nanos)
	})

	Convey("Now", t, func() {
		stubs := StubFunc(&timeNow, unixTime)
		defer stubs.Reset()
		protoRet := Now()
		So(protoRet.Seconds, ShouldEqual, protobufTime.Seconds)
		So(protoRet.Nanos, ShouldEqual, protobufTime.Nanos)
	})

}
