package auth

import (
	"github.com/CreativePhilip/backend/src/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func BuildRoutes(g *echo.Group) {
	g.POST("/login", handlers.New(LoginEndpoint{}))
}
