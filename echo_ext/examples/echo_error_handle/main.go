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
	engine.Use(middleware_ext.AccessLog())
	engine.HTTPErrorHandler = echo_ext.HttpErrorHandler
	engine.GET("/", func(context echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, "sdsfs").SetInternal(errors.New("asf"))
	})
	_ = engine.Start(":8080")
}
