package model

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
