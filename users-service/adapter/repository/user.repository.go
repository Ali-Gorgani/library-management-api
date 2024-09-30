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

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() ports.UserRepository {
	return &UserRepository{
		db: database.P().DB,
	}
}

// AddUser implements ports.UserRepository.
func (u UserRepository) AddUser(ctx context.Context, user domain.AddUserParam) (domain.User, error) {
	hashedPassword := util.HashedPassword(user.Password)
	var addedUser domain.User
	query := "INSERT INTO users (username, email, hashed_password) VALUES ($1, $2, $3) RETURNING *"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword)
	err := row.Scan(&addedUser.ID, &addedUser.Username, &addedUser.Email, &addedUser.HashedPassword, &addedUser.Role, &addedUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return addedUser, nil
}

// GetUsers implements ports.UserRepository.
func (u UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	var user domain.User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUser implements ports.UserRepository.
func (u UserRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE id=$1"
	row := u.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword, &user.Role, &user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// UpdateUser implements ports.UserRepository.
func (u UserRepository) UpdateUser(ctx context.Context, id int, user domain.UpdateUserParam) (domain.User, error) {
	var updatedUser domain.User
	hashedPassword := util.HashedPassword(user.Password)
	query := "UPDATE users SET username=$1, email=$2, hashed_password=$3, role=$4 WHERE id=$5 RETURNING *"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword, user.Role, id)
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.HashedPassword, &updatedUser.Role, &updatedUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return updatedUser, nil
}

// DeleteUser implements ports.UserRepository.
func (u UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// Login implements ports.UserRepository.
func (u UserRepository) Login(ctx context.Context, user domain.UserLoginParam) (domain.User, error) {
	var loggedUser domain.User
	query := "SELECT * FROM users WHERE username=$1"
	row := u.db.QueryRow(query, user.Username)
	err := row.Scan(&loggedUser.ID, &loggedUser.Username, &loggedUser.Email, &loggedUser.HashedPassword, &loggedUser.Role, &loggedUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	if !util.ComparePassword(loggedUser.HashedPassword, user.Password) {
		return domain.User{}, errors.New("invalid credentials")
	}

	return loggedUser, nil
}
