package user

import "library-management-api/auth-service/pkg/proto"

func MapDtoGetUserReqToPbGetUserReq(req GetUserReq) *proto.GetUserReq {
	return &proto.GetUserReq{
		Username: req.UserName,
	}
}

func MapPbGetUserResToDtoGetUserRes(res *proto.UserRes) UserRes {
	return UserRes{
		ID:        uint(res.Id),
		Username:  res.Username,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: res.CreatedAt.AsTime(),
	}
}
