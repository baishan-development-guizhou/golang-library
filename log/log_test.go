package log

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestLoggerContextKey(t *testing.T) {
	Convey("test", t, func() {
		ctx := context.Background()
		logger := Global()
		withValue := context.WithValue(ctx, loggerKey{}, logger)
		Convey("get logger from ctx", func() {
			value := ctx.Value(loggerKey{})
			ShouldEqual(fmt.Sprintf("%T", value), fmt.Sprintf("%T", logger))
		})
		Convey("get logger from withValue context", func() {
			ShouldEqual(withValue.Value(loggerKey{}), logger)
		})
		Convey("get logger from child context", func() {
			child := context.WithValue(withValue, "", "")
			ShouldEqual(child.Value(loggerKey{}), logger)
		})
	})
}
func TestLoggerFunc(t *testing.T) {
	Convey("test", t, func() {
		Convey("new Context", func() {
			parent := context.Background()
			ctx, logger := Context(parent)
			So(ctx, ShouldNotBeNil)
			So(ctx, ShouldNotEqual, parent)
			So(logger, ShouldNotBeNil)
			Convey("get logger from ctx", func() {
				child, loggerEq := Context(ctx)
				So(child, ShouldEqual, ctx)
				So(logger, ShouldEqual, loggerEq)
			})
			Convey("get logger from new ctx", func() {
				_, loggerNotEq := Context(context.TODO())
				So(logger, ShouldNotEqual, loggerNotEq)
			})
		})
	})
}

func TestRoutine(t *testing.T) {
	var wg sync.WaitGroup
	Convey("test", t, func(c C) {
		wg.Add(1)
		BindRoutine(Routine().With("test"))
		So(Routine(), ShouldEqual, Routine())
		logger := Routine()
		go func() {
			c.So(Routine(), ShouldNotEqual, logger)
			wg.Done()
		}()
		wg.Wait()
	})
}

func TestContext(t *testing.T) {
	var wg sync.WaitGroup
	Convey("test", t, func(c C) {
		wg.Add(1)
		ctx, logger := Context(context.TODO())
		ctx2, iLogger := Context(ctx)
		So(ctx, ShouldEqual, ctx2)
		So(logger, ShouldEqual, iLogger)
	})
}
