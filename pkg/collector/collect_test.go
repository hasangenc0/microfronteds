package collector

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCollect(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Collect should return 61", t, func() {
		services := []Service{
			{
				host: "asdasd",
				port: "4460",
			},
		}

		host := Collect(services)
		So(host, ShouldEqual, 61)
	})
}
