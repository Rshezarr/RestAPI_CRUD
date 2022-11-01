package user

import (
	"crud/db"
	"fmt"
)

type User struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// docker run --name=crud-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
// docker exec -it crud-db bash

func (u *User) CreateUser() error {
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2);`
	_, err := db.DB.Exec(query, u.FirstName, u.LastName)
	if err != nil {
		return fmt.Errorf("create user - %v", err)
	}
	return nil
}

func (u *User) GetUserByID() (User, error) {
	query := `SELECT * FROM users WHERE id = $1;`
	var usr User
	if err := db.DB.QueryRow(query, u.ID).Scan(&usr.ID, &usr.FirstName, &usr.LastName); err != nil {
		return User{}, fmt.Errorf("get user - %v", err)
	}
	return usr, nil
}

func (u *User) UpdateUserByID() error {
	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;`
	_, err := db.DB.Exec(query, u.FirstName, u.LastName, u.ID)
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
