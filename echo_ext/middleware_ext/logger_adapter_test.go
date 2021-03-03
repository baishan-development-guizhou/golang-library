package middleware_ext

import (
	"github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLoggerAdapter(t *testing.T) {
	c := setup()
	handler := func(c echo.Context) error {
		logger := c.Logger()
		_, ok := logger.(log.Logger)
		assert.Equal(t, ok, true)
		return c.String(http.StatusOK, "test")
	}
	middlewareFunc := LoggerAdapter()
	_ = middlewareFunc(handler)(c)
}

func TestLoggerAdapterWithConfig(t *testing.T) {
	sourceLogger := log.G().With("test")
	c := setup()
	handler := func(c echo.Context) error {
		logger := c.Logger()
		_, ok := logger.(log.Logger)
		assert.Equal(t, ok, true)
		//assert.Equal(t, fmt.Sprintf("%v", loggerd), fmt.Sprintf("%v", sourceLogger))
		return c.String(http.StatusOK, "test")
	}
	middlewareFunc := LoggerAdapterWithConfig(LoggerAdapterConfig{Generator: func() log.Logger {
		return sourceLogger
	}})
	_ = middlewareFunc(handler)(c)
}
