package report

import (
	"runtime"
	"strconv"
	"time"
)

// ApplicationReporter will report goroutine number and memory with this application.
// duration is sending Interval.
func ApplicationReporter(duration time.Duration) *Reporter {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &Reporter{
		Point: func() Point {
			return Point{
				Tags: Tags{
					"goroutine": strconv.Itoa(runtime.NumGoroutine()),
					"memory":    strconv.FormatUint(m.Sys, 10),
				},
			}
		},
		Interval: duration,
	}
}
