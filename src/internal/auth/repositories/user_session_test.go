package repositories

import (
	testutils "github.com/CreativePhilip/backend/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbUserSessionRepository_Create(t *testing.T) {
	db, teardown := testutils.SetupIntegrationTest()
	defer teardown()

	users := DbUserRepository{Db: db}
	sessions := DbUserSessionRepository{Db: db}

	user, err := users.Create(User{
		FirstName: "john",
		LastName:  "john",
		Email:     "john@john.com",
		Password:  "password",
		IsStaff:   false,
	})

	assert.Nil(t, err)

	newSession, err := sessions.Create(user.Id)

	assert.Nil(t, err)
	assert.NotNil(t, newSession)
}

func TestDbUserSessionRepository_GetByCookieValue(t *testing.T) {
	db, teardown := testutils.SetupIntegrationTest()
	defer teardown()

	users := DbUserRepository{Db: db}
	sessions := DbUserSessionRepository{Db: db}

	user, err := users.Create(User{
		FirstName: "john",
		LastName:  "john",
		Email:     "john@john.com",
		Password:  "password",
		IsStaff:   false,
	})

	assert.Nil(t, err)

	newSession, err := sessions.Create(user.Id)

	assert.Nil(t, err)
	assert.NotNil(t, newSession)

	fetchedSession, err := sessions.GetByCookieValue(newSession.CookieValue)

	assert.Nil(t, err)
	assert.NotNil(t, fetchedSession)
}
