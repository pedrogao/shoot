package test

// api 层面的测试

import (
	"github.com/PedroGao/shoot/config"
	"github.com/PedroGao/shoot/model"
	"github.com/PedroGao/shoot/router"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func setupApp() *echo.Echo {
	err := config.Init("../conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	app := echo.New()
	// init db
	model.Init()
	model.CreateTables()
	//defer model.Close()

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
