package main

import (
	"fmt"
	"github.com/baishan-development-guizhou/golang-library/report"
	"sync"
	"time"
)

var wg sync.WaitGroup

// hook will config all hooks.
func hook() {
	report.Global().WithAfter(func(point report.Point, success bool) {
		fmt.Printf("Finished report, result is: %v. \n", success)
		wg.Done()
	}).WithBefore(func(point report.Point) report.Point {
		point.Name = point.Name + "_test"
		return point
	}).WithErrorHandler(func(point report.Point, status int) {
		fmt.Printf("Error report, status is: %v. \n", status)
	}).Send(report.Point{
		Name:     "monitor-bigdata-test",
		Endpoint: "bigdata-vm-172-18-2-29",
		Value:    12,
		Step:     60,
		Fields:   report.Fields{"status": "ok", "love": "ya", "name": "balabala"},
		Tags:     report.Tags{"type": "memory"},
		Time:     time.Now().Unix(),
	})
}

func main() {
	wg.Add(1)
	hook()
	wg.Wait()
}
