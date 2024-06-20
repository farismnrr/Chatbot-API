package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
	"context"
	"strconv"
)

type sessionRepository struct {
	redis *database.Redis
}

type SessionRepository interface {
	CreateSession(ctx context.Context, session *model.Session) error
	GetSession(ctx context.Context, sessionID int) (string, error)
	DeleteSession(ctx context.Context, sessionID int) error
}

func NewSessionRepository(redis *database.Redis) *sessionRepository {
	return &sessionRepository{redis: redis}
}

func (r *sessionRepository) CreateSession(ctx context.Context, session *model.Session) error {
	key := "session:" + strconv.Itoa(session.UserID)
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
	return r.redis.Client.HSet(ctx, key, "token", session.Token, "expiry", session.Expiry).Err()
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
	return r.redis.Client.Del(ctx, key).Err()
}