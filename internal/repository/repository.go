package repository

import (
	"database/sql"
)

// Repository структура для работы с базой данных
type Repository struct {
	db *sql.DB
}

// New создает новый экземпляр repository
func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
