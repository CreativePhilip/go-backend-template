package repositories

import (
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/utils"
	"time"
)

type User struct {
	Id           uint      `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	IsStaff      bool      `db:"is_staff"`
	LastLoggedIn time.Time `db:"last_logged_in"`
	CreatedAt    time.Time `db:"created_at"`
}

type UserRepository interface {
	GetById(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user User) (*User, error)
}

type DbUserRepository struct {
	Db db.CommonDbTx
}

func (r *DbUserRepository) GetById(id uint) (*User, error) {
	var user User

	query := readFromFs("queries/user/by_id.sql")
	err := r.Db.Get(&user, query, id)

	return &user, err
}

func (r *DbUserRepository) GetByEmail(email string) (*User, error) {
	var user User
	query := readFromFs("queries/user/by_email.sql")
	err := r.Db.Get(&user, query, email)

	return &user, err
}

func (r *DbUserRepository) Create(user User) (*User, error) {
	query := readFromFs("queries/user/create.sql")
	err := r.Db.QueryRowx(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		utils.HashPassword(user.Password),
		user.IsStaff,
		time.Now(),
		time.Now(),
	).Scan(&user.Id, &user.LastLoggedIn, &user.CreatedAt)

	return &user, err
}
