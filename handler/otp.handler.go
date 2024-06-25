package handler

import (
	"capstone-project/api"
	"capstone-project/model"
	"capstone-project/service"
	"net/http"
	"strconv"

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

	otp, err := h.otpService.GenerateOTP(c, dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = api.SendMailSimple("Verification Code", otp.OTPCode, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewOTPResponse(http.StatusOK, "OTP sent successfully", []model.OTP{{UserID: dbUser.ID, Expiry: otp.Expiry}}))
}

func (h *otpHandler) VerifyOTP(c *gin.Context) {
	var otp model.OTP
	err := c.ShouldBindJSON(&otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	if otp.OTPCode == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "OTP code is required"))
		return
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	err = h.userService.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	_, err = h.userService.GetUserTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	getOTP, err := h.otpService.GetOTP(c, userID, otp.OTPCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	if getOTP != otp.OTPCode {
		c.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "OTP code is invalid"))
		return
	}

	c.JSON(http.StatusOK, model.NewOTPResponse(http.StatusOK, "OTP verified successfully", []model.OTP{{UserID: userID, OTPCode: getOTP}}))
}
