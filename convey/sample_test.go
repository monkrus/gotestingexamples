package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExampleCleanup(t *testing.T) {
	Convey("it works great", t, func() {
		So(1, ShouldEqual, 1)
		So(2*2, ShouldEqual, 4)
		//note that `t` is not needed
		Convey("another test", func() {
			So(2, ShouldEqual, 2)
			So(3*3, ShouldEqual, 9)
		})
	})
}

func TestExampleCleanup2(t *testing.T) {
	x := 0 //0
	Convey("A", t, func() {
		x++ //1
		Convey("A-B", func() {
			x++ //2
			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
			}) //4
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
			}) //6
			Convey("A-B-C3", func() {
				So(x, ShouldEqual, 6)
			})

		})
		Reset(func() {
			t.Log("finish")

		})
	})
}
