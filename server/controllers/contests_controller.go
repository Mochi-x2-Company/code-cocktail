package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// /contest
func ConIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest": "index",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

// /contest/:id
func ConShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

// /contest/team
func TeamIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest/team": "index",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

// /contest/team/:id
func TeamShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"contest/team": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
