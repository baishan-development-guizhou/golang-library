package report

import (
	"bytes"
	"encoding/json"
	"github.com/baishan-development-guizhou/golang-library/ocommon/oticker"
	"net/http"
	"time"
)

// AddReporter will add a report to the collection of report.
// IMPORTANT: This method must be called before the Start.
// Support chain call.
func (o *Options) AddReporter(point func() Point, interval time.Duration) *Options {
	o.Reporters = append(o.Reporters, Reporter{
		Point:    point,
		Interval: interval,
	})
	return o
}

// SendReporter will add a Point to chan with Interval.
// Support chain call.
func (o *Options) SendReporter(reporter Reporter) *Options {
	o.Reporters = append(o.Reporters, reporter)
	oticker.TickerCtx(o.context, reporter.Interval, func() {
		o.Send(reporter.Point())
	})
	return o
}

// SendReporterPayLoad will add a Point to chan with Interval.
// Support chain call.
func (o *Options) SendReporterPayLoad(point func() Point, interval time.Duration) *Options {
	o.Reporters = append(o.Reporters, Reporter{
		Point:    point,
		Interval: interval,
	})
	oticker.TickerCtx(o.context, interval, func() {
		o.Send(point())
	})
	return o
}

// Send will add a Point to chan, and only add once.
// Support chain call.
func (o *Options) Send(point Point) *Options {
	o.points <- point
	return o
}

// SendPayLoad will add a Point to chan with some fields, and only add once.
// Support chain call.
func (o *Options) SendPayLoad(fields Fields, tags Tags, value float64) *Options {
	return o.Send(o.assemble(fields, tags, value))
}

// assemble wrap Point entity.
func (o *Options) assemble(fields Fields, tags Tags, value float64) Point {
	//value copy
	point := *o.defaultValue
	point.Fields = fields
	point.Tags = tags
	point.Value = value
	point.Time = time.Now().Unix()
	return point
}

// SendPayLoadWithPoint will add a Point to chan with another Point, and only add once.
func (o *Options) SendPayLoadWithPoint(point Point) *Options {
	o.Send(o.assemblePoint(point))
	return o
}

// assemblePoint wrap Point struct by another Point.
func (o *Options) assemblePoint(point Point) Point {
	defaultValue := *o.defaultValue
	if point.Name != "" {
		defaultValue.Name = point.Name
	}
	if point.Endpoint != "" {
		defaultValue.Endpoint = point.Endpoint
	}
	if point.Value != 0 {
		defaultValue.Value = point.Value
	}
	if point.Step != 0 {
		defaultValue.Step = point.Step
	}
	if point.Fields != nil {
		defaultValue.Fields = point.Fields
	}
	if point.Tags != nil {
		defaultValue.Tags = point.Tags
	}
	if point.Time != 0 {
		defaultValue.Time = point.Time
	}
	return defaultValue
}

// Start will start report.
func (o *Options) Start() {
	for _, reporter := range o.Reporters {
		oticker.TickerCtx(o.context, reporter.Interval, func() {
			o.Send(reporter.Point())
		})
	}
	go o.report()
}

// report will do hooks and listen queue.
func (o *Options) report() {
	for {
		point := <-o.points
		// before hook.
		if o.before != nil {
			point = o.before(point)
		}
		// Request.
		status := o.request(point)
		// after hook.
		if o.after != nil {
			o.after(point, status)
		}
	}
}

// request is report main.
func (o *Options) request(point Point) bool {
	jsonStr, jsonErr := json.Marshal([]Point{point})
	if jsonErr != nil {
		o.log.Errorf("Json serialization error: %v.", jsonErr.Error())
		o.error(point, http.StatusInternalServerError)
		return false
	}
	resp, err := o.client.Post(o.url, "application/json; charset=UTF-8", bytes.NewBuffer(jsonStr))
	if err != nil {
		o.log.Errorf("Report Error: %v.", err.Error())
		o.error(point, http.StatusServiceUnavailable)
		return false
	}
	if resp == nil {
		o.log.Errorf("Report Error: Can not get response.")
		o.error(point, http.StatusInternalServerError)
		return false
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusIMUsed {
		o.log.Errorf("Report Error: status code: %v.", resp.StatusCode)
		o.error(point, resp.StatusCode)
		return false
	}
	return true
}

// error call simple errorHandler.
func (o *Options) error(point Point, status int) {
	if o.errorHandler == nil {
		return
	}
	o.errorHandler(point, status)
}
