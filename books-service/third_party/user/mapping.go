package user

import (
	"library-management-api/books-service/pkg/user/pb"
	"library-management-api/users-service/api/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapAddUserReqToPb(req *http.AddUserReq) *pb.AddUserReq {
	return &pb.AddUserReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapUserResToDto(res *pb.UserRes) *http.UserRes {
	return &http.UserRes{
		ID:        int(res.Id),
		Username:  res.Username,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: mapTimestampToTime(res.CreatedAt),
	}
}

func MapUsersResToDto(res grpc.ServerStreamingClient[pb.UserRes]) []*http.UserRes {
	var users []*http.UserRes
	for {
		user, err := res.Recv()
		if err != nil {
			break
		}
		users = append(users, MapUserResToDto(user))
	}
	return users
}

func MapGetUserReqToPb(req *http.GetUserReq) *pb.GetUserReq {
	return &pb.GetUserReq{
		Id: int32(req.ID),
	}
}

func MapUpdateUserReqToPb(req *http.UpdateUserReq) *pb.UpdateUserReq {
	return &pb.UpdateUserReq{
		Id:       int32(req.ID),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}
}

func MapDeleteUserReqToPb(req *http.DeleteUserReq) *pb.DeleteUserReq {
	return &pb.DeleteUserReq{
		Id: int32(req.ID),
	}
}

func mapTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	if err := ts.CheckValid(); err != nil {
		return time.Time{}
	}
	t := ts.AsTime()
	return t
}
