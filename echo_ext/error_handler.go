package echo_ext

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

//HttpErrorHandler
func HttpErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  http.StatusText(http.StatusInternalServerError),
			Internal: err,
		}
	}

	code := he.Code
	message := he.Message
	c.Logger().Infof("error handle message:%+v,error:%+v", message, he.Internal)
	if m, ok := he.Message.(string); ok {
		if c.Echo().Debug {
			message = echo.Map{"message": m, "error": fmt.Sprintf("%+v", he.Internal)}
		} else {
			message = echo.Map{"message": m, "error": he.Internal.Error()}
		}
	}
	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Logger().Error(he.Error())
		}
	}
}
