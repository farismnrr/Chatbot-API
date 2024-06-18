package service

import (
	"capstone-project/model"
	"capstone-project/repository"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
)

type userService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

type UserService interface {
	Register(user model.User) error
	Login(user model.User) (token *string, err error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUserById(id int, user model.User) error
	DeleteUserById(id int) error
	CheckPassLength(pass string) bool
	CheckEmail(data string) bool
	CheckUsername(data string) bool
	CheckFullName(data string) bool
	HasUpperLetter(password string) bool
	HasLowerLetter(password string) bool
	HasNumber(password string) bool
	HasSpecialChar(password string) bool
	GenerateHash(password string) string
}

func NewUserService(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) *userService {
	return &userService{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (s *userService) Register(user model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) Login(user model.User) (token *string, err error) {
	fetchedUser, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		Username: fetchedUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		Token:    tokenString,
		Username: fetchedUser.Username,
		Expiry:   expirationTime,
	}

	if _, err := s.sessionRepo.GetSessionByUsername(fetchedUser.Username); err != nil {
		s.sessionRepo.CreateSession(session)
	} else {
		s.sessionRepo.UpdateSession(session)
	}

	return &tokenString, nil
}

func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserById(id int) (*model.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) UpdateUserById(id int, user model.User) error {
	return s.userRepo.UpdateUserById(id, user)
}

func (s *userService) DeleteUserById(id int) error {
	return s.userRepo.DeleteUserById(id)
}

func (s *userService) CheckPassLength(pass string) bool {
	return len(pass) >= 8
}

func (s *userService) CheckEmail(data string) bool {
	return strings.Contains(data, "@") && strings.HasSuffix(data, ".com")
}

func (s *userService) CheckUsername(data string) bool {
	return !strings.Contains(data, " ")
}

func (s *userService) CheckFullName(data string) bool {
	for _, char := range data {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == ' ') {
			return false
		}
	}
	return true
}

func (s *userService) HasUpperLetter(password string) bool {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func (s *userService) HasLowerLetter(password string) bool {
	for _, c := range password {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func (s *userService) HasNumber(password string) bool {
	for _, c := range password {
		if unicode.IsNumber(c) {
			return true
		}
	}
	return false
}

func (s *userService) HasSpecialChar(password string) bool {
	for _, c := range password {
		if unicode.IsPunct(c) || unicode.IsSymbol(c) {
			return true
		}
	}
	return false
}

func (s *userService) GenerateHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
