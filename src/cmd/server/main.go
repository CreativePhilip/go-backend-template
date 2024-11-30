package main

import (
	"github.com/CreativePhilip/backend/src/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	internal.BuildTopLevelRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
