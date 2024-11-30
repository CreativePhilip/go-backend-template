package utils

import (
	"crypto/sha1"
	"github.com/CreativePhilip/backend/src/pkg/config"
	"golang.org/x/crypto/pbkdf2"
)

func HashPassword(password string) []byte {
	cfg := config.GetConfig()

	hash := pbkdf2.Key(
		[]byte(password),
		[]byte(cfg.PasswordSalt),
		4096,
		32,
		sha1.New,
	)

	return hash
}
