<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

# Table of Content

- [Name](#name)
- [Description](#description)
- [Author](#author)
- [Copyright and License](#copyright-and-license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Name

golang-library-echo_ext

a packet extending echo,Add `circuit breaker` and `error handler` etc...

# Description

### echo_ext

it includes `logger` `error_handler` `context_enhance` and extends echo

`logger` use a `log.Logger` to bridge `echo.Logger`,have access to log.Logger Unified printing.

`error_handler` in debug mode, use `pkg/error` to return all error messages.

`context_enhance` added some methods of context


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

# Author

RuiFG (樊国睿) <guorui.fan@baishancloud.com>

# Copyright and License

The MIT License (MIT)

Copyright (c) 2021 RuiFG (樊国睿) <guorui.fan@baishancloud.com>