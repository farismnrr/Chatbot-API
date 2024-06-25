package repository

import (
	database "capstone-project/database"
	"context"
	"strconv"
)

type sessionRepository struct {
	redis *database.Redis
}

type SessionRepository interface {
	CreateSession(ctx context.Context, sessionID int, token string) error
	GetSession(ctx context.Context, sessionID int) (string, error)
	DeleteSession(ctx context.Context, sessionID int) error
}

func NewSessionRepository(redis *database.Redis) *sessionRepository {
	return &sessionRepository{redis: redis}
}

func (r *sessionRepository) CreateSession(ctx context.Context, sessionID int, token string) error {
	key := "session:" + strconv.Itoa(sessionID)
	return r.redis.Client.HSet(ctx, key, "token", token).Err()
}

func (r *sessionRepository) GetSession(ctx context.Context, sessionID int) (string, error) {
	key := "session:" + strconv.Itoa(sessionID)
	session, err := r.redis.Client.HGet(ctx, key, "token").Result()
	if err != nil {
		return "", err
	}
	return session, nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, sessionID int) error {
	key := "session:" + strconv.Itoa(sessionID)
	err := r.redis.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	tokenKey := "token:" + strconv.Itoa(sessionID)
	return r.redis.Client.Del(ctx, tokenKey).Err()
}
