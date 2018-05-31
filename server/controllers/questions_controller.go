package controllers

import (
	"../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func QueIndex(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ques := models.QueAll()
		db.Debug().Find(&ques)
		// fmt.Println("data")
		// fmt.Println(data)

		// fmt.Println("range")
		// for _, que := range ques {
		// 	// fmt.Println(que.Id)
		// 	fmt.Println(que.Title)
		// 	fmt.Println(que.Body)
		// }

		fmt.Println("ques")
		fmt.Println(ques)

		// jsonMap := map[string]string{
		// 	"question": "index",
		// }
		// return c.JSON(http.StatusOK, jsonMap)

		return c.JSON(http.StatusOK, ques)
	}
}

func QueShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"question": "show",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
