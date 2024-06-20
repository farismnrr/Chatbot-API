package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
)

type conversationRepository struct {
	DB *database.Database
}

type ConversationRepository interface {
	GetConversation(id int) (*model.Conversation, error)
	CreateConversation(conversation model.Conversation) error
	DeleteConversation(id int) error
}

func NewConversationRepository(db *database.Database) *conversationRepository {
	return &conversationRepository{DB: db}
}

func (r *conversationRepository) GetConversation(id int) (*model.Conversation, error) {
	// query := `
	// 		SELECT u.full_name, 
	// 			CASE 
	// 				WHEN c.status = 'online' THEN 'online'
	// 				ELSE 'offline'
	// 			END AS status,
	// 			c.created_at
	// 		FROM Users u
	// 		LEFT JOIN Conversations c ON u.id = c.user_id
	// 		WHERE c.id = $1
	// 		ORDER BY status DESC;
	// 	`
	// row := r.DB.DB.QueryRow(query, id)
	// var conversation model.Conversation
	// err := row.Scan(&fullName, &status)
	// if err != nil {
	// 	return nil, err
	// }
	// return &conversation, nil
	return &model.Conversation{}, nil
}

func (r *conversationRepository) CreateConversation(conversation model.Conversation) error {
	query := "INSERT INTO Conversations (user_id) VALUES (?)"
	_, err := r.DB.DB.Exec(query, conversation.UserID)
	return err
}

func (r *conversationRepository) DeleteConversation(id int) error {
	query := "DELETE FROM conversations WHERE id = $1"
	_, err := r.DB.DB.Exec(query, id)
	return err
}
