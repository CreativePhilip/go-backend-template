package auth

import (
	"crypto/subtle"
	"errors"
	"github.com/CreativePhilip/backend/src/internal/auth/repositories"
	"github.com/CreativePhilip/backend/src/utils"
)

var (
	ErrCouldNotLogin = errors.New("user not found or invalid password")
)

func LoginService(payload *LoginEndpointPayload, users repositories.UserRepository, sessions repositories.UserSessionRepository) (*repositories.UserSession, error) {
	user, err := users.GetByEmail(payload.Email)

	if err != nil {
		return nil, ErrCouldNotLogin
	}

	testHash := utils.HashPassword(payload.Password)
	userHash := []byte(user.Password)

	if subtle.ConstantTimeCompare(testHash, userHash) == 1 {
		return nil, ErrCouldNotLogin
	}

	session := utils.Must(sessions.Create(user.Id))

	return &session, nil
}
