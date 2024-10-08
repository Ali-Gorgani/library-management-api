package user

import (
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/third_party/user"
)

func MapDomainUserToDtoGetUserReq(req domain.Auth) user.GetUserReq {
	return user.GetUserReq{
		UserName: req.Username,
	}
}

func MapDtoUserResToDomainUser(res user.UserRes) domain.User {
	return domain.User{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: res.CreatedAt,
	}
}
