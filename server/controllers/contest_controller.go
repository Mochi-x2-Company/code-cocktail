package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func ConIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest": "index",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

func ConShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
