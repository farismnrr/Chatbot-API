package model

import (
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret-key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type JWTSuccessResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []JWTResponse `json:"data"`
}

type JWTResponse struct {
	Token string `json:"token"`
}

func NewJWTSuccessResponse(code int, message string, data []JWTResponse) *JWTSuccessResponse {
	return &JWTSuccessResponse{Code: code, Message: message, Data: data}
}
