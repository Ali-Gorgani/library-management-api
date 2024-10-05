package repository

import (
	"database/sql"
	"library-management-api/users-service/core/domain"
)

type User struct {
	ID             uint
	Username       sql.NullString
	HashedPassword sql.NullString
	Email          sql.NullString
	IsAdmin        sql.NullBool
	CreatedAt      sql.NullTime
}

func MapUserEntityToUserDomain(user User) domain.User {
	return domain.User{
		ID:        user.ID,
		Username:  user.Username.String,
		Password:  user.HashedPassword.String,
		Email:     user.Email.String,
		IsAdmin:   user.IsAdmin.Bool,
		CreatedAt: user.CreatedAt.Time,
	}
}

func MapUsersEntityToUsersDomain(users []User) []domain.User {
	var usersDomain []domain.User
	for _, user := range users {
		usersDomain = append(usersDomain, MapUserEntityToUserDomain(user))
	}
	return usersDomain
}

func MapUserDomainToUserEntity(user domain.User) User {
	return User{
		ID:             user.ID,
		Username:       sql.NullString{String: user.Username, Valid: user.Username != ""},
		HashedPassword: sql.NullString{String: user.Password, Valid: user.Password != ""},
		Email:          sql.NullString{String: user.Email, Valid: user.Email != ""},
		IsAdmin:        sql.NullBool{Bool: user.IsAdmin, Valid: true},
		CreatedAt:      sql.NullTime{Time: user.CreatedAt, Valid: true},
	}
}
