package service

import (
	"capstone-project/model"
	"capstone-project/repository"
)

type sessionService struct {
	repository repository.SessionRepository
}

type SessionService interface {
}

func NewSessionService(repository repository.SessionRepository) *sessionService {
	return &sessionService{repository: repository}
}

func (s *sessionService) GetSessionByUsername(username string) (model.Session, error) {
	return s.repository.GetSessionByUsername(username)
}
