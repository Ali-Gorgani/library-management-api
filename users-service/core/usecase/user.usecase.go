package usecase

import (
	"context"
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/adapter/service/auth"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/util/errorhandler"
)

type UserUseCase struct {
	userRepository ports.UserRepository
	authService    *auth.AuthService
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{
		userRepository: repository.NewUserRepository(),
		authService:    auth.NewAuthService(),
	}
}

// AddUser handles logic for adding a new user
func (u *UserUseCase) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPasswordReq := domain.Auth{
		Password: user.Password,
	}
	hashedPasswordRes, err := u.authService.HashedPassword(ctx, hashedPasswordReq)
	hashedPassword, err := u.authService.HashedPassword(ctx, hashedPasswordRes)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPassword.Password
	newUser, err := u.userRepository.AddUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

// GetUsers handles logic for retrieving all users
func (u *UserUseCase) GetUsers(ctx context.Context) ([]domain.User, error) {
	contextToken := ctx.Value("token").(string)

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := u.authService.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return []domain.User{}, errorhandler.ErrInvalidSession
	}
	claims := verifyTokenRes.Claims

	if !claims.IsAdmin {
		return []domain.User{}, errorhandler.ErrForbidden
	}

	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByUsername handles logic for retrieving a single user
func (u *UserUseCase) GetUserByUsername(ctx context.Context, user domain.User) (domain.User, error) {
	contextToken := ctx.Value("token").(string)

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := u.authService.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return domain.User{}, errorhandler.ErrInvalidSession
	}
	claims := verifyTokenRes.Claims

	if !claims.IsAdmin {
		return domain.User{}, errorhandler.ErrForbidden
	}

	foundUser, err := u.userRepository.GetUserByUsername(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return foundUser, nil
}

// UpdateUser handles logic for updating a user
func (u *UserUseCase) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	contextToken := ctx.Value("token").(string)

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := u.authService.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return domain.User{}, errorhandler.ErrInvalidSession
	}
	claims := verifyTokenRes.Claims

	if claims.ID != user.ID && !claims.IsAdmin {
		return domain.User{}, errorhandler.ErrForbidden
	}

	hashedPasswordReq := domain.Auth{
		Password: user.Password,
	}
	hashedPasswordRes, err := u.authService.HashedPassword(ctx, hashedPasswordReq)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPasswordRes.Password
	updatedUser, err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return updatedUser, nil
}

// DeleteUser handles logic for deleting a user
func (u *UserUseCase) DeleteUser(ctx context.Context, user domain.User) error {
	contextToken := ctx.Value("token").(string)

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := u.authService.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return errorhandler.ErrInvalidSession
	}
	claims := verifyTokenRes.Claims

	if claims.ID != user.ID && !claims.IsAdmin {
		return errorhandler.ErrForbidden
	}

	err = u.userRepository.DeleteUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
