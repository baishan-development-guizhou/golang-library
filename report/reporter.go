package report

import (
	"runtime"
	"time"
)

// ApplicationReporter will report goroutine number and memory with this application.
// duration is sending Interval.
func ApplicationReporter(duration time.Duration) *Reporter {
	return ApplicationReporterWithAnother(func() Point {
		return Point{}
	}, duration)
}

// ApplicationReporterWithAnother will report goroutine number and memory with this application By another point.
// duration is sending Interval.
func ApplicationReporterWithAnother(pointFunc func() Point, duration time.Duration) *Reporter {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &Reporter{
		Point: func() Point {
			point := pointFunc()
			point.Fields["goroutine"] = runtime.NumGoroutine()
			point.Fields["memory"] = m.Sys
			return point
		},
		Interval: duration,
	}
}
