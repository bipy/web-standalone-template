package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"web-standalone-template/pkg/utils"
)

func MustAuthorize(c echo.Context) (int, error) {
	token := c.QueryParam("token")
	if token == "" {
		token = c.FormValue("token")
	}
	userID, err := utils.Verify(token)
	if err != nil {
		if e := c.JSON(http.StatusOK, utils.FailResponse("Unauthorized")); e != nil {
			return 0, e
		}
		return 0, errors.New("invalid token")
	}
	return userID, nil
}

func Authorize(c echo.Context) (int, error) {
	token := c.QueryParam("token")
	if token == "" {
		token = c.FormValue("token")
	}
	userID, err := utils.Verify(token)
	if err != nil {
		return -1, errors.New("invalid token")
	}
	return userID, nil
}
