package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func UserIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"user": "index",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

func UserShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"user": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
