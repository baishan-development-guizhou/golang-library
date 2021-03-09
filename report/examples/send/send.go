package main

import (
	"fmt"
	"github.com/baishan-development-guizhou/golang-library/report"
	"sync"
	"time"
)

var wg sync.WaitGroup

func send() {
	report.WithAfter(func(_ report.Point, success bool) {
		fmt.Printf("Finished report, result is: %v. \n", success)
		wg.Done()
	})
	// Send once.
	report.Send(report.Point{Value: 100})
	// Send Reporter
	report.SendReporterPayLoad(func() report.Point {
		return report.Point{Value: 100}
	}, time.Hour*24)
	report.SendReporter(report.Reporter{
		Point: func() report.Point {
			return report.Point{}
		},
		Interval: time.Second,
	})

	// Send Point.
	report.SendPayLoad(report.Fields{"a": "1"}, report.Tags{"b": "1"}, 15)
	report.SendPayLoadWithPoint(report.Point{Value: 100})
	report.SendReporter(*report.ApplicationReporter(time.Hour))
}

func main() {
	wg.Add(5)
	send()
	wg.Wait()
}
