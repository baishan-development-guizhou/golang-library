package echo_ext

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
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
