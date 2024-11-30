package internal

import (
	"errors"
	appErrors "github.com/CreativePhilip/backend/src/pkg/app_errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"net/http"
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
