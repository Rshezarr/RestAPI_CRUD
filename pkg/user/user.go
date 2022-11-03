package user

import (
	"crud/db"
	"fmt"
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
	data := fmt.Sprintf("%s %s %s", u.Data.FirstName, u.Data.LastName, u.Data.Interests)
	_, err := db.DB.NamedExec(`INSERT INTO users (data) VALUES (:data);`, map[string]interface{}{"data": data})
	if err != nil {
		return fmt.Errorf("create user - %v", err)
	}
	return nil
}

func (u *User) GetUserByID() ([]string, error) {
	var data []string
	err := db.DB.Select(&data, `SELECT data FROM users WHERE id = $1;`, u.ID)
	if err != nil {
		return nil, fmt.Errorf("get user - %v", err)
	}

	return data, nil
}

func (u *User) UpdateUserByID() error {
	data := fmt.Sprintf("%s %s %s", u.Data.FirstName, u.Data.LastName, u.Data.Interests)
	_, err := db.DB.NamedExec(`UPDATE users SET data = :data WHERE id = :id;`, map[string]interface{}{"data": data, "id": u.ID})
	if err != nil {
		return fmt.Errorf("update user - %v", err)
	}
	return nil
}

func (u *User) DeleteUserByID() error {
	_, err := db.DB.NamedExec(`DELETE FROM users WHERE id = :id;`, map[string]interface{}{"id": u.ID})
	if err != nil {
		return fmt.Errorf("delete user - %v", err)
	}
	return nil
}
