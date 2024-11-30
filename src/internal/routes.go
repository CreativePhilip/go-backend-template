package internal

import (
	"github.com/CreativePhilip/backend/src/internal/auth"
	"github.com/labstack/echo/v4"
)

func BuildTopLevelRoutes(e *echo.Echo) {
	e.Use(ErrorHandlerMiddleware)

	auth.BuildRoutes(e.Group("/auth"))
}
