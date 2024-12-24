package repos

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepo struct {
	DB *sqlx.DB
}

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-" db:"password_hash"`

	CreatedAt time.Time `db:"created_at" json:"createadAt"`
}

func CreateUserRepo(db *sqlx.DB) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (repo UserRepo) CreateUser(username, passwordHash string) (User, error) {
	var user User

	err := repo.DB.Get(&user, `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
		RETURNING *
	`, username, passwordHash)

	if err != nil {
		return user, err
	}

	return user, nil
}
