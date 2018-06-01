package controllers

import (
	"../models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func QueIndex(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ques := models.QueAll()
		db.Debug().Find(&ques)
		return c.JSON(http.StatusOK, ques)
	}
}

func QueShow(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		num, _ := strconv.Atoi(c.Param("id"))
		id := uint(num)

		que := models.QueSingle()
		que.ID = id
		db.Debug().First(&que)
		return c.JSON(http.StatusOK, que)
	}
}

func QueCreate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := new(models.Question)

		if err := c.Bind(q); err != nil {
			fmt.Println("err")
			fmt.Println()
			fmt.Println(err)
			return nil
		}

		db.Debug().Create(&q)

		return c.JSON(http.StatusCreated, q)
	}
}
func QueUpdate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		num, _ := strconv.Atoi(c.Param("id"))
		id := uint(num)
		q := new(models.Question)

		if err := c.Bind(q); err != nil {
			fmt.Println("err")
			fmt.Println()
			fmt.Println(err)
			return nil
		}

		que := models.QueSingle()
		que.ID = id
		db.Debug().Model(que).Update(&q)

		jsonMap := map[string]string{
			"question": "update",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

func QueDelete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		num, _ := strconv.Atoi(c.Param("id"))
		id := uint(num)

		que := models.QueSingle()
		que.ID = id
		db.Debug().First(&que)
		db.Debug().Delete(&que)

		jsonMap := map[string]string{
			"question": "delete",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
