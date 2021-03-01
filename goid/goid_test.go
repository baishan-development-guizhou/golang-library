package goid

import (
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestGoID(t *testing.T) {
	Convey("test", t, func() {
		So(GoID(), ShouldEqual, GoID())
	})
}
func TestConcurrentGetID(t *testing.T) {
	Convey("test", t, func(c C) {
		var wg sync.WaitGroup
		concurrent := 1000
		wg.Add(concurrent)
		otherID := GoID()
		for i := 0; i < concurrent; i++ {
			go func() {
				c.So(GoID(), ShouldEqual, GoID())
				c.So(otherID, ShouldNotEqual, GoID())
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func BenchmarkGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoID()
	}
}
