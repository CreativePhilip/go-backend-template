package auth

import (
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/internal/auth/repositories"
	appErrors "github.com/CreativePhilip/backend/src/pkg/app_errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginEndpointPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginEndpointResponse struct {
	Status string `json:"status"`
}

type LoginEndpoint struct{}

var ErrEndpointCouldNotLogin = appErrors.Error{
	ErrorCode: http.StatusBadRequest,
	Errors:    []appErrors.ErrorBody{{Message: "Invalid login or password"}},
}

func (e LoginEndpoint) Handler(payload *LoginEndpointPayload, c echo.Context) (*LoginEndpointResponse, error) {
	tx := db.Client().MustBegin()

	users := repositories.DbUserRepository{Db: tx}
	sessions := repositories.DbUserSessionRepository{Db: tx}

	session, err := LoginService(payload, &users, &sessions)

	if err != nil {
		return nil, ErrEndpointCouldNotLogin
	}

	c.SetCookie(&http.Cookie{
		Name:     "app-session",
		Value:    session.CookieValue,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

	return &LoginEndpointResponse{
		Status: "success",
	}, nil
}

func (e LoginEndpoint) ValidateInput(in *LoginEndpointPayload) error {
	return validation.ValidateStruct(in,
		validation.Field(&in.Email, validation.Required, is.Email),
		validation.Field(&in.Password, validation.Required),
	)
}

func (e LoginEndpoint) ValidateOutput(out *LoginEndpointResponse) error {
	return nil
}
