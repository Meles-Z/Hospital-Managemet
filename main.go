package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/router"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	database.InitDB()
	database.AutoMigrateDB()
	router.SetUpRouter(e)

	e.Start(":1010")
}
