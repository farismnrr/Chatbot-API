package repository

import (
	database "capstone-project/database"
	model "capstone-project/model"
	"log"
)

type messageRepository struct {
	DB *database.Database
}

type MessageRepository interface {
	CreateMessage(conversation_id int, user_id int, message string, role string) error
	GetMessage(conversation_id int) ([]*model.RequestMessage, error)
	GetMessageById(id int) (*model.RequestMessage, error)
	DeleteMessage(id int) error
}

func NewMessageRepository(db *database.Database) *messageRepository {
	return &messageRepository{DB: db}
}

func (r *messageRepository) CreateMessage(conversation_id int, user_id int, message string, role string) error {
	query := "INSERT INTO Messages (conversation_id, user_id, message, role) VALUES (?, ?, ?, ?)"
	_, err := r.DB.DB.Exec(query, conversation_id, user_id, message, role)
	return err
}

func (r *messageRepository) GetMessage(conversation_id int) ([]*model.RequestMessage, error) {
	log.Println("GetMessage called with conversation ID:", conversation_id)
	query := "SELECT id, conversation_id, user_id, role, message, created_at FROM Messages WHERE conversation_id = ?"
	log.Println("Query:", query)
	rows, err := r.DB.DB.Query(query, conversation_id)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()
	var messages []*model.RequestMessage
	for rows.Next() {
		var message model.RequestMessage
		err = rows.Scan(&message.ID, &message.ConversationID, &message.UserID, &message.Role, &message.Message, &message.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		messages = append(messages, &message)
	}
	log.Println("Returning", len(messages), "messages")
	return messages, err
}

func (r *messageRepository) GetMessageById(id int) (*model.RequestMessage, error) {
	query := "SELECT id, conversation_id, user_id, role, message, created_at FROM Messages WHERE id = ?"
	row := r.DB.DB.QueryRow(query, id)
	var message model.RequestMessage
	err := row.Scan(&message.ID, &message.ConversationID, &message.UserID, &message.Role, &message.Message, &message.CreatedAt)
	return &message, err
}

func (r *messageRepository) DeleteMessage(id int) error {
	query := "DELETE FROM Messages WHERE id = ?"
	_, err := r.DB.DB.Exec(query, id)
	return err
}
