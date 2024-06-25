package model

import "time"

type OTP struct {
	UserID  int       `json:"user_id,omitempty"`
	OTPCode string    `json:"otp_code,omitempty"`
	Expiry  time.Time `json:"expiry,omitempty"`
}

type OTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []OTP  `json:"data,omitempty"`
}

func NewOTPResponse(code int, message string, data []OTP) *OTPResponse {
	return &OTPResponse{Code: code, Message: message, Data: data}
}
