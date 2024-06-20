package model

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Conversation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Status    string    `json:"status"`
}

type Message struct {
	ID             int       `json:"id"`
	ConversationID int       `json:"conversation_id"`
	UserID         int       `json:"user_id"`
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Intent struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Entity struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Response struct {
	ID        int       `json:"id"`
	IntentID  int       `json:"intent_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Session struct {
	UserID   int       `json:"user_id"`
	Token    string    `json:"token"`
	Expiry   time.Time `json:"expiry"`
}

func NewUser(fullName string, username string, password string, email string) *User {
	return &User{FullName: fullName, Username: username, Password: password, Email: email}
}

func NewConversation(userID int) *Conversation {
	return &Conversation{UserID: userID}
}

func NewMessage(conversationID int, userID int, message string) *Message {
	return &Message{ConversationID: conversationID, UserID: userID, Message: message}
}

func NewIntent(name string, description string) *Intent {
	return &Intent{Name: name, Description: description}
}

func NewEntity(name string, value string) *Entity {
	return &Entity{Name: name, Value: value}
}

func NewResponse(intentID int, message string) *Response {
	return &Response{IntentID: intentID, Message: message}
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
