package service

import (
	model "capstone-project/model"
	repository "capstone-project/repository"
)

type messageService struct {
	repo repository.MessageRepository
}

type MessageService interface {
	CreateMessage(conversation_id int, user_id int, message string, role string) error
	GetMessage(conversation_id int) ([]*model.RequestMessage, error)
	GetMessageById(id int) (*model.RequestMessage, error)
	DeleteMessage(id int) error
}

func NewMessageService(repo repository.MessageRepository) *messageService {
	return &messageService{repo: repo}
}

func (s *messageService) CreateMessage(conversation_id int, user_id int, message string, role string) error {
	return s.repo.CreateMessage(conversation_id, user_id, message, role)
}

func (s *messageService) GetMessage(conversation_id int) ([]*model.RequestMessage, error) {
	return s.repo.GetMessage(conversation_id)
}

func (s *messageService) GetMessageById(id int) (*model.RequestMessage, error) {
	return s.repo.GetMessageById(id)
}

func (s *messageService) DeleteMessage(id int) error {
	return s.repo.DeleteMessage(id)
}
