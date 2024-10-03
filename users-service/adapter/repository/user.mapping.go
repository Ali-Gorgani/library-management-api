package repository

import "library-management-api/users-service/core/domain"

func MapUserResToUser(user *UserRes) *domain.User {
	return &domain.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
	}
}

func MapUsersResToUsers(user []*UserRes) []*domain.User {
	var users []*domain.User
	for _, u := range user {
		users = append(users, &domain.User{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   u.IsAdmin,
			CreatedAt: u.CreatedAt,
		})
	}
	return users
}
