package repository

import (
	database "capstone-project/database"
	"capstone-project/model"
)

type userRepository struct {
	DB *database.Database
}

type UserRepository interface {
	GetUserTable() (*model.User, error)
	CreateUser(user model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUserByEmail(email string, user model.User) error
	DeleteUserById(id int) error
}

func NewUserRepository(db *database.Database) *userRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetUserTable() (*model.User, error) {
	var user model.User
	err := r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Users").Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (r *userRepository) CreateUser(user model.User) error {
	return r.DB.DB.QueryRow("INSERT INTO Users (full_name, username, password, email) VALUES (?, ?, ?, ?)", user.FullName, user.Username, user.Password, user.Email).Err()
}

func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Users WHERE username = ?", username).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Users WHERE email = ?", email).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (r *userRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.DB.DB.QueryRow("SELECT id, full_name, username, password, email, created_at, updated_at FROM Users WHERE id = ?", id).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (r *userRepository) UpdateUserByEmail(email string, user model.User) error {
	return r.DB.DB.QueryRow("UPDATE Users SET password = ? WHERE email = ?", user.Password, email).Err()
}

func (r *userRepository) DeleteUserById(id int) error {
	return r.DB.DB.QueryRow("DELETE FROM Users WHERE id = ?", id).Err()
}