package internal

import "github.com/labstack/echo/v4"

type Endpoint[T any] interface {
	Handler(T) (T, error)

	ValidateInput(T) error
	ValidateOutput(T) error
}

func createHandler[T any](e Endpoint[T]) echo.HandlerFunc {
	return func(e echo.Context) error {
		return nil
	}
}
