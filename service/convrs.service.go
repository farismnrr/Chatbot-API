package service

import (
	"capstone-project/model"
	"capstone-project/repository"
)

type conversationService struct {
	repo repository.ConversationRepository
}

type ConversationService interface {
	GetAllConversations(user_id int) ([]*model.Conversation, error)
	GetConversation(user_id int) (*model.Conversation, error)
	GetConversationById(id int) (*model.Conversation, error)
	CreateConversation(user_id int, title string) error
	DeleteConversation(id int) error
	DeleteAllConversation(user_id int) error
}

func NewConversationService(repo repository.ConversationRepository) *conversationService {
	return &conversationService{repo: repo}
}

func (s *conversationService) GetAllConversations(user_id int) ([]*model.Conversation, error) {
	return s.repo.GetAllConversations(user_id)
}

func (s *conversationService) GetConversation(user_id int) (*model.Conversation, error) {
	return s.repo.GetConversation(user_id)
}

func (s *conversationService) GetConversationById(id int) (*model.Conversation, error) {
	return s.repo.GetConversationById(id)
}

func (s *conversationService) CreateConversation(user_id int, title string) error {
	return s.repo.CreateConversation(user_id, title)
}

func (s *conversationService) DeleteConversation(id int) error {
	return s.repo.DeleteConversation(id)
}

func (s *conversationService) DeleteAllConversation(user_id int) error {
	return s.repo.DeleteAllConversation(user_id)
}
