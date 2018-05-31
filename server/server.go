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
	// サンプル
	e.GET("/", controllers.Sample())

	// index
	e.GET("/user", controllers.UserIndex())
	e.GET("/question", controllers.QueIndex())
	e.GET("/contest", controllers.ConIndex())
	e.GET("/contest/team", controllers.TeamIndex())

	// show
	e.GET("/user/:id", controllers.UserShow())
	e.GET("/question/:id", controllers.QueShow())
	e.GET("/contest/:id", controllers.ConShow())
	e.GET("/contest/team/:id", controllers.TeamShow())

	e.Logger.Fatal(e.Start(":3000"))
}
