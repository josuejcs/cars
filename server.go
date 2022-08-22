package main

import (
	"cars/configs"
	"cars/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	routes.UserRoute(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
