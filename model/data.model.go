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
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MessageData struct {
	ID             int       `json:"id,omitempty"`
	ConversationID int       `json:"conversation_id,omitempty"`
	UserID         int       `json:"user_id,omitempty"`
	Message        string    `json:"message,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type RequestBody struct {
	Model    string `json:"model"`
	User     string `json:"user"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
		Name    string `json:"name"`
	} `json:"messages"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    MessageData `json:"data,omitempty"`
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

func NewMessage(conversationID int, userID int, message string) *MessageData {
	return &MessageData{ConversationID: conversationID, UserID: userID, Message: message}
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}

func NewSuccessResponse(code int, message string) *SuccessResponse {
	return &SuccessResponse{Code: code, Message: message}
}

func NewSession(token string, userID int, expiry time.Time) *Session {
	return &Session{Token: token, UserID: userID, Expiry: expiry}
}
