package repositories

import (
	"github.com/CreativePhilip/backend/src/db"
	"github.com/gofrs/uuid"
	"time"
)

type UserSession struct {
	Id          uint      `db:"id"`
	UserId      uint      `db:"user_id"`
	CookieValue string    `db:"cookie_value"`
	ExpiresAt   time.Time `db:"expires_at"`
	CreatedAt   time.Time `db:"created_at"`
}

type UserSessionRepository interface {
	CreateSession(userId uint) (UserSession, error)

	GetSessionByCookieValue(cookieValue string) (*UserSession, error)
}

type DbUserSessionRepository struct {
	Db db.CommonDbTx
}

func (r *DbUserSessionRepository) CreateSession(userId uint) (UserSession, error) {
	query := readFromFs("queries/user_session/create.sql")
	uuidValue, err := uuid.NewV4()

	if err != nil {
		panic(err)
	}

	session := UserSession{
		UserId:      userId,
		CookieValue: uuidValue.String(),
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 7),
		CreatedAt:   time.Now(),
	}

	err = r.Db.QueryRowx(
		query,
		session.UserId,
		session.CookieValue,
		session.ExpiresAt,
		session.CreatedAt,
	).Scan(&session.Id)

	return session, err
}

func (r *DbUserSessionRepository) GetSessionByCookieValue(cookieValue string) (*UserSession, error) {
	query := readFromFs("queries/user_session/by_cookie_value.sql")

	var session UserSession
	err := r.Db.Get(&session, query, cookieValue)

	return &session, err
}
