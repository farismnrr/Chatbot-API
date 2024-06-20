package handler

import (
	"bytes"
	"capstone-project/api"
	"capstone-project/helper"
	"capstone-project/model"
	"capstone-project/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService    service.UserService
	sessionService service.SessionService
}

type UserHandler interface {
	GetServer(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	ResetPassword(c *gin.Context)
	RemoveUser(c *gin.Context)
	RequestAPI(c *gin.Context)
}

func NewUserHandler(userService service.UserService, sessionService service.SessionService) *userHandler {
	return &userHandler{userService: userService, sessionService: sessionService}
}

func (h *userHandler) GetServer(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func (h *userHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username is required"))
		return
	} else if user.Email == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is required"))
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
		return
	} else if user.FullName == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name is required"))
		return
	}

	if !helper.CheckPassLength(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must be at least 8 characters"))
		return
	} else if !helper.HasUpperLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 uppercase letter"))
		return
	} else if !helper.HasLowerLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 lowercase letter"))
		return
	} else if !helper.HasNumber(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 number"))
		return
	} else if !helper.HasSpecialChar(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 special character"))
		return
	}

	if !helper.CheckUsername(user.Username) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username should not contain spaces"))
		return
	} else if !helper.CheckEmail(user.Email) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is not valid"))
		return
	} else if !helper.CheckFullName(user.FullName) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name should not contain numbers or special characters"))
		return
	}

	hashedPassword := helper.GenerateHash(user.Password)
	user.Password = hashedPassword

	err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse(http.StatusCreated, "User registered successfully"))
}

func (h *userHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Username == "" && user.Email == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username or email is required"))
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
		return
	}

	if user.Username != "" {
		err := h.userService.GetUserByUsername(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Username Incorrect"))
			return
		}
	} else if user.Email != "" {
		err := h.userService.GetUserByEmail(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, "Email Incorrect"))
			return
		}
	}

	dbUser, err := h.userService.GetUserTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	hashedPassword := helper.GenerateHash(user.Password)
	user.Password = hashedPassword

	if user.Password != dbUser.Password {
		c.JSON(http.StatusForbidden, model.NewErrorResponse(http.StatusForbidden, "Incorrect password"))
		return
	}

	_, err = h.sessionService.GenerateSession(c, dbUser.ID, dbUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	//Dummy Get Token
	token, err := h.sessionService.GetSession(c, dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewJWTSuccessResponse(http.StatusOK, "Login successful", []model.JWTResponse{{Token: token}}))
}

func (h *userHandler) ResetPassword(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is required"))
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
		return
	}

	if !helper.CheckPassLength(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must be at least 8 characters"))
		return
	} else if !helper.HasUpperLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 uppercase letter"))
		return
	} else if !helper.HasLowerLetter(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 lowercase letter"))
		return
	} else if !helper.HasNumber(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 number"))
		return
	} else if !helper.HasSpecialChar(user.Password) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 special character"))
		return
	}

	hashedPassword := helper.GenerateHash(user.Password)
	user.Password = hashedPassword

	err := h.userService.UpdateUserByEmail(user.Email, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Password reset successful"))
}

func (h *userHandler) RemoveUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	err = h.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	_, err = h.userService.GetUserTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = h.sessionService.DeleteSession(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = h.userService.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "User removed successfully"))
}

func (h *userHandler) RequestAPI(c *gin.Context) {
	var response model.Message
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	body, err := api.FetchAPI(response.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, prettyJSON.String()))
}