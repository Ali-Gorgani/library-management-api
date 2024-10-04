package user

import (
	"library-management-api/users-service/api/http"
	"library-management-api/users-service/core/domain"
)

func MapDomainUserToAddUserReq(req *domain.User) *http.AddUserReq {
	return &http.AddUserReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapUserResToDomainUser(res *http.UserRes) *domain.User {
	return &domain.User{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: res.CreatedAt,
	}
}

func MapUsersResToDomainUsers(res []*http.UserRes) []*domain.User {
	var users []*domain.User
	for _, user := range res {
		users = append(users, MapUserResToDomainUser(user))
	}
	return users
}

func MapDomainUserToGetUserReq(req *domain.User) *http.GetUserReq {
	return &http.GetUserReq{
		ID: req.ID,
	}
}

func MapDomainUserToUpdateUserReq(req *domain.User) *http.UpdateUserReq {
	return &http.UpdateUserReq{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapDomainUserToDeleteUserReq(req *domain.User) *http.DeleteUserReq {
	return &http.DeleteUserReq{
		ID: req.ID,
	}
}
