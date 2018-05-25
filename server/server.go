package main

import (
	"./config"
	"./controllers"
	"./models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// 初期準備
	e := echo.New()
	db := config.PQCon()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 前処理
	models.Migrate(db)

	// ルーティング
	e.GET("/", controllers.QueIndex())

	e.Logger.Fatal(e.Start(":3000"))
}
