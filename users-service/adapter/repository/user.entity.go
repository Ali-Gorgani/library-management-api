package repository

import (
	"database/sql"
	"library-management-api/users-service/core/domain"
)

type User struct {
	ID        uint
	Username  sql.NullString
	Email     sql.NullString
	IsAdmin   sql.NullBool
	CreatedAt sql.NullTime
}

func MapUserEntityToUserDomain(user User) domain.User {
	return domain.User{
		ID:        user.ID,
		Username:  user.Username.String,
		Email:     user.Email.String,
		IsAdmin:   user.IsAdmin.Bool,
		CreatedAt: user.CreatedAt.Time,
	}
}

func MapUsersEntityToUsersDomain(user []User) []domain.User {
	var users []domain.User
	for _, u := range user {
		users = append(users, domain.User{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   u.IsAdmin,
			CreatedAt: u.CreatedAt,
		})
	}
	return users
}

func MapUserDomainToUserEntity(user domain.User) User {
	return User{
		ID:        user.ID,
		Username:  sql.NullString{String: user.Username, Valid: user.Username != ""},
		Email:     sql.NullString{String: user.Email, Valid: user.Email != ""},
		IsAdmin:   sql.NullBool{Bool: user.IsAdmin, Valid: true},
		CreatedAt: sql.NullTime{Time: user.CreatedAt, Valid: !user.CreatedAt.IsZero()},
	}
}
