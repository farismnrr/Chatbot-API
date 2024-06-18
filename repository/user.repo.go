package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
)

type Repository struct {
	DB *database.Database
}

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByUsername(user model.User) error
	GetUserByEmail(user model.User) error
	GetUserById(user model.User) error
	UpdateUserById(user model.User) error
	DeleteUserById(user model.User) error
}

func NewUserRepository(db *database.Database) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user model.User) error {
	return r.DB.DB.QueryRow("INSERT INTO Users (full_name, username, password, email) VALUES (?, ?, ?, ?)", user.FullName, user.Username, user.Password, user.Email).Err()
}

func (r *Repository) GetUserByUsername(user model.User) error {
	return r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Uers WHERE username = ?", user.Username).Err()
}

func (r *Repository) GetUserByEmail(user model.User) error {
	return r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Uers WHERE email = ?", user.Email).Err()
}

func (r *Repository) GetUserById(user model.User) error {
	return r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Users WHERE id = ?", user.ID).Err()
}

func (r *Repository) UpdateUserById(user model.User) error {
	return r.DB.DB.QueryRow("UPDATE Users SET full_name = ?, email = ?, password = ? WHERE id = ?", user.FullName, user.Email, user.Password, user.ID).Err()
}

func (r *Repository) DeleteUserById(user model.User) error {
	return r.DB.DB.QueryRow("DELETE FROM Users WHERE id = ?", user.ID).Err()
}
