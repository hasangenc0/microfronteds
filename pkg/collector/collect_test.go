package collector

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCollect(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Collect should return 61", t, func() {
		services := []Gateway{
			{
				Name: "Gateway",
				Host: "localhost",
				Port: "4461",
			},
		}

		host := Collect(services)
		So(host, ShouldEqual, 61)
	})
}
