package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
	"time"
)

type sessionRepository struct {
	DB *database.Database
}

type SessionRepository interface {
	CreateSession(session model.Session) error
	UpdateSession(session model.Session) error
	DeleteSession(token string) error
	GetSessionByUsername(username string) (model.Session, error)
	GetSessionByToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
	TokenValidity(token string) (model.Session, error)
}

func NewSessionRepository(db *database.Database) *sessionRepository {
	return &sessionRepository{DB: db}
}

func (r *sessionRepository) CreateSession(session model.Session) error {
	return r.DB.DB.QueryRow("INSERT INTO Sessions (username, token, expiry) VALUES (?, ?, ?)", session.Username, session.Token, session.Expiry).Err()
}

func (r *sessionRepository) UpdateSession(session model.Session) error {
	return r.DB.DB.QueryRow("UPDATE Sessions SET token = ?, expiry = ? WHERE username = ?", session.Token, session.Expiry, session.Username).Err()
}

func (r *sessionRepository) DeleteSession(token string) error {
	return r.DB.DB.QueryRow("DELETE FROM Sessions WHERE token = ?", token).Err()
}

func (r *sessionRepository) GetSessionByUsername(username string) (model.Session, error) {
	var session model.Session
	err := r.DB.DB.QueryRow("SELECT username, token, expiry FROM Sessions WHERE username = ?", username).Scan(&session.Username, &session.Token, &session.Expiry)
	return session, err
}

func (r *sessionRepository) GetSessionByToken(token string) (model.Session, error) {
	var session model.Session
	err := r.DB.DB.QueryRow("SELECT username, token, expiry FROM Sessions WHERE token = ?", token).Scan(&session.Username, &session.Token, &session.Expiry)
	return session, err
}

func (r *sessionRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

func (r *sessionRepository) TokenValidity(token string) (model.Session, error) {
	session, err := r.GetSessionByToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if r.TokenExpired(session) {
		err := r.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}
