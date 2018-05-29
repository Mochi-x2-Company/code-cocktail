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
	// db := config.PQCon()
	db := config.PQPro()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 前処理
	models.Migrate(db)

	// ルーティング
	e.GET("/", controllers.Sample())
	e.GET("/question", controllers.QueIndex())
	e.GET("/user", controllers.UserIndex())
	e.GET("/contest", controllers.ConIndex())
	e.GET("/question/:id", controllers.QueShow())
	e.GET("/user/:id", controllers.UserShow())
	e.GET("/contest/:id", controllers.ConShow())

	e.Logger.Fatal(e.Start(":3000"))
}
