package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strings"
	"user-api/internal/errors"
	"user-api/internal/handler/types"
	"user-api/internal/service"
	"user-api/internal/storage"
	"user-api/pkg/utils"
)

type Handler struct {
	Repository *storage.Repository
	Service    *service.Service
}

func NewHandler(baseGroup *echo.Group, repo *storage.Repository, service *service.Service) {

	h := Handler{Repository: repo, Service: service}

	baseGroup.POST("users", h.CreateUser)
	baseGroup.GET("users", h.GetAllUser)
	baseGroup.GET("users/:userId", h.GetUser)
	baseGroup.PUT("users/:userId", h.UpdateUser)
	baseGroup.DELETE("users/:userId", h.DeleteUser)

}

// CreateUser
// @Summary Create User
// @Tags    Users
// @Accept  json
// @Param   User 		body    types.UserCreateRequest  true  "name"
// @Success 200        {object} types.User
// @Failure 409        {object} errors.Error
// @Failure 400        {object} errors.Error
// @Failure 500        {object} errors.Error
// @Router  /users [POST]
func (h Handler) CreateUser(c echo.Context) error {

	req := new(types.UserCreateRequest)

	if err := c.Bind(req); err != nil {
		panic(errors.InvalidCreateUserBody)
	}

	createdUser := h.Service.CreateUser(req)

	return c.JSON(200, createdUser)

}

// GetUser
// @Summary Get User With User Id
// @Tags    Users
// @Accept  json
// @Param   userId   path  string  true  "userId"
// @Success 200        {object} types.User
// @Failure 400        {object} errors.Error
// @Failure 404        {object} errors.Error
// @Failure 500        {object} errors.Error
// @Router  /users/{userId} [GET]
func (h Handler) GetUser(c echo.Context) error {

	userId := strings.TrimSpace(c.Param("userId"))

	if _, err := uuid.Parse(userId); err != nil {
		panic(errors.UUIDParseError.WrapDesc(err.Error()))
	}

	user := h.Service.GetUser(userId)

	return c.JSON(200, user)
}

// UpdateUser
// @Summary Update User With User Id
// @Tags    Users
// @Accept  json
// @Param   userId path    string  true  "userId"
// @Param   User 		body    types.UserUpdateRequest  true  "name"
// @Success 200        {object} types.User
// @Failure 400        {object} errors.Error
// @Failure 409        {object} errors.Error
// @Failure 500        {object} errors.Error
// @Router  /users/{userId} [PUT]
func (h Handler) UpdateUser(c echo.Context) error {

	userId := strings.TrimSpace(c.Param("userId"))

	if _, err := uuid.Parse(userId); err != nil {
		panic(errors.UUIDParseError.WrapDesc(err.Error()))
	}

	req := new(types.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		panic(errors.InvalidUpdateUserBody)
	}

	updatedUser := h.Service.UpdateUser(userId, req)

	return c.JSON(200, updatedUser)
}

// DeleteUser
// @Summary Delete User With User Id
// @Tags    Users
// @Accept  json
// @Param   userId path    string  true  "userId"
// @Success 200        {string}  string    "ok"
// @Failure 400        {object} errors.Error
// @Failure 404        {object} errors.Error
// @Failure 500        {object} errors.Error
// @Router  /users/{userId} [DELETE]
func (h Handler) DeleteUser(c echo.Context) error {

	userId := strings.TrimSpace(c.Param("userId"))

	if _, err := uuid.Parse(userId); err != nil {
		panic(errors.UUIDParseError.WrapDesc(err.Error()))
	}

	h.Service.DeleteUser(userId)

	return c.JSON(200, "User Deleted")
}

// GetAllUser
// @Summary Get All User
// @Tags    Users
// @Accept  json
// @Success 200        {object} types.GetAllUsersResponse
// @Failure 500        {object} errors.Error
// @Router  /users [GET]
func (h Handler) GetAllUser(c echo.Context) error {

	limitParam := strings.TrimSpace(c.Param("limit"))
	offsetParam := strings.TrimSpace(c.Param("offset"))

	limit, offset := utils.PaginationWithDefaults(limitParam, offsetParam)

	users := h.Service.GetAll(limit, offset)

	return c.JSON(200, users)
}
