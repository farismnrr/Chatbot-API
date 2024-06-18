package handler

import (
	"capstone-project/model"
	"capstone-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	service service.UserService
}

type UserHandler interface {
	GetServer(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	RemoveUser(c *gin.Context)
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

	username, err := h.service.GetUserByUsername(user.Username)
	if err != nil && username != nil {
		c.JSON(http.StatusConflict, model.NewErrorResponse(http.StatusConflict, "Username is already taken"))
		return
	}

	email, err := h.service.GetUserByEmail(user.Email)
	if err != nil && email != nil {
		c.JSON(http.StatusConflict, model.NewErrorResponse(http.StatusConflict, "Email is already taken"))
		return
	}

	user.Password = h.service.GenerateHash(user.Password)

	if err := h.service.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to register user"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse(http.StatusCreated, "User registered successfully"))
}

func (h *Handler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username or email is required"))
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
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

	if !h.service.CheckUsername(user.Username) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username should not contain spaces"))
		return
	}

	dbUser, err := h.service.GetUserByUsername(user.Username)
	if err != nil && dbUser == nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	user.Password = h.service.GenerateHash(user.Password)

	if user.Password != dbUser.Password {
		c.JSON(http.StatusForbidden, model.NewErrorResponse(http.StatusForbidden, "Incorrect password"))
		return
	}

	tokenString, err := h.service.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Failed to generate user token"))
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	token, err := jwt.ParseWithClaims(*tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "error internal server"))
		return
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "invalid token"))
		return
	}
	claims.ExpiresAt = expirationTime.Unix()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   *tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Login successful"))
}

func (h *Handler) RemoveUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	_, err = h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	err = h.service.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Internal server error"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "User removed successfully"))
}
