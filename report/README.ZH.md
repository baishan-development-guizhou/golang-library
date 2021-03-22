# `reporter`

[英文/English](README.md)

## 描述

在我们的服务器中拥有一个收集信息的接口(端点) —— 它允许我们报告所有我们想要在 Grafana 中进行可视化的信息，例如 *内存*、*Goroutine* 或者其他的信息。

此模块可以为你的应用自动上报信息，或者使用你自己自定义的 `reporter` 去上报信息。

> 重要： 目前只至此 `v2` 版本的端点。参照 [文档](http://jr.baishancloud.com:8090/pages/viewpage.action?pageId=188953838).

## 使用

你可以使用全局的 `report`.

> 如果需要请记得使用 `With...` 开头的方法进行更新全局的相关配置.

```go
package main
import "github.com/baishan-development-guizhou/golang-library/report"

func main() {
    report.Global().WithUrl("http://127.0.0.1:10699/v1/push").
    	Send(report.Point{Value:100})
}
```

或者开启一个新的 `report`.

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

更多示例参考 `examples` 文件夹。

## 配置

你可以自己配置 `report.Options`.

| 名称 | 描述 | 默认值
| ---  | ---| ---
| `Reporters`| **只能在新的 report 中使用**, 只有在应用调用 `Start` 前配置才有效，`Start` 启动时将会自动上报。| []
| `addr`| 通过调用 `WithAddr(string)` 进行配置，服务器上报地址， | http://127.0.0.1:10699/v2/push
| `defaultValue`| 通过调用 `WithDefaultValue(string)` 进行配置. `DefaultValue` 是发送消息时的一些字段的默认值，例如 SendPayLoad, SendPayLoadWithPoint.| `Point{Metric: "monitor-bigdata-test", Endpoint: "", Step: 60, Tags: Fields{"name=": "default"}}`
| `context`| 通过调用 `WithContext(context)` 进行配置. 上下文将绑定到定时器 `ticker`。| `context.Background()`
| `client`| 通过调用 `WithClient(client)` 进行配置. 请求端点的客户端.| `http.Client{Timeout: time.Second * 20}`
| `log`| 通过调用 `WithLog(log)` 进行配置 .| `log.G()`
| `before`,`after`, `errorHandler`| 钩子函数 `WithBefore`, `WithAfter`, `WithErrorHandler`.| `func`