package service

import (
	"capstone-project/repository"
	"time"
)

type OTPService interface {
	SetOTP(userID string, otpCode string, expiry time.Time) error
	ValidateOTP(userID string, otpCode string) (bool, error)
}

type otpService struct {
	repository repository.OTPRepository
}

func NewOTPService(repository repository.OTPRepository) OTPService {
	return &otpService{repository: repository}
}

func (s *otpService) SetOTP(userID string, otpCode string, expiry time.Time) error {
	return s.repository.SetOTP(userID, otpCode, expiry)
}

func (s *otpService) ValidateOTP(userID string, otpCode string) (bool, error) {
	return s.repository.ValidateOTP(userID, otpCode)
}