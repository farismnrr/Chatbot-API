package service

import (
	"capstone-project/helper"
	"capstone-project/model"
	"capstone-project/repository"
	"context"
	"time"
)

type sessionService struct {
	repository repository.SessionRepository
}

type SessionService interface {
	GenerateSession(ctx context.Context, userID int, username string) (*model.Session, error)
	GetSession(ctx context.Context, sessionID int) (string, error)
	DeleteSession(ctx context.Context, sessionID int) error
}

func NewSessionService(repository repository.SessionRepository) *sessionService {
	return &sessionService{repository: repository}
}

func (s *sessionService) GenerateSession(ctx context.Context, userID int, username string) (*model.Session, error) {
	token, err := helper.GenerateToken(username, "user", "active")
	if err != nil {
		return nil, err
	}
	session := &model.Session{
		UserID: userID,
		Token:  token,
		Expiry: time.Now().Add(time.Hour * 24),
	}
	err = s.repository.CreateSession(ctx, session.UserID, session.Token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) GetSession(ctx context.Context, sessionID int) (string, error) {
	session, err := s.repository.GetSession(ctx, sessionID)
	if err != nil {
		return "", err
	}
	return session, nil
}

func (s *sessionService) DeleteSession(ctx context.Context, sessionID int) error {
	return s.repository.DeleteSession(ctx, sessionID)
}
