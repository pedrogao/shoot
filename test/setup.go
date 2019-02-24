package test

import (
	"github.com/PedroGao/shoot/router"
	"github.com/labstack/echo"
)

func setupApp() *echo.Echo {
	app := echo.New()
	// load middleware and routes
	router.Load(app)
	// test api
	app.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"msg": "greeting from pedro",
		})
	})
	return app
}
