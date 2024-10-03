package http

import (
	"errors"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/users-service/core/usecase"
	"library-management-api/util/errorhandler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase *usecase.UserUsecase
}

func NewUserController() *UserController {
	return &UserController{
		userUseCase: usecase.NewUserUseCase(),
	}
}

func (uc *UserController) AddUser(c *gin.Context) {
	var user AddUserReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	addedUser, err := uc.userUseCase.AddUser(c.Request.Context(), MapAddUserReqToUser(&user))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	res := MapUserToUserRes(addedUser)

	c.JSON(http.StatusCreated, res)

}

// GetUsers handles GET requests for retrieving all users
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapUsersToUsersRes(users)

	c.JSON(http.StatusOK, res)
}

// GetUser handles GET requests for retrieving a single user by ID
func (uc *UserController) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	getUserReq := &GetUserReq{
		ID: userID,
	}
	
	user, err := uc.userUseCase.GetUser(c.Request.Context(), MapGetUserReqToUser(getUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapUserToUserRes(user)

	c.JSON(http.StatusOK, res)
}

// UpdateUser handles PUT requests for updating a user
func (uc *UserController) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)
	if claims.ID != userID && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		return
	}

	var user UpdateUserReqToBind
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	updateUserReq := &UpdateUserReq{
		ID:       userID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
	}

	updatedUser, err := uc.userUseCase.UpdateUser(c.Request.Context(), MapUpdateUserReqToUser(updateUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapUserToUserRes(updatedUser)

	c.JSON(http.StatusOK, res)
}

// DeleteUser handles DELETE requests for deleting a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)
	if claims.ID != userID && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		return
	}
	
	deleteUserReq := &DeleteUserReq{
		ID: userID,
	}

	err = uc.userUseCase.DeleteUser(c.Request.Context(), MapDeleteUserReqToUser(deleteUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
