package helper

import (
	"capstone-project/model"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := os.Getenv(key)
	if env == "" {
		panic("Environment variable not found")
	}

	return env
}

func UseCertificate() (key string, err error) {
	filePath := "ca.pem"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read the certificate file: %v", err)
		return "", err
	}
	key = string(fileContent)
	return key, nil
}

func GenerateToken(username string, role string, status string) (string, error) {
	claims := model.Claims{
		Username: username,
		Role:     role,
		Status:   status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
			Issuer:    "AskIN!",
			IssuedAt:  time.Now().Unix(),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := UseCertificate()
	if err != nil {
		return "", err
	}
	token, err := unsignedToken.SignedString([]byte(jwtToken))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*model.Claims, error) {
	claims := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtKey, err := UseCertificate()
		if err != nil {
			return nil, err
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	if _, ok := model.TokenBlacklist[token]; ok {
		return nil, errors.New("token has been used before")
	}

	return claims, nil
}

func ValidateUser(token string) (string, string, string, error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return "", "", "", err
	}
	return claims.Username, claims.Role, claims.Status, nil
}

func GenerateOTPCode() string {
	rand.Seed(time.Now().UnixNano())
	digits := []rune("0123456789")
	otpCode := make([]rune, 6)
	for i := range otpCode {
		otpCode[i] = digits[rand.Intn(len(digits))]
	}
	return string(otpCode)
}
