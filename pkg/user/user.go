package user

import (
	"crud/db"
	"fmt"
	"strings"
)

type User struct {
	ID   int
	Data `json:"data"`
}

type Data struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Interests string `json:"interests"`
}

// docker run --name=crud-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
// docker exec -it crud-db bash

func (u *User) CreateUser() error {
	query := `INSERT INTO users (data) VALUES ($1);`
	temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	data := strings.Join(temp, " ")
	_, err := db.DB.Exec(query, data)
	if err != nil {
		return fmt.Errorf("create user - %v", err)
	}
	return nil
}

func (u *User) GetUserByID() (string, error) {
	query := `SELECT data FROM users WHERE id = $1;`
	var data string
	if err := db.DB.QueryRow(query, u.ID).Scan(&data); err != nil {
		return "", fmt.Errorf("get user - %v", err)
	}

	return data, nil
}

func (u *User) UpdateUserByID() error {
	query := `UPDATE users SET data = $1 WHERE id = $2;`
	temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	data := strings.Join(temp, " ")

	_, err := db.DB.Exec(query, data, u.ID)
	if err != nil {
		return fmt.Errorf("update user - %v", err)
	}
	return nil
}

func (u *User) DeleteUserByID() error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.DB.Exec(query, u.ID)
	if err != nil {
		return fmt.Errorf("delete user - %v", err)
	}
	return nil
}
