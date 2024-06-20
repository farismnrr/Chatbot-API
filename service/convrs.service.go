package service

import (
	"capstone-project/model"
	"capstone-project/repository"
)

type conversationService struct {
	repo repository.ConversationRepository
}

type ConversationService interface {
	GetConversation(id int) (*model.Conversation, error)
	CreateConversation(conversation model.Conversation) error
	DeleteConversation(id int) error
}

func NewConversationService(repo repository.ConversationRepository) *conversationService {
	return &conversationService{repo: repo}
}

func (s *conversationService) GetConversation(id int) (*model.Conversation, error) {
	return s.repo.GetConversation(id)
}

func (s *conversationService) CreateConversation(conversation model.Conversation) error {
	return s.repo.CreateConversation(conversation)
}

func (s *conversationService) DeleteConversation(id int) error {
	return s.repo.DeleteConversation(id)
}
