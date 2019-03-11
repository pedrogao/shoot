package router

import (
	"github.com/PedroGao/shoot/controller"
	"github.com/PedroGao/shoot/libs/context"
	err2 "github.com/PedroGao/shoot/libs/err"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Load(app *echo.Echo, mw ...echo.MiddlewareFunc) *echo.Echo {
	app.HTTPErrorHandler = err2.HTTPErrHandler
	app.Validator = &Validator{validator: validator.New()}

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	// replace default context by extendedContext
	app.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := context.ExtendedContext{c}
			return h(cc)
		}
	})

	// apply middleware
	app.Use(mw...)

	//app.Use(middleware.ErrHandler)

	// mount routes
	app.POST("/login", controller.Login)
	app.POST("/register", controller.Register)

	user := app.Group("/user")
	user.GET("/", controller.GetUsers)

	book := app.Group("/book")
	book.GET("/search", controller.SearchBook)

	book.GET("/id/:id", controller.GetBook)

	return app
}
