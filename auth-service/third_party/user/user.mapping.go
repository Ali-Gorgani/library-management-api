package user

import "library-management-api/pkg/proto/user"

func MapDtoGetUserReqToPbGetUserReq(req GetUserReq) *user.GetUserReq {
	return &user.GetUserReq{
		Username: req.Username,
	}
}

func MapPbGetUserResToDtoGetUserRes(res *user.UserRes) UserRes {
	return UserRes{
		ID:        uint(res.Id),
		Username:  res.Username,
		Password:  res.Password,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: res.CreatedAt.AsTime(),
	}
}
