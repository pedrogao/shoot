package main

import (
	err2 "github.com/PedroGao/shoot/libs/err"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type ExtendedContext struct {
	echo.Context
}

func (ctx *ExtendedContext) Err(e error) error {
	err, ok := e.(err2.HTTPErr)
	// if err is httpErr which we defined , we can destructure it and return the information
	if ok {
		return ctx.JSON(err.HTTPCode, echo.Map{
			"msg":        err.Msg,
			"error_code": err.ErrorCode,
			"url":        ctx.Request().URL.Path,
		})
	}
	// else we did not define it , it will be the inline err that cause by other libraries or language itself
	// so we return the unknown err
	return ctx.JSON(err2.UnKnown.ErrorCode, echo.Map{
		"msg":        err2.UnKnown.Msg,
		"error_code": err2.UnKnown.ErrorCode,
		"url":        ctx.Request().URL.Path,
	})
}

func HTTPErrHandler(err error, c echo.Context) {
	var (
		code      = http.StatusInternalServerError
		msg       interface{}
		errorCode = err2.UnKnownCode
	)
	if he, ok := err.(*err2.HTTPErr); ok {
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

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	app := echo.New()
	app.HideBanner = true
	app.Debug = true
	app.Validator = &CustomValidator{validator: validator.New()}
	app.HTTPErrorHandler = HTTPErrHandler

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// replace default context by extendedContext
	app.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := ExtendedContext{c}
			return h(cc)
		}
	})

	app.POST("/", func(c echo.Context) error {
		var (
			err error
		)
		u := new(User)
		if err = c.Bind(u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	app.GET("/err", func(c echo.Context) error {
		return err2.BookNotFound
	})

	app.Logger.Fatal(app.Start(":3000"))
}
