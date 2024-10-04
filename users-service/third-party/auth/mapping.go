package auth

import (
	"library-management-api/auth-service/api/http"
	"library-management-api/users-service/pkg/auth/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapPbAuthLoginResToDto(res *pb.AuthLoginRes) *http.AuthLoginRes {
	return &http.AuthLoginRes{
		SessionID:             res.SessionId,
		AccessToken:           res.AccessToken,
		RefreshToken:          res.RefreshToken,
		AccessTokenExpiresAt:  mapTimestampToTime(res.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: mapTimestampToTime(res.RefreshTokenExpiresAt),
		User: http.UserRes{
			ID:        int(res.User.Id),
			Username:  res.User.Username,
			Email:     res.User.Email,
			IsAdmin:   res.User.IsAdmin,
			CreatedAt: mapTimestampToTime(res.User.CreatedAt),
		},
	}
}

func MapDtoAuthLoginReqToPb(req *http.AuthLoginReq) *pb.AuthLoginReq {
	return &pb.AuthLoginReq{
		Username: req.Username,
		Password: req.Password,
	}
}

func MapDtoAuthLogoutReqToPb(req *http.AuthLogoutReq) *pb.AuthLogoutReq {
	return &pb.AuthLogoutReq{
		SessionId: req.SessionID,
	}
}

func MapDtoAuthRefreshTokenReqToPb(req *http.AuthRefreshTokenReq) *pb.AuthRefreshTokenReq {
	return &pb.AuthRefreshTokenReq{
		RefreshToken: req.RefreshToken,
		UserId:       int32(req.UserID),
	}
}

func MapPbAuthRefreshTokenResToDto(res *pb.AuthRefreshTokenRes) *http.AuthRefreshTokenRes {
	return &http.AuthRefreshTokenRes{
		AccessToken:          res.AccessToken,
		AccessTokenExpiresAt: mapTimestampToTime(res.AccessTokenExpiresAt),
	}
}

func MapDtoAuthRevokeTokenReqToPb(req *http.AuthRevokeTokenReq) *pb.AuthRevokeTokenReq {
	return &pb.AuthRevokeTokenReq{
		SessionId: req.SessionID,
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
