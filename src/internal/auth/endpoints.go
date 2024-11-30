package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
}

type LoginEndpoint struct{}

func (e LoginEndpoint) Handler(payload *LoginPayload) (*LoginResponse, error) {
	return nil, nil
}

func (e LoginEndpoint) ValidateInput(in *LoginPayload) error {
	return validation.ValidateStruct(in,
		validation.Field(&in.Email, validation.Required, is.Email),
		validation.Field(&in.Password, validation.Required),
	)
}

func (e LoginEndpoint) ValidateOutput(out *LoginResponse) error {
	return nil
}
