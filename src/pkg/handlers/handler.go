package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Endpoint[TIn any, TOut any, Context any] interface {
	Handler(*TIn, Context) (*TOut, error)

	ValidateInput(*TIn) error
	ValidateOutput(*TOut) error
}

func New[TIn any, TOut any, Context any](e Endpoint[TIn, TOut, Context]) echo.HandlerFunc {
	return func(echoContext echo.Context) error {
		payload := new(TIn)

		if err := echoContext.Bind(payload); err != nil {
			return err
		}

		err := e.ValidateInput(payload)

		if err != nil {
			return err
		}

		ctx := e.(Context)
		response, err := e.Handler(payload, ctx)

		if err != nil {
			return err
		}

		return echoContext.JSON(http.StatusOK, response)
	}
}
