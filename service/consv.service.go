package service

import (
	"capstone-project/model"
	"capstone-project/repository"
)

type conversationService struct {
	repo repository.ConversationRepository
}

type ConversationService interface {
	GetConversation(id int) (model.Conversation, error)
	CreateConversation(status string, conversation model.Conversation) error
	DeleteConversation(id int) error
}

func NewConversationService(repo repository.ConversationRepository) *conversationService {
	return &conversationService{repo: repo}
}

func (s *conversationService) GetConversation(id int, fullName string, status string) (*model.Conversation, error) {
	return s.repo.GetConversation(id, fullName, status)
}

func (s *conversationService) CreateConversation(status string, conversation model.Conversation) error {
	return s.repo.CreateConversation(status, conversation)
}

func (s *conversationService) DeleteConversation(id int) error {
	return s.repo.DeleteConversation(id)
}