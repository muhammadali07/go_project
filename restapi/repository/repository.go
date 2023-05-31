package repository

import (
	"github.com/muhammadali07/go_project/db"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(users *[]User) error {
	db := db.GetDB()
	err := db.Select(users, "SELECT * FROM users")
	return err
}

func GetUser(id string) (*User, error) {
	db := db.GetDB()
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return &user, err
}

func CreateUser(user *User) error {
	db := db.GetDB()
	err := db.Get(user, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *", user.Name, user.Email)
	return err
}

func UpdateUser(id string, user *User) error {
	db := db.GetDB()
	err := db.Get(user, "UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING *", user.Name, user.Email, id)
	return err
}

func DeleteUser(id string) error {
	db := db.GetDB()
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
