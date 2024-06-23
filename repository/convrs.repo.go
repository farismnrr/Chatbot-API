package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
)

type conversationRepository struct {
	DB *database.Database
}

type ConversationRepository interface {
	GetAllConversations(user_id int) ([]*model.Conversation, error)
	GetConversation(user_id int) (*model.Conversation, error)
	GetConversationById(id int) (*model.Conversation, error)
	CreateConversation(user_id int, title string) error
	DeleteConversation(id int) error
	DeleteAllConversation(user_id int) error
}

func NewConversationRepository(db *database.Database) *conversationRepository {
	return &conversationRepository{DB: db}
}

func (r *conversationRepository) GetAllConversations(user_id int) ([]*model.Conversation, error) {
	query := "SELECT id, user_id, title, created_at, updated_at FROM Conversations WHERE user_id = ?"
	rows, err := r.DB.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var conversations []*model.Conversation
	for rows.Next() {
		var conversation model.Conversation
		err := rows.Scan(&conversation.ID, &conversation.UserID, &conversation.Title, &conversation.CreatedAt, &conversation.UpdatedAt)
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, &conversation)
	}
	return conversations, nil
}

func (r *conversationRepository) GetConversation(user_id int) (*model.Conversation, error) {
	query := "SELECT id, user_id, title, created_at, updated_at FROM Conversations WHERE user_id = ? ORDER BY id DESC LIMIT 1"
	row := r.DB.DB.QueryRow(query, user_id)
	var conversation model.Conversation
	err := row.Scan(&conversation.ID, &conversation.UserID, &conversation.Title, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) GetConversationById(id int) (*model.Conversation, error) {
	query := "SELECT id, user_id, title, created_at, updated_at FROM Conversations WHERE id = ?"
	row := r.DB.DB.QueryRow(query, id)
	var conversation model.Conversation
	err := row.Scan(&conversation.ID, &conversation.UserID, &conversation.Title, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) CreateConversation(user_id int, title string) error {
	query := "INSERT INTO Conversations (user_id, title) VALUES (?, ?)"
	_, err := r.DB.DB.Exec(query, user_id, title)
	return err
}

func (r *conversationRepository) DeleteConversation(id int) error {
	// Delete associated messages
	query := "DELETE FROM Messages WHERE conversation_id = ?"
	_, err := r.DB.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Delete conversation
	query = "DELETE FROM Conversations WHERE id = ?"
	_, err = r.DB.DB.Exec(query, id)
	return err
}

func (r *conversationRepository) DeleteAllConversation(user_id int) error {
	// Delete related messages
	query := "DELETE FROM Messages WHERE conversation_id IN (SELECT id FROM Conversations WHERE user_id = ?)"
	_, err := r.DB.DB.Exec(query, user_id)
	if err != nil {
		return err
	}

	// Then delete the conversations
	query = "DELETE FROM Conversations WHERE user_id = ?"
	_, err = r.DB.DB.Exec(query, user_id)
	return err
}
