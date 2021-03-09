package report

import (
	"context"
	"github.com/baishan-development-guizhou/golang-library/log"
	"net/http"
	"time"
)

var global = Configure()

func init() {
	global.Start()
}

// Global returns global options.
func Global() *Options {
	return global
}

// SendReporter will add a Point to chan with Interval.
// Support chain call.
func SendReporter(reporter Reporter) *Options {
	return global.SendReporter(reporter)
}

// SendReporterPayLoad will add a Point to chan with Interval.
// Support chain call.
func SendReporterPayLoad(point func() Point, interval time.Duration) *Options {
	return global.SendReporterPayLoad(point, interval)
}

// Send will add a Point to chan, and only add once.
// Support chain call.
func Send(point Point) *Options {
	return global.Send(point)
}

// SendPayLoad will add a Point to chan with some fields, and only add once.
// Support chain call.
func SendPayLoad(fields Fields, tags Tags, value float64) *Options {
	return global.SendPayLoad(fields, tags, value)
}

// SendPayLoadWithPoint will add a Point to chan with another Point, and only add once.
// Support chain call.
func SendPayLoadWithPoint(point Point) *Options {
	return global.SendPayLoadWithPoint(point)
}

// WithUrl can only update `url` parameters.
// Support chain call.
func WithUrl(url string) *Options {
	global.url = url
	return global
}

// WithUrl can only update `defaultValue` parameters.
// Support chain call.
func WithDefaultValue(defaultValue *Point) *Options {
	global.defaultValue = defaultValue
	return global
}

// WithLog can update log.
// Support chain call.
func WithLog(log log.Logger) *Options {
	global.log = log
	return global
}

// WithBefore can custom before hook.
// Support chain call.
func WithBefore(before func(point Point) Point) *Options {
	global.before = before
	return global
}

// WithAfter can custom after hook.
// Support chain call.
func WithAfter(after func(point Point, success bool)) *Options {
	global.after = after
	return global
}

// WithErrorHandler can custom error handle.
// Support chain call.
func WithErrorHandler(errorHandler func(point Point, status int)) *Options {
	global.errorHandler = errorHandler
	return global
}

// WithContext will update contest.
// Support chain call.
func WithContext(context context.Context) *Options {
	global.context = context
	return global
}

// WithClient can custom http client.
// Support chain call.
func WithClient(client *http.Client) *Options {
	global.client = client
	return global
}
