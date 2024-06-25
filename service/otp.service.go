package service

import (
	"capstone-project/helper"
	"capstone-project/model"
	"capstone-project/repository"
	"context"
	"time"
)

type OTPService interface {
	GenerateOTP(ctx context.Context, userID int) (*model.OTP, error)
	GetOTP(ctx context.Context, userID int, otpCode string) (string, error)
}

type otpService struct {
	repository repository.OTPRepository
}

func NewOTPService(repository repository.OTPRepository) OTPService {
	return &otpService{repository: repository}
}

func (s *otpService) GenerateOTP(ctx context.Context, userID int) (*model.OTP, error) {
	otp := &model.OTP{
		UserID:  userID,
		OTPCode: helper.GenerateOTPCode(),
		Expiry:  time.Now().Add(time.Minute * 5),
	}
	err := s.repository.SetOTP(ctx, otp)
	return otp, err
}

func (s *otpService) GetOTP(ctx context.Context, userID int, otpCode string) (string, error) {
	otp, err := s.repository.GetOTP(ctx, userID, otpCode)
	if err != nil {
		return "", err
	}
	return otp, nil
}