package middleware_ext

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sony/gobreaker"
	"net/http"
)

var (
	_defaultCircuitBreakJudge = func(ctx echo.Context) bool {
		return ctx.Response().Status >= http.StatusInternalServerError
	}
	ErrJudgmentNotPassed = errors.New("judgement not passed")
)

type CircuitBreakConfig struct {
	//Judge whether the circuit breaker needs counting
	Judge func(ctx echo.Context) bool
	//CircuitBreaker at work
	CircuitBreaker *gobreaker.CircuitBreaker
	// Skipper defines a function to skip middleware_ext.
	Skipper middleware.Skipper
}

//CircuitBreakWithConfig Returns the middleware_ext of the circuit break
func CircuitBreakWithConfig(config CircuitBreakConfig) echo.MiddlewareFunc {
	if config.CircuitBreaker == nil {
		panic("nil breaker")
	}
	if config.Judge == nil {
		config.Judge = _defaultCircuitBreakJudge
	}
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if config.Skipper(context) {
				return next(context)
			}
			_, err := config.CircuitBreaker.Execute(func() (interface{}, error) {
				err := next(context)
				if err == nil {
					if config.Judge(context) {
						return nil, ErrJudgmentNotPassed
					}
				}
				return nil, err
			})
			return err
		}
	}

}
