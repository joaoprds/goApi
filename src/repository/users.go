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
	return 0, nil
}
