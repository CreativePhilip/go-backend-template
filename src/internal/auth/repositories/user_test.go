package repositories

import (
	"github.com/CreativePhilip/backend/src/utils"
	"testing"
	"time"
)

func TestDbUserRepository_Create(t *testing.T) {
	db, teardown := utils.SetupIntegrationTest()
	defer teardown()

	users := DbUserRepository{Db: db}
	user, err := users.Create(User{
		Id:           1,
		FirstName:    "John",
		LastName:     "Wick",
		Email:        "email@email.com",
		Password:     "",
		IsStaff:      true,
		LastLoggedIn: time.Time{},
		CreatedAt:    time.Time{},
	})

	if err != nil {
		t.Error(err)
	}

	if user.Id != 1 {
		t.Errorf("User id should be 1, was %v", user.Id)
	}

	if user.Email != "email@email.com" {
		t.Errorf("User email should be email@email.com, was %v", user.Email)
	}

	if user.IsStaff != true {
		t.Errorf("User should be staff, was %v", user.IsStaff)
	}
}

func TestDbUserRepository_GetById(t *testing.T) {
	db, teardown := utils.SetupIntegrationTest()
	defer teardown()

	users := DbUserRepository{Db: db}
	user, err := users.Create(User{
		Id:           1,
		FirstName:    "John",
		LastName:     "Wick",
		Email:        "email@email.com",
		Password:     "",
		IsStaff:      true,
		LastLoggedIn: time.Time{},
		CreatedAt:    time.Time{},
	})

	if err != nil {
		t.Error(err)
	}

	dbUser, err := users.GetById(user.Id)

	if err != nil {
		t.Error(err)
	}

	if dbUser.FirstName != user.FirstName {
		t.Errorf("First name should be %v, was %v", dbUser.FirstName, user.FirstName)
	}
}

func TestDbUserRepository_GetByEmail(t *testing.T) {
	db, teardown := utils.SetupIntegrationTest()
	defer teardown()

	users := DbUserRepository{Db: db}
	user, err := users.Create(User{
		Id:           1,
		FirstName:    "John",
		LastName:     "Wick",
		Email:        "email@email.com",
		Password:     "",
		IsStaff:      true,
		LastLoggedIn: time.Time{},
		CreatedAt:    time.Time{},
	})

	if err != nil {
		t.Error(err)
	}

	dbUser, err := users.GetByEmail(user.Email)

	if err != nil {
		t.Error(err)
	}

	if dbUser.FirstName != user.FirstName {
		t.Errorf("First name should be %v, was %v", dbUser.FirstName, user.FirstName)
	}
}
