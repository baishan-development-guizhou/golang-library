package middleware_ext

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAccessLog(t *testing.T) {
	c := setup()
	handler := func(c echo.Context) error {

		return c.String(http.StatusOK, "test")
	}
	middlewareFunc := AccessLog()
	err := middlewareFunc(handler)(c)
	assert.Equal(t, err, nil)
}
