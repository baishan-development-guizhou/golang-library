# `echo_ext`

[中文/Chinese](README.ZH.md)

## Description

A packet extending [echo](https://github.com/labstack/echo). Add `circuit breaker` and `error handler` etc...

### echo_ext

It includes `logger` `error_handler` `context_enhance` and extends echo

|     |  description | 
| --- |  ------- |
| `logger` | use a `log.Logger` to bridge `echo.Logger`,have access to log.Logger Unified printing. | 
| `error_handler` | in debug mode, use `pkg/error` to return all error messages. |
| `context_enhance` | added some methods of context

### middleware_ext

it adds `access_log` `circuit_breaker` `logger_adapter`  echo middleware

## Synopsis

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
