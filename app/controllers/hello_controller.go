package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"web-standalone-template/pkg/utils"
)

func GetHello(c echo.Context) error {
	id, err := Authorize(c)
	if err != nil {
		c.Logger().Printf(err.Error())
		return c.JSON(http.StatusUnauthorized, utils.FailResponse("Unauthorized"))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("Hello! "+strconv.Itoa(id)))
}
