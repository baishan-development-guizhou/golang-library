# `echo_ext`

[英文/English](README.md)

## 描述

一个对 [echo](https://github.com/labstack/echo) 框架的扩展，添加了 `路由熔断`、`错误处理` 等。

### echo_ext

包含了 `logger`， `error_handler`，`context_enhance` 等。

|                   | 描述 |
| ----------------- | ---  |
| `logger`          | 使用了 `log.Logger` 桥接 `echo.Logger`，能够使用其进行统一打印 |
| `error_handler`   | 在 `debug` 模式下，`pkg/error` 将会返回所有的错误信息。        |
| `context_enhance` | 扩展了 `context` |
| `middleware_ext`  | 添加了 `访问日志`、`路由熔断`、`日志适配器` 的 echo 中间件。     |

## 使用

### access_log

```go
package main

import (
	"github.com/baishan-development-guizhou/golang-library/echo_ext/middleware_ext"
	"github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	engine := echo.New()
	log.ReplaceG(log.Configure().WithOutputEncoder(log.ConsoleOutputEncoder).Init())
	engine.Use(middleware_ext.LoggerAdapter())
	engine.Use(middleware_ext.AccessLog())
	engine.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, echo.Map{"message": "ok"})
	})
	_ = engine.Start(":8080")

}


```

### error_handler

```go
package main

import (
	"errors"
	"github.com/baishan-development-guizhou/golang-library/echo_ext"
	"github.com/baishan-development-guizhou/golang-library/echo_ext/middleware_ext"
	"github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	engine := echo.New()
	engine.Debug = true
	log.ReplaceG(log.Configure().WithOutputEncoder(log.ConsoleOutputEncoder).WithNamed("[nefarian]").Init())
	engine.Use(middleware_ext.LoggerAdapter())
	engine.HTTPErrorHandler = echo_ext.HttpErrorHandler
	engine.GET("/", func(context echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, "sdsfs").SetInternal(errors.New("asf"))
	})
	_ = engine.Start(":8080")
}

```
