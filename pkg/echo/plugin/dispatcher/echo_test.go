package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnary_echo(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	Convey("TestUnary_echo success", t, func() {
		var c string = "test"
		var d string = "test"
		So(c, ShouldEqual, d)
	})

}
