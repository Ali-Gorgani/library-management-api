package usecase

import (
	"encoding/json"
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserUsecase struct {
	UserRepository ports.UserRepository
}

func NewUserUseCase() *UserUsecase {
	return &UserUsecase{
		UserRepository: repository.NewUserRepository(),
	}
}

// AddUser implements ports.UserRepository.
func (u UserUsecase) AddUser(w http.ResponseWriter, r *http.Request) {
	var user domain.AddUserParam
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	addedUser, err := u.UserRepository.AddUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(addedUser); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUsers implements ports.UserRepository.
func (u UserUsecase) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserRepository.GetUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser implements ports.UserRepository.
func (u UserUsecase) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := u.UserRepository.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateUser implements ports.UserRepository.
func (u UserUsecase) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user domain.UpdateUserParam
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedUser, err := u.UserRepository.UpdateUser(r.Context(), userID, user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser implements ports.UserRepository.
func (u UserUsecase) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = u.UserRepository.DeleteUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Login implements ports.UserRepository.
func (u UserUsecase) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.UserLoginParam
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	loggedInUser, err := u.UserRepository.Login(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	// TODO: Implement JWT token generation

	if err := json.NewEncoder(w).Encode(loggedInUser); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}