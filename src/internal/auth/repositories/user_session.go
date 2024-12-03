package repositories

import (
	"fmt"
	"github.com/CreativePhilip/backend/src/db"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
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
	Create(userId uint) (UserSession, error)
	DeleteByIds(ids []uint) error

	GetByCookieValue(cookieValue string) (*UserSession, error)
	GetAllExpired() ([]*UserSession, error)
}

type DbUserSessionRepository struct {
	Db db.CommonDbTx
}

func (r *DbUserSessionRepository) Create(userId uint) (UserSession, error) {
	query := readFromFs("queries/user_session/create.sql")
	uuidValue, err := uuid.NewV7()

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

func (r *DbUserSessionRepository) DeleteByIds(ids []uint) error {
	query := readFromFs("queries/user_session/delete_ids.sql")
	query, args, err := sqlx.In(query, ids)

	fmt.Printf("query: %s, args: %v\n", query, ids)

	if err != nil {
		return err
	}

	query = r.Db.Rebind(query)

	// TODO: Maybe use normal exec, but that needs extending the shared db tx interface which I don't want to do now
	_ = r.Db.MustExec(query, args...)

	return nil
}

func (r *DbUserSessionRepository) GetByCookieValue(cookieValue string) (*UserSession, error) {
	query := readFromFs("queries/user_session/by_cookie_value.sql")

	var session UserSession
	err := r.Db.Get(&session, query, cookieValue)

	return &session, err
}

func (r *DbUserSessionRepository) GetAllExpired() ([]*UserSession, error) {
	query := readFromFs("queries/user_session/all_expired.sql")

	var sessions []*UserSession
	err := r.Db.Select(&sessions, query)

	return sessions, err
}
