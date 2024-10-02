package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
	"library-management-api/users-service/pkg/util"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrForbidden          = errors.New("you are not allowed to access this resource")
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
func (u UserRepository) AddUser(ctx context.Context, user domain.AddUserReq) (domain.UserRes, error) {
	hashedPassword := util.HashedPassword(user.Password)
	var addedUser domain.UserRes
	query := "INSERT INTO users (username, email, hashed_password) VALUES ($1, $2, $3) RETURNING id, username, email, is_admin, created_at"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword)
	err := row.Scan(&addedUser.ID, &addedUser.Username, &addedUser.Email, &addedUser.IsAdmin, &addedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserRes{}, ErrUserNotFound
		}
		return domain.UserRes{}, err
	}
	return addedUser, nil
}

// GetUsers implements ports.UserRepository.
func (u UserRepository) GetUsers(ctx context.Context) ([]domain.UserRes, error) {
	var users []domain.UserRes
	var user domain.UserRes

	rows, err := u.db.Query("SELECT id, username, email, is_admin, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUser implements ports.UserRepository.
func (u UserRepository) GetUser(ctx context.Context, id int) (domain.UserRes, error) {
	var user domain.UserRes
	query := "SELECT id, username, email, is_admin, created_at FROM users WHERE id=$1"
	row := u.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserRes{}, ErrUserNotFound
		}
		return domain.UserRes{}, err
	}
	return user, nil
}

// UpdateUser implements ports.UserRepository.
func (u UserRepository) UpdateUser(ctx context.Context, id int, user domain.UpdateUserReq) (domain.UserRes, error) {
	var updatedUser domain.UserRes
	hashedPassword := util.HashedPassword(user.Password)
	query := "UPDATE users SET username=$1, email=$2, hashed_password=$3, is_admin=$4 WHERE id=$5 RETURNING id, username, email, is_admin, created_at"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword, user.IsAdmin, id)
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.IsAdmin, &updatedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserRes{}, ErrUserNotFound
		}
		return domain.UserRes{}, err
	}

	return updatedUser, nil
}

// DeleteUser implements ports.UserRepository.
func (u UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}

// Login implements ports.UserRepository.
func (u UserRepository) Login(ctx context.Context, user domain.UserLoginReq) (domain.UserRes, error) {
	var loggedUser domain.UserRes
	var hashedPassword string
	query := "SELECT * FROM users WHERE username=$1"
	row := u.db.QueryRow(query, user.Username)
	err := row.Scan(&loggedUser.ID, &loggedUser.Username, &loggedUser.Email, &hashedPassword, &loggedUser.IsAdmin, &loggedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserRes{}, ErrUserNotFound
		}
		return domain.UserRes{}, err
	}

	if !util.ComparePassword(hashedPassword, user.Password) {
		return domain.UserRes{}, ErrInvalidCredentials
	}

	return loggedUser, nil
}
