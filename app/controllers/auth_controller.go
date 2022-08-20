package controllers

import (
	"github.com/labstack/echo/v4"
	"web-standalone-template/pkg/utils"
)

func Authorize(c echo.Context) (int, error) {
	token, err := utils.RequestLookUp(c, "token", "query", "form")
	if err != nil {
		return 0, err
	}
	return utils.Verify(token)
}
