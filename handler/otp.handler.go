package handler

import (
	"capstone-project/helper"
	"capstone-project/model"
	"capstone-project/service"
	"capstone-project/smtp"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type otpHandler struct {
	otpService  service.OTPService
	userService service.UserService
}

type OTPHandler interface {
	SendOTP(c *gin.Context)
	VerifyOTP(c *gin.Context)
}

func NewOTPHandler(otpService service.OTPService, userService service.UserService) *otpHandler {
	return &otpHandler{otpService: otpService, userService: userService}
}

func (h *otpHandler) SendOTP(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is required"))
		return
	}

	if user.Email != "" {
		err := h.userService.GetUserByEmail(user)
		if err != nil {
			c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "Email not found"))
			return
		}
	}

	dbUser, err := h.userService.GetUserTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	otpCode := helper.GenerateOTPCode()
	err = h.otpService.SetOTP(user.Email, otpCode, time.Now().Add(time.Minute*5))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = smtp.SendMailSimple("Verification Code", otpCode, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewOTPResponse(http.StatusOK, "OTP sent successfully", []model.OTP{{UserID: dbUser.ID}}))
}

func (h *otpHandler) VerifyOTP(c *gin.Context) {
	// var otp model.OTP
	// err := c.ShouldBindJSON(&otp)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
	// 	return
	// }

	// if otp.OTPCode == "" {
	// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "OTP code is required"))
	// 	return
	// }

	// userID, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
	// 	return
	// }

	// err = h.userService.GetUserById(userID)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
	// 	return
	// }

	// _, err = h.userService.GetUserTable()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	// 	return
	// }

	// verified, err := h.otpService.ValidateOTP(strconv.Itoa(userID), otp.OTPCode)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	// 	return
	// }

	// if !verified {
	// 	c.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "OTP code is invalid"))
	// 	return
	// }

	// c.JSON(http.StatusOK, model.NewOTPResponse(http.StatusOK, "OTP verified successfully", []model.OTP{{UserID: userID}})
}
