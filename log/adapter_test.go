package log

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestZapLoggerAdapter_With(t *testing.T) {
	Convey("test", t, func() {
		Convey("new logger from with func", func() {
			logger := Global()
			withLogger := logger.With("")
			So(withLogger, ShouldNotBeNil)
			So(withLogger, ShouldNotEqual, logger)
		})
	})
}
