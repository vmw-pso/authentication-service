package data

import "database/sql"

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
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func (u *User) Insert(db *sql.DB) (int, error) {
	return 1, nil
}
