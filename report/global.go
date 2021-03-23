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
	return global.WithUrl(url)
}

// WithUrl can only update `defaultValue` parameters.
// Support chain call.
func WithDefaultValue(defaultValue *Point) *Options {
	return global.WithDefaultValue(defaultValue)
}

// WithLog can update log.
// Support chain call.
func WithLog(log log.Logger) *Options {
	return global.WithLog(log)
}

// WithBefore can custom before hook.
// Support chain call.
func WithBefore(before func(point Point) Point) *Options {
	return global.WithBefore(before)
}

// WithAfter can custom after hook.
// Support chain call.
func WithAfter(after func(point Point, success bool)) *Options {
	return global.WithAfter(after)
}

// WithErrorHandler can custom error handle.
// Support chain call.
func WithErrorHandler(errorHandler func(point Point, status int)) *Options {
	return global.WithErrorHandler(errorHandler)
}

// WithContext will update contest.
// Support chain call.
func WithContext(context context.Context) *Options {
	return global.WithContext(context)
}

// WithClient can custom http client.
// Support chain call.
func WithClient(client *http.Client) *Options {
	return global.WithClient(client)
}

// WithDev will output information.
func WithDev(dev bool) *Options {
	global.Dev = dev
	return global
}
