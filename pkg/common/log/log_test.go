package log_helper

import (
	//"errors"
	"testing" //. "github.com/prashantv/gostub"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLogger(t *testing.T) {
	Convey("TestNewLogger success", t, func() {
		log_ret, err := NewLogger("test", "debug")
		So(log_ret, ShouldNotBeEmpty)
		So(err, ShouldBeEmpty)
	})

	Convey("TestNewLogger fail", t, func() {
		log_ret, err := NewLogger("test", "debug1")
		So(log_ret, ShouldBeEmpty)
		So(err, ShouldNotBeEmpty)
	})
}
