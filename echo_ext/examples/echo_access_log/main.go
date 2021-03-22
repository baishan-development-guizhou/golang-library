package main

import (
	"github.com/baishan-development-guizhou/golang-library/echo_ext/middleware_ext"
	"github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	engine := echo.New()
	engine.Debug = true
	log.ReplaceG(log.Configure().WithOutputEncoder(log.ConsoleOutputEncoder).Init())
	engine.Use(middleware_ext.LoggerAdapter())
	engine.Use(middleware_ext.AccessLog())
	engine.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, echo.Map{"message": "ok"})
	})
	_ = engine.Start(":8080")
}
