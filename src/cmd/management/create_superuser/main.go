package main

import (
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/internal/auth/repositories"
	"log"
	"time"
)

func main() {
	d := db.Client()

	users := repositories.DbUserRepository{Db: d}

	newUser := repositories.User{
		FirstName:    "Admin",
		LastName:     "Admin",
		Email:        "admin@admin.com",
		Password:     "admin",
		IsStaff:      true,
		LastLoggedIn: time.Time{},
		CreatedAt:    time.Time{},
	}

	_, err := users.Create(newUser)

	if err != nil {
		log.Fatal(err)
	}
}
