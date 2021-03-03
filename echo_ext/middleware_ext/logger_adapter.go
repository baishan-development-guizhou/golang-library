package middleware_ext

import (
	"github.com/baishan-development-guizhou/golang-library/echo_ext"
	"github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _defaultLoggerAdapterGenerator = func() log.Logger {
	//use log global
	return log.G()
}

type LoggerAdapterConfig struct {
	//Generator Generate a log logger
	Generator func() log.Logger
	Skipper   middleware.Skipper
}

//LoggerAdapter use default config middleware_ext
func LoggerAdapter() echo.MiddlewareFunc {
	return LoggerAdapterWithConfig(LoggerAdapterConfig{})
}

//LoggerAdapterWithConfig use custom config middleware_ext
func LoggerAdapterWithConfig(config LoggerAdapterConfig) echo.MiddlewareFunc {
	if config.Generator == nil {
		config.Generator = _defaultLoggerAdapterGenerator
	}
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if config.Skipper(context) {
				return next(context)
			}
			context.SetLogger(echo_ext.NewEchoLoggerWithGLogger(config.Generator()))
			return next(context)
		}
	}
}
