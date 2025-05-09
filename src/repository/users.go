package repository

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(usuario models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users (name, nickname, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(usuario.Name, usuario.Nick, usuario.Email, usuario.Password)
	if err != nil {
		return 0, err
	}

	lastUserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastUserID), nil

}
