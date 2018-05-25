package main

import (
	"./controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.QueIndex())

	e.Logger.Fatal(e.Start(":3000"))
}
