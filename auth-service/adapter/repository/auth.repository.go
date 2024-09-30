package repository

import (
	"database/sql"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/init/database"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() ports.AuthRepository {
	return &AuthRepository{
		db: database.P().DB,
	}
}
