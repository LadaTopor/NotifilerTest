package users

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateNewUser(id, name, email, password, token string) error {
	_, err := r.db.Exec(`INSERT INTO users (id, name, email, hashed_password, created_at, token) VALUES ($1, $2, $3, $4, $5, $6)`, id, name, email, password, time.Now(), token)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) VerifyingUserData(email, password string) error {
	_, err := r.db.Exec(`SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
      AND hashed_password = $2);`, email, password)
	if err != nil {
		return err
	}

	return nil
}
