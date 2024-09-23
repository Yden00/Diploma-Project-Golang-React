package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int
	Username string
	Password string
}

func CreateUser(db *sql.DB, username, password string) (*User, error) {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &User{Username: username, Password: password}, nil
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}
