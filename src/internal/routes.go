package internal

import (
	"github.com/CreativePhilip/backend/src/internal/auth"
	"github.com/labstack/echo/v4"
)

func BuildTopLevelRoutes(e *echo.Echo) {
	auth.BuildRoutes(e.Group("/auth"))
}
