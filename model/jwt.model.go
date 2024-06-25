package model

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	jwt.StandardClaims
}

type JWTSuccessResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []JWTResponse `json:"data"`
}

type JWTResponse struct {
	UserID int    `json:"user_id"`
}

func NewJWTSuccessResponse(code int, message string, data []JWTResponse) *JWTSuccessResponse {
	return &JWTSuccessResponse{Code: code, Message: message, Data: data}
}

var TokenBlacklist = make(map[string]bool)