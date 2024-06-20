package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
	"context"
	"strconv"
)

type OTPRepository interface {
	SetOTP(ctx context.Context, otp *model.OTP) error
	GetOTP(ctx context.Context, userID int, otpCode string) (string, error)
}

type otpRepository struct {
	redis *database.Redis
}

func NewOTPRepository(redis *database.Redis) OTPRepository {
	return &otpRepository{redis: redis}
}

func (r *otpRepository) SetOTP(ctx context.Context, otp *model.OTP) error {
	key := "otp:" + strconv.Itoa(otp.UserID)
	exists, err := r.redis.Client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists != 0 {
		err := r.redis.Client.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return r.redis.Client.HSet(ctx, key, "otp_code", otp.OTPCode, "expiry", otp.Expiry).Err()
}

func (r *otpRepository) GetOTP(ctx context.Context, userID int, otpCode string) (string, error) {
	key := "otp:" + strconv.Itoa(userID)
	otp, err := r.redis.Client.HGet(ctx, key, "otp_code").Result()
	if err != nil {
		return "", err
	}

	if otp != otpCode {
		return "", nil
	}
	return otp, nil
}