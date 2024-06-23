package model

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type Conversation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type RequestMessage struct {
	ID             int       `json:"id,omitempty"`
	ConversationID int       `json:"conversation_id,omitempty"`
	UserID         int       `json:"user_id,omitempty"`
	Role           string    `json:"role,omitempty"`
	Message        string    `json:"message,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Session struct {
	UserID int       `json:"user_id"`
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}

func NewUser(fullName string, username string, password string, email string) *User {
	return &User{FullName: fullName, Username: username, Password: password, Email: email}
}

func NewConversation(userID int) Conversation {
	return Conversation{UserID: userID}
}

func NewRequestMessage(conversationID int, userID int, message string) *RequestMessage {
	return &RequestMessage{ConversationID: conversationID, UserID: userID, Message: message}
}

func NewSession(token string, userID int, expiry time.Time) *Session {
	return &Session{Token: token, UserID: userID, Expiry: expiry}
}
