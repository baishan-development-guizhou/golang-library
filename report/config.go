package report

import (
	"context"
	"github.com/baishan-development-guizhou/golang-library/log"
	"net/http"
	"time"
)

type Fields map[string]string
type Tags map[string]interface{}

// Point is an immutable structure.
type Point struct {
	Name     string  `json:"name"`
	Endpoint string  `json:"endpoint"`
	Value    float64 `json:"value"`
	Step     int     `json:"step"`
	Fields   Fields  `json:"fields"`
	Tags     Tags    `json:"tags"`
	Time     int64   `json:"time"`
}

// Reporter will send Point at regular intervals.
type Reporter struct {
	Point    func() Point
	Interval time.Duration
}

// Options can be custom configured by Configure or WithConfigure.
type Options struct {
	// Reporters is an init value. It will Send when call Start method.
	Reporters []Reporter

	// url is request url.
	url string
	// defaultValue is the default value in shortcut methods, like SendPayLoad, SendPayLoadWithPoint..
	defaultValue *Point

	// points is request queue.
	points chan Point

	// context will be bound to ticker.
	context context.Context

	// client is used request url server.
	client *http.Client

	// log can custom logger instance.
	log log.Logger

	// Hooks, will run when conditions are met.
	before       func(point Point) Point
	after        func(point Point, success bool)
	errorHandler func(point Point, status int)
}

// Configure report options and return a new Options.
func Configure() *Options {
	return &Options{
		url:          "http://127.0.0.1:10699/v2/push",
		defaultValue: &Point{Name: "monitor-bigdata-test", Endpoint: "", Step: 60, Tags: Tags{"dev": "test"}},
		points:       make(chan Point, 64),
		client:       &http.Client{Timeout: time.Second * 20},
		context:      context.Background(),
		log:          log.G(),
	}
}

// WithConfigure report options and return a exist Options.
func WithConfigure(options *Options) *Options {
	return options
}

// WithConfigure report options and return a exist Options.
func (o *Options) WithConfigure(options Options) *Options {
	if options.url != "" {
		o.url = options.url
	}
	if options.defaultValue != nil {
		o.defaultValue = options.defaultValue
	}
	if options.client != nil {
		o.client = options.client
	}
	if options.log != nil {
		o.log = options.log
	}
	if options.before != nil {
		o.before = options.before
	}
	if options.errorHandler != nil {
		o.errorHandler = options.errorHandler
	}
	if options.after != nil {
		o.after = options.after
	}
	if options.context != nil {
		o.context = options.context
	}
	return o
}

// WithUrl can only update `url` parameters.
// Support chain call.
func (o *Options) WithUrl(url string) *Options {
	o.url = url
	return o
}

// WithUrl can only update `defaultValue` parameters.
// Support chain call.
func (o *Options) WithDefaultValue(defaultValue *Point) *Options {
	o.defaultValue = defaultValue
	return o
}

// WithLog can update log.
// Support chain call.
func (o *Options) WithLog(log log.Logger) *Options {
	o.log = log
	return o
}

// WithBefore can custom before hook.
// Support chain call.
func (o *Options) WithBefore(before func(point Point) Point) *Options {
	o.before = before
	return o
}

// WithAfter can custom after hook.
// Support chain call.
func (o *Options) WithAfter(after func(point Point, success bool)) *Options {
	o.after = after
	return o
}

// WithErrorHandler can custom error handle.
// Support chain call.
func (o *Options) WithErrorHandler(errorHandler func(point Point, status int)) *Options {
	o.errorHandler = errorHandler
	return o
}

// WithContext will update contest.
// Support chain call.
func (o *Options) WithContext(context context.Context) *Options {
	o.context = context
	return o
}

// WithClient can custom http client.
// Support chain call.
func (o *Options) WithClient(client *http.Client) *Options {
	o.client = client
	return o
}
