package usecase

import (
	"context"
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
)

type UserUsecase struct {
	userRepository ports.UserRepository
}

func NewUserUseCase() *UserUsecase {
	return &UserUsecase{
		userRepository: repository.NewUserRepository(),
	}
}

// AddUser handles logic for adding a new user
func (u *UserUsecase) AddUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	newUser, err := u.userRepository.AddUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}
	return newUser, nil
}

// GetUsers handles logic for retrieving all users
func (u *UserUsecase) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser handles logic for retrieving a single user
func (u *UserUsecase) GetUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	getUser, err := u.userRepository.GetUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}
	return getUser, nil
}

// UpdateUser handles logic for updating a user
func (u *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	updatedUser, err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}
	return updatedUser, nil
}

// DeleteUser handles logic for deleting a user
func (u *UserUsecase) DeleteUser(ctx context.Context, user *domain.User) error {
	err := u.userRepository.DeleteUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
