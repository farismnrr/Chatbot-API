package repository

import (
	database "capstone-project/database"
	"time"
)

type OTPRepository interface {
	SetOTP(userID string, otpCode string, expiry time.Time) error
	ValidateOTP(userID string, otpCode string) (bool, error)
}

type otpRepository struct {
	redis *database.Redis
}

func NewOTPRepository(redis *database.Redis) OTPRepository {
	return &otpRepository{redis: redis}
}

func (r *otpRepository) SetOTP(userID string, otpCode string, expiry time.Time) error {
	return r.redis.Client.Set(r.redis.Context, "otp:"+userID, otpCode, time.Until(expiry)).Err()
}

func (r *otpRepository) ValidateOTP(userID string, otpCode string) (bool, error) {
	otp, err := r.redis.Client.Get(r.redis.Context, "otp:"+userID).Result()
	if err != nil {
		return false, err
	}

	if otp != otpCode {
		return false, nil
	}

	err = r.redis.Client.Del(r.redis.Context, "otp:"+userID).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
