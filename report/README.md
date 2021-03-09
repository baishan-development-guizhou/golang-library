# `reporter`

## Description

There are an interface(endpoint) that collect information in our servers —— It allows us to report all the information we want to visualize in Grafana, such as *Memory*, *Goroutine* or other information.

This module can auto report information with your application. Or use custom a `reporter` to report by yourself.

> IMPORTANT： Only support `v2` endpoint.

## Usage

You can use global report. ()
> Remember update config by `Global()` method when use global report.

```go
package main
import "github.com/baishan-development-guizhou/golang-library/report"

func main() {
    report.Global().WithUrl("http://127.0.0.1:10699/v1/push").
    	Send(report.Point{Value:100})
}
```

Also create a new reporter.

```go
package main
import (
    "github.com/baishan-development-guizhou/golang-library/report"
    "time"
)

func main() {
	report.Configure().
		AddReporter(report.Point{Value:100},time.Minute).
		Start()
}
```

More examples see `examples` dir.

## Configure

`report.Options` can config by yourself.

|Name |Description |Default
| --- | ---| ---
| `Reporters`| **Only use new report**, the new report will auto request when the application call `Start`.| []
| `addr`|Config this by calling `WithAddr(string)`. Server endpoint.| http://127.0.0.1:10699/v2/push
| `defaultValue`|Config this by calling `WithDefaultValue(string)`. DefaultValue is the default value in shortcut methods, like SendPayLoad, SendPayLoadWithPoint.| `Point{Metric: "monitor-bigdata-test", Endpoint: "", Step: 60, Tags: Fields{"name=": "default"}}`
| `context`| Config this by calling `WithContext(context)`. Context will be bound to ticker.| `context.Background()`
| `client`| Config this by calling `WithClient(client)`. Request endpoint by this client.| `http.Client{}`
| `log`| Config this by calling `WithLog(log)`.| `log.G()`
| `before`,`after`, `errorHandler`| Config hook by `WithBefore`, `WithAfter`, `WithErrorHandler`.| func