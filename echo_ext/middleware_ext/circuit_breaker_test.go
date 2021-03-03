package middleware_ext

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func TestCircuitBreakWithConfig(t *testing.T) {
	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{OnStateChange: nil, MaxRequests: 0, Interval: 10000000, Timeout: 500000, ReadyToTrip: nil})
	c := setup()
	//Always error
	cb := CircuitBreakWithConfig(CircuitBreakConfig{Judge: func(ctx echo.Context) bool {
		return true
	}, CircuitBreaker: breaker})
	successHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	assert.Equal(t, cb(successHandler)(c), ErrJudgmentNotPassed)
	errorHandler := func(c echo.Context) error {
		return errors.New("error")
	}
	//will return  the original error
	assert.Equal(t, cb(errorHandler)(c).Error(), "error")
}
