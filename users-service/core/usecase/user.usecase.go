package usecase

import (
	"context"
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/util/errorhandler"
)

type UserUseCase struct {
	userRepository ports.UserRepository
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{
		userRepository: repository.NewUserRepository(),
	}
}

// AddUser handles logic for adding a new user
func (u *UserUseCase) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	// TODO: hash password before adding user data to database with gRPC from auth-service
	newUser, err := u.userRepository.AddUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}

// GetUsers handles logic for retrieving all users
func (u *UserUseCase) GetUsers(ctx context.Context) ([]domain.User, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return []domain.User{}, errorhandler.ErrInvalidSession
	}

	if !claims.IsAdmin {
		return []domain.User{}, errorhandler.ErrForbidden
	}

	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser handles logic for retrieving a single user
func (u *UserUseCase) GetUser(ctx context.Context, user domain.User) (domain.User, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.User{}, errorhandler.ErrInvalidSession
	}

	if !claims.IsAdmin {
		return domain.User{}, errorhandler.ErrForbidden
	}

	getUser, err := u.userRepository.GetUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return getUser, nil
}

// UpdateUser handles logic for updating a user
func (u *UserUseCase) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.User{}, errorhandler.ErrInvalidSession
	}

	if claims.ID != user.ID && !claims.IsAdmin {
		return domain.User{}, errorhandler.ErrForbidden
	}

	// TODO: hash password before updating user data to database with gRPC from auth-service
	updatedUser, err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return updatedUser, nil
}

// DeleteUser handles logic for deleting a user
func (u *UserUseCase) DeleteUser(ctx context.Context, user domain.User) error {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return errorhandler.ErrInvalidSession
	}

	if claims.ID != user.ID && !claims.IsAdmin {
		return errorhandler.ErrForbidden
	}

	err = u.userRepository.DeleteUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
