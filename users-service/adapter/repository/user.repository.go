package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
	"log"

	"golang.org/x/crypto/bcrypt"
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
	hashedPassword := HashedPassword(user.Password)
	var addedUser domain.User
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *"
	row := u.db.QueryRow(query, user.Username, user.Email, hashedPassword)
	err := row.Scan(&addedUser.ID, &addedUser.Username, &addedUser.Email, &addedUser.HashedPassword, &addedUser.CreatedAt)
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedAt)
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
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// UpdateUser implements ports.UserRepository.
func (u UserRepository) UpdateUser(ctx context.Context, id int, user domain.UpdateUserParam) (domain.User, error) {
	var updatedUser domain.User
	query := "UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4 RETURNING *"
	row := u.db.QueryRow(query, user.Username, user.Email, user.Password, id)
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.HashedPassword, &updatedUser.CreatedAt)
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
	err := row.Scan(&loggedUser.ID, &loggedUser.Username, &loggedUser.Email, &loggedUser.HashedPassword, &loggedUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	if !ComparePassword(loggedUser.HashedPassword, user.Password) {
		return domain.User{}, errors.New("invalid credentials")
	}

	return loggedUser, nil
}

// HashedPassword implements ports.UserRepository.
func HashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}

// ComparePassword implements ports.UserRepository.
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Fatalf("Failed to compare password: %v", err)
		return false
	}
	return true
}
