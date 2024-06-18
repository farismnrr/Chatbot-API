package service

import (
	"capstone-project/model"
	"capstone-project/repository"
	"strings"
	"unicode"
)

type Service struct {
	repository repository.UserRepository
}

type UserService interface {
	Register(user model.User) error
	Login(user model.User) error
	GetUserByUsername(user model.User) error
	GetUserByEmail(user model.User) error
	GetUserById(user model.User) error
	UpdateUserById(user model.User) error
	DeleteUserById(user model.User) error
	CheckPassLength(pass string) bool
	CheckEmail(data string) bool
	CheckUsername(data string) bool
	CheckFullName(data string) bool
	HasUpperLetter(password string) bool
	HasLowerLetter(password string) bool
	HasNumber(password string) bool
	HasSpecialChar(password string) bool
}

func NewUserService(repository repository.UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Register(user model.User) error {
	return s.repository.CreateUser(user)
}

func (s *Service) Login(user model.User) error {
	return s.repository.GetUserByUsername(user)
}

func (s *Service) GetUserByUsername(user model.User) error {
	return s.repository.GetUserByUsername(user)
}

func (s *Service) GetUserByEmail(user model.User) error {
	return s.repository.GetUserByEmail(user)
}

func (s *Service) GetUserById(user model.User) error {
	return s.repository.GetUserById(user)
}

func (s *Service) UpdateUserById(user model.User) error {
	return s.repository.UpdateUserById(user)
}

func (s *Service) DeleteUserById(user model.User) error {
	return s.repository.DeleteUserById(user)
}

func (s *Service) CheckPassLength(pass string) bool {
	return len(pass) >= 8
}

func (s *Service) CheckEmail(data string) bool {
	return strings.Contains(data, "@") && strings.HasSuffix(data, ".com")
}

func (s *Service) CheckUsername(data string) bool {
	return !strings.Contains(data, " ")
}

func (s *Service) CheckFullName(data string) bool {
	for _, char := range data {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == ' ') {
			return false
		}
	}
	return true
}

func (s *Service) HasUpperLetter(password string) bool {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func (s *Service) HasLowerLetter(password string) bool {
	for _, c := range password {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func (s *Service) HasNumber(password string) bool {
	for _, c := range password {
		if unicode.IsNumber(c) {
			return true
		}
	}
	return false
}

func (s *Service) HasSpecialChar(password string) bool {
	for _, c := range password {
		if unicode.IsPunct(c) || unicode.IsSymbol(c) {
			return true
		}
	}
	return false
}
