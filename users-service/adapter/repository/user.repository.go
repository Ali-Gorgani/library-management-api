package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/auth-service/pkg/util"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
	"library-management-api/util/errorhandler"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() ports.UserRepository {
	return &UserRepository{
		db: database.P().DB,
	}
}

// AddUser implements ports.UserRepository.
func (u UserRepository) AddUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var addedUser UserRes
	hashedPassword := util.HashedPassword(user.Password)

	query := "INSERT INTO users (username, email, hashed_password, is_admin) VALUES ($1, $2, $3, $4) RETURNING id, username, email, is_admin, created_at"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword, user.IsAdmin)
	err := row.Scan(&addedUser.ID, &addedUser.Username, &addedUser.Email, &addedUser.IsAdmin, &addedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, errorhandler.ErrUserNotFound
		}
		return &domain.User{}, err
	}
	res := MapUserResToUser(&addedUser)
	return res, nil
}

// GetUsers implements ports.UserRepository.
func (u UserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*UserRes

	rows, err := u.db.Query("SELECT id, username, email, is_admin, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(UserRes)
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	res := MapUsersResToUsers(users)
	return res, nil
}

// GetUser implements ports.UserRepository.
func (u UserRepository) GetUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var getUser UserRes
	query := "SELECT id, username, email, is_admin, created_at FROM users WHERE id=$1"
	row := u.db.QueryRow(query, user.ID)
	err := row.Scan(&getUser.ID, &getUser.Username, &getUser.Email, &getUser.IsAdmin, &getUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, errorhandler.ErrUserNotFound
		}
		return &domain.User{}, err
	}
	res := MapUserResToUser(&getUser)
	return res, nil
}

// UpdateUser implements ports.UserRepository.
func (u UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var updatedUser UserRes
	hashedPassword := util.HashedPassword(user.Password)

	query := "UPDATE users SET username=$1, email=$2, hashed_password=$3, is_admin=$4 WHERE id=$5 RETURNING id, username, email, is_admin, created_at"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword, user.IsAdmin, user.ID)
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.IsAdmin, &updatedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, errorhandler.ErrUserNotFound
		}
		return &domain.User{}, err
	}
	res := MapUserResToUser(&updatedUser)
	return res, nil
}

// DeleteUser implements ports.UserRepository.
func (u UserRepository) DeleteUser(ctx context.Context, user *domain.User) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Exec(query, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrUserNotFound
		}
		return err
	}
	return nil
}
