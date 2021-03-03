package echo_ext

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PlaygroundValidator struct {
	validate *validator.Validate
}

func (c *PlaygroundValidator) Validate(i interface{}) error {
	return c.validate.Struct(i)
}

func NewPlaygroundValidator() echo.Validator {
	return &PlaygroundValidator{validate: validator.New()}

}
