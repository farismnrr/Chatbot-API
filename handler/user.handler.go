package handler

import (
	"capstone-project/model"
	"capstone-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.UserService
}

type UserHandler interface {
	GetServer(c *gin.Context)
	Register(c *gin.Context)
}

func NewUserHandler(service service.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetServer(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username is required"))
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
		return
	} else if user.FullName == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name is required"))
		return
	} else if user.Email == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is required"))
		return
	}

	if !h.service.CheckPassLength(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must be at least 8 characters"))
		return
	} else if !h.service.HasUpperLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 uppercase letter"))
		return
	} else if !h.service.HasLowerLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 lowercase letter"))
		return
	} else if !h.service.HasNumber(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 number"))
		return
	} else if !h.service.HasSpecialChar(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 special character"))
		return
	}

	if !h.service.CheckEmail(user.Email) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is not valid"))
		return
	} else if !h.service.CheckUsername(user.Username) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username should not contain spaces"))
		return
	} else if !h.service.CheckFullName(user.FullName) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name should not contain numbers or special characters"))
		return
	}

	if err := h.service.GetUserByUsername(user); err == nil {
		c.JSON(http.StatusConflict, model.NewErrorResponse(http.StatusConflict, "Username is already taken"))
		return
	} else if err := h.service.GetUserByEmail(user); err == nil {
		c.JSON(http.StatusConflict, model.NewErrorResponse(http.StatusConflict, "Email is already taken"))
		return
	}

	if err := h.service.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Internal server error"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse(http.StatusCreated, "User registered successfully"))
}
