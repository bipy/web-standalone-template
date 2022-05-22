package routes

import (
	"github.com/labstack/echo/v4"
	"web-standalone-template/app/controllers"
)

func PublicRoutes(a *echo.Echo) {
	g := a.Group("/api")

	g.GET("/hello", controllers.GetHello)

}
