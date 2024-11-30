package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Endpoint[TIn any, TOut any] interface {
	Handler(*TIn) (*TOut, error)

	ValidateInput(*TIn) error
	ValidateOutput(*TOut) error
}

func New[TIn any, TOut any](e Endpoint[TIn, TOut]) echo.HandlerFunc {
	return func(echo echo.Context) error {
		payload := new(TIn)

		if err := echo.Bind(payload); err != nil {
			return err
		}

		err := e.ValidateInput(payload)

		if err != nil {
			return err
		}

		response, err := e.Handler(payload)

		if err != nil {
			return err
		}

		return echo.JSON(http.StatusOK, response)
	}
}
