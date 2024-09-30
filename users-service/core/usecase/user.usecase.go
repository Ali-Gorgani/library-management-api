package usecase

import (
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	userRepository ports.UserRepository
}

func NewUserUseCase() *UserUsecase {
	return &UserUsecase{
		userRepository: repository.NewUserRepository(),
	}
}

// errorResponse returns error details in JSON format.
func errorResponse(statusCode int, err error) gin.H {
	return gin.H{"status": statusCode, "error": err.Error()}
}

// AddUser handles POST requests for adding a new user
func (u *UserUsecase) AddUser(c *gin.Context) {
	var user domain.AddUserParam
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	addedUser, err := u.userRepository.AddUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, addedUser)
}

// GetUsers handles GET requests for retrieving all users
func (u *UserUsecase) GetUsers(c *gin.Context) {
	users, err := u.userRepository.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser handles GET requests for retrieving a single user by ID
func (u *UserUsecase) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	user, err := u.userRepository.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT requests for updating a user
func (u *UserUsecase) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	var user domain.UpdateUserParam
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	updatedUser, err := u.userRepository.UpdateUser(c.Request.Context(), userID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE requests for deleting a user
func (u *UserUsecase) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	err = u.userRepository.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Login handles POST requests for user login
func (u *UserUsecase) Login(c *gin.Context) {
	var user domain.UserLoginParam
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	loggedInUser, err := u.userRepository.Login(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	// TODO: Implement JWT token generation

	c.JSON(http.StatusOK, loggedInUser)
}
