package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Models struct {
	User User
}

func New() Models {
	return Models{
		User: User{},
	}
}

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func (u *User) GetByUsername(db *sql.DB, username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, email, passwordHash from users where username = $1`

	var user User
	row := db.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) PasswordMatches(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func (u *User) Insert(db *sql.DB) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	stmt := `insert into users (username, email, passwordHash) values ($1, $2, $3)`

	err := db.QueryRowContext(ctx, stmt, u.Username, u.Username, u.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
