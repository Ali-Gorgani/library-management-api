package http

import (
	"errors"
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

	addedUser, err := uc.userUseCase.AddUser(c.Request.Context(), MapDtoAddUserReqToDomainUser(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapDomainUserToDtoUserRes(addedUser)
	c.JSON(http.StatusCreated, res)
}

// GetUsers handles GET requests for retrieving all users
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userUseCase.GetUsers(c.Request.Context())
	if err != nil {
		if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapDomainUsersToDtoUsersRes(users)
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

	getUserReq := GetUserReq{
		ID: uint(userID),
	}

	user, err := uc.userUseCase.GetUser(c.Request.Context(), MapDtoGetUserReqToDomainUser(getUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapDomainUserToDtoUserRes(user)

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

	var updateUserReq UpdateUserReq
	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	updateUserReq.ID = uint(userID)

	updatedUser, err := uc.userUseCase.UpdateUser(c.Request.Context(), MapDtoUpdateUserReqToDomainUser(updateUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapDomainUserToDtoUserRes(updatedUser)
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

	deleteUserReq := DeleteUserReq{
		ID: uint(userID),
	}

	err = uc.userUseCase.DeleteUser(c.Request.Context(), MapDtoDeleteUserReqToDomainUser(deleteUserReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
