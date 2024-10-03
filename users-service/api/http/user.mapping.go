package http

import "library-management-api/users-service/core/domain"

func MapUserToUserRes(user *domain.User) *UserRes {
	return &UserRes{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
	}
}

func MapUsersToUsersRes(users []*domain.User) []*UserRes {
	var usersRes []*UserRes
	for _, u := range users {
		usersRes = append(usersRes, &UserRes{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			IsAdmin:   u.IsAdmin,
			CreatedAt: u.CreatedAt,
		})
	}
	return usersRes
}

func MapAddUserReqToUser(req *AddUserReq) *domain.User {
	return &domain.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapGetUserReqToUser(req *GetUserReq) *domain.User {
	return &domain.User{
		ID: req.ID,
	}
}

func MapUpdateUserReqToUser(req *UpdateUserReq) *domain.User {
	return &domain.User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapDeleteUserReqToUser(req *DeleteUserReq) *domain.User {
	return &domain.User{
		ID: req.ID,
	}
}
