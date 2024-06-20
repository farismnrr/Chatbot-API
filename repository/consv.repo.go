package repository

import (
	"capstone-project/model"
	"database/sql"
)

type conversationRepository struct {
	db *sql.DB
}

type ConversationRepository interface {
	GetConversation(id int, fullName string, status string) (*model.Conversation, error)
	CreateConversation(status string, conversation model.Conversation) error
	DeleteConversation(id int) error
}

func NewConversationRepository(db *sql.DB) *conversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) GetConversation(id int, fullName string, status string) (*model.Conversation, error) {
	query := `
			SELECT u.full_name, 
				CASE 
					WHEN c.status = 'online' THEN 'online'
					ELSE 'offline'
				END AS status,
				c.created_at
			FROM Users u
			LEFT JOIN Conversations c ON u.id = c.user_id
			WHERE c.id = $1
			ORDER BY status DESC;
		`
	row := r.db.QueryRow(query, id)
	var conversation model.Conversation
	err := row.Scan(&fullName, &status)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) CreateConversation(status string, conversation model.Conversation) error {
	query := "INSERT INTO conversations (user_id, status) VALUES ($1, $2) RETURNING id"
	_, err := r.db.Exec(query, conversation.UserID, status)
	return err
}

func (r *conversationRepository) DeleteConversation(id int) error {
	query := "DELETE FROM conversations WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
