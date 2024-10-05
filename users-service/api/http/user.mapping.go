package http

import "library-management-api/users-service/core/domain"

func MapDomainUserToDtoUserRes(user domain.User) UserRes {
	return UserRes{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
	}
}

func MapDomainUsersToDtoUsersRes(users []domain.User) []UserRes {
	var usersRes []UserRes
	for _, user := range users {
		usersRes = append(usersRes, MapDomainUserToDtoUserRes(user))
	}
	return usersRes
}

func MapDtoAddUserReqToDomainUser(req AddUserReq) domain.User {
	return domain.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapDtoGetUserReqToDomainUser(req GetUserReq) domain.User {
	return domain.User{
		ID: req.ID,
	}
}

func MapDtoUpdateUserReqToDomainUser(req UpdateUserReq) domain.User {
	return domain.User{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapDtoDeleteUserReqToDomainUser(req DeleteUserReq) domain.User {
	return domain.User{
		ID: req.ID,
	}
}
