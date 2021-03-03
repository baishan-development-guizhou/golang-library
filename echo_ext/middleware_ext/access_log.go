package middleware_ext

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
	"time"
)

func stringNeedOne(args ...string) string {
	for _, arg := range args {
		if arg != "" {
			return arg
		}
	}
	return ""
}

var (
	_defaultLoggerAccessor = func(ctx echo.Context, startTime time.Time, endTime time.Time) {
		logger := ctx.Logger()
		request := ctx.Request()
		response := ctx.Response()
		requestID := stringNeedOne(request.Header.Get(echo.HeaderXRequestID), response.Header().Get(echo.HeaderXRequestID), "-")
		requestContentLength := stringNeedOne(request.Header.Get(echo.HeaderContentLength), "0")
		responseBodySize := strconv.FormatInt(response.Size, 10)
		responseContentLength := stringNeedOne(response.Header().Get(echo.HeaderContentLength), responseBodySize)
		logger.Infof("%-10s - %-10s %s:%-7s %s %3d - %s %s - %s %13v %s %s %s",
			request.RemoteAddr,
			request.Host,
			request.Method,
			request.RequestURI,
			request.Proto,
			response.Status,
			requestContentLength,
			responseContentLength,
			responseBodySize,
			endTime.Sub(startTime).String(),
			requestID,
			request.Referer(),
			request.UserAgent(),
		)

	}
)

type LoggerConfig struct {
	Accessor func(ctx echo.Context, startTime time.Time, endTime time.Time)
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
			err = next(context)
			stop := time.Now()
			config.Accessor(context, start, stop)
			return err
		}
	}
}
