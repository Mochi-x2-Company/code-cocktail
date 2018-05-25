package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func QueIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"hoge": "hogehoge",
			"piyo": "piyopiyo",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
