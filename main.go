package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"web-standalone-template/pkg/routes"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.JWT())

	routes.PublicRoutes(app)
	routes.GeneralRoutes(app)

	app.Logger.Fatal(app.Start(":80"))
}
