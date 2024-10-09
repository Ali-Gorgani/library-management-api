package repository

import (
	"context"
	"database/sql"
	"errors"
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
func (u *UserRepository) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	var addedUser User
	mappedUser := MapUserDomainToUserEntity(user)

	query := "INSERT INTO users (username, hashed_password, email, is_admin) VALUES ($1, $2, $3, $4) RETURNING *"
	row := u.db.QueryRow(query, mappedUser.Username, mappedUser.HashedPassword, mappedUser.Email, mappedUser.IsAdmin)
	err := row.Scan(&addedUser.ID, &addedUser.Username, &addedUser.HashedPassword, &addedUser.Email, &addedUser.IsAdmin, &addedUser.CreatedAt)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)" {
			return domain.User{}, errorhandler.ErrDuplicateUsername
		}
		return domain.User{}, err
	}
	res := MapUserEntityToUserDomain(addedUser)
	return res, nil
}

// GetUsers implements ports.UserRepository.
func (u *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []User
	query := "SELECT * FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Email, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}
	res := MapUsersEntityToUsersDomain(users)
	return res, nil
}

// GetUserByID implements ports.UserRepository.
func (u *UserRepository) GetUserByID(ctx context.Context, user domain.User) (domain.User, error) {
	var foundUser User
	mappedUser := MapUserDomainToUserEntity(user)

	query := "SELECT * FROM users WHERE id=$1"
	row := u.db.QueryRow(query, mappedUser.ID)
	err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.HashedPassword, &foundUser.Email, &foundUser.IsAdmin, &foundUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errorhandler.ErrUserNotFound
		}
		return domain.User{}, err
	}
	res := MapUserEntityToUserDomain(foundUser)
	return res, nil
}

// GetUserByUsername implements ports.UserRepository.
func (u *UserRepository) GetUserByUsername(ctx context.Context, user domain.User) (domain.User, error) {
	var foundUser User
	mappedUser := MapUserDomainToUserEntity(user)

	query := "SELECT * FROM users WHERE username=$1"
	row := u.db.QueryRow(query, mappedUser.Username.String)
	err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.HashedPassword, &foundUser.Email, &foundUser.IsAdmin, &foundUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errorhandler.ErrUserNotFound
		}
		return domain.User{}, err
	}
	res := MapUserEntityToUserDomain(foundUser)
	return res, nil
}

// UpdateUser implements ports.UserRepository.
func (u *UserRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	var updatedUser User
	mappedUser := MapUserDomainToUserEntity(user)

	query := "UPDATE users SET username=$1, hashed_password=$2, email=$3, is_admin=$4 WHERE id=$5 RETURNING *"
	row := u.db.QueryRow(query, mappedUser.Username, mappedUser.HashedPassword, mappedUser.Email, mappedUser.IsAdmin, mappedUser.ID)
	err := row.Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.HashedPassword, &updatedUser.Email, &updatedUser.IsAdmin, &updatedUser.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errorhandler.ErrUserNotFound
		}
		return domain.User{}, err
	}
	res := MapUserEntityToUserDomain(updatedUser)
	return res, nil
}

// DeleteUser implements ports.UserRepository.
func (u *UserRepository) DeleteUser(ctx context.Context, user domain.User) error {
	mappedUser := MapUserDomainToUserEntity(user)
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Exec(query, mappedUser.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrUserNotFound
		}
		return err
	}
	return nil
}
