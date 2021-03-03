package echo_ext

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
}

//NewHttpError Enhance context
func (c *Context) NewHttpError(code int, err error, message ...string) error {
	var result *echo.HTTPError
	if len(message) > 0 {
		result = echo.NewHTTPError(code).SetInternal(err)
	} else {
		result = echo.NewHTTPError(code, message[0]).SetInternal(err)
	}
	return result
}

func (c *Context) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

//ContextEnhance make context easy to use
func ContextEnhance(ctx echo.Context) *Context {
	return &Context{ctx}
}
