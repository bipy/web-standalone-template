package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"web-standalone-template/pkg/utils"
)

func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.SuccessResponse("Hello World!"))
}
