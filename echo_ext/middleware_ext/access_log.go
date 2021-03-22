package middleware_ext

import (
	"fmt"
	"github.com/baishan-development-guizhou/golang-library/ocommon/ostring"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

var (
	_defaultLoggerAccessor = func(ctx echo.Context, startTime time.Time, endTime time.Time) error {
		logger := ctx.Logger()
		request := ctx.Request()
		response := ctx.Response()
		logger.Infof("%s %s %s %s %s %d %s %s %d %s %s %s %s",
			request.RemoteAddr,
			request.Host,
			request.Method,
			request.RequestURI,
			request.Proto,
			response.Status,
			ostring.DefaultIfEmpty(request.Header.Get(echo.HeaderContentLength), "-"),
			ostring.DefaultIfEmpty(response.Header().Get(echo.HeaderContentLength), "-"),
			response.Size,
			fmt.Sprintf("%fs", endTime.Sub(startTime).Seconds()),
			ostring.FirstNotEmpty(request.Header.Get(echo.HeaderXRequestID), response.Header().Get(echo.HeaderXRequestID), "-"),
			ostring.DefaultIfEmpty(request.Referer(), "-"),
			ostring.DefaultIfEmpty(request.UserAgent(), "-"),
		)
		return nil
	}
)

type LoggerConfig struct {
	Accessor func(ctx echo.Context, startTime time.Time, endTime time.Time) error
	// Skipper defines a function to skip middleware_ext.
	Skipper middleware.Skipper
}

//AccessLog returns the middleware_ext of AccessLog using the default logger generator
func AccessLog() echo.MiddlewareFunc {
	return AccessLogWithConfig(LoggerConfig{})
}

//AccessLogWithConfig returns the middleware_ext of AccessLog with a custom configuration
func AccessLogWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	if config.Accessor == nil {
		config.Accessor = _defaultLoggerAccessor
	}
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			var err error
			if config.Skipper(context) {
				return next(context)
			}
			start := time.Now()
			if err = next(context); err != nil {
				context.Error(err)
			}
			stop := time.Now()
			return config.Accessor(context, start, stop)
		}
	}
}
