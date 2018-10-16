package net_helper

import (
	"errors"
	"testing"

	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetLocalIP(t *testing.T) {
	Convey("TestGetLocalIP addr nil", t, func() {
		stubs := StubFunc(&InterfaceAddrs, nil, nil)
		defer stubs.Reset()
		So(GetLocalIP(), ShouldBeEmpty)
	})

	Convey("TestGetLocalIP addr err", t, func() {
		stubs := StubFunc(&InterfaceAddrs, nil, errors.New("error"))
		defer stubs.Reset()
		So(GetLocalIP(), ShouldBeEmpty)
	})

	Convey("TestGetLocalIP true env", t, func() {
		//this maybe failed in env have no ip
		So(GetLocalIP(), ShouldNotBeEmpty)
	})
}
