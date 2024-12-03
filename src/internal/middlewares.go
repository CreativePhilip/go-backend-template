package internal

import (
	"errors"
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/internal/auth"
	"github.com/CreativePhilip/backend/src/internal/auth/repositories"
	appErrors "github.com/CreativePhilip/backend/src/pkg/app_errors"
	"github.com/CreativePhilip/backend/src/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"net/http"
	"slices"
	"time"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err == nil {
			return nil
		}

		appErr := appErrors.Error{}
		if errors.As(err, &appErr) {
			return handleAppError(&appErr, c)
		}

		validationErr := validation.Errors{}
		if errors.As(err, &validationErr) {
			return handleValidationError(validationErr, c)
		}

		res := appErrors.Error{
			ErrorCode: http.StatusInternalServerError,
			Errors: []appErrors.ErrorBody{
				{Message: err.Error()},
			},
		}

		return c.JSON(res.ErrorCode, res)
	}
}

func handleAppError(err *appErrors.Error, c echo.Context) error {
	return c.JSON(err.ErrorCode, err)
}

func handleValidationError(err validation.Errors, c echo.Context) error {
	outErr := appErrors.Error{
		ErrorCode: http.StatusBadRequest,
		Errors:    []appErrors.ErrorBody{},
	}

	for field, fieldErr := range err {
		outErr.Errors = append(outErr.Errors, appErrors.ErrorBody{
			Field:   &field,
			Message: fieldErr.Error(),
		})
	}

	return handleAppError(&outErr, c)
}

var ErrCookieAuth = appErrors.Error{
	ErrorCode: http.StatusUnauthorized,
	Errors: []appErrors.ErrorBody{
		{
			Field:   nil,
			Message: "Auth cookie is missing or invalid",
		},
	},
}

var ErrCookieUnauthorized = appErrors.Error{
	ErrorCode: http.StatusUnauthorized,
	Errors: []appErrors.ErrorBody{
		{
			Field:   nil,
			Message: "Unauthorized",
		},
	},
}

type CookieAuthMiddlewareConfig struct {
	SkipPaths []string
}

type CookieAuthMiddlewareContext struct {
	echo.Context

	session repositories.UserSession
}

func CookieAuthMiddlewareWithConfig(config CookieAuthMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if slices.Contains(config.SkipPaths, c.Path()) {
				return next(c)
			}

			cookie, err := c.Cookie(auth.CookieName)

			if err != nil {
				return ErrCookieAuth
			}

			d := db.Client()
			sessions := repositories.DbUserSessionRepository{Db: d}
			session := utils.Must(sessions.GetByCookieValue(cookie.Value))

			if time.Now().After(session.ExpiresAt) {
				return ErrCookieUnauthorized
			}

			return next(c)
		}
	}
}
