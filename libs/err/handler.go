package err

import (
	"github.com/labstack/echo"
	"net/http"
)

func HTTPErrHandler(err error, c echo.Context) {
	var (
		code      = http.StatusInternalServerError
		msg       interface{}
		errorCode = UnKnownCode
	)
	if he, ok := err.(*HTTPErr); ok {
		code = he.HTTPCode
		msg = he.Msg
		errorCode = he.ErrorCode
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	// committed represent the response has committed to the client
	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, echo.Map{
				"msg":        msg,
				"error_code": errorCode,
				"url":        c.Request().URL.Path,
			})
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
